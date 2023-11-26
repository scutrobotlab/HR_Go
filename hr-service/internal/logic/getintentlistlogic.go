package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntentListLogic {
	return &GetIntentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIntentListLogic) GetIntentList(in *hr_service.ApplicantIdReq) (*hr_service.GetIntentListResp, error) {
	i := l.svcCtx.Query.Intent
	at := l.svcCtx.Query.ApplicantTime

	intents, err := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(in.ApplicantId)).Find()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	return &hr_service.GetIntentListResp{
		Intents: common.NotNullList(lo.Map(intents, func(item *model.Intent, _ int) *hr_service.IntentItem {
			finished := false
			applicantTime, err := at.WithContext(l.ctx).Where(at.ApplicantID.Eq(in.ApplicantId), at.Group_.Eq(item.Group_)).First()
			if err == nil && applicantTime.RoomID != 0 {
				finished = true
			}
			return &hr_service.IntentItem{
				Group:       item.Group_,
				ApplicantId: item.ApplicantID,
				Finished:    finished,
			}
		})),
	}, nil
}
