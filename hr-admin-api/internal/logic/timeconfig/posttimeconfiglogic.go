package timeconfig

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostTimeConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostTimeConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTimeConfigLogic {
	return &PostTimeConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostTimeConfigLogic) PostTimeConfig(req *types.PostTimeConfigReq) (resp *types.PostTimeConfigResp, err error) {
	timeConfigResp, err := l.svcCtx.AdminService.PostTimeConfig(l.ctx, &hr_admin_service.PostTimeConfigReq{
		TimeConfigs: logic.TimeConfigApiToGrpc(req.TimeConfigs),
	})
	if err != nil {
		return nil, err
	}

	return &types.PostTimeConfigResp{
		TimeConfigs: logic.TimeConfigGrpcToApi(timeConfigResp.TimeConfigs),
	}, nil
}
