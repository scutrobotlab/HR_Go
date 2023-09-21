package statistics

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDailyNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDailyNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDailyNewLogic {
	return &GetDailyNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDailyNewLogic) GetDailyNew() (resp *types.DailyNewResp, err error) {
	dailyNew, err := l.svcCtx.AdminService.GetStatisticsDailyNew(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.DailyNewResp{
		Applicants: lo.Map(dailyNew.Applicants, func(item *hr_admin_service.StatisticsDaily, index int) types.DailyNewApplicant {
			return types.DailyNewApplicant{
				Date:   item.Date,
				Counts: item.Count,
			}
		}),
	}, nil
}
