package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreLogic) GetScore(in *hr_admin_service.GetScoreReq) (*hr_admin_service.GetScoreResp, error) {
	s := l.svcCtx.Query.Score
	score, err := s.WithContext(l.ctx).Where(s.ApplicantID.Eq(in.ApplicantId), s.AdminID.Eq(in.AdminId), s.Group_.Eq(in.Group)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	admin, err := getAbstractAdmin(l.ctx, l.svcCtx, score.AdminID)
	if err != nil {
		return nil, err
	}
	return &hr_admin_service.GetScoreResp{
		Score: &hr_admin_service.Score{
			AdminId:          score.AdminID,
			ApplicantId:      score.ApplicantID,
			Group:            score.Group_,
			Score:            score.Score,
			StandardId:       score.StandardID,
			EvaluationDetail: score.EvaluationDetails,
			Admin:            admin,
		},
	}, nil
}
