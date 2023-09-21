package evaluation

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEvaluationInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationInfoLogic {
	return &GetEvaluationInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEvaluationInfoLogic) GetEvaluationInfo(req *types.GetEvaluationInfoReq) (resp *types.GetEvaluationInfoResp, err error) {
	infoResp, err := l.svcCtx.AdminService.GetEvaluationInfo(l.ctx, &hr_admin_service.GetEvaluationInfoReq{EvaluationId: req.Id})
	if err != nil {
		if logic.ErrorIsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	eva := infoResp.Evaluation

	return &types.GetEvaluationInfoResp{
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
