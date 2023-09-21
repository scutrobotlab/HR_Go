package times

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExportLogic {
	return &GetExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExportLogic) GetExport(req *types.GetExportReq) ([]byte, error) {
	timesResp, err := l.svcCtx.AdminService.ExportTimes(l.ctx, &hr_admin_service.ExportTimesReq{Groups: req.Groups})
	if err != nil {
		return nil, err
	}

	return timesResp.File, nil
}
