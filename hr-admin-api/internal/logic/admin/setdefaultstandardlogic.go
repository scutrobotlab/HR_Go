package admin

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultStandardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDefaultStandardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultStandardLogic {
	return &SetDefaultStandardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDefaultStandardLogic) SetDefaultStandard(req *types.SetDefaultStandardReq) error {
	_, err := l.svcCtx.AdminService.SetDefaultStandard(l.ctx, &hr_admin_service.SetDefaultStandardReq{
		AdminId:    common.GetUserInfo(l.ctx).Id,
		StandardId: req.StandardId,
	})
	if err != nil {
		return err
	}

	return nil
}
