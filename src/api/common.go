package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/middleware"
)

//GetUid ...
func GetUid(ctx *gear.Context) string {
	uid, _ := ctx.Value(middleware.UidKey{}).(string)
	return uid
}
