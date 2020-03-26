package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Group ...
type Group struct {
	services *service.Services
}

// BatchAdd ...
func (a *Group) BatchAdd(ctx context.Context, groups []*tpl.GroupBody) error {
	gb := make([]*urbssetting.GroupBody, len(groups))
	for i, g := range groups {
		gb[i] = &urbssetting.GroupBody{
			UID:  g.UID,
			Kind: g.Kind,
			Desc: g.Desc,
		}
	}
	_, err := a.services.UrbsSetting.GroupBatchAdd(ctx, gb)
	if err != nil {
		return nil
	}
	for _, g := range gb {
		users := []string{}
		_, err := a.services.UrbsSetting.GroupBatchAddMembers(ctx, g.UID, users)
		if err != nil {
			logger.Err(ctx, err.Error())
		}
	}
	return nil
}
