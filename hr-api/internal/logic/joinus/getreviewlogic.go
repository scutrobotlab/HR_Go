package joinus

import (
	"HR_Go/common"
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReviewLogic {
	return &GetReviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReviewLogic) GetReview() (resp *types.GetReviewResp, err error) {
	id := common.GetUserInfo(l.ctx).Id
	reviewResp, err := l.svcCtx.HrService.GetReview(l.ctx, &hr_service.ApplicantIdReq{ApplicantId: id})
	if err != nil {
		logx.Errorf("获取面试回顾失败: %s", err)
		return nil, nil
	}

	return &types.GetReviewResp{
		Text: reviewResp.Text,
	}, nil
}
