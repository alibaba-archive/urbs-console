package middleware

import (
	"context"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

// Verify ...
func Verify(services *service.Services) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		body := &thrid.UserVerifyReq{}
		body.Cookie, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey)
		body.Singed, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey + ".sig")
		body.Token = util.TokenExtractor(ctx)

		uid, err := services.UserAuth.Verify(ctx, body)
		if err != nil {
			return gear.ErrUnauthorized.WithMsg(err.Error())
		}
		_ctx := context.WithValue(ctx.Context(), util.UidKey{}, uid)
		ctx.WithContext(_ctx)
		return nil
	}
}

// CheckSuperAdmin ...
func CheckSuperAdmin(blls *bll.Blls) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		err := blls.UrbsAcAcl.CheckSuperAdmin(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

// CheckViewer ...
func CheckViewer(blls *bll.Blls) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		err := blls.UrbsAcAcl.CheckViewer(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
