package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetApplicantTimeExtendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetApplicantTimeExtendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetApplicantTimeExtendLogic {
	return &SetApplicantTimeExtendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetApplicantTimeExtendLogic) SetApplicantTimeExtend(in *hr_admin_service.SetApplicantTimeExtendReq) (*hr_admin_service.StatusResp, error) {
	at := l.svcCtx.Query.ApplicantTime
	_, err := at.WithContext(l.ctx).Where(at.ApplicantID.Eq(in.ApplicantId), at.Group_.Eq(in.Group)).UpdateColumn(at.Extend, in.Extend)
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
