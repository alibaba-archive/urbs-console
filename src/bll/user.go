package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/service"
)

// User ...
type User struct {
	services *service.Services
}

// BatchAdd 批量添加用户
func (a *User) BatchAdd(ctx context.Context, users []string) error {
	_, err := a.services.UrbsSetting.UserBatchAdd(ctx, users)
	if err != nil {
		return err
	}
	return nil
}
