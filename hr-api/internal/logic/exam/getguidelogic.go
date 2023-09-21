package exam

import (
	"HR_Go/hr-service/hrservice"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGuideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGuideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGuideLogic {
	return &GetGuideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGuideLogic) GetGuide(req *types.GetGuideReq) (resp *types.GetGuideResp, err error) {
	guide, err := l.svcCtx.HrService.GetGuide(l.ctx, &hrservice.GetGuideReq{
		Group: req.Group,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetGuideResp{
		Guide: guide.Guide,
	}, nil
}
