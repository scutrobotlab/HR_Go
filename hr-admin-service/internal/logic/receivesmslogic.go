package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReceiveSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveSmsLogic {
	return &ReceiveSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReceiveSmsLogic) ReceiveSms(in *hr_admin_service.ReceiveSmsReq) (*hr_admin_service.ReceiveSmsResp, error) {
	rs := l.svcCtx.Query.ReceiveSm

	m := &model.ReceiveSm{
		Phone:   in.Phone,
		Content: in.Content,
	}
	err := rs.WithContext(l.ctx).Create(m)
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.ReceiveSmsResp{
		Code: 0,
	}, nil
}
