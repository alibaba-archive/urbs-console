package dao

import (
	"context"

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

	sql := "insert ignore into `operation_log` ( `operator`, `object`,`action`,`content`, `description`) values (?, ?, ?, ?, ?)"

	args := []interface{}{obj.Operator, obj.Object, obj.Action, obj.Content, obj.Desc}

	return a.DB.Exec(sql, args...).Error
}

// FindByObject ...
func (a *OperationLog) FindByObject(ctx context.Context, object string, pg *tpl.ConsolePagination) ([]*dto.OperationLog, error) {
	sql := "SELECT a.id, a.created_at, a.operator, a.object, a.action, a.content, a.description, b.`name` FROM operation_log a LEFT JOIN urbs_ac_user b ON a.operator=b.uid WHERE a.object = ? ORDER BY a.id DESC LIMIT ?,?"

	row, err := a.DB.Raw(sql, object, pg.Skip, pg.PageSize+1).Rows()
	defer row.Close()
	if err != nil {
		return nil, err
	}
	data := make([]*dto.OperationLog, 0)

	for row.Next() {
		log := &dto.OperationLog{}
		err := row.Scan(&log.ID, &log.CreatedAt, &log.Operator, &log.Object, &log.Action, &log.Content, &log.Desc, &log.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, log)
	}
	return data, nil
}

// FindOneByID ...
func (a *OperationLog) FindOneByID(ctx context.Context, id int64) (*schema.OperationLog, error) {
	sql := "select * from operation_log where id = ?"

	log := &schema.OperationLog{}
	err := a.DB.Raw(sql, id).Scan(log).Error
	if err != nil {
		return nil, err
	}
	return log, nil
}

// FindOneByObject ...
func (a *OperationLog) FindOneByObject(ctx context.Context, object string) (*schema.OperationLog, error) {
	sql := "select * from operation_log where object = ? limit 1"

	log := &schema.OperationLog{}
	err := a.DB.Raw(sql, object).Scan(log).Error
	if err != nil {
		return nil, err
	}
	return log, nil
}

// FindAll ...
func (a *OperationLog) FindAll(ctx context.Context) ([]schema.OperationLog, error) {
	sql := "select * from operation_log limit 10"

	logs := []schema.OperationLog{}
	err := a.DB.Raw(sql).Scan(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// TxDelete ...
func (a *OperationLog) TxDelete(ctx context.Context, id int64, msgSql string) error {
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec(msgSql).Error
		if err != nil {
			return err
		}
		sql := "delete from operation_log where id = ?"
		err = tx.Exec(sql, id).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Delete ...
func (a *OperationLog) Delete(ctx context.Context, id int64) error {
	sql := "delete from operation_log where id = ?"
	return a.DB.Exec(sql, id).Error
}

// CountByObject ...
func (a *OperationLog) CountByObject(ctx context.Context, object string) (int, error) {
	sql := "select count(1) as count from operation_log where object = ?"

	res := &schema.CountResult{}
	err := a.DB.Raw(sql, object).Scan(res).Error
	return res.Count, err
}

// FindByObjectWithHandler ...
func (a *OperationLog) FindByObjectWithHandler(ctx context.Context, object string, handler func(*schema.OperationLog)) error {
	sql := "select * from operation_log where object = ?"
	cursor, err := a.DB.Raw(sql, object).Rows()
	if err != nil {
		return err
	}
	defer cursor.Close()
	for cursor.Next() {
		log := &schema.OperationLog{}
		err := cursor.Scan(log)
		if err != nil {
			return err
		}
		handler(log)
	}
	return nil
}
