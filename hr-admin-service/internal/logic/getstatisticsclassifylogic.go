package logic

import (
	"HR_Go/common"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"github.com/samber/lo"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatisticsClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStatisticsClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsClassifyLogic {
	return &GetStatisticsClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStatisticsClassifyLogic) GetStatisticsClassify(in *hr_admin_service.GetStatisticsClassifyReq) (*hr_admin_service.GetStatisticsClassifyResp, error) {
	a := l.svcCtx.Query.Applicant
	i := l.svcCtx.Query.Intent
	applicants, err := a.WithContext(l.ctx).Find()
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	classCountMap := make(map[string]int)
	for _, it := range applicants {
		form := make(map[string]string)
		err = json.Unmarshal([]byte(it.Form), &form)
		if err != nil {
			continue
		}
		value := form[in.Key]
		count, _ := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(it.ID), i.Group_.Eq(in.Group)).Count()
		if count > 0 {
			classCountMap[value]++
		}
	}

	return &hr_admin_service.GetStatisticsClassifyResp{
		Class: lo.MapToSlice(classCountMap, func(key string, value int) *hr_admin_service.StatisticsClassify {
			return &hr_admin_service.StatisticsClassify{
				Key:   key,
				Count: int32(value),
			}
		}),
	}, nil
}
