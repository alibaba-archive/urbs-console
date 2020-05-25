package schema

import (
	"encoding/json"
	"time"
)

// OperationLog 详见 ./sql/schema.sql table `operation_log`
// 操作日志
type OperationLog struct {
	ID        int64     `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Operator string `gorm:"column:operator"` // 操作人
	Object   string `gorm:"column:object"`   // 操作对象
	Action   string `gorm:"column:action"`   // 操作行为
	// Content 操作内容
	Content string `gorm:"column:content"`
	Desc    string `gorm:"column:description"` // 操作说明
}

// TableName retuns table name
func (OperationLog) TableName() string {
	return "operation_log"
}

// String ...
func (a *OperationLog) String() string {
	log, _ := json.Marshal(a)
	return string(log)
}
