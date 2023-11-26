package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"
	"strconv"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadTimesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadTimesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadTimesLogic {
	return &UploadTimesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadTimesLogic) UploadTimes(in *hr_admin_service.UploadTimesReq) (*hr_admin_service.StatusResp, error) {
	t := l.svcCtx.Query.Time

	for _, item := range in.Data {
		it := item.Data
		group := it["group"]
		if !lo.Contains(in.Groups, group) {
			// 不写入没有选中的组别
			continue
		}

		time_, err := time.ParseInLocation(time.DateTime, it["date"]+" "+it["time"], time.Local)
		if err != nil {
			continue
		}
		time_ = time_.AddDate(1, 0, 0)
		rank, err := strconv.Atoi(it["rank"])
		if err != nil {
			continue
		}
		totalCnt, err := strconv.Atoi(it["total_cnt"])
		if err != nil {
			continue
		}

		m := &model.Time{
			Group_:   group,
			Time:     time_,
			Rank:     int32(rank),
			Location: it["location"],
			Campus:   it["campus"],
			TotalCnt: int32(totalCnt),
			Grade:    it["grade"],
		}
		err = t.WithContext(l.ctx).Create(m)
		if err != nil {
			return nil, common.GrpcErrorInternal(err)
		}
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
