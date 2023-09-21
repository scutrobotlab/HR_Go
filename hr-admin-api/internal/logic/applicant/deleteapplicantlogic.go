package applicant

import (
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApplicantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApplicantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApplicantLogic {
	return &DeleteApplicantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApplicantLogic) DeleteApplicant(req *types.DeleteApplicantReq) error {
	_, err := l.svcCtx.AdminService.DeleteApplicant(l.ctx, &hr_admin_service.ApplicantIdReq{ApplicantId: req.Id})
	if err != nil {
		return err
	}

	return nil
}
