package dto

import "github.com/teambition/urbs-console/src/schema"

// OperationLog ...
type OperationLog struct {
	schema.OperationLog
	Name string `gorm:"column:name"`
}
