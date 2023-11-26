package logic

import (
	"HR_Go/common"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApplicantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApplicantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApplicantLogic {
	return &DeleteApplicantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteApplicantLogic) DeleteApplicant(in *hr_admin_service.ApplicantIdReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Applicant
	_, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).Delete()
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
