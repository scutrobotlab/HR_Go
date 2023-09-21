package config

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetConfigLogic {
	return &SetConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetConfigLogic) SetConfig(req *types.SetConfigReq) (resp *types.SetConfigResp, err error) {
	config, err := l.svcCtx.AdminService.UpdateConfig(l.ctx, &hr_admin_service.UpdateConfigReq{
		Key:   req.Key,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &types.SetConfigResp{
		Key:   config.Key,
		Value: config.Value,
	}, nil
}
