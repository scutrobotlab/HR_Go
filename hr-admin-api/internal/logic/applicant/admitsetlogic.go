package applicant

import (
	"HR_Go/common"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdmitSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdmitSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdmitSetLogic {
	return &AdmitSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdmitSetLogic) AdmitSet(req *types.AdmitSetReq) error {
	_, err := l.svcCtx.AdminService.SetApplicantAdmit(l.ctx, &hr_admin_service.SetApplicantAdmitReq{
		ApplicantId: req.ApplicantId,
		Group:       req.Group,
		Admit:       true,
		AdminId:     common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return err
	}

	return nil
}
