package dao

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsAcUser table `urbs_ac_user`
type UrbsAcUser struct {
	DB *gorm.DB
}

// BatchAdd 批量添加用户
func (a *UrbsAcUser) BatchAdd(ctx context.Context, users []*schema.UrbsAcUser) error {
	if len(users) == 0 {
		return nil
	}
	var buf bytes.Buffer
	fmt.Fprint(&buf, "insert ignore into `urbs_ac_user` (`uid`,`name`) values")
	for _, user := range users {
		fmt.Fprintf(&buf, " ('%s' ,'%s'),", user.UID, user.Name)
	}
	b := buf.Bytes()
	b[len(b)-1] = ';'
	return a.DB.Exec(string(b)).Error
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
