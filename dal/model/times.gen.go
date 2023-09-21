// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTime = "times"

// Time mapped from table <times>
type Time struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	Group_    string         `gorm:"column:group;not null;comment:组别" json:"group"`
	Time      time.Time      `gorm:"column:time;not null;default:curtime();comment:时间" json:"time"`
	Rank      int32          `gorm:"column:rank;not null;comment:顺序" json:"rank"`
	Location  string         `gorm:"column:location;comment:地点" json:"location"`
	TotalCnt  int32          `gorm:"column:total_cnt;not null;default:1;comment:总数" json:"total_cnt"`
	Campus    string         `gorm:"column:campus;comment:校区" json:"campus"`
	Grade     string         `gorm:"column:grade;comment:年级" json:"grade"`
	MeetingID string         `gorm:"column:meeting_id;comment:会议ID" json:"meeting_id"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

// TableName Time's table name
func (*Time) TableName() string {
	return TableNameTime
}