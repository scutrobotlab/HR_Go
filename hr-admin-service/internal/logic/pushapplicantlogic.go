package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushApplicantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushApplicantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushApplicantLogic {
	return &PushApplicantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushApplicantLogic) PushApplicant(in *hr_admin_service.PushApplicantReq) (*hr_admin_service.StatusResp, error) {
	at := l.svcCtx.Query.ApplicantTime
	r := l.svcCtx.Query.Room

	applicantTime, err := at.WithContext(l.ctx).Where(at.ID.Eq(in.ApplicantTimeId)).First()
	if err != nil {
		return nil, err
	}

	room, err := r.WithContext(l.ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Where(r.ID.Eq(in.RoomId)).First()
	if err != nil {
		return nil, err
	}
	if room.Status != RoomWaiting {
		return nil, fmt.Errorf("房间状态不是等待中")
	}

	room.Status = RoomInUse
	room.ApplicantTimeID = in.ApplicantTimeId
	room.UpdatedBy = in.AdminId
	room.StatusUpdatedAt = time.Now()
	room.GroupLabel = applicantTime.Group_
	_, err = r.WithContext(l.ctx).Where(r.ID.Eq(room.ID)).Updates(room)
	if err != nil {
		return nil, err
	}
	_, err = at.WithContext(l.ctx).Where(at.ID.Eq(in.ApplicantTimeId)).UpdateColumn(at.RoomID, room.ID)
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
