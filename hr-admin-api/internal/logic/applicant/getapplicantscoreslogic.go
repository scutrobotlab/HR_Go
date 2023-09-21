package applicant

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplicantScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantScoresLogic {
	return &GetApplicantScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplicantScoresLogic) GetApplicantScores(req *types.GetScoresReq) (resp *types.GetScoresResp, err error) {
	scoresResp, err := l.svcCtx.AdminService.GetApplicantScores(l.ctx, &hr_admin_service.ApplicantIdReq{ApplicantId: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.GetScoresResp{
		Scores: logic.ScoresGrpcToApi(scoresResp.Scores),
	}, nil
}
