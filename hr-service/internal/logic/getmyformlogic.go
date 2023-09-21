package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/util"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyFormLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFormLogic {
	return &GetMyFormLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyFormLogic) GetMyForm(in *hr_service.ApplicantIdReq) (*hr_service.GetMyFormResp, error) {
	a := l.svcCtx.Query.Applicant
	i := l.svcCtx.Query.Intent

	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}
	intents, err := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(applicant.ID)).Find()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_service.GetMyFormResp{
		Id:       applicant.ID,
		WechatId: applicant.WechatID,
		Name:     applicant.Name,
		Gender:   common.GenderIntToString(applicant.Gender),
		Phone:    applicant.Phone,
		Avatar:   applicant.Avatar,
		Form:     applicant.Form,
		Intents: util.NotNullList(lo.Map(intents, func(item *model.Intent, _ int) *hr_service.IntentItem {
			return &hr_service.IntentItem{
				Group:       item.Group_,
				ApplicantId: item.ApplicantID,
			}
		})),
	}, nil
}
