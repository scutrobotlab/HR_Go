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

type GetResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResultLogic {
	return &GetResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResultLogic) GetResult() (resp *types.GetResultResp, err error) {
	resultResp, err := l.svcCtx.HrService.GetMyResult(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetResultResp{
		Status: resultResp.Status,
		Times: common.NotNullList(lo.Map(resultResp.Times, func(item *hr_service.TimeItem, index int) types.SavedTime {
			return types.SavedTime{
				Id:       item.Id,
				Group:    item.Group,
				Date:     item.Date,
				Time:     item.Time,
				Rank:     item.Rank,
				Location: item.Location,
				TotalCnt: item.TotalCnt,
				Grade:    item.Grade,
				Campus:   item.Campus,
				Pivot: types.Pivot{
					ApplicantId: item.ApplicantId,
					TimeId:      item.Id,
				},
				ApplicantId: item.ApplicantId,
				MeetingId:   item.MeetingId,
			}
		})),
		Groups: common.NotNullList(lo.Map(resultResp.Groups, func(item *hr_service.AdmitGroup, index int) types.AdmitGroup {
			return types.AdmitGroup{
				Group: item.Group,
				Hint:  item.Hint,
			}
		})),
	}, nil
}
