package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(in *hr_admin_service.DeleteQuestionReq) (*hr_admin_service.StatusResp, error) {
	q := l.svcCtx.Query.Question
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.QuestionId)).Delete()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
