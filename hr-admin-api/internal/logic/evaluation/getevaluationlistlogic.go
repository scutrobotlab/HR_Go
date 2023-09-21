package evaluation

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEvaluationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationListLogic {
	return &GetEvaluationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEvaluationListLogic) GetEvaluationList() (resp *types.GetEvaluationListResp, err error) {
	evaluations, err := l.svcCtx.AdminService.GetEvaluations(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetEvaluationListResp{
		List: logic.EvaluationsGrpcToApi(evaluations.Evaluations),
	}, nil
}
