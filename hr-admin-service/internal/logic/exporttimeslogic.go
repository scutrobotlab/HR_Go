package logic

import (
	"HR_Go/util"
	"context"
	"fmt"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTimesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExportTimesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTimesLogic {
	return &ExportTimesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExportTimesLogic) ExportTimes(in *hr_admin_service.ExportTimesReq) (*hr_admin_service.ExportTimesResp, error) {
	t := l.svcCtx.Query.Time
	times, err := t.WithContext(l.ctx).Where(t.Group_.In(in.Groups...)).Find()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}
	file := "组别,日期,时间,地点,第几志愿,人数,年级,校区,会议号\n"
	for _, it := range times {
		date := it.Time.Format("2006-01-02")
		time := it.Time.Format("15:04:05")
		file += fmt.Sprintf("%s,%s,%s,%s,%d,%d,%s,%s,%s\n", it.Group_, date, time, it.Location, it.Rank, it.TotalCnt, it.Grade, it.Campus, it.MeetingID)
	}

	return &hr_admin_service.ExportTimesResp{
		File: []byte(file),
	}, nil
}
