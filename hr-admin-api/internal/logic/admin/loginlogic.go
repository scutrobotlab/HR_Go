package admin

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.AdminService.AdminLogin(l.ctx, &hr_admin_service.AdminLoginReq{Code: req.Code})
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
