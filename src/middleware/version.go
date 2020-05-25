package middleware

import (
	"net/http"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/util"
)

// Version ...
func Version(ctx *gear.Context) error {
	if ctx.Path == "/version" { // used for health check, so ingore logger
		ctx.Set(gear.HeaderContentType, gear.MIMEApplicationJSONCharsetUTF8)
		return ctx.End(http.StatusOK, util.GetVersion())
	}
	return nil
}
