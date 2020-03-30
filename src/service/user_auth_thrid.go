package service

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/thrid"
	"github.com/teambition/urbs-console/src/util/request"
)

// UserAuthThrid ...
type UserAuthThrid struct {
}

// Verify ...
func (a *UserAuthThrid) Verify(ctx *gear.Context, body *thrid.UserVerifyReq) error {
	resp, err := request.Post(conf.Config.Thrid.UserAuth.URL).Header(genThridHeader()).Body(body).Do()
	if err := HanderResponse(resp, err); err != nil {
		return err
	}
	return nil
}
