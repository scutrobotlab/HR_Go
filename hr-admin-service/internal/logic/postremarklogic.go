package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRemarkLogic {
	return &PostRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostRemarkLogic) PostRemark(in *hr_admin_service.PostRemarkReq) (*hr_admin_service.PostRemarkResp, error) {
	r := l.svcCtx.Query.Remark
	err := r.WithContext(l.ctx).Where(r.ApplicantID.Eq(in.ApplicantId), r.AdminID.Eq(in.AdminId)).Save(&model.Remark{
		AdminID:     in.AdminId,
		ApplicantID: in.ApplicantId,
		Remark:      in.Remark,
	})
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.PostRemarkResp{
		ApplicantId: in.ApplicantId,
		AdminId:     in.AdminId,
		Remark:      in.Remark,
	}, nil
}
