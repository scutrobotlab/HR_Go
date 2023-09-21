package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"HR_Go/util"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIntentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntentsLogic {
	return &GetIntentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIntentsLogic) GetIntents(req *types.GetIntentsReq) (resp *types.GetIntentsResp, err error) {
	intentResp, err := l.svcCtx.HrService.GetIntentList(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetIntentsResp{
		Intents: util.NotNullList(lo.Map(intentResp.Intents, func(item *hr_service.IntentItem, index int) types.Intent {
			return types.Intent{
				Group:       item.Group,
				ApplicantId: item.ApplicantId,
				Finished:    item.Finished,
			}
		})),
	}, nil
}
