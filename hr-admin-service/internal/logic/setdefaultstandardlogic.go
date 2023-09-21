package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultStandardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetDefaultStandardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultStandardLogic {
	return &SetDefaultStandardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetDefaultStandardLogic) SetDefaultStandard(in *hr_admin_service.SetDefaultStandardReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Admin
	_, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).UpdateColumn(a.DefaultStandardID, in.StandardId)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
