package admin

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(req *types.UpdateAdminReq) (resp *types.UpdateAdminResp, err error) {
	_, err = l.svcCtx.AdminService.UpdateAdmin(l.ctx, &hr_admin_service.UpdateAdminReq{
		AdminId: req.AdminId,
		Admin: &hr_admin_service.Admin{
			Id:   req.AdminId,
			Name: req.Name,
			Groups: lo.Map(req.Groups, func(item string, index int) *hr_admin_service.AdminGroup {
				return &hr_admin_service.AdminGroup{Name: item}
			}),
			DefaultStandardId: 0,
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
