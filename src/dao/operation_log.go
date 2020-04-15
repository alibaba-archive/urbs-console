package dao

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// OperationLog table `operation_log`
type OperationLog struct {
	DB *gorm.DB
}

// Add ...
func (a *OperationLog) Add(ctx context.Context, obj *schema.OperationLog) error {
	sql := "insert ignore into `operation_log` (`created_at`, `operator`, `object`,`action`,`content`, `description`) values (?, ?, ?, ?, ?, ?)"

	args := []interface{}{obj.CreatedAt, obj.Operator, obj.Object, obj.Action, obj.Content, obj.Desc}

	_, err := a.DB.DB().Exec(sql, args...)

	return err
}

// FindByObject ...
func (a *OperationLog) FindByObject(ctx context.Context, object string, pg *tpl.Pagination) ([]*schema.OperationLog, error) {
	operationLogs := make([]*schema.OperationLog, 0)
	where := "object = ?"
	orderBy := "id desc"

	err := a.DB.Where(where, object).Order(orderBy).Offset(pg.Skip).Limit(pg.PageSize + 1).Find(&operationLogs).Error
	if err != nil {
		return nil, err
	}
	return operationLogs, nil
}
