package logic

import (
	"HR_Go/common"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminPasswordLogic {
	return &UpdateAdminPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminPasswordLogic) UpdateAdminPassword(in *hr_admin_service.UpdateAdminPasswordReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Admin

	encryptedPassword := common.EncryptPassword(l.svcCtx.Common.Encrypt.Salt, in.AdminId, in.Password)
	_, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).UpdateColumn(a.Password, encryptedPassword)
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
