package middleware

import (
	"context"
	"time"

	"github.com/teambition/gear"
	auth "github.com/teambition/gear-auth"
	authjwt "github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	keys := conf.Config.AuthKeys
	if len(keys) > 0 {
		Auther = auth.New(authjwt.StrToKeys(keys...)...)
		Auther.JWT().SetExpiresIn(time.Minute * 10)
	} else {
		logger.Default.Warningf("`auth_keys` is empty, Auth middleware will not be executed.")
	}
}

// Auther 是基于 JWT 的身份验证，当 config.auth_keys 配置了才会启用
var Auther *auth.Auth

// Verify ...
func Verify(services *service.Services) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		var uid string
		if Auther != nil {
			xToken := util.XAuthExtractor(ctx)
			if xToken != "" {
				claims, err := Auther.JWT().Verify(xToken)
				if err != nil {
					return gear.ErrUnauthorized.WithMsg(err.Error())
				}
				uid = claims.Get("uid").(string)
			} else {
				body := &thrid.UserVerifyReq{}
				body.Cookie, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey)
				body.Singed, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey + ".sig")
				body.Token = util.AuthorizationExtractor(ctx)
				if body.Cookie == "" && body.Token == "" {
					return gear.ErrUnauthorized.WithMsg("invalid authorization")
				}
				var err error
				uid, err = services.UserAuth.Verify(ctx, body)
				if err != nil {
					return gear.ErrUnauthorized.WithMsg(err.Error())
				}
			}
		} else if len(conf.Config.SuperAdmins) > 0 {
			uid = conf.Config.SuperAdmins[0]
		}
		if uid == "" {
			return gear.ErrUnauthorized.WithMsg("invalid uid")
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
