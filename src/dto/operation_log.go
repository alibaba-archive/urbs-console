package dto

import "github.com/teambition/urbs-console/src/schema"

// OperationLog ...
type OperationLog struct {
	schema.OperationLog
	Name string `gorm:"column:name"`
}

// OperationLogContent ...
type OperationLogContent struct {
	Users   []string
	Groups  []string
	Desc    string
	Value   string
	Kind    string
	Percent *int

	Release int64
}
