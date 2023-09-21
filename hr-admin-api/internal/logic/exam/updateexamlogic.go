package exam

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExamLogic {
	return &UpdateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExamLogic) UpdateExam(req *types.UpdateExamReq) (resp *types.UpdateExamResp, err error) {
	_, err = l.svcCtx.AdminService.UpdateQuestion(l.ctx, &hr_admin_service.UpdateQuestionReq{
		Question: &hr_admin_service.Question{
			Id:       req.Id,
			Group:    req.Group,
			Question: req.Question,
		}})
	if err != nil {
		return nil, err
	}

	return &types.UpdateExamResp{
		Id: req.Id,
	}, nil
}
