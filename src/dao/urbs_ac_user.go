package dao

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcUser table `urbs_ac_user`
type UrbsAcUser struct {
	DB *gorm.DB
}

// Add ...
func (a *UrbsAcUser) Add(ctx context.Context, obj *schema.UrbsAcUser) error {
	obj.CreatedAt = time.Now().UTC()

	sql := "insert ignore into `urbs_ac_user` (`created_at`, `uid`, `name`) values (?, ?, ?)"

	args := []interface{}{obj.CreatedAt, obj.UID, obj.Name}

	_, err := a.DB.DB().Exec(sql, args...)

	return err
}

// FindByUID ...
func (a *UrbsAcUser) FindByUID(ctx context.Context, uid string) (*schema.UrbsAcUser, error) {
	urbsAcUser := &schema.UrbsAcUser{}

	where := "uid = ?"

	err := a.DB.Where(where, uid).Find(&urbsAcUser).Error
	if err != nil {
		return nil, err
	}
	return urbsAcUser, nil
}

// FindByUIDS ...
func (a *UrbsAcUser) FindByUIDS(ctx context.Context, uids []string) ([]*schema.UrbsAcUser, error) {
	urbsAcUsers := make([]*schema.UrbsAcUser, 0)

	where := "uid in ( ? )"

	err := a.DB.Where(where, uids).Find(&urbsAcUsers).Error
	if err != nil {
		return nil, err
	}
	return urbsAcUsers, nil
}

// RemoveByUID ...
func (a *UrbsAcUser) RemoveByUID(ctx context.Context, uid string) error {
	sql := "delete from `urbs_ac_user` where uid = ?"

	_, err := a.DB.DB().Exec(sql, uid)

	return err
}
