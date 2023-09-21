package logic

import (
	"HR_Go/common"
	"HR_Go/common/sms"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"time"
)

type GetApplicantSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantSmsLogic {
	return &GetApplicantSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantSmsLogic) GetApplicantSms(in *hr_admin_service.GetApplicantSmsReq) (*hr_admin_service.GetApplicantSmsResp, error) {
	type Result struct {
		ID       int64
		Name     string
		Phone    string
		ReplyCnt int32
		Form     string
	}

	limit := int(in.PageSize)
	offset := int(in.PageSize * (in.Page - 1))
	if in.PageSize == -1 {
		limit = math.MaxInt
		offset = 0
	}

	db := l.svcCtx.Db
	var result []*Result
	var total int64
	query := db.WithContext(l.ctx).Table("applicants as a").
		Select("a.id as id, a.name, a.phone, count(rs.id) as reply_cnt, a.form").
		Joins("left join receive_sms as rs on a.phone = rs.phone and rs.deleted_at is null").
		Where("a.name != ?", "").
		Where("a.deleted_at is null").
		Group("a.id").
		Order("a.id")
	query.Count(&total)
	err := query.Limit(limit).Offset(offset).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	tc := l.svcCtx.Query.TimeConfig
	pubResStTime, err := tc.WithContext(l.ctx).Where(tc.Key.Eq(common.PublishResultStart)).First()

	return &hr_admin_service.GetApplicantSmsResp{
		ApplicantSms: lo.Map(result, func(item *Result, index int) *hr_admin_service.ApplicantSms {
			applicant := &model.Applicant{
				ID:    item.ID,
				Name:  item.Name,
				Phone: item.Phone,
				Form:  item.Form,
			}
			ctx := context.WithValue(l.ctx, "applicant", applicant)
			ctx = context.WithValue(ctx, "publish_result_start", time.Now().After(pubResStTime.Value))
			applicantSms := &hr_admin_service.ApplicantSms{
				ApplicantId: item.ID,
				Name:        item.Name,
				Phone:       item.Phone,
				Status:      sms.DoNotSend,
				Time:        "",
				Args:        []string{},
				Content:     "",
				ReplyCnt:    item.ReplyCnt,
			}
			getSms, err := sms.GenerateSms(ctx, l.svcCtx.Query, in.Typ, item.ID)
			if err != nil {
				return applicantSms
			}
			if getSms.Time != nil {
				applicantSms.Time = getSms.Time.Format("2006-01-02 15:04:05")
			}
			applicantSms.Status = int32(getSms.Status)
			applicantSms.Args = getSms.Args
			applicantSms.Content = getSms.Content
			return applicantSms
		}),
		Total: int32(total),
	}, nil
}
