package logic

import (
	"HR_Go/common/sms"
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"sort"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistorySmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistorySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistorySmsLogic {
	return &GetHistorySmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHistorySmsLogic) GetHistorySms(in *hr_admin_service.ApplicantIdReq) (*hr_admin_service.GetHistorySmsResp, error) {
	as := l.svcCtx.Query.ApplicantSm
	a := l.svcCtx.Query.Applicant
	rs := l.svcCtx.Query.ReceiveSm

	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if err != nil {
		return nil, err
	}
	applicantSms, err := as.WithContext(l.ctx).Where(as.ApplicantID.Eq(in.ApplicantId)).Find()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	receiveSms, err := rs.WithContext(l.ctx).Where(rs.Phone.Eq(applicant.Phone)).Find()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	historySms := make([]*hr_admin_service.HistorySms, 0)

	for _, item := range applicantSms {
		if item.Status != sms.HaveSent {
			continue
		}
		historySms = append(historySms, &hr_admin_service.HistorySms{
			SenderName:    sms.Signature,
			SenderPhone:   "",
			ReceiverName:  applicant.Name,
			ReceiverPhone: applicant.Phone,
			Content:       item.Content,
			Time:          item.CreatedAt.Format("2006-01-02 15:04:05"),
			Right:         true,
		})
	}

	for _, item := range receiveSms {
		historySms = append(historySms, &hr_admin_service.HistorySms{
			SenderName:    applicant.Name,
			SenderPhone:   applicant.Phone,
			ReceiverName:  sms.Signature,
			ReceiverPhone: "",
			Content:       item.Content,
			Time:          item.CreatedAt.Format("2006-01-02 15:04:05"),
			Right:         false,
		})
	}

	sort.Slice(historySms, func(i, j int) bool {
		time1, _ := time.Parse("2006-01-02 15:04:05", historySms[i].Time)
		time2, _ := time.Parse("2006-01-02 15:04:05", historySms[j].Time)
		return time1.Before(time2)
	})

	return &hr_admin_service.GetHistorySmsResp{
		HistorySms: historySms,
	}, nil
}
