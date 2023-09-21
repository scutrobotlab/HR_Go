package sms

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveSmsLogic {
	return &ReceiveSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveSmsLogic) ReceiveSms(req *types.ReceiveSmsReq) (resp *types.ReceiveSmsResp, err error) {
	smsResp, err := l.svcCtx.AdminService.ReceiveSms(l.ctx, &hr_admin_service.ReceiveSmsReq{
		Phone:   req.Phone,
		Content: req.Content,
	})
	if err != nil {
		return &types.ReceiveSmsResp{Code: 1}, nil
	}

	return &types.ReceiveSmsResp{
		Code: smsResp.Code,
	}, nil
}
