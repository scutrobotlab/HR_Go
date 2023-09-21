package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/hrservice"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostScoreLogic {
	return &PostScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostScoreLogic) PostScore(in *hr_admin_service.PostScoreReq) (*hr_admin_service.PostScoreResp, error) {
	s := l.svcCtx.Query.Score
	err := s.WithContext(l.ctx).Where(s.ApplicantID.Eq(in.ApplicantId), s.AdminID.Eq(in.AdminId), s.Group_.Eq(in.Group)).Save(&model.Score{
		AdminID:           in.AdminId,
		ApplicantID:       in.ApplicantId,
		Score:             in.Score,
		Group_:            in.Group,
		StandardID:        in.StandardId,
		EvaluationDetails: in.EvaluationDetails,
	})
	if err != nil {
		return nil, err
	}

	admin, _ := getAbstractAdmin(l.ctx, l.svcCtx, in.AdminId)
	return &hr_admin_service.PostScoreResp{
		Score: &hrservice.Score{
			AdminId:          in.AdminId,
			ApplicantId:      in.ApplicantId,
			Group:            in.Group,
			Score:            in.Score,
			StandardId:       in.StandardId,
			EvaluationDetail: in.EvaluationDetails,
			Admin:            admin,
		},
	}, nil
}
