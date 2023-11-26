package admin

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdminLogic {
	return &CreateAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAdminLogic) CreateAdmin(req *types.CreateAdminReq) (resp *types.CreateAdminResp, err error) {
	createAdminResp, err := l.svcCtx.AdminService.CreateAdmin(l.ctx, &hr_admin_service.CreateAdminReq{
		Admin: &hr_admin_service.Admin{
			Name: req.Name,
			Groups: lo.Map(req.Groups, func(item string, index int) *hr_admin_service.AdminGroup {
				return &hr_admin_service.AdminGroup{Name: item}
			}),
		},
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateAdminResp{
		Id: createAdminResp.Admin.Id,
	}, nil
}
