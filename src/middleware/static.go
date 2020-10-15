package middleware

import (
	"strings"

	"github.com/teambition/gear"
)

// StaticFile ...
func StaticFile(staticServer gear.Middleware) func(ctx *gear.Context) error {
	return func(ctx *gear.Context) error {
		if staticServer == nil {
			return nil
		}
		if !strings.HasPrefix(ctx.Path, "/api") && !strings.HasPrefix(ctx.Path, "/v1") && !strings.HasPrefix(ctx.Path, "/inner") {
			if !strings.HasPrefix(ctx.Path, "/umi.css") && !strings.HasPrefix(ctx.Path, "/umi.js") {
				ctx.Path = "/"
			}
			return staticServer(ctx)
		}
		return nil
	}
}
