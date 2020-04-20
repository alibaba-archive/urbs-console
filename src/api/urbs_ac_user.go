package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsAcUser ..
type UrbsAcUser struct {
	blls *bll.Blls
}

// Add 添加权限
func (a *UrbsAcUser) Add(ctx *gear.Context) error {
	body := tpl.UrbsAcUsersBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcUser.Add(ctx, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(struct{}{})
}
