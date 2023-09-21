package remarks

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostRemarksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostRemarksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRemarksLogic {
	return &PostRemarksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostRemarksLogic) PostRemarks(req *types.PostRemarksReq) (resp *types.RemarksResp, err error) {
	remarkResp, err := l.svcCtx.AdminService.PostRemark(l.ctx, &hr_admin_service.PostRemarkReq{
		ApplicantId: req.Id,
		AdminId:     common.GetUserInfo(l.ctx).Id,
		Remark:      req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.RemarksResp{
		AdminId:     remarkResp.AdminId,
		ApplicantId: remarkResp.ApplicantId,
		Remark:      remarkResp.Remark,
	}, nil
}
