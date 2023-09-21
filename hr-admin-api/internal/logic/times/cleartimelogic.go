package times

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTimeLogic {
	return &ClearTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearTimeLogic) ClearTime(req *types.ClearTimeReq) error {
	_, err := l.svcCtx.AdminService.ClearTimes(l.ctx, &hr_admin_service.ClearTimesReq{Group: req.Groups})
	if err != nil {
		return err
	}

	return nil
}
