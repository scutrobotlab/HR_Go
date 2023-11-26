package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantNameListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantNameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantNameListLogic {
	return &GetApplicantNameListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantNameListLogic) GetApplicantNameList(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetApplicantNameListResp, error) {
	a := l.svcCtx.Query.Applicant
	applicants, err := a.WithContext(l.ctx).Find()
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.GetApplicantNameListResp{
		NameList: lo.Map(applicants, func(item *model.Applicant, index int) *hr_admin_service.NameListItem {
			return &hr_admin_service.NameListItem{
				Id:   item.ID,
				Name: item.Name,
			}
		}),
	}, nil
}
