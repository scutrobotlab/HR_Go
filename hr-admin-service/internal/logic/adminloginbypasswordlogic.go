package logic

import (
	"HR_Go/common"
	"context"
	"fmt"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginByPasswordLogic {
	return &AdminLoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginByPasswordLogic) AdminLoginByPassword(in *hr_admin_service.AdminLoginByPasswordReq) (*hr_admin_service.AdminLoginResp, error) {
	a := l.svcCtx.Query.Admin

	admin, err := a.WithContext(l.ctx).Where(a.Name.Eq(in.Name)).First()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	if !common.VerifyPassword(l.svcCtx.Common.Encrypt.Salt, admin.ID, in.Password, admin.Password) {
		return nil, common.GrpcErrorPermissionDenied(fmt.Errorf("password error"))
	}

	return &hr_admin_service.AdminLoginResp{
		Admin: &hr_admin_service.Admin{
			Id:                admin.ID,
			Name:              admin.Name,
			Groups:            []*hr_admin_service.AdminGroup{},
			DefaultStandardId: admin.DefaultStandardID,
			SmsEnabled:        admin.SmsEnabled,
		},
	}, nil
}
