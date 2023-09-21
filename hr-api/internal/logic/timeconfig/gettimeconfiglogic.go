package timeconfig

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimeConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTimeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeConfigLogic {
	return &GetTimeConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTimeConfigLogic) GetTimeConfig() (resp *types.GetTimeConfigResp, err error) {
	timeConfig, err := l.svcCtx.HrService.GetTimeConfig(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetTimeConfigResp{
		TimeConfigs: lo.Map(timeConfig.TimeConfigs, func(item *hr_service.TimeConfig, index int) types.TimeConfigItem {
			return types.TimeConfigItem{
				Key:   item.Key,
				Value: item.Value,
			}
		}),
	}, nil
}
