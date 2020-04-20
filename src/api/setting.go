package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// Setting ..
type Setting struct {
	blls *bll.Blls
}

// List ..
func (a *Setting) List(ctx *gear.Context) error {
	req := tpl.ProductModuleURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.List(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Create ..
func (a *Setting) Create(ctx *gear.Context) error {
	req := tpl.ProductModuleURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.NameDescBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.Create(ctx, &req, &body)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// Update ..
func (a *Setting) Update(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.SettingUpdateBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.Update(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Offline ..
func (a *Setting) Offline(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.Offline(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Assign ..
func (a *Setting) Assign(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.UsersGroupsBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module)
	if err != nil {
		return err
	}

	res, err := a.blls.Setting.Assign(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Recall ..
func (a *Setting) Recall(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	err = a.blls.Setting.Recall(ctx, req.Product, req.Module, req.Setting)
	if err != nil {
		return err
	}
	return ctx.OkJSON(tpl.BoolRes{Result: true})
}

// Logs 返回操作日志列表
func (a *Setting) Logs(ctx *gear.Context) error {
	req := &tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.OperationLog.List(ctx, req.Product+req.Module+req.Setting, &req.Pagination)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
