package config

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsServiceConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsServiceConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsServiceConfigLogic {
	return &GetSmsServiceConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsServiceConfigLogic) GetSmsServiceConfig() (resp *types.GetSmsServiceConfigResp, err error) {
	config, err := l.svcCtx.AdminService.GetSmsServiceConfig(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetSmsServiceConfigResp{
		Debug:    config.Debug,
		Enabled:  config.Enabled,
		SendUrl:  config.SendUrl,
		UserName: config.UserName,
	}, nil
}
