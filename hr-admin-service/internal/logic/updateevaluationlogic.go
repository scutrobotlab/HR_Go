package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEvaluationLogic {
	return &UpdateEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEvaluationLogic) UpdateEvaluation(in *hr_admin_service.UpdateEvaluationReq) (*hr_admin_service.UpdateEvaluationResp, error) {
	es := l.svcCtx.Query.EvaluationStandard
	s := l.svcCtx.Query.Score

	standard, err := json.Marshal(in.Standard)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	m := &model.EvaluationStandard{
		ID:              in.EvaluationId,
		Name:            in.Name,
		Standard:        string(standard),
		LastEditAdminID: in.AdminId,
	}
	err = es.WithContext(l.ctx).Where(es.ID.Eq(in.EvaluationId)).Save(m)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}
	updatedBy, err := getAbstractAdmin(l.ctx, l.svcCtx, m.LastEditAdminID)
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}
	count, _ := s.WithContext(l.ctx).Where(s.StandardID.Eq(in.EvaluationId)).Count()

	return &hr_admin_service.UpdateEvaluationResp{
		Evaluation: &hr_admin_service.Evaluation{
			Id:          m.ID,
			Name:        m.Name,
			ScoresCount: int32(count),
			UpdatedById: m.LastEditAdminID,
			UpdatedBy:   updatedBy,
			Standard:    in.Standard,
		},
	}, nil
}
