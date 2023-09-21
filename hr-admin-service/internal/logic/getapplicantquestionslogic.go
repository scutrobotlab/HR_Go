package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantQuestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantQuestionsLogic {
	return &GetApplicantQuestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantQuestionsLogic) GetApplicantQuestions(in *hr_admin_service.ApplicantIdReq) (*hr_admin_service.GetApplicantQuestionsResp, error) {
	db := l.svcCtx.Db
	var result []*hr_admin_service.QuestionAndAnswer
	err := db.WithContext(l.ctx).Table("applicant_questions as aq").
		Select("q.id, q.group, q.question, aq.answer").
		Joins("left join questions as q on aq.question_id = q.id").
		Where("aq.applicant_id = ?", in.ApplicantId).
		Where("aq.deleted_at is null").
		Where("q.deleted_at is null").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return &hr_admin_service.GetApplicantQuestionsResp{Questions: result}, nil
}
