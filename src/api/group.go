package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// Group ..
type Group struct {
	blls *bll.Blls
}

// List ..
func (a *Group) List(ctx *gear.Context) error {
	req := tpl.GroupsURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.Group.List(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListLables ..
func (a *Group) ListLables(ctx *gear.Context) error {
	req := tpl.UIDPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.ListLables(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListMembers ..
func (a *Group) ListMembers(ctx *gear.Context) error {
	req := tpl.UIDPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.Group.ListMembers(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettings ..
func (a *Group) ListSettings(ctx *gear.Context) error {
	req := tpl.UIDPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.ListSettings(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// CheckExists ..
func (a *Group) CheckExists(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	if err != nil {
		return err
	}
	res, err := a.blls.Group.CheckExists(ctx, req.UID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// BatchAdd ..
func (a *Group) BatchAdd(ctx *gear.Context) error {
	req := tpl.GroupsBody{}
	if err := ctx.ParseBody(&req); err != nil {
		return err
	}
	if err := a.blls.Group.BatchAdd(ctx, req.Groups); err != nil {
		return err
	}

	return ctx.OkJSON(tpl.BoolRes{Result: true})
}

// Update ..
func (a *Group) Update(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.GroupUpdateBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}

	res, err := a.blls.Group.Update(ctx, req.UID, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Delete ..
func (a *Group) Delete(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.Delete(ctx, req.UID)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// BatchAddMembers ..
func (a *Group) BatchAddMembers(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.UsersBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	res, err := a.blls.Group.BatchAddMembers(ctx, req.UID, body.Users)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// RemoveMembers ..
func (a *Group) RemoveMembers(ctx *gear.Context) error {
	req := tpl.GroupMembersURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.RemoveMembers(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// RemoveLable ..
func (a *Group) RemoveLable(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.RemoveLable(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RollbackSetting 回退当前设置值到上一个值
// 更新值请用 POST /products/:product/modules/:module/settings/:setting+:assign 接口
func (a *Group) RollbackSetting(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.RollbackSetting(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RemoveSetting ..
func (a *Group) RemoveSetting(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Group.RemoveSetting(ctx, &req)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
