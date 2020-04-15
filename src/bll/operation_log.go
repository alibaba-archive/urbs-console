package bll

import (
	"context"
	"time"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// OperationLog table `operation_log`
type OperationLog struct {
	dao *dao.Daos
}

// Add ...
func (a *OperationLog) Add(ctx context.Context, obj *schema.OperationLog) error {
	obj.CreatedAt = time.Now().UTC()
	return a.dao.OperationLog.Add(ctx, obj)
}

// List 返回操作日志列表
func (a *OperationLog) List(ctx context.Context, object string, pg *tpl.Pagination) ([]*schema.OperationLog, error) {
	return a.dao.OperationLog.FindByObject(ctx, object, pg)
}
