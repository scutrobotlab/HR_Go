package room

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPushableListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPushableListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPushableListLogic {
	return &GetPushableListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPushableListLogic) GetPushableList() (resp *types.GetPushableListResp, err error) {
	pushableList, err := l.svcCtx.AdminService.GetPushableList(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetPushableListResp{
		Applicants: logic.PushableApplicantGrpcToApi(pushableList.Applicants),
	}, nil
}
