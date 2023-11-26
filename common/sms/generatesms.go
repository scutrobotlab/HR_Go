package sms

import (
	"HR_Go/common"
	"HR_Go/dal/query"
	"context"
	"encoding/json"
	"errors"
	"time"
)

const (
	DoNotSend    = iota // 不发送
	InSchedule          // 定时发送
	HaveSent            // 已发送
	SendFailed          // 发送失败
	Expired             // 已过期
	Manual              // 手动发送
	SendRightNow        // 立即发送
)

type Sms struct {
	Status  int
	Time    *time.Time
	Content string
	Args    []string
}

func GenerateSms(ctx context.Context, query *query.Query, typ string, applicantId int64) (*Sms, error) {
	as := query.ApplicantSm

	appSms, err := as.WithContext(ctx).Where(as.ApplicantID.Eq(applicantId), as.Typ.Eq(typ)).First()
	if err == nil {
		args := make([]string, 0)
		err := json.Unmarshal([]byte(appSms.Args), &args)
		if err != nil {
			return nil, err
		}
		if appSms.Status == DoNotSend || appSms.Status == InSchedule || appSms.Status == Expired {
			return nil, common.GrpcErrorUnknown(errors.New("没有发送短信但存在记录"))
		}
		// Status == HaveSent or Status == SendFailed or Status == Manual or Status == SendRightNow
		return &Sms{
			Status:  int(appSms.Status),
			Time:    &appSms.UpdatedAt,
			Content: appSms.Content,
			Args:    args,
		}, nil
	}

	var handler Handler
	switch typ {
	case "选择时间":
		handler = SelectTimeSmsHandler{}
	case "面试邀请1":
		handler = InterviewSmsHandler{Rank: 0, ExpireTime: 1 * time.Hour}
	case "面试邀请2":
		handler = InterviewSmsHandler{Rank: 1, ExpireTime: 1 * time.Hour}
	case "录取结果":
		handler = ResultSmsHandler{}
	default:
		return nil, common.GrpcErrorUnknown(errors.New("短信类型错误"))
	}
	getSms, err := handler.GetSms(ctx, query, applicantId)
	if err != nil {
		return nil, err
	}

	return &Sms{
		Status:  getSms.Status,
		Time:    getSms.Time,
		Content: getSms.Content,
		Args:    getSms.Args,
	}, nil
}
