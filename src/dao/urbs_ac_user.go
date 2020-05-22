package dao

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
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
	sql := "select * from urbs_ac_user where uid = ?"

	res := &schema.UrbsAcUser{}

	err := a.DB.Raw(sql, uid).Scan(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindByUIDs ...
func (a *UrbsAcUser) FindByUIDs(ctx context.Context, uids []string) ([]schema.UrbsAcUser, error) {
	sql := "select * from urbs_ac_user where uid in ( ? )"

	res := []schema.UrbsAcUser{}

	err := a.DB.Where(sql, uids).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteByUID ...
func (a *UrbsAcUser) DeleteByUID(ctx context.Context, uid string) error {
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		sql := "delete from `urbs_ac_user` where uid = ?"
		err := tx.Exec(sql, uid).Error
		if err != nil {
			return err
		}
		sql = "delete from `urbs_ac_acl` where subject = ?"
		return tx.Exec(sql, uid).Error
	})
	return err
}

// UpdateByUID ...
func (a *UrbsAcUser) UpdateByUID(ctx context.Context, name, uid string) error {

	sql := "update `urbs_ac_user` set name=? where uid = ?"

	return a.DB.Exec(sql, name, uid).Error
}

// List ...
func (a *UrbsAcUser) List(ctx context.Context, pg *tpl.Pagination) ([]*schema.UrbsAcUser, error) {
	sql := "select * from urbs_ac_user order by id desc limit ?,?"

	urbsAcUsers := make([]*schema.UrbsAcUser, 0)
	err := a.DB.Raw(sql, pg.Skip, pg.PageSize+1).Scan(&urbsAcUsers).Error
	if err != nil {
		return nil, err
	}
	return urbsAcUsers, nil
}

// Search ...
func (a *UrbsAcUser) Search(ctx context.Context, key string) ([]*schema.UrbsAcUser, error) {
	sql := "select * from urbs_ac_user where name like ? or uid like ? limit 10"

	urbsAcUsers := make([]*schema.UrbsAcUser, 0)
	err := a.DB.Raw(sql, "%"+key+"%", "%"+key+"%").Scan(&urbsAcUsers).Error
	if err != nil {
		return nil, err
	}
	return urbsAcUsers, nil
}

// Count 用户数量
func (a *UrbsAcUser) Count(ctx context.Context) (int, error) {
	sql := "select count(1) as count from urbs_ac_user"

	res := &schema.CountResult{}

	err := a.DB.Raw(sql).Scan(res).Error
	return res.Count, err
}
