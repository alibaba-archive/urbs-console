package service

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dto/thrid"
)

// UserAuthInterface ...
type UserAuthInterface interface {
	Verify(ctx *gear.Context, body *thrid.UserVerifyReq) (string, error)
}
