package dto

import "github.com/teambition/urbs-console/src/schema"

// UrbsAcAcl ...
type UrbsAcAcl struct {
	schema.UrbsAcAcl
	Name string `gorm:"column:name"`
}
