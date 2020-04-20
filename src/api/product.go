package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/constant"
	"github.com/teambition/urbs-console/src/tpl"
)

// Product ..
type Product struct {
	blls *bll.Blls
}

// List ..
func (a *Product) List(ctx *gear.Context) error {
	req := new(tpl.Pagination)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.Product.List(ctx, req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// Create ..
func (a *Product) Create(ctx *gear.Context) error {
	body := new(tpl.NameDescBody)
	if err := ctx.ParseBody(body); err != nil {
		return err
	}
	res, err := a.blls.Product.Create(ctx, body)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// Update ..
func (a *Product) Update(ctx *gear.Context) error {
	req := tpl.ProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := new(tpl.ProductUpdateBody)
	if err := ctx.ParseBody(body); err != nil {
		return err
	}

	err := a.blls.UrbsAcAcl.Check(ctx, req.Product, constant.PermissionAll)
	if err != nil {
		return err
	}

	res, err := a.blls.Product.Update(ctx, req.Product, body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Offline ..
func (a *Product) Offline(ctx *gear.Context) error {
	req := tpl.ProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	err := a.blls.UrbsAcAcl.Check(ctx, req.Product, constant.PermissionAll)
	if err != nil {
		return err
	}

	res, err := a.blls.Product.Offline(ctx, req.Product)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Delete ..
func (a *Product) Delete(ctx *gear.Context) error {
	req := tpl.ProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Product.Delete(ctx, req.Product)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
