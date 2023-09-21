package times

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScheduleLogic {
	return &GetScheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScheduleLogic) GetSchedule(req *types.GetScheduleReq) (resp *types.GetScheduleResp, err error) {
	schedule, err := l.svcCtx.AdminService.GetSchedule(l.ctx, &hr_admin_service.GetScheduleReq{
		Group:           req.Group,
		Date:            req.Date,
		Type:            req.Type,
		ShowNotSelected: req.ShowNotSelected,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetScheduleResp{
		TimeTable:  logic.TimeTableItemGrpcToApi(schedule.TimeTable),
		Categories: schedule.Categories,
		Focus:      schedule.Focus,
	}, nil
}
