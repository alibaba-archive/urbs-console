package bll

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/schema/urbssetting"
	"github.com/teambition/urbs-console/src/util"
)

// User ...
type User struct {
}

// BatchAdd 批量添加用户
func (b *User) BatchAdd(ctx context.Context, users []string) error {
	url := fmt.Sprintf("%s/%s", conf.Config.UrbsSetting.Addr, "/v1/users:batch")
	body := new(urbssetting.UsersBody)
	body.Users = users

	err := util.RequestPost(ctx, url, nil, body, nil)
	if err != nil {
		return err
	}
	return nil
}
