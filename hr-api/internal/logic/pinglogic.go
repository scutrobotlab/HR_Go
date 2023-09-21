package logic

import (
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingResp, err error) {
	response, err := l.svcCtx.HrService.Ping(l.ctx, &hr_service.PingRequest{Ping: "ping"})
	if err != nil {
		return nil, err
	}
	return &types.PingResp{Ping: response.Pong}, nil
}
