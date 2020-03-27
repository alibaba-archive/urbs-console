package middleware

import (
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

var defaultLogger = logging.Default()

// Logger logger middleware
func Logger(ctx *gear.Context) error {
	err := defaultLogger.Serve(ctx)

	ctx.OnEnd(func() {
		statusCode := ctx.Res.Status()
		if statusCode >= 300 {
			resBody := string(ctx.Res.Body())
			defaultLogger.SetTo(ctx, "resBody", resBody)
		}
	})
	return err
}
