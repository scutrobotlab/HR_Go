package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTimeLogic {
	return &SelectTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectTimeLogic) SelectTime(req *types.SelectTimeReq) (resp *types.SelectTimeResp, err error) {
	time, err := l.svcCtx.HrService.SelectTime(l.ctx, &hr_service.SelectTimeReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		TimeId:      req.TimeId,
		Group:       req.Group,
	})
	if err != nil {
		return nil, err
	}

	return &types.SelectTimeResp{
		Time: types.SelectTime{
			ApplicantId: time.ApplicantId,
			Group:       time.Group,
			TimeId:      time.TimeId,
		},
	}, nil
}
