package logic

import (
	"context"
	"fmt"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *hr_admin_service.PingRequest) (*hr_admin_service.PingResponse, error) {
	fmt.Printf("Ping = %s\n", in.Ping)
	return &hr_admin_service.PingResponse{Pong: "pong"}, nil
}
