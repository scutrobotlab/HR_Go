package applicant

import (
	"HR_Go/hr-admin-api/internal/logic"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplicantInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantInfoLogic {
	return &GetApplicantInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplicantInfoLogic) GetApplicantInfo(req *types.GetInfoReq) (resp *types.GetInfoResp, err error) {
	info, err := l.svcCtx.AdminService.GetApplicantInfo(l.ctx, &hr_admin_service.ApplicantIdReq{ApplicantId: req.Id})
	if err != nil {
		return nil, err
	}

	applicant := info.Applicant
	form := make(map[string]any)
	_ = json.Unmarshal([]byte(applicant.Form), &form)
	return &types.GetInfoResp{
		Applicant: types.Applicant{
			Id:        applicant.Id,
			WechatId:  applicant.WechatId,
			Name:      applicant.Name,
			Gender:    applicant.Gender,
			Phone:     applicant.Phone,
			Avatar:    applicant.Avatar,
			Form:      form,
			Intents:   logic.IntentsGrpcToApi(applicant.Intents),
			Times:     logic.TimesGrpcToApi(applicant.Times),
			Admits:    logic.AdmitsGrpcToApi(applicant.Admits),
			CreatedAt: applicant.CreatedAt,
			UpdatedAt: applicant.UpdatedAt,
		},
	}, nil
}
