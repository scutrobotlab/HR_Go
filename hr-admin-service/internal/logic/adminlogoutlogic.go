package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogoutLogic {
	return &AdminLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLogoutLogic) AdminLogout(in *hr_admin_service.AdminIdReq) (*hr_admin_service.StatusResp, error) {
	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
