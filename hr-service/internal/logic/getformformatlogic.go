package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"encoding/json"
	"github.com/samber/lo"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormFormatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFormFormatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormFormatLogic {
	return &GetFormFormatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFormFormatLogic) GetFormFormat(in *hr_service.ApplicantIdReq) (*hr_service.GetFormFormatResp, error) {
	f := l.svcCtx.Query.Form
	forms, err := f.WithContext(l.ctx).Find()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	return &hr_service.GetFormFormatResp{
		Forms: common.NotNullList(lo.Map(forms, func(item *model.Form, _ int) *hr_service.FormItem {
			var options []string
			_ = json.Unmarshal([]byte(item.Options), &options)
			return &hr_service.FormItem{
				Id:        item.ID,
				Name:      item.Name,
				Key:       item.Key,
				Type:      item.Type,
				Required:  item.Required,
				Options:   options,
				Regexp:    item.Regexp,
				MaxLength: item.MaxLength,
			}
		})),
	}, nil
}
