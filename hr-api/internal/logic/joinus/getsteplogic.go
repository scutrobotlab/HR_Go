package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStepLogic {
	return &GetStepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStepLogic) GetStep() (resp *types.GetStepResp, err error) {
	step, err := l.svcCtx.HrService.GetMyStep(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetStepResp{
		Step: step.Step,
	}, nil
}
