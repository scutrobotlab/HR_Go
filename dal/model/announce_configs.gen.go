// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAnnounceConfig = "announce_configs"

// AnnounceConfig mapped from table <announce_configs>
type AnnounceConfig struct {
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	Status string `gorm:"column:status;not null;comment:状态名称" json:"status"`
	Body   string `gorm:"column:body;not null;comment:状态内容" json:"body"`
}

// TableName AnnounceConfig's table name
func (*AnnounceConfig) TableName() string {
	return TableNameAnnounceConfig
}
