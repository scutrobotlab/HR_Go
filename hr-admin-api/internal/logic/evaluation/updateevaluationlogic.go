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

type UpdateEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEvaluationLogic {
	return &UpdateEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEvaluationLogic) UpdateEvaluation(req *types.UpdateEvaluationReq) (resp *types.UpdateEvaluationResp, err error) {
	evaluationResp, err := l.svcCtx.AdminService.UpdateEvaluation(l.ctx, &hr_admin_service.UpdateEvaluationReq{
		EvaluationId: req.Id,
		Name:         req.Name,
		Standard:     logic.StandardsApiToGrpc(req.Standard),
		AdminId:      common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}
	eva := evaluationResp.Evaluation

	return &types.UpdateEvaluationResp{
		EvaluationStandard: types.EvaluationStandard{
			Id:              eva.Id,
			Name:            eva.Name,
			ScoresCount:     eva.ScoresCount,
			LastEditAdminId: eva.UpdatedById,
			LastEditAdmin: types.EvaluationAdmin{
				Id:                eva.UpdatedBy.Id,
				Name:              eva.UpdatedBy.Name,
				DefaultStandardId: eva.UpdatedBy.DefaultStandardId,
			},
			Standard: logic.StandardsGrpcToApi(eva.Standard),
		},
	}, nil
}
