package exam

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExamGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamGroupLogic {
	return &GetExamGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExamGroupLogic) GetExamGroup(req *types.GetExamGroupReq) (resp *types.GetExamGroupResp, err error) {
	questionsResp, err := l.svcCtx.AdminService.GetQuestions(l.ctx, &hr_admin_service.GetQuestionsReq{Group: req.Group})
	if err != nil {
		if logic.ErrorIsNotFound(err) {
			return &types.GetExamGroupResp{
				Questions: make([]types.GetExamGroupQuestion, 0),
			}, nil
		}
		return nil, err
	}

	return &types.GetExamGroupResp{
		Questions: lo.Map(questionsResp.Questions, func(item *hr_admin_service.Question, index int) types.GetExamGroupQuestion {
			return types.GetExamGroupQuestion{
				Id:       item.Id,
				Question: item.Question,
			}
		}),
	}, nil
}
