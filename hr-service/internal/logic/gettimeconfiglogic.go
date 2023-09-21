package logic

import (
	"HR_Go/dal/model"
	"HR_Go/util"
	"context"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimeConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTimeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeConfigLogic {
	return &GetTimeConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTimeConfigLogic) GetTimeConfig(in *hr_service.ApplicantIdReq) (*hr_service.GetTimeConfigResp, error) {
	tc := l.svcCtx.Query.TimeConfig

	timeConfigs, err := tc.WithContext(l.ctx).Find()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_service.GetTimeConfigResp{
		TimeConfigs: util.NotNullList(lo.Map(timeConfigs, func(item *model.TimeConfig, _ int) *hr_service.TimeConfig {
			return &hr_service.TimeConfig{
				Key:   item.Key,
				Value: item.Value.Format(time.DateTime),
			}
		})),
	}, nil
}
