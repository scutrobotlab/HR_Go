package exam

import (
	"HR_Go/common"
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetExamLogic {
	return &SetExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetExamLogic) SetExam(req *types.SetExamReq) (resp *types.SetExamResp, err error) {
	_, err = l.svcCtx.HrService.PostExam(l.ctx, &hr_service.PostExamReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
		Answers: lo.Map(req.Answers, func(item types.Answer, index int) *hr_service.Answer {
			return &hr_service.Answer{
				QuestionId: item.QuestionId,
				Answer:     strconv.Itoa(int(item.Answer)),
			}
		}),
	})
	if err != nil {
		return nil, err
	}

	examResp, err := l.svcCtx.HrService.GetExam(l.ctx, &hr_service.GetExamReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
	})
	if err != nil {
		return nil, err
	}

	return &types.SetExamResp{
		ApplicantId: 0,
		Questions: common.NotNullList(lo.Map(examResp.Questions, func(item *hr_service.Question, index int) types.Question {
			return types.Question{
				Id:       item.Id,
				Question: item.Question,
				Answer:   item.Answer,
			}
		})),
	}, nil
}
