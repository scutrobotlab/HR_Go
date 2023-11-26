package logic

import (
	"HR_Go/common"
	"context"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantInfoLogic {
	return &GetApplicantInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantInfoLogic) GetApplicantInfo(in *hr_admin_service.ApplicantIdReq) (*hr_admin_service.GetApplicantInfoResp, error) {
	a := l.svcCtx.Query.Applicant
	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	intents, err := getIntentsByApplicantId(l.ctx, l.svcCtx, in.ApplicantId)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	times, err := getTimesByApplicantId(l.ctx, l.svcCtx, in.ApplicantId)
	if err != nil {
		return nil, err
	}

	admits, err := getAdmitByApplicantId(l.ctx, l.svcCtx, in.ApplicantId)
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.GetApplicantInfoResp{
		Applicant: &hr_admin_service.Applicant{
			Id:        applicant.ID,
			WechatId:  applicant.WechatID,
			Name:      applicant.Name,
			Gender:    common.GenderIntToString(applicant.Gender),
			Phone:     applicant.Phone,
			Avatar:    applicant.Avatar,
			Form:      applicant.Form,
			Intents:   intents,
			Times:     times,
			Admits:    admits,
			CreatedAt: applicant.CreatedAt.Format(time.DateTime),
			UpdatedAt: applicant.UpdatedAt.Format(time.DateTime),
		},
	}, nil
}
