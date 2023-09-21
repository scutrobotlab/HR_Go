package statistics

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassLogic {
	return &GetClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassLogic) GetClass(req *types.ClassReq) (resp *types.ClassResp, err error) {
	classify, err := l.svcCtx.AdminService.GetStatisticsClassify(l.ctx, &hr_admin_service.GetStatisticsClassifyReq{
		Group: req.Group,
		Key:   req.Key,
	})
	if err != nil {
		return nil, err
	}

	return &types.ClassResp{
		Class: logic.ClassGrpcToApi(classify.Class),
	}, nil
}
