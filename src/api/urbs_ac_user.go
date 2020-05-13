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

// List 获取用户列表
func (a *UrbsAcUser) List(ctx *gear.Context) error {
	req := new(tpl.Pagination)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.UrbsAcUser.List(ctx, req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Search 搜索用户
func (a *UrbsAcUser) Search(ctx *gear.Context) error {
	args := new(tpl.UrbsAcUserUrl)
	if err := ctx.ParseURL(args); err != nil {
		return err
	}
	res, err := a.blls.UrbsAcUser.Search(ctx, args.Key)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Delete ...
func (a *UrbsAcUser) Delete(ctx *gear.Context) error {
	args := tpl.UrbsAcUserUidUrl{}
	if err := ctx.ParseURL(&args); err != nil {
		return err
	}
	err := a.blls.UrbsAcUser.Delete(ctx, args.Uid)
	if err != nil {
		return err
	}
	return ctx.OkJSON(struct{}{})
}
