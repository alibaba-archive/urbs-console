package dao

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcAcl table `urbs_ac_acl`
type UrbsAcAcl struct {
	DB *gorm.DB
}

// Add ...
func (a *UrbsAcAcl) Add(ctx context.Context, obj *schema.UrbsAcAcl) error {
	sql := "insert ignore into `urbs_ac_acl` (`created_at`, `subject`, `object`,`permission`) values (?, ?, ?, ?)"

	args := []interface{}{obj.CreatedAt, obj.Subject, obj.Object, obj.Permission}

	_, err := a.DB.DB().Exec(sql, args...)

	return err
}

// FindOne ...
func (a *UrbsAcAcl) FindOne(ctx context.Context, subject, object, permission string) (*schema.UrbsAcAcl, error) {
	urbsAcAcl := &schema.UrbsAcAcl{}

	where := "subject = ? and object = ? and permission = ?"

	args := []interface{}{subject, object, permission}

	err := a.DB.Where(where, args...).Find(&urbsAcAcl).Error
	if err != nil {
		return nil, err
	}
	return urbsAcAcl, nil
}

// Remove ...
func (a *UrbsAcAcl) Remove(ctx context.Context, subject, object, permission string) error {
	sql := "delete from `urbs_ac_acl` where subject = ? and object = ? and permission = ?"

	args := []interface{}{subject, object, permission}

	_, err := a.DB.DB().Exec(sql, args...)

	return err
}
