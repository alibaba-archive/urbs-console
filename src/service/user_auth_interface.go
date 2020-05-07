package service

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/thrid"
)

// UserAuthInterface ...
type UserAuthInterface interface {
	Verify(ctx context.Context, body *thrid.UserVerifyReq) (string, error)
}
