package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	SmsServiceEnabled = "SmsServiceEnabled"
)

type GetSmsServiceConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsServiceConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsServiceConfigLogic {
	return &GetSmsServiceConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsServiceConfigLogic) GetSmsServiceConfig(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetSmsServiceConfigResp, error) {
	enabled := GetConfigOrDefault(l.ctx, l.svcCtx, SmsServiceEnabled, "0")
	smsBaoConfig := l.svcCtx.Config.SmsBaoConfig

	return &hr_admin_service.GetSmsServiceConfigResp{
		Debug:    smsBaoConfig.Debug,
		Enabled:  enabled == "1",
		SendUrl:  smsBaoConfig.SendUrl,
		UserName: smsBaoConfig.Username,
	}, nil
}
