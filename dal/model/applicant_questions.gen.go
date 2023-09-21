// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameApplicantQuestion = "applicant_questions"

// ApplicantQuestion mapped from table <applicant_questions>
type ApplicantQuestion struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	ApplicantID int64          `gorm:"column:applicant_id;not null;comment:申请ID" json:"applicant_id"`
	QuestionID  int64          `gorm:"column:question_id;not null;comment:问题ID" json:"question_id"`
	Answer      string         `gorm:"column:answer;comment:答案" json:"answer"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

// TableName ApplicantQuestion's table name
func (*ApplicantQuestion) TableName() string {
	return TableNameApplicantQuestion
}
