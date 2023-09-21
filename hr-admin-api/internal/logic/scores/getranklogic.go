package scores

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankLogic {
	return &GetRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankLogic) GetRank(req *types.GetRankReq) (resp *types.GetRankResp, err error) {
	rankResp, err := l.svcCtx.AdminService.GetRank(l.ctx, &hr_admin_service.GetRankReq{
		ApplicantId: req.Id,
		Group:       req.Group,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetRankResp{
		Rank:  rankResp.Rank,
		Total: rankResp.Total,
	}, nil
}
