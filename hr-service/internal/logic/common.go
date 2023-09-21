package logic

import (
	"HR_Go/hr-service/internal/svc"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"
)

type Profile struct {
	OpenId     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Language   string   `json:"language"`
	City       string   `json:"city"`
	Province   string   `json:"province"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
}

func GetTimeAndApplicantTime(ctx context.Context, svcCtx *svc.ServiceContext, applicantId int64) []*hr_service.TimeItem {
	db := svcCtx.Db
	var result []*hr_service.TimeItem
	err := db.WithContext(ctx).Table("applicant_times as at").
		Select("at.id as applicant_id, t.id, t.group, date(t.time) as date, time(t.time) as time, t.rank, t.location, t.total_cnt, t.grade, t.campus, t.meeting_id").
		Joins("left join times as t on t.id = at.time_id and t.deleted_at is null").
		Where("at.deleted_at is null").
		Where("at.applicant_id = ?", applicantId).
		Scan(&result).Error
	if err != nil {
		return nil
	}

	return result
}
