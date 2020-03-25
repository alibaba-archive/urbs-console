package bll

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util"
)

// Group ...
type Group struct {
}

// BatchAdd ...
func (a *Group) BatchAdd(ctx context.Context, groups []tpl.GroupBody) error {
	url := fmt.Sprintf("/v1/groups:batch/%s", conf.Config.UrbsSetting.Addr)
	resp := new(tpl.BoolRes)
	err := util.RequestPost(ctx, url, nil, groups, resp)
	if err != nil {
		return err
	}
	for _, group := range groups {
		memberAPI := fmt.Sprintf("%s/v1/groups/%s/members:batch", conf.Config.UrbsSetting.Addr, group.UID)
		body := new(urbssetting.UsersBody)
		// TODO: 拉取成员
		err := util.RequestPost(ctx, memberAPI, nil, body, resp)
		if err != nil {
			logger.Err(ctx, err.Error())
		}
	}
	return nil
}
