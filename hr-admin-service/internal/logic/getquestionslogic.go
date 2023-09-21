package logic

import (
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionsLogic {
	return &GetQuestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionsLogic) GetQuestions(in *hr_admin_service.GetQuestionsReq) (*hr_admin_service.GetQuestionsResp, error) {
	q := l.svcCtx.Query.Question
	questions, err := q.WithContext(l.ctx).Where(q.Group_.Eq(in.Group)).Find()
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.GetQuestionsResp{
		Questions: lo.Map(questions, func(item *model.Question, index int) *hr_admin_service.Question {
			return &hr_admin_service.Question{
				Id:       item.ID,
				Group:    item.Group_,
				Question: item.Question,
			}
		}),
	}, nil
}
