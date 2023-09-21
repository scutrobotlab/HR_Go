package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostTimeConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostTimeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTimeConfigLogic {
	return &PostTimeConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostTimeConfigLogic) PostTimeConfig(in *hr_admin_service.PostTimeConfigReq) (*hr_admin_service.PostTimeConfigResp, error) {
	tc := l.svcCtx.Query.TimeConfig

	for _, it := range in.TimeConfigs {
		timeValue, err := time.ParseInLocation(time.DateTime, it.Value, time.Local)
		if err != nil {
			return nil, util.GrpcErrorInvalidArgument(err)
		}
		count, _ := tc.WithContext(l.ctx).Where(tc.Key.Eq(it.Key)).Count()
		if count == 0 {
			err = tc.WithContext(l.ctx).Create(&model.TimeConfig{
				Key:   it.Key,
				Value: timeValue,
			})
		} else {
			_, err = tc.WithContext(l.ctx).Where(tc.Key.Eq(it.Key)).Update(tc.Value, timeValue)
		}
		if err != nil {
			return nil, util.GrpcErrorInternal(err)
		}
	}

	return &hr_admin_service.PostTimeConfigResp{
		TimeConfigs: in.TimeConfigs,
	}, nil
}
