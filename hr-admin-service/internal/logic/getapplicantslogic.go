package logic

import (
	"HR_Go/common"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"fmt"
	"github.com/samber/lo"
	"math"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplicantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplicantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplicantsLogic {
	return &GetApplicantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplicantsLogic) GetApplicants(in *hr_admin_service.GetApplicantsReq) (*hr_admin_service.GetApplicantsResp, error) {
	var PageUrl = fmt.Sprintf("%s/api/applicant/group", l.svcCtx.Common.Url.BaseUrl)
	const PerPage = 15
	currentPage := int(in.Page)
	offset := (currentPage - 1) * PerPage

	const CountScore = "count_score"
	const AvgScore = "avg_score"

	db := l.svcCtx.Db
	s := l.svcCtx.Query.Score
	a := l.svcCtx.Query.Applicant
	at := l.svcCtx.Query.ApplicantTime
	t := l.svcCtx.Query.Time
	ad := l.svcCtx.Query.Admit

	type Result struct {
		Id         int64
		Name       string
		Gender     int32
		Avatar     string
		CountScore int32
		AvgScore   float32
		Time       time.Time
		Location   string
		Extend     string
	}

	var err error
	var result []*Result
	var total int

	switch in.State {
	case "未评分":
		query := db.WithContext(l.ctx).Table("intents as i").
			Select("a.id, a.name, a.gender, a.avatar, t.time, t.location, at.extend").
			Joins("left join applicants as a on i.applicant_id = a.id").
			Joins("left join applicant_times as at on a.id = at.applicant_id and i.group = at.group and at.deleted_at is null").
			Joins("left join times as t on at.time_id = t.id and t.deleted_at is null").
			Joins("left join scores as s on s.applicant_id = i.applicant_id and s.group = i.group").
			Where("i.group = ?", in.Group).
			Where("i.deleted_at is null").
			Where("a.deleted_at is null").
			Where("s.id is null").
			Order("at.id is null").
			Order("abs(t.time - current_timestamp)").
			Order("a.id")
		var count int64
		query.Count(&count)
		total = int(count)
		err = query.Limit(PerPage).Offset(offset).Scan(&result).Error
	case "已评分":
		query := s.WithContext(l.ctx).Select(a.ID, a.Name, a.Gender, a.Avatar,
			s.Score.Count().As(CountScore), s.Score.Avg().As(AvgScore), t.Time, t.Location, at.Extend).
			LeftJoin(a, s.ApplicantID.EqCol(a.ID)).
			LeftJoin(at, a.ID.EqCol(at.ApplicantID), s.Group_.EqCol(at.Group_), at.DeletedAt.IsNull()).
			LeftJoin(t, at.TimeID.EqCol(t.ID), t.DeletedAt.IsNull()).
			Where(s.Group_.Eq(in.Group), a.DeletedAt.IsNull(), at.DeletedAt.IsNull(), t.DeletedAt.IsNull()).
			Group(s.ApplicantID, t.ID, at.ID).
			Order(s.Score.Avg().Desc())
		count, _ := query.Count()
		total = int(count)
		err = query.Limit(PerPage).Offset(offset).Scan(&result)
	case "已录取":
		query := ad.WithContext(l.ctx).Select(a.ID, a.Name, a.Gender, a.Avatar,
			s.Score.Count().As(CountScore), s.Score.Avg().As(AvgScore), t.Time, t.Location, at.Extend).
			LeftJoin(s, ad.ApplicantID.EqCol(s.ApplicantID), ad.Group_.EqCol(s.Group_)).
			LeftJoin(a, s.ApplicantID.EqCol(a.ID)).
			LeftJoin(at, a.ID.EqCol(at.ApplicantID), s.Group_.EqCol(at.Group_), at.DeletedAt.IsNull()).
			LeftJoin(t, at.TimeID.EqCol(t.ID), t.DeletedAt.IsNull()).
			Where(ad.DeletedAt.IsNull(), s.Group_.Eq(in.Group), a.DeletedAt.IsNull()).
			Group(s.ApplicantID, t.ID, at.ID).
			Order(s.Score.Avg().Desc())
		count, _ := query.Count()
		total = int(count)
		err = query.Limit(PerPage).Offset(offset).Scan(&result)
	case "临近面试":
		const CloseToMinutes = 40 // 40 minutes
		query := db.WithContext(l.ctx).Table("intents as i").
			Select("a.id, a.name, a.gender, a.avatar, t.time, t.location, at.extend").
			Joins("left join applicants as a on i.applicant_id = a.id").
			Joins("left join applicant_times as at on a.id = at.applicant_id and i.group = at.group and at.deleted_at is null").
			Joins("left join times as t on at.time_id = t.id and t.deleted_at is null").
			Where("i.group = ?", in.Group).
			Where("i.deleted_at is null").
			Where("a.deleted_at is null").
			Where("abs(timestampdiff(minute, t.time, current_timestamp)) < ?", CloseToMinutes).
			Order("at.id is null").
			Order("abs(t.time - current_timestamp)").
			Order("a.id")
		var count int64
		query.Count(&count)
		total = int(count)
		err = query.Limit(PerPage).Offset(offset).Scan(&result).Error
	case "需要评分":
		query := db.WithContext(l.ctx).Table("intents as i").
			Select("a.id, a.name, a.gender, a.avatar, t.time, t.location, at.extend").
			Joins("left join applicants as a on i.applicant_id = a.id").
			Joins("left join applicant_times as at on a.id = at.applicant_id and i.group = at.group and at.deleted_at is null").
			Joins("left join times as t on at.time_id = t.id and t.deleted_at is null").
			Joins("left join scores as s on s.applicant_id = i.applicant_id and s.group = i.group").
			Where("i.group = ?", in.Group).
			Where("i.deleted_at is null").
			Where("a.deleted_at is null").
			Where("s.id is null").
			Where("(JSON_EXTRACT(at.extend, '$.receive_status') = '准时签到' or JSON_EXTRACT(at.extend, '$.receive_status') = '迟到签到')").
			Order("at.id is null").
			Order("abs(t.time - current_timestamp)").
			Order("a.id")
		var count int64
		query.Count(&count)
		total = int(count)
		err = query.Limit(PerPage).Offset(offset).Scan(&result).Error
	default:
		return nil, fmt.Errorf("invalid state: %s", in.State)
	}
	if err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(PerPage)))
	if lastPage == 0 {
		lastPage = 1
	}
	nextPage := int(in.Page) + 1
	if nextPage > lastPage {
		nextPage = lastPage
	}
	prevPage := int(in.Page) - 1
	if prevPage < 1 {
		prevPage = 1
	}
	pageSize := PerPage
	if currentPage == lastPage {
		pageSize = total - (currentPage-1)*PerPage
	}

	data := lo.Map(result, func(item *Result, index int) *hr_admin_service.ApplicantListItem {
		date1, time1 := "", ""
		if !item.Time.IsZero() {
			date1 = item.Time.Format(time.DateOnly)
			time1 = item.Time.Format(time.TimeOnly)
		}
		extend := "{}"
		if item.Extend != "" {
			extend = item.Extend
		}
		admits, _ := getAdmitByApplicantId(l.ctx, l.svcCtx, item.Id)
		return &hr_admin_service.ApplicantListItem{
			Id:          item.Id,
			Name:        item.Name,
			Gender:      common.GenderIntToString(item.Gender),
			Avatar:      item.Avatar,
			ScoresCount: item.CountScore,
			AvgScore:    item.AvgScore,
			Date:        date1,
			Time:        time1,
			Location:    item.Location,
			Extend:      extend,
			Admits:      admits,
		}
	})
	links := make([]*hr_admin_service.Link, 0)
	for page := 1; page <= lastPage; page++ {
		links = append(links, &hr_admin_service.Link{
			Url:    fmt.Sprintf("%s?page=%d", PageUrl, page),
			Label:  strconv.Itoa(page),
			Active: page == currentPage,
		})
	}

	return &hr_admin_service.GetApplicantsResp{
		Data:         data,
		Links:        links,
		CurrentPage:  in.Page,
		Path:         PageUrl,
		PerPage:      PerPage,
		LastPage:     int32(lastPage),
		FirstPageUrl: fmt.Sprintf("%s?page=%d", PageUrl, 1),
		NextPageUrl:  fmt.Sprintf("%s?page=%d", PageUrl, nextPage),
		PrevPageUrl:  fmt.Sprintf("%s?page=%d", PageUrl, prevPage),
		LastPageUrl:  fmt.Sprintf("%s?page=%d", PageUrl, lastPage),
		From:         in.Page*PerPage + 1,
		To:           in.Page*PerPage + int32(pageSize),
		Total:        int32(total),
	}, nil
}
