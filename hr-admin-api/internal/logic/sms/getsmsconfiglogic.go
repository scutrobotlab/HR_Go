package sms

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsConfigLogic {
	return &GetSmsConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsConfigLogic) GetSmsConfig() (resp *types.GetSmsConfigResp, err error) {
	configResp, err := l.svcCtx.AdminService.GetSmsConfig(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetSmsConfigResp{
		Types:  configResp.Types,
		Alerts: configResp.Alerts,
	}, nil
}
