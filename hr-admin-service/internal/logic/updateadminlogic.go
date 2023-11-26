package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"github.com/samber/lo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(in *hr_admin_service.UpdateAdminReq) (*hr_admin_service.UpdateAdminResp, error) {
	a := l.svcCtx.Query.Admin

	groups, err := json.Marshal(lo.Map(in.Admin.Groups, func(item *hr_admin_service.AdminGroup, index int) string {
		return item.Name
	}))
	if err != nil {
		return nil, err
	}

	_, err = a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).Updates(model.Admin{
		Name:              in.Admin.Name,
		DefaultStandardID: in.Admin.DefaultStandardId,
		Profile:           "{}",
		SmsEnabled:        in.Admin.SmsEnabled,
		RoomID:            0,
		Groups:            string(groups),
	})
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.UpdateAdminResp{
		Admin: in.Admin,
	}, nil
}
