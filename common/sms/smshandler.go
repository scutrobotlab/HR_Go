package sms

import (
	"HR_Go/dal/model"
	"HR_Go/dal/query"
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"strings"
	"time"
)

type Handler interface {
	GetSms(ctx context.Context, query *query.Query, applicantId int64) (*Sms, error)
}

type (
	// SelectTimeSmsHandler 选择时间短信
	SelectTimeSmsHandler struct {
		Handler
	}

	// InterviewSmsHandler 面试邀请短信
	InterviewSmsHandler struct {
		Handler

		Rank       int
		ExpireTime time.Duration
	}

	// ResultSmsHandler 面试结果短信
	ResultSmsHandler struct {
		Handler
	}
)

func (h SelectTimeSmsHandler) GetSms(ctx context.Context, query *query.Query, applicantId int64) (*Sms, error) {
	i := query.Intent
	at := query.ApplicantTime

	intents, err := i.WithContext(ctx).Where(i.ApplicantID.Eq(applicantId)).Find()
	if err != nil {
		return nil, err
	}
	applicantTimes, err := at.WithContext(ctx).Where(at.ApplicantID.Eq(applicantId)).Find()
	if err != nil {
		return nil, err
	}

	if len(intents) > len(applicantTimes) {
		// 有志愿但没有选择时间，需要发送短信
		iGroups := lo.Map(intents, func(item *model.Intent, index int) string {
			return item.Group_ + "组"
		})
		atGroups := lo.Map(applicantTimes, func(item *model.ApplicantTime, index int) string {
			return item.Group_ + "组"
		})
		difference, _ := lo.Difference(iGroups, atGroups)
		groupStr := strings.Join(difference, "、")

		applicant := ctx.Value("applicant").(*model.Applicant)
		config := Configs["选择时间"][""]
		args := []string{applicant.Name, groupStr}
		content := config.GetContent(args...)
		return &Sms{
			Status:  Manual,
			Time:    nil,
			Content: content,
			Args:    args,
		}, nil
	} else {
		// 所有志愿都选择了时间，不需要发送短信
		return &Sms{
			Status:  DoNotSend,
			Time:    nil,
			Content: "",
			Args:    []string{},
		}, nil
	}
}

func (h InterviewSmsHandler) GetSms(ctx context.Context, query *query.Query, applicantId int64) (*Sms, error) {
	i := query.Intent
	at := query.ApplicantTime
	t := query.Time

	intent, err := i.WithContext(ctx).Where(i.ApplicantID.Eq(applicantId), i.Rank.Eq(int32(h.Rank))).First()
	if err != nil {
		// 没有志愿，不需要发送短信
		return &Sms{
			Status:  DoNotSend,
			Time:    nil,
			Content: "",
			Args:    []string{},
		}, nil
	}

	applicantTime, err := at.WithContext(ctx).Where(at.ApplicantID.Eq(applicantId), at.Group_.Eq(intent.Group_)).First()
	if err != nil {
		return nil, err
	}
	oneTime, err := t.WithContext(ctx).Where(t.ID.Eq(applicantTime.TimeID)).First()
	if err != nil {
		return nil, err
	}
	year, month, day := oneTime.Time.Date()
	hour, min, sec := oneTime.Time.Clock()
	interviewTime := time.Date(year, month, day, hour, min, sec, 0, time.Local)
	sendTime := interviewTime.Add(-h.ExpireTime)

	applicant := ctx.Value("applicant").(*model.Applicant)
	condition := "offline"
	if oneTime.Location == "飞书会议" || oneTime.Location == "线上会议" {
		condition = "online"
	}
	config := Configs["面试邀请"][condition]
	var args []string
	if condition == "online" {
		meetingId := strings.Replace(oneTime.MeetingID, "-", "", -1)
		args = []string{applicant.Name, intent.Group_, interviewTime.Format("1月2日15:04"), meetingId}
	} else {
		location := oneTime.Location
		if location == "五山33号楼" {
			location = "五山33号楼1层机器人实验室"
		} else if location == "大学城木棉" {
			location = "大学城木棉空间"
		}
		args = []string{applicant.Name, intent.Group_, interviewTime.Format("1月2日15:04"), location}
	}
	content := config.GetContent(args...)

	status := SendRightNow
	if time.Now().After(interviewTime) {
		// 当前时间在面试时间之后，不需要发送短信
		status = Expired
	} else if time.Now().Before(sendTime) {
		// 当前时间在发送时间之前，还未到发送时间
		status = InSchedule
	}

	return &Sms{
		Status:  status,
		Time:    &sendTime,
		Content: content,
		Args:    args,
	}, nil
}

func (h ResultSmsHandler) GetSms(ctx context.Context, query *query.Query, applicantId int64) (*Sms, error) {
	at := query.Admit
	status := Manual

	if !ctx.Value("publish_result_start").(bool) {
		status = DoNotSend
	}

	applicant := ctx.Value("applicant").(*model.Applicant)
	admits, err := at.WithContext(ctx).Where(at.ApplicantID.Eq(applicantId)).Find()
	if err != nil || len(admits) == 0 {
		var form map[string]interface{}
		_ = json.Unmarshal([]byte(applicant.Form), &form)
		if grade, ok := form["grade"].(string); ok {
			condition := "fail.not_freshman"
			if grade == "大一" {
				condition = "fail.freshman"
			}
			config := Configs["录取结果"][condition]
			args := []string{"不录取" + grade, applicant.Name}
			return &Sms{
				Status:  status,
				Time:    nil,
				Content: config.GetContent(args...),
				Args:    args,
			}, nil
		}
		return nil, fmt.Errorf("applicant %d has no grade", applicantId)
	}

	condition := "pass.group2"
	admitGroups := lo.Map(admits, func(item *model.Admit, index int) string {
		if item.Group_ == "机械" || item.Group_ == "电控" || item.Group_ == "视觉" {
			condition = "pass.group1"
		}
		return item.Group_ + "组"
	})
	groupStr := strings.Join(admitGroups, "、")
	config := Configs["录取结果"][condition]
	args := []string{"录取", applicant.Name, groupStr}

	return &Sms{
		Status:  status,
		Time:    nil,
		Content: config.GetContent(args...),
		Args:    args,
	}, nil
}
