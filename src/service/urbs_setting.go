package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/util"
)

// UrbsSetting ...
type UrbsSetting struct {
}

// UserBatchAdd ...
func (a *UrbsSetting) UserBatchAdd(ctx context.Context, users []string) (*urbssetting.BoolRes, error) {
	url := fmt.Sprintf("%s/%s", conf.Config.UrbsSetting.Addr, "/v1/users:batch")

	body := new(urbssetting.UsersBody)
	body.Users = users

	resp := new(urbssetting.BoolRes)

	err := util.RequestPost(ctx, url, UrbsSettingHeader(), body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
