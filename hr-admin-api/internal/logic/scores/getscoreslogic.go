package scores

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoresLogic {
	return &GetScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoresLogic) GetScores(req *types.GetScoreReq) (resp *types.Score, err error) {
	scoreResp, err := l.svcCtx.AdminService.GetScore(l.ctx, &hr_admin_service.GetScoreReq{
		ApplicantId: req.Id,
		AdminId:     common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
	})
	if err != nil {
		if logic.ErrorIsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	EvaluationDetails := make([]types.EvaluationDetail, 0)
	_ = json.Unmarshal([]byte(scoreResp.Score.EvaluationDetail), &EvaluationDetails)
	return &types.Score{
		AdminId:           scoreResp.Score.AdminId,
		ApplicantId:       scoreResp.Score.ApplicantId,
		Score:             float32(scoreResp.Score.Score),
		Group:             scoreResp.Score.Group,
		StandardId:        scoreResp.Score.StandardId,
		EvaluationDetails: EvaluationDetails,
		Admin: types.ScoresAdmin{
			Id:                scoreResp.Score.Admin.Id,
			Name:              scoreResp.Score.Admin.Name,
			DefaultStandardId: scoreResp.Score.Admin.DefaultStandardId,
		},
	}, nil
}
