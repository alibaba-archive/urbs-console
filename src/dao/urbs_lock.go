package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/schema"
)

// UrbsLock table `urbs_lock`
type UrbsLock struct {
	DB *gorm.DB
}

// Lock ...
func (a *UrbsLock) Lock(ctx context.Context, key string, expire time.Duration) error {
	now := time.Now().UTC()
	lock := &schema.UrbsLock{Name: key, ExpireAt: now.Add(expire)}
	err := a.DB.Create(lock).Error
	if err != nil {
		l := &schema.UrbsLock{}
		e := a.DB.Where("`name` = ?", key).First(l).Error
		if e == nil {
			if l.ExpireAt.Before(now) {
				a.Unlock(ctx, key) // 释放失效、异常的锁
				err = a.DB.Create(lock).Error
			} else {
				lock = l
			}
		}
	}
	if err != nil {
		err = fmt.Errorf("%s locked, should expire at: %v, error: %s", key, lock.ExpireAt, err.Error())
	}
	return err
}

// Unlock ...
func (a *UrbsLock) Unlock(ctx context.Context, key string) {
	err := a.DB.Where("`name` = ?", key).Delete(&schema.UrbsLock{}).Error
	if err != nil {
		logger.Default.Errf("unlock: key %s, error %v", key, err)
	}
}
