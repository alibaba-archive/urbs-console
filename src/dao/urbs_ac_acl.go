package dao

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/dto"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcAcl table `urbs_ac_acl`
type UrbsAcAcl struct {
	DB *gorm.DB
}

// Add ...
func (a *UrbsAcAcl) Add(ctx context.Context, obj *schema.UrbsAcAcl) error {

	sql := "insert ignore into `urbs_ac_acl` (`subject`, `object`,`permission`) values ( ?, ?, ?)"

	args := []interface{}{obj.Subject, obj.Object, obj.Permission}
	return a.DB.Exec(sql, args...).Error
}

// FindOne ...
func (a *UrbsAcAcl) FindOne(ctx context.Context, subject, object, permission string) (*schema.UrbsAcAcl, error) {
	sql := "SELECT a.id, a.created_at, a.subject, a.object, a.permission FROM urbs_ac_acl a INNER JOIN urbs_ac_user b ON a.`subject`=b.uid WHERE a.subject = ? and a.object = ? and a.permission = ?"
	row, err := a.DB.Raw(sql, subject, object, permission).Rows()
	defer row.Close()
	if err != nil {
		return nil, err
	}
	urbsAcAcl := &schema.UrbsAcAcl{}
	for row.Next() {
		err := row.Scan(&urbsAcAcl.ID, &urbsAcAcl.CreatedAt, &urbsAcAcl.Subject, &urbsAcAcl.Object, &urbsAcAcl.Permission)
		if err != nil {
			return nil, err
		}
	}
	if urbsAcAcl.ID == 0 {
		return nil, errors.New("not found")
	}
	return urbsAcAcl, nil
}

// FindBySubjects ...
func (a *UrbsAcAcl) FindBySubjects(ctx context.Context, subjects []string) ([]*schema.UrbsAcAcl, error) {
	where := "subject in ( ? )"

	urbsAcAcl := []*schema.UrbsAcAcl{}
	err := a.DB.Where(where, subjects).Find(&urbsAcAcl).Error
	if err != nil {
		return nil, err
	}
	return urbsAcAcl, nil
}

// DeleteByObject ...
func (a *UrbsAcAcl) DeleteByObject(ctx context.Context, object string) error {
	sql := "delete from `urbs_ac_acl` where object = ?"
	return a.DB.Exec(sql, object).Error
}

// Delete ...
func (a *UrbsAcAcl) Delete(ctx context.Context, subject, object, permission string) error {
	sql := "delete from `urbs_ac_acl` where subject = ? and object = ? and permission = ?"

	args := []interface{}{subject, object, permission}
	return a.DB.Exec(sql, args...).Error
}

// UpdateSubjects ...
func (a *UrbsAcAcl) UpdateSubjects(ctx context.Context, subjects []string, object string, permission string) error {
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		sql := "delete from `urbs_ac_acl` where object = ? and permission = ?"
		err := tx.Exec(sql, object, permission).Error
		if err != nil {
			return err
		}
		for _, subject := range subjects {
			sql := "insert ignore into `urbs_ac_acl` (`subject`,`object`,`permission`) values ( ?, ?, ?)"
			args := []interface{}{subject, object, permission}
			err = tx.Exec(sql, args...).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// FindByObjects ...
func (a *UrbsAcAcl) FindByObjects(ctx context.Context, objects []string) ([]*dto.UrbsAcAcl, error) {
	sql := `SELECT
				a.id,
				a.created_at,
				a.subject,
				a.object,
				a.permission,
				b.name
			FROM
				urbs_ac_acl a
				INNER JOIN urbs_ac_user b ON a.subject = b.uid 
			WHERE
				a.object IN (?)`
	row, err := a.DB.Raw(sql, objects).Rows()
	defer row.Close()
	if err != nil {
		return nil, err
	}
	data := make([]*dto.UrbsAcAcl, 0)

	for row.Next() {
		acl := &dto.UrbsAcAcl{}
		err := row.Scan(&acl.ID, &acl.CreatedAt, &acl.Subject, &acl.Object, &acl.Permission, &acl.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, acl)
	}
	return data, nil
}
