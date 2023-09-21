package remarks

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRemarksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarksLogic {
	return &DeleteRemarksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRemarksLogic) DeleteRemarks(req *types.DeleteRemarksReq) error {
	_, err := l.svcCtx.AdminService.DeleteRemark(l.ctx, &hr_admin_service.DeleteRemarkReq{
		ApplicantId: req.Id,
		AdminId:     common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return err
	}

	return nil
}
