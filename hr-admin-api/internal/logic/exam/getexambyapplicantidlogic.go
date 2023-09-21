package exam

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamByApplicantIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExamByApplicantIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamByApplicantIdLogic {
	return &GetExamByApplicantIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExamByApplicantIdLogic) GetExamByApplicantId(req *types.GetExamByApplicantIdReq) (resp *types.GetExamByApplicantIdResp, err error) {
	questionsResp, err := l.svcCtx.AdminService.GetApplicantQuestions(l.ctx, &hr_admin_service.ApplicantIdReq{ApplicantId: req.ApplicantId})
	if err != nil {
		return nil, err
	}

	return &types.GetExamByApplicantIdResp{
		ApplicantId: req.ApplicantId,
		Questions:   logic.QuestionAndAnswerGrpcToApi(questionsResp.Questions),
	}, nil
}
