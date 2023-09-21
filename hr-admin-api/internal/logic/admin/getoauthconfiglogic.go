package admin

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOAuthConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOAuthConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOAuthConfigLogic {
	return &GetOAuthConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOAuthConfigLogic) GetOAuthConfig() (resp *types.OAuthConfig, err error) {
	userInfo := common.GetUserInfo(l.ctx)
	authConfig, err := l.svcCtx.AdminService.GetOAuthConfig(l.ctx, &hr_admin_service.AdminIdReq{AdminId: userInfo.Id})
	if err != nil {
		return nil, err
	}

	clientSecret := "<敏感字段，仅允许超级管理员查看>"
	if userInfo.Permission == common.SuperAdmin {
		clientSecret = authConfig.ClientSecret
	}

	return &types.OAuthConfig{
		ClientId:     authConfig.ClientId,
		ClientSecret: clientSecret,
		AdminUri:     authConfig.AdminUri,
		WechatUri:    authConfig.WechatUri,
		RedirectUri:  authConfig.RedirectUri,
	}, nil
}
