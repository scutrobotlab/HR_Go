package logic

import (
	"HR_Go/common"
	"HR_Go/common/sms"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *hr_admin_service.SendSmsReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Applicant
	sm := l.svcCtx.Query.ApplicantSm
	adm := l.svcCtx.Query.Admin
	tc := l.svcCtx.Query.TimeConfig

	smsEnabled := GetConfigOrDefault(l.ctx, l.svcCtx, SmsServiceEnabled, "0")
	if smsEnabled != "1" {
		return nil, fmt.Errorf("sms service is not enabled")
	}

	admin, err := adm.WithContext(l.ctx).Where(adm.ID.Eq(in.AdminId)).First()
	if err != nil {
		return nil, err
	}
	if !admin.SmsEnabled {
		return nil, fmt.Errorf("sms not enabled for %s", admin.Name)
	}

	pubResStTime, err := tc.WithContext(l.ctx).Where(tc.Key.Eq(common.PublishResultStart)).First()

	for _, id := range in.ApplicantIds {
		applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(id)).First()
		if err != nil {
			continue
		}
		ctx := context.WithValue(l.ctx, "applicant", applicant)
		ctx = context.WithValue(ctx, "publish_result_start", time.Now().After(pubResStTime.Value))
		generateSms, err := sms.GenerateSms(ctx, l.svcCtx.Query, in.Typ, applicant.ID)
		if err != nil {
			continue
		}
		switch generateSms.Status {
		case sms.DoNotSend, sms.HaveSent:
			// 防止错误发送或重复发送
			continue
		case sms.InSchedule, sms.SendFailed, sms.Expired, sms.Manual, sms.SendRightNow:
			err := sms.SendSms(l.svcCtx.Config.SmsBaoConfig, applicant.Phone, generateSms.Content)
			newStatus := sms.HaveSent
			if err != nil {
				newStatus = sms.SendFailed
			}
			args, _ := json.Marshal(generateSms.Args)
			m := &model.ApplicantSm{
				ApplicantID: applicant.ID,
				Typ:         in.Typ,
				Status:      int32(newStatus),
				Args:        string(args),
				Content:     generateSms.Content,
				CreatedBy:   in.AdminId,
			}
			_, _ = sm.WithContext(l.ctx).Where(sm.ApplicantID.Eq(applicant.ID), sm.Typ.Eq(in.Typ)).Delete()
			_ = sm.WithContext(l.ctx).Create(m)
		default:
			logx.Errorf("unknown status: %d", generateSms.Status)
			continue
		}
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
