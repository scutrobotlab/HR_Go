package logic

import (
	"HR_Go/common"
	"HR_Go/common/sms"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type GetSmsConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsConfigLogic {
	return &GetSmsConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsConfigLogic) GetSmsConfig(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetSmsConfigResp, error) {
	tc := l.svcCtx.Query.TimeConfig
	pubResStTime, _ := tc.WithContext(l.ctx).Where(tc.Key.Eq(common.PublishResultStart)).First()

	return &hr_admin_service.GetSmsConfigResp{
		Types: sms.Types,
		Alerts: lo.Map(sms.Types, func(item string, index int) string {
			switch item {
			case "录取结果":
				if time.Now().Before(pubResStTime.Value) {
					return "公开录取结果的时间没有到。"
				}
			}
			return ""
		}),
	}, nil
}
