package exam

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExamLogic {
	return &DeleteExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteExamLogic) DeleteExam(req *types.DeleteExamReq) (resp *types.DeleteExamResp, err error) {
	_, err = l.svcCtx.AdminService.DeleteQuestion(l.ctx, &hr_admin_service.DeleteQuestionReq{QuestionId: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DeleteExamResp{
		Id: req.Id,
	}, nil
}
