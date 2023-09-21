package logic

import (
	"testing"
	"time"
)

func TestGetApplicantsLogic(t1 *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()

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

	group := "视觉"
	db := svcCtx.Db
	//i := svcCtx.Query.Intent
	s := svcCtx.Query.Score
	a := svcCtx.Query.Applicant
	at := svcCtx.Query.ApplicantTime
	t := svcCtx.Query.Time
	ad := svcCtx.Query.Admit

	var result []*Result
	var err error

	sql1 := db.Table("intents as i").
		Select("a.id, a.name, a.gender, a.avatar, t.time, t.location, at.extend").
		Joins("left join applicants as a on i.applicant_id = a.id").
		Joins("left join applicant_times as at on a.id = at.applicant_id and i.group = at.group").
		Joins("left join times as t on at.time_id = t.id").
		Where("i.group = ?", group).
		Where("i.deleted_at is null").
		Where("a.deleted_at is null").
		Where("not exists (select 1 from scores as s where s.applicant_id = i.applicant_id and s.group = i.group)").
		Order("at.id is null").
		Order("abs(t.time - current_timestamp)").
		Order("a.id")

	//sql1 := i.WithContext(ctx).Select(a.ID, a.Name, a.Gender, a.Gender, t.Time, t.Location, at.Extend).
	//	LeftJoin(a, i.ApplicantID.EqCol(a.ID)).
	//	LeftJoin(at, a.ID.EqCol(at.ApplicantID), i.Group_.EqCol(at.Group_)).
	//	LeftJoin(t, at.TimeID.EqCol(t.ID)).
	//	Where(i.Group_.Eq(group), a.DeletedAt.IsNull(), at.DeletedAt.IsNull(), t.DeletedAt.IsNull(),
	//		i.Not(i.Exists(s.WithContext(ctx).Where(s.ApplicantID.EqCol(i.ApplicantID), s.Group_.EqCol(i.Group_))))).
	//	Order(at.ID.IsNull(), t.Time, a.ID)

	//total, _ := sql1.ScoreCount()
	//err := sql1.Limit(15).Offset(0).Scan(&result)

	var total int64
	sql1.Count(&total)

	sql1.Limit(15).Offset(0).Scan(&result)

	t1.Logf("total: %d", total)
	for _, r := range result {
		t1.Logf("%+v", r)
	}

	sql2 := s.WithContext(ctx).Select(a.ID, a.Name, a.Gender, a.Gender,
		s.Score.Count().As("count_score"), s.Score.Avg().As("avg_score"), t.Time, t.Location, at.Extend).
		LeftJoin(a, s.ApplicantID.EqCol(a.ID)).
		LeftJoin(at, a.ID.EqCol(at.ApplicantID), s.Group_.EqCol(at.Group_)).
		LeftJoin(t, at.TimeID.EqCol(t.ID)).
		Where(s.Group_.Eq(group), a.DeletedAt.IsNull(), at.DeletedAt.IsNull(), t.DeletedAt.IsNull()).
		Group(s.ApplicantID, t.ID, at.ID).
		Order(s.Score.Avg().Desc())

	total, _ = sql2.Count()
	err = sql2.Limit(15).Offset(0).Scan(&result)

	if err != nil {
		return
	}

	t1.Logf("total: %d", total)
	for _, r := range result {
		t1.Logf("%+v", r)
	}

	sql3 := ad.WithContext(ctx).Select(a.ID, a.Name, a.Gender, a.Gender,
		s.Score.Count().As("count_score"), s.Score.Avg().As("avg_score"), t.Time, t.Location, at.Extend).
		LeftJoin(s, ad.ApplicantID.EqCol(s.ApplicantID), ad.Group_.EqCol(s.Group_)).
		LeftJoin(a, s.ApplicantID.EqCol(a.ID)).
		LeftJoin(at, a.ID.EqCol(at.ApplicantID), s.Group_.EqCol(at.Group_)).
		LeftJoin(t, at.TimeID.EqCol(t.ID)).
		Where(ad.DeletedAt.IsNull(), s.Group_.Eq(group), a.DeletedAt.IsNull(), at.DeletedAt.IsNull(), t.DeletedAt.IsNull()).
		Group(s.ApplicantID, t.ID, at.ID).
		Order(s.Score.Avg().Desc())

	total, _ = sql3.Count()
	err = sql3.Limit(15).Offset(0).Scan(&result)

	if err != nil {
		return
	}

	t1.Logf("total: %d", total)
	for _, r := range result {
		t1.Logf("%+v", r)
	}
}
