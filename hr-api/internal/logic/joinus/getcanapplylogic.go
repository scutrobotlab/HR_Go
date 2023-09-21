package joinus

import (
	"HR_Go/hr-service/hrservice"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCanApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCanApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCanApplyLogic {
	return &GetCanApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCanApplyLogic) GetCanApply() (resp *types.CanApplyResp, err error) {
	able, err := l.svcCtx.HrService.GetCanJoin(l.ctx, &hrservice.GetCanJoinReq{
		Key: "can-apply",
	})
	if err != nil {
		return nil, err
	}

	return &types.CanApplyResp{
		Status: able.Status,
		Start:  able.Start,
		End:    able.End,
	}, nil
}
