package joinus

import (
	"HR_Go/hr-service/hrservice"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCanSelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCanSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCanSelectLogic {
	return &GetCanSelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCanSelectLogic) GetCanSelect() (resp *types.CanSelectResp, err error) {
	able, err := l.svcCtx.HrService.GetCanJoin(l.ctx, &hrservice.GetCanJoinReq{
		Key: "can-select",
	})
	if err != nil {
		return nil, err
	}

	return &types.CanSelectResp{
		Status: able.Status,
		Start:  able.Start,
		End:    able.End,
	}, nil
}
