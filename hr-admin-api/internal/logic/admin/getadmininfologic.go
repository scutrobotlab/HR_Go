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

type GetAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminInfoLogic {
	return &GetAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminInfoLogic) GetAdminInfo() (resp *types.GetAdminResp, err error) {
	info, err := l.svcCtx.AdminService.GetAdminInfo(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetAdminResp{
		Id:     info.Admin.Id,
		UUID:   info.Admin.Uuid,
		Name:   info.Admin.Name,
		Email:  info.Admin.Email,
		Avatar: info.Admin.Avatar,
		Groups: lo.Map(info.Admin.Groups, func(item *hr_admin_service.AdminGroup, index int) types.AdminGroup {
			return types.AdminGroup{
				Id:   item.Id,
				Name: item.Name,
				Pivot: types.AdminGroupPivot{
					UserId:  info.Admin.Id,
					GroupId: item.Id,
					Power:   item.Power,
				},
			}
		}),
		DefaultStandardId: info.Admin.DefaultStandardId,
		SmsEnabled:        info.Admin.SmsEnabled,
	}, nil
}
