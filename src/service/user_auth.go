package service

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// UserAuth ...
type UserAuth struct {
}

// Verify ...
func (a *UserAuth) Verify(ctx *gear.Context, body *thrid.UserVerifyReq) (string, error) {
	if conf.Config.Thrid.UserAuth.URL == "" {
		logger.Warning(ctx, "`user_auth.url` is empty, verify user will not be executed.")
		return "", nil
	}
	result := new(tpl.StringRes)
	resp, err := request.Post(conf.Config.Thrid.UserAuth.URL).Header(genThridHeader()).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return "", err
	}
	return result.Result, nil
}
