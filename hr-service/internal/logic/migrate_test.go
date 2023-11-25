package logic

import (
	"HR_Go/dal/model"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	_, svcCtx := GetCtxAndSvcCtxForTest()
	db := svcCtx.Db

	err := db.AutoMigrate(
		&model.Admin{},
		&model.Admit{},
		&model.AnnounceConfig{},
		&model.ApplicantQuestion{},
		&model.ApplicantSm{},
		&model.ApplicantTime{},
		&model.Applicant{},
		&model.Configuration{},
		&model.EvaluationStandard{},
		&model.Form{},
		&model.Guide{},
		&model.Intent{},
		&model.Question{},
		&model.ReceiveSm{},
		&model.Remark{},
		&model.Room{},
		&model.Score{},
		&model.TimeConfig{},
		&model.Time{},
	)
	if err != nil {
		t.Error(err)
		return
	}
}
