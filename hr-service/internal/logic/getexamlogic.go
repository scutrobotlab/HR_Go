package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamLogic {
	return &GetExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExamLogic) GetExam(in *hr_service.GetExamReq) (*hr_service.GetExamResp, error) {
	q := l.svcCtx.Query.Question
	aq := l.svcCtx.Query.ApplicantQuestion

	applicantQuestions, err := aq.WithContext(l.ctx).Where(aq.ApplicantID.Eq(in.ApplicantId)).Find()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}
	var aqMap = lo.Associate(applicantQuestions, func(item *model.ApplicantQuestion) (int64, *model.ApplicantQuestion) {
		return item.QuestionID, item
	})

	questions, err := q.WithContext(l.ctx).Where(q.Group_.Eq(in.Group)).Find()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	return &hr_service.GetExamResp{
		ApplicantId: in.ApplicantId,
		Questions: common.NotNullList(lo.Map(questions, func(item *model.Question, _ int) *hr_service.Question {
			aqItem := aqMap[item.ID]
			var answer = ""
			if aqItem != nil {
				answer = aqItem.Answer
			}
			return &hr_service.Question{
				Id:       item.ID,
				Question: item.Question,
				Answer:   answer,
			}
		})),
	}, nil
}
