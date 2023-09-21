package logic

import (
	"context"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFormGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormGroupsLogic {
	return &GetFormGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFormGroupsLogic) GetFormGroups(in *hr_service.ApplicantIdReq) (*hr_service.GetFormGroupsResp, error) {
	return &hr_service.GetFormGroupsResp{
		Groups: l.svcCtx.Common.Groups,
	}, nil
}
