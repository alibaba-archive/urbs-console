package service

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dto/thrid"
)

// UserAuth ...
type UserAuth interface {
	Verify(ctx *gear.Context, body *thrid.UserVerifyReq) error
}
