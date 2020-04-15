package dao

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// OperationLog table `operation_log`
type OperationLog struct {
	DB *gorm.DB
}

// Add ...
func (a *OperationLog) Add(ctx context.Context, obj *schema.OperationLog) error {
	obj.CreatedAt = time.Now().UTC()

	sql := "insert ignore into `operation_log` (`created_at`, `operator`, `object`,`action`,`content`, `description`) values (?, ?, ?, ?, ?, ?)"

	args := []interface{}{obj.CreatedAt, obj.Operator, obj.Object, obj.Action, obj.Content, obj.Desc}

	_, err := a.DB.DB().Exec(sql, args...)

	return err
}

// FindByObject ...
func (a *OperationLog) FindByObject(ctx context.Context, object string, pg *tpl.Pagination) ([]*dto.OperationLog, error) {
	sql := "SELECT a.id,a.created_at,a.operator,a.object,a.action,a.content,a.description,b.`name` FROM operation_log a LEFT JOIN urbs_ac_user b ON a.operator=b.uid WHERE a.object = ? ORDER BY a.id DESC LIMIT ?,?"
	row, err := a.DB.Raw(sql, object, pg.Skip, pg.PageSize+1).Rows()
	if err != nil {
		return nil, err
	}
	data := make([]*dto.OperationLog, 0)

	for row.Next() {
		log := &dto.OperationLog{}
		err := row.Scan(&log.ID, &log.CreatedAt, &log.Operator, &log.Object, &log.Action, &log.Desc, &log.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, log)
	}
	return data, nil
}
