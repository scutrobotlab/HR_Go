package logic

import (
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllRoomLogic {
	return &GetAllRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllRoomLogic) GetAllRoom(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetAllRoomResp, error) {
	db := l.svcCtx.Db
	a := l.svcCtx.Query.Admin
	r := l.svcCtx.Query.Room

	admin, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).First()
	if err != nil {
		return nil, err
	}
	myRoomName, myRoomId := "", int64(0)
	room, err := r.WithContext(l.ctx).Where(r.ID.Eq(admin.RoomID)).First()
	if room != nil {
		myRoomName = room.Name
		myRoomId = room.ID
	}

	type Result struct {
		ID                 int64
		Name               string
		Location           string
		Status             int32
		Group              string
		ApplicantId        int64
		ApplicantName      string
		Time               time.Time
		StatusUpdatedAt    time.Time
		UpdatedAt          time.Time
		UpdatedBy          int64
		InterviewerComment string
		ReceiverComment    string
		GroupLabel         string
	}

	var result []Result
	err = db.WithContext(l.ctx).Table("rooms as r").
		Select("r.id, r.name, r.location, r.status, at.group," +
			" a.id as applicant_id, a.name as applicant_name, t.time, r.status_updated_at, " +
			"r.updated_at, r.updated_by, r.interviewer_comment, r.receiver_comment, r.group_label").
		Joins("left join applicant_times as at on r.applicant_time_id = at.id and at.deleted_at is null").
		Joins("left join applicants as a on at.applicant_id = a.id and a.deleted_at is null").
		Joins("left join times as t on at.time_id = t.id and t.deleted_at is null").
		Where("r.deleted_at is null").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.GetAllRoomResp{
		Rooms: lo.Map(result, func(item Result, index int) *hr_admin_service.Room {
			admins, err := a.WithContext(l.ctx).Where(a.RoomID.Eq(item.ID)).Find()
			if err != nil {
				admins = []*model.Admin{}
			}
			applicantId, applicant, group, time1 := int64(0), "", "", ""
			if item.Status == RoomInUse {
				group = item.Group
				applicantId = item.ApplicantId
				applicant = item.ApplicantName
				time1 = item.Time.Format(time.DateTime)
			}
			updatedBy, _ := a.WithContext(l.ctx).Where(a.ID.Eq(item.UpdatedBy)).First()
			updatedByName := ""
			if updatedBy != nil {
				updatedByName = updatedBy.Name
			}
			return &hr_admin_service.Room{
				Id:       item.ID,
				Name:     item.Name,
				Location: item.Location,
				Status:   item.Status,
				Admins: lo.Map(admins, func(item1 *model.Admin, index int) string {
					return item1.Name
				}),
				Applicant:          applicant,
				Group:              group,
				Time:               time1,
				UpdatedAt:          item.StatusUpdatedAt.Format(time.DateTime),
				ApplicantId:        applicantId,
				UpdatedBy:          updatedByName,
				InterviewerComment: item.InterviewerComment,
				ReceiverComment:    item.ReceiverComment,
				GroupLabel:         item.GroupLabel,
			}
		}),
		MyRoomName: myRoomName,
		MyRoomId:   myRoomId,
	}, nil
}
