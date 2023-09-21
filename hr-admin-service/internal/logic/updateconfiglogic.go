package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigLogic {
	return &UpdateConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigLogic) UpdateConfig(in *hr_admin_service.UpdateConfigReq) (*hr_admin_service.UpdateConfigResp, error) {
	c := l.svcCtx.Query.Configuration
	err := c.WithContext(l.ctx).Where(c.Key.Eq(in.Key)).Save(&model.Configuration{
		Key:   in.Key,
		Value: in.Value,
	})
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.UpdateConfigResp{
		Key:   in.Key,
		Value: in.Value,
	}, nil
}
