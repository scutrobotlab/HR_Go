package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostExamLogic {
	return &PostExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostExamLogic) PostExam(in *hr_service.PostExamReq) (*hr_service.PostExamResp, error) {
	aq := l.svcCtx.Query.ApplicantQuestion

	lo.Map(in.Answers, func(item *hr_service.Answer, _ int) *hr_service.Question {
		m := &model.ApplicantQuestion{
			ApplicantID: in.ApplicantId,
			QuestionID:  item.QuestionId,
			Answer:      item.Answer,
		}

		err := aq.WithContext(l.ctx).Where(aq.ID.Eq(item.QuestionId), aq.ApplicantID.Eq(in.ApplicantId)).Save(m)
		if err != nil {
			return nil
		}

		return nil
	})

	return &hr_service.PostExamResp{
		ApplicantId: in.ApplicantId,
		Questions:   nil,
	}, nil
}
