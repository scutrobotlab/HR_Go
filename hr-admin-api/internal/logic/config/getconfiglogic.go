package config

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigLogic {
	return &GetConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigLogic) GetConfig(req *types.GetConfigReq) (resp *types.GetConfigResp, err error) {
	config, err := l.svcCtx.AdminService.GetConfig(l.ctx, &hr_admin_service.GetConfigReq{Key: req.Key})
	if err != nil {
		return nil, err
	}

	return &types.GetConfigResp{
		Key:   config.Key,
		Value: config.Value,
	}, nil
}
