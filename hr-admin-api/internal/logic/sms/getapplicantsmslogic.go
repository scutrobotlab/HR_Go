package sms

import (
	"HR_Go/hr-admin-api/internal/logic"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplicantSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantSmsLogic {
	return &GetApplicantSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplicantSmsLogic) GetApplicantSms(req *types.GetApplicantSmsReq) (resp *types.GetApplicantSmsResp, err error) {
	smsResp, err := l.svcCtx.AdminService.GetApplicantSms(l.ctx, &hr_admin_service.GetApplicantSmsReq{
		Typ:      req.Type,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetApplicantSmsResp{
		ApplicantSms: logic.ApplicantSmsGrpcToApi(smsResp.ApplicantSms),
		Total:        smsResp.Total,
	}, nil
}
