package form

import (
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"HR_Go/util"
	"context"
	"github.com/samber/lo"

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

func (l *GetFormLogic) GetForm() (resp *types.GetFormResq, err error) {
	formFormat, err := l.svcCtx.HrService.GetFormFormat(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: 0,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetFormResq{
		Form: lo.Map(formFormat.Forms, func(item *hr_service.FormItem, index int) types.FormItem {
			required := int32(0)
			if item.Required {
				required = 1
			}
			return types.FormItem{
				Id:        item.Id,
				Name:      item.Name,
				Key:       item.Key,
				Type:      item.Type,
				Required:  required,
				Options:   util.NotNullList(item.Options),
				Regexp:    item.Regexp,
				MaxLength: item.MaxLength,
			}
		}),
	}, nil
}
