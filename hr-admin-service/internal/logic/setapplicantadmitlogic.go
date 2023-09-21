package logic

import (
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetApplicantAdmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetApplicantAdmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetApplicantAdmitLogic {
	return &SetApplicantAdmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetApplicantAdmitLogic) SetApplicantAdmit(in *hr_admin_service.SetApplicantAdmitReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Admit

	if in.Admit {
		err := a.WithContext(l.ctx).Where(a.ApplicantID.Eq(in.ApplicantId), a.Group_.Eq(in.Group)).Save(&model.Admit{
			Group_:      in.Group,
			AdminID:     in.AdminId,
			ApplicantID: in.ApplicantId,
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err := a.WithContext(l.ctx).Where(a.ApplicantID.Eq(in.ApplicantId), a.Group_.Eq(in.Group)).Delete()
		if err != nil {
			return nil, err
		}
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
