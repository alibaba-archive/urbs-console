package schema

import "time"

// UrbsAcUser 详见 ./sql/schema.sql table `urbs_ac_user`
// 用户表
type UrbsAcUser struct {
	ID        int64     `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`

	UID  string `gorm:"object:uid"`
	Name string `gorm:"column:name"`
}

// TableName retuns table name
func (UrbsAcUser) TableName() string {
	return "urbs_ac_user"
}
