package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarkLogic {
	return &GetRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRemarkLogic) GetRemark(in *hr_admin_service.GetRemarkReq) (*hr_admin_service.GetRemarkResp, error) {
	r := l.svcCtx.Query.Remark
	remark, err := r.WithContext(l.ctx).Where(r.ApplicantID.Eq(in.ApplicantId), r.AdminID.Eq(in.AdminId)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_admin_service.GetRemarkResp{
		ApplicantId: remark.ApplicantID,
		AdminId:     remark.AdminID,
		Remark:      remark.Remark,
	}, nil
}
