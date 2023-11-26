package admin

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginByPasswordLogic {
	return &AdminLoginByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginByPasswordLogic) AdminLoginByPassword(req *types.AdminLoginByPasswordReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.AdminService.AdminLoginByPassword(l.ctx, &hr_admin_service.AdminLoginByPasswordReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	userInfo := common.UserInfo{
		Id:         loginResp.Admin.Id,
		Permission: common.Admin,
	}
	if userInfo.Id == 1 {
		userInfo.Permission = common.SuperAdmin
	}
	token, err := common.GetLoginJwtToken(l.svcCtx.Config.Auth, userInfo)
	if err != nil {
		return nil, err
	}

	admin := loginResp.Admin
	return &types.LoginResp{
		Id:     admin.Id,
		Name:   admin.Name,
		UUID:   admin.Uuid,
		Email:  admin.Email,
		Avatar: admin.Avatar,
		Groups: lo.Map(admin.Groups, func(item *hr_admin_service.AdminGroup, index int) types.AdminGroup {
			return types.AdminGroup{
				Id:   item.Id,
				Name: item.Name,
				Pivot: types.AdminGroupPivot{
					UserId:  admin.Id,
					GroupId: item.Id,
					Power:   item.Power,
				},
			}
		}),
		Token: token,
	}, nil
}
