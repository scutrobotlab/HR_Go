package common

import (
	"HR_Go/dal/model"
	"HR_Go/dal/query"
	"github.com/samber/lo"
	"time"
)

const (
	ApplyFormStart         = "apply_form_start"
	ApplyFormEnd           = "apply_form_end"
	SelectTimeStart        = "select_time_start"
	SelectTimeEnd          = "select_time_end"
	PublishResultStart     = "publish_result_start"
	ShowInterviewQuestions = "show_interview_questions"
	ReviewThisSeason       = "review_this_season"
	FirstInterviewDate     = "first_interview_date"
)

func GetTimeConfig(q *query.Query, key string) time.Time {
	timeConfigs, err := q.TimeConfig.Find()
	if err != nil {
		return time.Time{}
	}

	find, b := lo.Find(timeConfigs, func(item *model.TimeConfig) bool {
		return item.Key == key
	})
	if b {
		return find.Value
	} else {
		return time.Time{}
	}
}
