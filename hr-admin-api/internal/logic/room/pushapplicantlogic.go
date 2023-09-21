package room

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushApplicantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushApplicantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushApplicantLogic {
	return &PushApplicantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushApplicantLogic) PushApplicant(req *types.PushApplicantReq) error {
	_, err := l.svcCtx.AdminService.PushApplicant(l.ctx, &hr_admin_service.PushApplicantReq{
		ApplicantId:     req.ApplicantId,
		ApplicantTimeId: req.ApplicantTimeId,
		RoomId:          req.RoomId,
		AdminId:         common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return err
	}

	return nil
}
