package logic

import (
	"HR_Go/util"
	"context"
	"github.com/samber/lo"
	"sort"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatisticsDailyNewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStatisticsDailyNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsDailyNewLogic {
	return &GetStatisticsDailyNewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStatisticsDailyNewLogic) GetStatisticsDailyNew(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetStatisticsDailyNewResp, error) {
	a := l.svcCtx.Query.Applicant
	applicants, err := a.WithContext(l.ctx).Find()
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	timeCountMap := make(map[string]int)
	for _, it := range applicants {
		if it.Name == "" || it.Form == "{}" {
			continue
		}
		date := it.CreatedAt.Format(time.DateOnly)
		timeCountMap[date]++
	}
	aSlice := lo.MapToSlice(timeCountMap, func(key string, value int) *hr_admin_service.StatisticsDaily {
		return &hr_admin_service.StatisticsDaily{
			Date:  key,
			Count: int32(value),
		}
	})
	sort.Slice(aSlice, func(i, j int) bool {
		date1, _ := time.Parse(time.DateOnly, aSlice[i].Date)
		date2, _ := time.Parse(time.DateOnly, aSlice[j].Date)
		return date1.Before(date2)
	})

	return &hr_admin_service.GetStatisticsDailyNewResp{
		Applicants: aSlice,
	}, nil
}
