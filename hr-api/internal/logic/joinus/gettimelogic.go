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

type GetTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeLogic {
	return &GetTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTimeLogic) GetTime(req *types.GetTimeReq) (resp *types.GetTimeResp, err error) {
	time, err := l.svcCtx.HrService.GetTime(l.ctx, &hr_service.GetTimeReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Group:       req.Group,
		Rank:        req.Rank,
	})
	if err != nil {
		return nil, err
	}

	var savedTime types.SavedTime
	if time.SavedTime != nil {
		savedTime = types.SavedTime{
			Id:       time.SavedTime.Id,
			Group:    time.SavedTime.Group,
			Date:     time.SavedTime.Date,
			Time:     time.SavedTime.Time,
			Location: time.SavedTime.Location,
			TotalCnt: time.SavedTime.TotalCnt,
			Grade:    time.SavedTime.Grade,
			Campus:   time.SavedTime.Campus,
			Pivot: types.Pivot{
				ApplicantId: time.SavedTime.ApplicantId,
				TimeId:      time.SavedTime.Id,
			},
			ApplicantId: time.SavedTime.ApplicantId,
		}
	}

	return &types.GetTimeResp{
		Group:     req.Group,
		SavedTime: savedTime,
		Time: util.NotNullList(lo.Map(time.Times, func(item *hr_service.TimeItem, index int) types.ChoosableTime {
			return types.ChoosableTime{
				Id:              item.Id,
				Group:           item.Group,
				Date:            item.Date,
				Time:            item.Time,
				Location:        item.Location,
				TotalCnt:        item.TotalCnt,
				Grade:           item.Grade,
				Campus:          item.Campus,
				ApplicantsCount: item.TotalCnt - item.LeftCnt,
				LeftCnt:         item.LeftCnt,
			}
		})),
	}, nil
}
