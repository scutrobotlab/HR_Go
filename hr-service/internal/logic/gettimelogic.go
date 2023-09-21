package logic

import (
	"HR_Go/dal/model"
	"HR_Go/util"
	"context"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeLogic {
	return &GetTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTimeLogic) GetTime(in *hr_service.GetTimeReq) (*hr_service.GetTimeResp, error) {
	at := l.svcCtx.Query.ApplicantTime
	t := l.svcCtx.Query.Time

	var times []*model.Time
	times, err := t.WithContext(l.ctx).Where(t.Group_.Eq(in.Group)).Find()
	if err != nil {
		return nil, util.GrpcErrorInternal(err)
	}
	applicantTime, _ := at.WithContext(l.ctx).Where(at.ApplicantID.Eq(in.ApplicantId), at.Group_.Eq(in.Group)).First()

	var savedTime *hr_service.TimeItem
	var timesDto = lo.Map(times, func(item *model.Time, _ int) *hr_service.TimeItem {
		count, _ := at.WithContext(l.ctx).Where(at.TimeID.Eq(item.ID)).Count()
		oneDto := &hr_service.TimeItem{
			Id:          item.ID,
			Group:       item.Group_,
			Date:        item.Time.Format(time.DateOnly),
			Time:        item.Time.Format(time.TimeOnly),
			Rank:        item.Rank,
			Location:    item.Location,
			TotalCnt:    item.TotalCnt,
			Grade:       item.Grade,
			Campus:      item.Campus,
			ApplicantId: in.ApplicantId,
			LeftCnt:     item.TotalCnt - int32(count),
		}
		if applicantTime != nil && item.ID == applicantTime.TimeID {
			savedTime = oneDto
		}
		return oneDto
	})

	return &hr_service.GetTimeResp{
		Group:     in.Group,
		SavedTime: savedTime,
		Times:     util.NotNullList(timesDto),
	}, nil
}
