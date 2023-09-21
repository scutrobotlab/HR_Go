package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(in *hr_admin_service.UpdateQuestionReq) (*hr_admin_service.UpdateQuestionResp, error) {
	q := l.svcCtx.Query.Question
	err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Question.Id)).Save(&model.Question{
		ID:       in.Question.Id,
		Question: in.Question.Question,
		Group_:   in.Question.Group,
	})
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.UpdateQuestionResp{
		Question: in.Question,
	}, nil
}
