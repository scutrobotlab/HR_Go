package exam

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExamLogic {
	return &CreateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExamLogic) CreateExam(req *types.CreateExamReq) (resp *types.CreateExamResp, err error) {
	questionResp, err := l.svcCtx.AdminService.CreateQuestion(l.ctx, &hr_admin_service.CreateQuestionReq{
		Group:    req.Group,
		Question: req.Question,
	})
	if err != nil {
		return nil, err
	}
	question := questionResp.Question

	return &types.CreateExamResp{
		Question: types.Question{
			Id:       question.Id,
			Group:    question.Group,
			Question: question.Question,
		},
	}, nil
}
