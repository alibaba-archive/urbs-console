package schema

import "time"

// UrbsLock 详见 ./sql/schema.sql table `urbs_lock`
type UrbsLock struct {
	ID       int64     `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	ExpireAt time.Time `gorm:"column:expire_at"`
}

// TableName retuns table name
func (UrbsLock) TableName() string {
	return "urbs_lock"
}
