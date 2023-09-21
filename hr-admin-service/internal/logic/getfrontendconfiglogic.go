package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrontendConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFrontendConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrontendConfigLogic {
	return &GetFrontendConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFrontendConfigLogic) GetFrontendConfig(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetFrontendConfigResp, error) {
	common := l.svcCtx.Common

	return &hr_admin_service.GetFrontendConfigResp{
		ClientId:    common.OAuth.ClientId,
		RedirectUri: common.OAuth.RedirectUrl,
		AdminUri:    common.Url.AdminUrl,
		WechatUri:   common.Url.WechatUrl,
		BaseUri:     common.Url.BaseUrl,
		Debug:       common.Debug,
	}, nil
}
