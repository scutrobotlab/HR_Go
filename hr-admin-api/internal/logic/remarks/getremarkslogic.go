package remarks

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRemarksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarksLogic {
	return &GetRemarksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRemarksLogic) GetRemarks(req *types.GetRemarksReq) (resp *types.RemarksResp, err error) {
	remarkResp, err := l.svcCtx.AdminService.GetRemark(l.ctx, &hr_admin_service.GetRemarkReq{
		ApplicantId: req.Id,
		AdminId:     common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		if logic.ErrorIsNotFound(err) {
			return &types.RemarksResp{
				AdminId:     common.GetUserInfo(l.ctx).Id,
				ApplicantId: req.Id,
				Remark:      "",
			}, nil
		}
		return nil, err
	}

	return &types.RemarksResp{
		AdminId:     remarkResp.AdminId,
		ApplicantId: remarkResp.ApplicantId,
		Remark:      remarkResp.Remark,
	}, nil
}
