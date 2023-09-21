package form

import (
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	"HR_Go/hr-service/hrservice"
	"HR_Go/util"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFormGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormGroupsLogic {
	return &GetFormGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFormGroupsLogic) GetFormGroups() (resp *types.GetFormGroupsResq, err error) {
	groups, err := l.svcCtx.HrService.GetFormGroups(l.ctx, &hrservice.ApplicantIdReq{
		ApplicantId: 0,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetFormGroupsResq{
		Groups: util.NotNullList(groups.Groups),
	}, nil
}
