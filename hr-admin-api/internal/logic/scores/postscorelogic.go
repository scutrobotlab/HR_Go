package scores

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostScoreLogic {
	return &PostScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostScoreLogic) PostScore(req *types.PostScoreReq) (resp *types.Score, err error) {
	evaDetails, err := json.Marshal(req.EvaluationDetails)
	if err != nil {
		return nil, err
	}
	scoreResp, err := l.svcCtx.AdminService.PostScore(l.ctx, &hr_admin_service.PostScoreReq{
		ApplicantId:       req.Id,
		Score:             float64(req.Score),
		Group:             req.Group,
		StandardId:        req.StandardId,
		EvaluationDetails: string(evaDetails),
		AdminId:           common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
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
