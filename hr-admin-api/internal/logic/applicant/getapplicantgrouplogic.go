package applicant

import (
	"HR_Go/hr-admin-api/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"github.com/samber/lo"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplicantGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantGroupLogic {
	return &GetApplicantGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplicantGroupLogic) GetApplicantGroup(req *types.GetGroupReq) (resp *types.GetGroupResp, err error) {
	getApplicantsResp, err := l.svcCtx.AdminService.GetApplicants(l.ctx, &hr_admin_service.GetApplicantsReq{
		Group: req.Group,
		State: req.State,
		Page:  req.Page,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetGroupResp{
		CurrentPage: getApplicantsResp.CurrentPage,
		Data: lo.Map(getApplicantsResp.Data, func(item *hr_admin_service.ApplicantListItem, index int) types.GetGroupData {
			var extend map[string]interface{}
			_ = json.Unmarshal([]byte(item.Extend), &extend)
			return types.GetGroupData{
				Id:          item.Id,
				Name:        item.Name,
				Gender:      item.Gender,
				Avatar:      item.Avatar,
				ScoresCount: item.ScoresCount,
				AvgScore:    item.AvgScore,
				Date:        item.Date,
				Time:        item.Time,
				Location:    item.Location,
				Extend:      extend,
				Admits:      logic.AdmitsGrpcToApi(item.Admits),
			}
		}),
		FirstPageUrl: getApplicantsResp.FirstPageUrl,
		LastPage:     getApplicantsResp.LastPage,
		LastPageUrl:  getApplicantsResp.LastPageUrl,
		Links: lo.Map(getApplicantsResp.Links, func(item *hr_admin_service.Link, index int) types.Link {
			return types.Link{
				Url:    item.Url,
				Label:  item.Label,
				Active: item.Active,
			}
		}),
		NextPageUrl: getApplicantsResp.NextPageUrl,
		Path:        getApplicantsResp.Path,
		PerPage:     getApplicantsResp.PerPage,
		PrevPageUrl: getApplicantsResp.PrevPageUrl,
		From:        getApplicantsResp.From,
		To:          getApplicantsResp.To,
		Total:       getApplicantsResp.Total,
	}, nil
}
