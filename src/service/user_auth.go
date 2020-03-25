package service

import "github.com/teambition/gear"

// UserAuth ...
type UserAuth interface {
	Verify(ctx *gear.Context) error
}
