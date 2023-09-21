package logic

import (
	"HR_Go/util"
	"context"
	"encoding/json"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEvaluationInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationInfoLogic {
	return &GetEvaluationInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEvaluationInfoLogic) GetEvaluationInfo(in *hr_admin_service.GetEvaluationInfoReq) (*hr_admin_service.GetEvaluationInfoResp, error) {
	s := l.svcCtx.Query.Score
	es := l.svcCtx.Query.EvaluationStandard

	evaluationStandard, err := es.WithContext(l.ctx).Where(es.ID.Eq(in.EvaluationId)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	updatedBy, err := getAbstractAdmin(l.ctx, l.svcCtx, evaluationStandard.LastEditAdminID)
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}
	count, _ := s.WithContext(l.ctx).Where(s.StandardID.Eq(in.EvaluationId)).Count()
	standard := make([]*hr_admin_service.Standard, 0)
	err = json.Unmarshal([]byte(evaluationStandard.Standard), &standard)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.GetEvaluationInfoResp{
		Evaluation: &hr_admin_service.Evaluation{
			Id:          evaluationStandard.ID,
			Name:        evaluationStandard.Name,
			ScoresCount: int32(count),
			UpdatedById: evaluationStandard.LastEditAdminID,
			UpdatedBy:   updatedBy,
			Standard:    standard,
		},
	}, nil
}
