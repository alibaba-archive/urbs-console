package schema

import "time"

// UrbsAcUser 详见 ./sql/schema.sql table `urbs_ac_user`
// 用户表
type UrbsAcUser struct {
	ID        int64     `gorm:"column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`

	UID  string `gorm:"column:uid" json:"uid"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName retuns table name
func (UrbsAcUser) TableName() string {
	return "urbs_ac_user"
}
