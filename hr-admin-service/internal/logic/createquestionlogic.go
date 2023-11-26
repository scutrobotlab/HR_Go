package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateQuestionLogic {
	return &CreateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateQuestionLogic) CreateQuestion(in *hr_admin_service.CreateQuestionReq) (*hr_admin_service.CreateQuestionResp, error) {
	q := l.svcCtx.Query.Question
	m := &model.Question{
		Question: in.Question,
		Group_:   in.Group,
	}
	err := q.WithContext(l.ctx).Create(m)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.CreateQuestionResp{
		Question: &hr_admin_service.Question{
			Id:       m.ID,
			Group:    m.Group_,
			Question: m.Question,
		},
	}, nil
}
