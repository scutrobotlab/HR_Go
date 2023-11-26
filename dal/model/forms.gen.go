// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameForm = "forms"

// Form mapped from table <forms>
type Form struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(32);not null;comment:显示名称" json:"name"`                                                 // 显示名称
	Key       string         `gorm:"column:key;type:varchar(32);not null;uniqueIndex:forms_key_unique,priority:1;comment:存储键" json:"key"`            // 存储键
	Type      string         `gorm:"column:type;type:varchar(32);not null;default:text-field;comment:表单类型 {text-field,select,textarea}" json:"type"` // 表单类型 {text-field,select,textarea}
	Required  bool           `gorm:"column:required;type:tinyint(1);not null;comment:是否必填 0-否 1-是" json:"required"`                                  // 是否必填 0-否 1-是
	Options   string         `gorm:"column:options;type:json;comment:选项列表（若type=select）" json:"options"`                                             // 选项列表（若type=select）
	Regexp    string         `gorm:"column:regexp;type:varchar(64);comment:正则校验" json:"regexp"`                                                      // 正则校验
	MaxLength int32          `gorm:"column:max_length;type:int;comment:最大长度" json:"max_length"`                                                      // 最大长度
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
}

// TableName Form's table name
func (*Form) TableName() string {
	return TableNameForm
}
