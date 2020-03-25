package service

import (
	"net/http"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/gear-auth/jwt"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/util"
)

// UserAuthThrid ...
type UserAuthThrid struct {
}

// Verify ...
func (a *UserAuthThrid) Verify(ctx *gear.Context) error {
	j := jwt.New(conf.Config.UserAuth.Keys)
	token, err := j.Sign(conf.Config.UserAuth.UserAuthThrid.TokenKV, time.Hour)
	if err != nil {
		return err
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+token)

	body := make(map[string]interface{})
	for k, v := range conf.Config.UserAuth.UserAuthThrid.BodyKK {
		if conf.Config.UserAuth.UserAuthThrid.From == "header" {
			body[k] = ctx.GetHeader(v)
		} else {
			body[k], _ = ctx.Cookies.Get(v)
		}
	}
	err = util.RequestPost(ctx, conf.Config.UserAuth.UserAuthThrid.URL, header, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
