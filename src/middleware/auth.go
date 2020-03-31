package middleware

import (
	"context"
	"strings"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

// UidKey ...
type UidKey struct{}

// Auth ...
func Auth(services *service.Services, ignoreURLs []string, memberURLs []string) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		for _, u := range ignoreURLs {
			if strings.HasPrefix(ctx.Req.URL.Path, u) {
				return nil
			}
		}
		body := &thrid.UserVerifyReq{}
		body.Cookie, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey)
		body.Singed, _ = ctx.Cookies.Get(conf.Config.Thrid.UserAuth.CookieKey + ".sig")
		body.Token = util.TokenExtractor(ctx)
		body.Role = "admin"

		for _, u := range memberURLs {
			if strings.HasPrefix(ctx.Req.URL.Path, u) {
				body.Role = "member"
			}
		}
		uid, err := services.UserAuth.Verify(ctx, body)
		_ctx := context.WithValue(ctx.Context(), UidKey{}, uid)
		ctx.WithContext(_ctx)
		return err
	}
}
