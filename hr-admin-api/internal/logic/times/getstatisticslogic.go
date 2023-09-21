package times

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsLogic {
	return &GetStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatisticsLogic) GetStatistics() (resp *types.GetStatisticsResp, err error) {
	statisticsResp, err := l.svcCtx.AdminService.GetTimesStatistics(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetStatisticsResp{
		Times: logic.TimeStatisticsGrpcToApi(statisticsResp.Times),
	}, nil
}
