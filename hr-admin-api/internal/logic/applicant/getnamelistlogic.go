package applicant

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNameListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNameListLogic {
	return &GetNameListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNameListLogic) GetNameList() (resp *types.GetNameListResp, err error) {
	list, err := l.svcCtx.AdminService.GetApplicantNameList(l.ctx, &hr_admin_service.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetNameListResp{
		List: logic.NameListGrpcToApi(list.NameList),
	}, nil
}
