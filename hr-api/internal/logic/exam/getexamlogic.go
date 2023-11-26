package exam

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamLogic {
	return &GetExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExamLogic) GetExam(req *types.GetExamReq) (resp *types.GetExamResp, err error) {
	exam, err := l.svcCtx.HrService.GetExam(l.ctx, &hr_service.GetExamReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetExamResp{
		ApplicantId: exam.ApplicantId,
		Questions: common.NotNullList(lo.Map(exam.Questions, func(item *hr_service.Question, _ int) types.Question {
			return types.Question{
				Id:       item.Id,
				Question: item.Question,
				Answer:   item.Answer,
			}
		})),
	}, nil
}
