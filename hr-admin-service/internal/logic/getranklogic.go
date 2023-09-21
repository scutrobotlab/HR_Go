package logic

import (
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"
	"errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"sort"
)

type GetRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankLogic {
	return &GetRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRankLogic) GetRank(in *hr_admin_service.GetRankReq) (*hr_admin_service.GetRankResp, error) {
	s := l.svcCtx.Query.Score
	scores, err := s.WithContext(l.ctx).Where(s.Group_.Eq(in.Group)).Find()
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	var scoreMap = make(map[int64]float64)
	var countMap = make(map[int64]int)
	for _, score := range scores {
		scoreMap[score.ApplicantID] += score.Score
		countMap[score.ApplicantID]++
	}

	type AvgScore struct {
		ApplicantID int64
		AvgScore    float64
	}
	var avgScores []AvgScore
	for i, f := range scoreMap {
		avgScores = append(avgScores, AvgScore{
			ApplicantID: i,
			AvgScore:    f / float64(countMap[i]),
		})
	}

	// 从大到小排序
	sort.Slice(avgScores, func(i, j int) bool {
		return avgScores[i].AvgScore > avgScores[j].AvgScore
	})
	myScore, ok := lo.Find(avgScores, func(item AvgScore) bool {
		return item.ApplicantID == in.ApplicantId
	})
	if !ok {
		return nil, util.GrpcErrorNotFound(errors.New("该申请者没有评分"))
	}

	var rank int
	for i, item := range avgScores {
		if item.AvgScore == myScore.AvgScore {
			rank = i + 1
			break
		}
	}

	return &hr_admin_service.GetRankResp{
		Rank:  int32(rank),
		Total: int32(len(avgScores)),
	}, nil
}
