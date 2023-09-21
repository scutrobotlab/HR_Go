package form

import (
	"HR_Go/hr-service/hrservice"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormIntentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFormIntentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormIntentLogic {
	return &GetFormIntentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFormIntentLogic) GetFormIntent() (resp *types.GetFormIntentResq, err error) {
	intentResp, err := l.svcCtx.HrService.GetFormIntent(l.ctx, &hrservice.ApplicantIdReq{
		ApplicantId: 0,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetFormIntentResq{
		Intent: types.IntentConfig{
			Min:      intentResp.Min,
			Max:      intentResp.Max,
			Parallel: intentResp.Parallel,
		},
	}, nil
}
