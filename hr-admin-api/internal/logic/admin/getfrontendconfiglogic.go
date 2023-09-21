package admin

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrontendConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFrontendConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrontendConfigLogic {
	return &GetFrontendConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFrontendConfigLogic) GetFrontendConfig() (resp *types.FrontendConfig, err error) {
	configResp, err := l.svcCtx.AdminService.GetFrontendConfig(l.ctx, &hr_admin_service.AdminIdReq{AdminId: 0})
	if err != nil {
		return nil, err
	}

	return &types.FrontendConfig{
		Debug: configResp.Debug,
		OAuth: types.OAuth{
			ClientId:    configResp.ClientId,
			RedirectUri: configResp.RedirectUri,
		},
		Url: types.Url{
			AdminUri:  configResp.AdminUri,
			WechatUri: configResp.WechatUri,
			BaseUri:   configResp.BaseUri,
		},
	}, nil
}
