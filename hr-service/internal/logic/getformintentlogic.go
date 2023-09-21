package logic

import (
	"context"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormIntentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFormIntentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormIntentLogic {
	return &GetFormIntentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFormIntentLogic) GetFormIntent(in *hr_service.ApplicantIdReq) (*hr_service.GetFormIntentResp, error) {
	intents := l.svcCtx.Common.Intents
	return &hr_service.GetFormIntentResp{
		Min:      intents.Min,
		Max:      intents.Max,
		Parallel: intents.Parallel,
	}, nil
}
