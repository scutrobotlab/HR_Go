package remarks

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarksListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRemarksListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarksListLogic {
	return &GetRemarksListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRemarksListLogic) GetRemarksList(req *types.GetRemarksListReq) (resp *types.GetRemarksListResp, err error) {
	remarkList, err := l.svcCtx.AdminService.GetRemarkList(l.ctx, &hr_admin_service.GetRemarkListReq{ApplicantId: req.Id})
	if err != nil {
		if logic.ErrorIsNotFound(err) {
			return &types.GetRemarksListResp{
				Remarks: make([]types.RemarkItem, 0),
			}, nil
		}
		return nil, err
	}

	return &types.GetRemarksListResp{
		Remarks: logic.RemarkGrpcToApi(remarkList.Remarks),
	}, nil
}
