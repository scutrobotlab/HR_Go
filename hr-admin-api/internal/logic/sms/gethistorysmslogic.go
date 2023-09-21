package sms

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistorySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistorySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistorySmsLogic {
	return &GetHistorySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistorySmsLogic) GetHistorySms(req *types.GetHistorySmsReq) (resp *types.GetHistorySmsResp, err error) {
	smsResp, err := l.svcCtx.AdminService.GetHistorySms(l.ctx, &hr_admin_service.ApplicantIdReq{ApplicantId: req.ApplicantId})
	if err != nil {
		return nil, err
	}

	return &types.GetHistorySmsResp{
		HistorySms: logic.HistorySmsGrpcToApi(smsResp.HistorySms),
	}, nil
}
