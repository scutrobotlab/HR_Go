package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEvaluationLogic {
	return &AddEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddEvaluationLogic) AddEvaluation(in *hr_admin_service.AddEvaluationReq) (*hr_admin_service.AddEvaluationResp, error) {
	es := l.svcCtx.Query.EvaluationStandard
	standard, err := json.Marshal(in.Standard)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	m := &model.EvaluationStandard{
		Name:            in.Name,
		Standard:        string(standard),
		LastEditAdminID: in.AdminId,
	}
	err = es.WithContext(l.ctx).Create(m)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}
	updatedBy, err := getAbstractAdmin(l.ctx, l.svcCtx, m.LastEditAdminID)
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_admin_service.AddEvaluationResp{
		Evaluation: &hr_admin_service.Evaluation{
			Id:          m.ID,
			Name:        m.Name,
			ScoresCount: 0,
			UpdatedById: m.LastEditAdminID,
			UpdatedBy:   updatedBy,
			Standard:    in.Standard,
		},
	}, nil
}
