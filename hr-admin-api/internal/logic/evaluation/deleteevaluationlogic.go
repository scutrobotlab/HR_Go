package evaluation

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEvaluationLogic {
	return &DeleteEvaluationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEvaluationLogic) DeleteEvaluation(req *types.DeleteEvaluationReq) error {
	_, err := l.svcCtx.AdminService.DeleteEvaluation(l.ctx, &hr_admin_service.DeleteEvaluationReq{EvaluationId: req.Id})
	if err != nil {
		return err
	}

	return nil
}
