package evaluation

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEvaluationLogic {
	return &AddEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddEvaluationLogic) AddEvaluation(req *types.AddEvaluationReq) (resp *types.AddEvaluationResp, err error) {
	evaluationResp, err := l.svcCtx.AdminService.AddEvaluation(l.ctx, &hr_admin_service.AddEvaluationReq{
		Name: req.Name,
		Standard: lo.Map(req.Standard, func(item types.Standard, index int) *hr_admin_service.Standard {
			return &hr_admin_service.Standard{
				Name:   item.Name,
				Weight: item.Weight,
			}
		}),
		AdminId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	eva := evaluationResp.Evaluation
	return &types.AddEvaluationResp{
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
