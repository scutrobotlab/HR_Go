package logic

import (
	"HR_Go/dal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"sort"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGetScheduleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGetScheduleLogic {
	return &GetGetScheduleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGetScheduleLogic) GetSchedule(in *hr_admin_service.GetScheduleReq) (*hr_admin_service.GetScheduleResp, error) {
	type Result struct {
		ID         int64
		Name       string
		Time       time.Time
		TotalCnt   int64
		ScoreCount int64
		AtNull     bool
		RoomID     int64
		Extend     string
	}

	db := l.svcCtx.Db
	r := l.svcCtx.Query.Room
	tc := l.svcCtx.Query.TimeConfig

	rooms, err := r.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}
	roomMap := lo.SliceToMap(rooms, func(item *model.Room) (int64, *model.Room) {
		return item.ID, item
	})

	var timeTable []*hr_admin_service.TimeTableItem
	roomCategories := make(map[string]struct{})
	addCategories := make(map[string]struct{})
	usedTime := make(map[string]int)
	totalCntMap := make(map[string]int)

	recordUsedTime := func(theTime time.Time) string {
		timeStr := theTime.Format(time.DateTime)
		if used, ok := usedTime[timeStr]; ok {
			usedTime[timeStr] = used + 1
		} else {
			usedTime[timeStr] = 1
		}
		return fmt.Sprintf("未分配%d", usedTime[timeStr])
	}

	focus := time.Now()
	if in.Date != "" {
		parse, err := time.Parse("2006-01-02", in.Date)
		if err == nil {
			focus = parse
		}
	} else {
		timeConfig, err := tc.WithContext(l.ctx).Where(tc.Key.Eq("first_interview_date")).First()
		if err != nil {
			timeConfig = &model.TimeConfig{
				Key:   "first_interview_date",
				Value: time.Now(),
			}
			_ = tc.WithContext(l.ctx).Save(timeConfig)
		}
		focus, err = time.Parse("2006-01-02", timeConfig.Value.Format("2006-01-02"))
	}

	startDate, endDate := focus, focus
	if in.Type == "category" {
		endDate = endDate.Add(24 * time.Hour)
	} else if in.Type == "4day" {
		endDate = endDate.Add(4 * 24 * time.Hour)
	} else {
		return nil, fmt.Errorf("invalid type: %s", in.Type)
	}

	var result []*Result
	err = db.WithContext(l.ctx).Table("times as t").
		Select("a.id, a.name, t.time, t.total_cnt, count(s.id) as score_count, at.id is null as at_null, at.room_id, at.extend").
		Joins("left join applicant_times as at on t.id = at.time_id and at.deleted_at is null").
		Joins("left join applicants as a on at.applicant_id = a.id and a.deleted_at is null").
		Joins("left join scores as s on a.id = s.applicant_id and at.group = s.group and s.deleted_at is null").
		Where("t.deleted_at is null").
		Where("t.group = ?", in.Group).
		Where("t.time between ? and ?", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly)).
		Group("a.id, a.name, t.time, t.total_cnt, at.id, at.room_id, at.extend").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		start := item.Time.Format(time.DateTime)
		end := item.Time.Add(InterviewDuration).Format(time.DateTime)
		totalCntMap[start] = int(item.TotalCnt)
		if item.AtNull {
			continue
		}
		name, color, desc, category, id := "", "", "", "", int64(0)
		if item.RoomID != 0 {
			color, desc = "done", "已面试"
			recordUsedTime(item.Time)
			category = roomMap[item.RoomID].Name
			roomCategories[category] = struct{}{}
		} else {
			color, desc = getColorAndDesc(item.Extend)
			category = recordUsedTime(item.Time)
			addCategories[category] = struct{}{}
		}
		if item.ScoreCount > 0 {
			color, desc = "success", "已评分"
		}
		name, id = item.Name, item.ID
		if desc != "" {
			name += fmt.Sprintf("(%s)", desc)
		}
		timeTable = append(timeTable, &hr_admin_service.TimeTableItem{
			Name:     name,
			Start:    start,
			End:      end,
			Color:    color,
			Category: category,
			Id:       id,
		})
	}

	if in.ShowNotSelected {
		for datetime, count := range totalCntMap {
			start, _ := time.Parse(time.DateTime, datetime)
			endStr := start.Add(InterviewDuration).Format(time.DateTime)
			remainCount := count - usedTime[datetime]
			for i := 0; i < remainCount; i++ {
				timeTable = append(timeTable, &hr_admin_service.TimeTableItem{
					Name:     "未分配",
					Start:    datetime,
					End:      endStr,
					Color:    "none",
					Category: recordUsedTime(start),
					Id:       0,
				})
			}
		}
	}

	categories := lo.MapToSlice(roomCategories, func(key string, value struct{}) string {
		return key
	})
	for key := range addCategories {
		categories = append(categories, fmt.Sprintf("%s", key))
	}
	sort.Slice(categories, func(i, j int) bool {
		return categories[i] < categories[j]
	})
	return &hr_admin_service.GetScheduleResp{
		TimeTable:  timeTable,
		Categories: categories,
		Focus:      focus.Format("2006-01-02"),
	}, nil
}

func getColorAndDesc(extendStr string) (color, desc string) {
	color = "default"
	var extend map[string]interface{}
	_ = json.Unmarshal([]byte(extendStr), &extend)
	if receiveStatus, ok := extend["receive_status"].(string); ok {
		switch receiveStatus {
		case "准时签到", "迟到签到":
			color = "ready"
		case "已请假", "已改期", "无故缺席":
			color = "danger"
		}
		desc = receiveStatus
	}
	return color, desc
}
