package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
)

// User ..
type User struct {
	blls *bll.Blls
}

// ListLables 返回 user 的 labels，按照 label 指派时间正序，支持分页
func (a *User) ListLables(ctx *gear.Context) error {
	req := new(urbssetting.UIDPaginationURL)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}

	res, err := a.blls.User.ListLables(ctx, req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// BatchAdd ..
func (a *User) BatchAdd(ctx *gear.Context) error {
	req := tpl.UsersBody{}
	if err := ctx.ParseBody(&req); err != nil {
		return err
	}

	if err := a.blls.User.BatchAdd(ctx, req.Users); err != nil {
		return err
	}

	return ctx.OkJSON(tpl.BoolRes{Result: true})
}
