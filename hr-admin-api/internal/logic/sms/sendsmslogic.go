package sms

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsReq) error {
	_, err := l.svcCtx.AdminService.SendSms(l.ctx, &hr_admin_service.SendSmsReq{
		AdminId:      common.GetUserInfo(l.ctx).Id,
		Typ:          req.Typ,
		ApplicantIds: req.ApplicanrIds,
	})
	if err != nil {
		return err
	}

	return nil
}
