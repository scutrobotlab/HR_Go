package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearTimesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTimesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTimesLogic {
	return &ClearTimesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearTimesLogic) ClearTimes(in *hr_admin_service.ClearTimesReq) (*hr_admin_service.StatusResp, error) {
	t := l.svcCtx.Query.Time
	for _, g := range in.Group {
		_, _ = t.WithContext(l.ctx).Where(t.Group_.Eq(g)).Delete()
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
