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

// ListByProduct ..
func (a *Setting) ListByProduct(ctx *gear.Context) error {
	req := tpl.ProductPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.ListByProduct(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
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

// Get ..
func (a *Setting) Get(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.Get(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListGroups ..
func (a *Setting) ListGroups(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.ListGroups(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RollbackGroupSetting ..
func (a *Setting) RollbackGroupSetting(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.RollbackGroupSetting(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteGroup ..
func (a *Setting) DeleteGroup(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.DeleteGroup(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListUsers ..
func (a *Setting) ListUsers(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.ListUsers(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RollbackUserSetting ..
func (a *Setting) RollbackUserSetting(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.RollbackUserSetting(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteUser ..
func (a *Setting) DeleteUser(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.DeleteUser(ctx, &req)
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

	body := tpl.SettingBody{}
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
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
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
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
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
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
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
	body := tpl.RecallBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.Recall(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
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

// ListRules ..
func (a *Setting) ListRules(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Setting.ListRules(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// CreateRule ..
func (a *Setting) CreateRule(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.SettingRuleBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.CreateRule(ctx, &req, &body)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// UpdateRule ..
func (a *Setting) UpdateRule(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.SettingRuleBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.UpdateRule(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteRule ..
func (a *Setting) DeleteRule(ctx *gear.Context) error {
	req := tpl.ProductModuleSettingHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Module+req.Setting)
	if err != nil {
		return err
	}
	res, err := a.blls.Setting.DeleteRule(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
