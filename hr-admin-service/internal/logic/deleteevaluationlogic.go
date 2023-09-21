package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEvaluationLogic {
	return &DeleteEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEvaluationLogic) DeleteEvaluation(in *hr_admin_service.DeleteEvaluationReq) (*hr_admin_service.StatusResp, error) {
	es := l.svcCtx.Query.EvaluationStandard
	_, err := es.WithContext(l.ctx).Where(es.ID.Eq(in.EvaluationId)).Delete()
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
