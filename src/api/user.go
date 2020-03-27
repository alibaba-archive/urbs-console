package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// User ..
type User struct {
	blls *bll.Blls
}

// RefreshCachedLables 强制更新 user 的 labels 缓存
func (a *User) RefreshCachedLables(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.User.RefreshCachedLables(ctx, req.UID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// ListLables 返回 user 的 labels，按照 label 指派时间正序，支持分页
func (a *User) ListLables(ctx *gear.Context) error {
	req := new(tpl.UIDPaginationURL)
	if err := ctx.ParseURL(req); err != nil {
		return err
	}

	res, err := a.blls.User.ListLables(ctx, req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettings 返回 user 的 settings，按照 setting 设置时间正序，支持分页
func (a *User) ListSettings(ctx *gear.Context) error {
	req := tpl.UIDProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.User.ListSettings(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// ListSettingsUnionAll 返回 user 的 settings，按照 setting 设置时间反序，支持分页
// 包含了 user 从属的 group 的 settings
func (a *User) ListSettingsUnionAll(ctx *gear.Context) error {
	req := tpl.MySettingsQueryURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.User.ListSettingsUnionAll(ctx, &req)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// CheckExists ..
func (a *User) CheckExists(ctx *gear.Context) error {
	req := tpl.UIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	res, err := a.blls.User.CheckExists(ctx, req.UID)
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

	res, err := a.blls.User.BatchAdd(ctx, req.Users)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// RemoveLable ..
func (a *User) RemoveLable(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.RemoveLable(ctx, req.UID, req.HID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RollbackSetting 回退当前设置值到上一个值
// 更新值请用 POST /products/:product/modules/:module/settings/:setting+:assign 接口
func (a *User) RollbackSetting(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.RollbackSetting(ctx, req.UID, req.HID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// RemoveSetting ..
func (a *User) RemoveSetting(ctx *gear.Context) error {
	req := tpl.UIDHIDURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.User.RemoveSetting(ctx, req.UID, req.HID)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
