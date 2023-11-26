package common

import (
	"HR_Go/dal/model"
	"gorm.io/gorm"
)

var Models = []any{
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
}

// AutoMigrate 自动迁移
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(Models...)
}
