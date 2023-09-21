package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPushableListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPushableListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPushableListLogic {
	return &GetPushableListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPushableListLogic) GetPushableList(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetPushableListResp, error) {
	db := l.svcCtx.Db
	var result []*hr_admin_service.PushableApplicant
	err := db.WithContext(l.ctx).Table("applicant_times as at").
		Select("a.name, t.group, time(t.time) as time, a.id as applicant_id, at.id as applicant_time_id").
		Joins("left join scores as s on at.applicant_id = s.applicant_id and at.group = s.group").
		Joins("left join applicants as a on at.applicant_id = a.id").
		Joins("left join times as t on at.time_id = t.id and t.deleted_at is null").
		Where("at.deleted_at is null and at.room_id = 0").
		Where("(JSON_EXTRACT(at.extend, '$.receive_status') = '准时签到' or JSON_EXTRACT(at.extend, '$.receive_status') = '迟到签到')").
		Where("s.id is null").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.GetPushableListResp{
		Applicants: result,
	}, nil
}
