package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigLogic {
	return &GetConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigLogic) GetConfig(in *hr_admin_service.GetConfigReq) (*hr_admin_service.GetConfigResp, error) {
	c := l.svcCtx.Query.Configuration
	config, err := c.WithContext(l.ctx).Where(c.Key.Eq(in.Key)).First()
	if err != nil {
		return &hr_admin_service.GetConfigResp{
			Key:   config.Key,
			Value: "",
		}, nil
	}

	return &hr_admin_service.GetConfigResp{
		Key:   config.Key,
		Value: config.Value,
	}, nil
}
