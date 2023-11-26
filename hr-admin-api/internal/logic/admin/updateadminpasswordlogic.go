package admin

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"fmt"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPasswordLogic {
	return &UpdateAdminPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminPasswordLogic) UpdateAdminPassword(req *types.UpdateAdminPasswordReq) error {
	userInfo := common.GetUserInfo(l.ctx)
	// 只有超级管理员和自己可以修改密码
	if userInfo.Permission != common.SuperAdmin && userInfo.Id != req.AdminId {
		return common.GrpcErrorPermissionDenied(fmt.Errorf("permission denied"))
	}

	_, err := l.svcCtx.AdminService.UpdateAdminPassword(l.ctx, &hr_admin_service.UpdateAdminPasswordReq{
		AdminId:  req.AdminId,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	return nil
}
