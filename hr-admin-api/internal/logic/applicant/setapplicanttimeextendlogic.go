package applicant

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetApplicantTimeExtendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetApplicantTimeExtendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetApplicantTimeExtendLogic {
	return &SetApplicantTimeExtendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetApplicantTimeExtendLogic) SetApplicantTimeExtend(req *types.SetApplicantTimeExtendReq) error {
	bytes, err := json.Marshal(req.Extend)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.AdminService.SetApplicantTimeExtend(l.ctx, &hr_admin_service.SetApplicantTimeExtendReq{
		ApplicantId: req.AppicantId,
		AdminId:     common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
		Extend:      string(bytes),
	})
	if err != nil {
		return err
	}

	return nil
}
