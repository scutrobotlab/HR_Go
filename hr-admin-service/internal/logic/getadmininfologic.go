package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"
	"encoding/json"
	"github.com/samber/lo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminInfoLogic {
	return &GetAdminInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminInfoLogic) GetAdminInfo(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetAdminInfoResp, error) {
	a := l.svcCtx.Query.Admin
	admin, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	profile := &AdminProfile{}
	err = json.Unmarshal([]byte(admin.Profile), profile)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.GetAdminInfoResp{
		Admin: &hr_admin_service.Admin{
			Id:     admin.ID,
			Name:   admin.Name,
			Uuid:   profile.UUID,
			Email:  profile.Email,
			Avatar: profile.Avatar,
			Groups: lo.Map(profile.Groups, func(item Group, index int) *hr_admin_service.AdminGroup {
				return &hr_admin_service.AdminGroup{
					Id:    item.Id,
					Name:  item.Name,
					Power: int32(item.Power),
				}
			}),
			DefaultStandardId: admin.DefaultStandardID,
			SmsEnabled:        admin.SmsEnabled,
		},
	}, nil
}
