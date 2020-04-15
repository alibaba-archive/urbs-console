package schema

import "time"

// UrbsAcAcl 详见 ./sql/schema.sql table `urbs_ac_acl`
// 用户访问控制
type UrbsAcAcl struct {
	ID        int64     `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Subject    string `gorm:"column:subject"`    // who
	Object     string `gorm:"object:object"`     // what
	Permission string `gorm:"column:permission"` // How
}

// TableName retuns table name
func (UrbsAcAcl) TableName() string {
	return "urbs_ac_acl"
}
