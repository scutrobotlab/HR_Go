package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdminLogic {
	return &CreateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAdminLogic) CreateAdmin(in *hr_admin_service.CreateAdminReq) (*hr_admin_service.CreateAdminResp, error) {
	a := l.svcCtx.Query.Admin

	groups, err := json.Marshal(lo.Map(in.Admin.Groups, func(item *hr_admin_service.AdminGroup, index int) string {
		return item.Name
	}))
	if err != nil {
		return nil, err
	}

	admin := &model.Admin{
		Name:              in.Admin.Name,
		DefaultStandardID: in.Admin.DefaultStandardId,
		Profile:           "{}",
		SmsEnabled:        in.Admin.SmsEnabled,
		RoomID:            0,
		Groups:            string(groups),
	}
	err = a.WithContext(l.ctx).Create(admin)
	if err != nil {
		return nil, err
	}

	password := common.EncryptPassword(l.svcCtx.Common.Encrypt.Salt, admin.ID, in.Password)
	_, err = a.WithContext(l.ctx).Where(a.ID.Eq(admin.ID)).UpdateColumn(a.Password, password)
	if err != nil {
		return nil, err
	}

	in.Admin.Id = admin.ID
	return &hr_admin_service.CreateAdminResp{
		Admin: in.Admin,
	}, nil
}
