package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/schema"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsAcUser ...
type UrbsAcUser struct {
	daos *dao.Daos
}

// Add ...
func (a *UrbsAcUser) Add(ctx context.Context, body *tpl.UrbsAcUsersBody) error {
	users := make([]*schema.UrbsAcUser, len(body.Users))
	for i, user := range body.Users {
		users[i] = &schema.UrbsAcUser{
			UID:  user.Uid,
			Name: user.Name,
		}
	}
	return a.daos.UrbsAcUser.BatchAdd(ctx, users)
}
