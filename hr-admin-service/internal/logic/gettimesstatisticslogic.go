package logic

import (
	"HR_Go/util"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimesStatisticsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTimesStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimesStatisticsLogic {
	return &GetTimesStatisticsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTimesStatisticsLogic) GetTimesStatistics(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetTimesStatisticsResp, error) {
	t := l.svcCtx.Query.Time
	times, err := t.WithContext(l.ctx).Find()
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	timeCountMap := make(map[string]int)
	for _, time := range times {
		timeCountMap[time.Group_]++
	}

	return &hr_admin_service.GetTimesStatisticsResp{
		Times: lo.MapToSlice(timeCountMap, func(key string, value int) *hr_admin_service.TimeStatistics {
			return &hr_admin_service.TimeStatistics{
				Group: key,
				Count: int32(value),
			}
		}),
	}, nil
}
