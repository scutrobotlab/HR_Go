package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarkLogic {
	return &DeleteRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRemarkLogic) DeleteRemark(in *hr_admin_service.DeleteRemarkReq) (*hr_admin_service.StatusResp, error) {
	r := l.svcCtx.Query.Remark
	_, err := r.WithContext(l.ctx).Where(r.ApplicantID.Eq(in.ApplicantId), r.AdminID.Eq(in.AdminId)).Delete()
	if err != nil {
		return &hr_admin_service.StatusResp{Ok: false}, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.StatusResp{Ok: true}, nil
}
