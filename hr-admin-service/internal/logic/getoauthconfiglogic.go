package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOAuthConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOAuthConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAuthConfigLogic {
	return &GetOAuthConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOAuthConfigLogic) GetOAuthConfig(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetOAuthConfigResp, error) {
	common := l.svcCtx.Common

	return &hr_admin_service.GetOAuthConfigResp{
		ClientId:     common.OAuth.ClientId,
		ClientSecret: common.OAuth.ClientSecret,
		AdminUri:     common.Url.AdminUrl,
		WechatUri:    common.Url.WechatUrl,
		RedirectUri:  common.OAuth.RedirectUrl,
	}, nil
}
