package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/hrservice"
	"HR_Go/hr-admin-service/internal/svc"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"
	"time"
)

type (
	Group struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Power int64  `json:"power"`
	}

	AdminProfile struct {
		UUID   string  `json:"uuid"`
		Email  string  `json:"email"`
		Avatar string  `json:"avatar"`
		Groups []Group `json:"groups"`
	}
)

const (
	RoomStopped  = 0 // 房间已停用
	RoomRelaxing = 1 // 房间休息中
	RoomWaiting  = 2 // 房间等待中
	RoomInUse    = 3 // 房间使用中

	InterviewerComment = 1 // 面试官留言
	ReceiverComment    = 2 // 接待者留言
)

const InterviewDuration = 20 * time.Minute // 面试时长

func getAbstractAdmin(ctx context.Context, svcCtx *svc.ServiceContext, id int64) (*hr_admin_service.AbstractAdmin, error) {
	a := svcCtx.Query.Admin
	admin, err := a.WithContext(ctx).Where(a.Where(a.ID.Eq(id))).First()
	if err != nil {
		return nil, err
	}
	return &hrservice.AbstractAdmin{
		Id:                admin.ID,
		Name:              admin.Name,
		DefaultStandardId: admin.DefaultStandardID,
	}, err
}

func getIntentsByApplicantId(ctx context.Context, svcCtx *svc.ServiceContext, applicantId int64) ([]*hr_admin_service.Intent, error) {
	i := svcCtx.Query.Intent
	intents, err := i.WithContext(ctx).Where(i.ApplicantID.Eq(applicantId)).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(intents, func(item *model.Intent, index int) *hr_admin_service.Intent {
		var parallel = int32(0)
		if item.Parallel {
			parallel = 1
		}
		return &hr_admin_service.Intent{
			Group:       item.Group_,
			ApplicantId: item.ApplicantID,
			Rank:        item.Rank,
			Parallel:    parallel,
		}
	}), nil
}

func getTimesByApplicantId(ctx context.Context, svcCtx *svc.ServiceContext, applicantId int64) ([]*hr_admin_service.Time, error) {
	type Result struct {
		Id              int64
		Group           string
		Time            time.Time
		Rank            int32
		Location        string
		TotalCnt        int32
		Grade           string
		Campus          string
		Extend          string
		ApplicantTimeId int64
	}

	db := svcCtx.Db
	var result []*Result
	err := db.WithContext(ctx).Debug().
		Table("applicant_times as at").
		Select("t.id, t.group, t.time, t.rank, t.location, t.total_cnt, t.grade, t.campus, at.extend, at.id as applicant_time_id").
		Joins("left join times as t on at.time_id = t.id").
		Where("at.applicant_id = ?", applicantId).
		Where("at.deleted_at is null").
		Where("t.deleted_at is null").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return lo.Map(result, func(item *Result, index int) *hr_admin_service.Time {
		return &hr_admin_service.Time{
			Id:              item.Id,
			Group:           item.Group,
			Date:            item.Time.Format(time.DateOnly),
			Time:            item.Time.Format(time.TimeOnly),
			Rank:            item.Rank,
			Location:        item.Location,
			TotalCnt:        item.TotalCnt,
			Grade:           item.Grade,
			Campus:          item.Campus,
			ApplicantId:     applicantId,
			Extend:          item.Extend,
			ApplicantTimeId: item.ApplicantTimeId,
		}
	}), nil
}

func getAdmitByApplicantId(ctx context.Context, svcCtx *svc.ServiceContext, applicantId int64) ([]*hr_admin_service.Admit, error) {
	a := svcCtx.Query.Admit
	admits, err := a.WithContext(ctx).Where(a.ApplicantID.Eq(applicantId)).Find()
	if err != nil {
		return nil, err
	}

	return lo.Map(admits, func(item *model.Admit, index int) *hr_admin_service.Admit {
		admin, err := getAbstractAdmin(ctx, svcCtx, item.AdminID)
		if err != nil {
			return nil
		}

		return &hr_admin_service.Admit{
			AdminId:     item.AdminID,
			ApplicantId: item.ApplicantID,
			Group:       item.Group_,
			Admin:       admin,
			UpdatedAt:   item.UpdatedAt.Format(time.DateTime),
		}
	}), nil
}

func GetConfigOrDefault(ctx context.Context, svcCtx *svc.ServiceContext, key, defaultValue string) string {
	c := svcCtx.Query.Configuration
	config, err := c.WithContext(ctx).Where(c.Key.Eq(key)).First()
	if err != nil {
		_ = c.WithContext(ctx).Create(&model.Configuration{
			Key:   key,
			Value: defaultValue,
		})
		return defaultValue
	}

	return config.Value
}
