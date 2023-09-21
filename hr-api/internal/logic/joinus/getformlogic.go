package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormLogic {
	return &GetFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFormLogic) GetForm() (resp *types.GetFormResp, err error) {
	form, err := l.svcCtx.HrService.GetMyForm(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	var intents = lo.Map(form.Intents, func(item *hr_service.IntentItem, index int) types.Intent {
		return types.Intent{
			Group:       item.Group,
			ApplicantId: item.ApplicantId,
		}
	})

	return &types.GetFormResp{
		Applicant: types.Applicant{
			Id:       form.Id,
			WechatId: form.WechatId,
			Name:     form.Name,
			Gender:   form.Gender,
			Phone:    form.Phone,
			Avatar:   form.Avatar,
			Form:     form.Form,
			Intents:  intents,
		},
	}, nil
}
