package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsAcAcl ..
type UrbsAcAcl struct {
	blls *bll.Blls
}

// Add ...
func (a *UrbsAcAcl) Add(ctx *gear.Context) error {
	req := tpl.UrbsAcAclURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	body := tpl.UrbsAcAclAddBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}
	err := a.blls.UrbsAcAcl.AddByReq(ctx, &req, &body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(struct{}{})
}

// Check ...
func (a *UrbsAcAcl) Check(ctx *gear.Context) error {
	body := tpl.UrbsAcAclCheckBody{}
	if err := ctx.ParseURL(&body); err != nil {
		return err
	}
	object := body.Product + body.Label + body.Module + body.Setting
	var err error
	if object == "" {
		err = a.blls.UrbsAcAcl.CheckSuperAdmin(ctx)
	} else {
		err = a.blls.UrbsAcAcl.CheckAdmin(ctx, object)
	}
	return ctx.OkJSON(&tpl.BoolRes{Result: err == nil})
}
