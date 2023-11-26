package logic

import (
	"HR_Go/common"
	"context"
	"errors"
	"time"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCanJoinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCanJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCanJoinLogic {
	return &GetCanJoinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCanJoinLogic) GetCanJoin(in *hr_service.GetCanJoinReq) (*hr_service.GetCanJoinResp, error) {
	var startKey string
	var endKey string
	switch in.Key {
	case "can-apply":
		startKey = common.ApplyFormStart
		endKey = common.ApplyFormEnd
	case "can-select":
		startKey = common.SelectTimeStart
		endKey = common.SelectTimeEnd
	default:
		return nil, common.GrpcErrorInvalidArgument(errors.New(in.Key))
	}
	tc := l.svcCtx.Query.TimeConfig
	start, err := tc.WithContext(l.ctx).Where(tc.Key.Eq(startKey)).First()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}
	end, err := tc.WithContext(l.ctx).Where(tc.Key.Eq(endKey)).First()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	return &hr_service.GetCanJoinResp{
		Status: time.Now().After(start.Value) && time.Now().Before(end.Value),
		Start:  start.Value.Format(time.DateTime),
		End:    end.Value.Format(time.DateTime),
	}, nil
}
