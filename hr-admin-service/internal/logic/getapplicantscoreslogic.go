package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantScoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantScoresLogic {
	return &GetApplicantScoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantScoresLogic) GetApplicantScores(in *hr_admin_service.ApplicantIdReq) (*hr_admin_service.GetApplicantScoresResp, error) {
	db := l.svcCtx.Db
	var result []*hr_admin_service.Score
	err := db.WithContext(l.ctx).Table("scores as s").
		Select("s.id, s.admin_id, s.applicant_id, s.group, s.score, s.standard_id, es.standard").
		Joins("left join evaluation_standards as es on s.standard_id = es.id and es.deleted_at is null").
		Where("s.applicant_id = ?", in.ApplicantId).
		Where("s.deleted_at is null").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	for i, score := range result {
		admin, err := getAbstractAdmin(l.ctx, l.svcCtx, score.AdminId)
		if err != nil {
			return nil, err
		}
		result[i].Admin = admin
	}

	return &hr_admin_service.GetApplicantScoresResp{Scores: result}, nil
}
