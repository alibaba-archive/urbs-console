package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// Label ..
type Label struct {
	blls *bll.Blls
}

// ListGroups ..
func (a *Label) ListGroups(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.ListGroups(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteGroup ..
func (a *Label) DeleteGroup(ctx *gear.Context) error {
	req := tpl.ProductLabelUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.DeleteGroup(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListUsers ..
func (a *Label) ListUsers(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.ListUsers(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteUser ..
func (a *Label) DeleteUser(ctx *gear.Context) error {
	req := tpl.ProductLabelUIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.DeleteUser(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Create ..
func (a *Label) Create(ctx *gear.Context) error {
	req := tpl.ProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := &tpl.LabelBody{}
	if err := ctx.ParseBody(body); err != nil {
		return err
	}

	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product)
	if err != nil {
		return err
	}

	res, err := a.blls.Label.Create(ctx, req.Product, body)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// List ..
func (a *Label) List(ctx *gear.Context) error {
	req := &tpl.ProductPaginationURL{}
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.Label.List(ctx, req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Update ..
func (a *Label) Update(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := &tpl.LabelUpdateBody{}
	if err := ctx.ParseBody(body); err != nil {
		return err
	}

	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product)
	if err != nil {
		return err
	}

	res, err := a.blls.Label.Update(ctx, req.Product, req.Label, body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Offline ..
func (a *Label) Offline(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.Offline(ctx, req.Product, req.Label)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Assign ..
func (a *Label) Assign(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.UsersGroupsBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	result, err := a.blls.Label.Assign(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(result)
}

// Delete ..
func (a *Label) Delete(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.Delete(ctx, req.Product, req.Label)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Recall ..
func (a *Label) Recall(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.RecallBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.Recall(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Logs 返回操作日志列表
func (a *Label) Logs(ctx *gear.Context) error {
	req := &tpl.LogProductLabelPaginationURL{}
	if err := ctx.ParseURL(req); err != nil {
		return err
	}
	res, err := a.blls.OperationLog.List(ctx, req.Product+req.Label, &req.ConsolePagination)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// CreateRule ..
func (a *Label) CreateRule(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.LabelRuleBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.CreateRule(ctx, &req, &body)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListRules ..
func (a *Label) ListRules(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.ListRules(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// UpdateRule ..
func (a *Label) UpdateRule(ctx *gear.Context) error {
	req := tpl.ProductLabelHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.LabelRuleBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.UpdateRule(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// DeleteRule ..
func (a *Label) DeleteRule(ctx *gear.Context) error {
	req := tpl.ProductLabelHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.DeleteRule(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// CleanUp ..
func (a *Label) CleanUp(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckAdmin(ctx, req.Product+req.Label)
	if err != nil {
		return err
	}
	res, err := a.blls.Label.CleanUp(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
