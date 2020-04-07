package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/teambition/gear/middleware/cors"

	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/requestid"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/middleware"
	"github.com/teambition/urbs-console/src/util"
)

// NewApp ...
func NewApp() *gear.App {
	app := gear.New()

	app.Set(gear.SetTrustedProxy, true)
	app.Set(gear.SetBodyParser, gear.DefaultBodyParser(2<<22)) // 8MB
	// ignore TLS handshake error
	app.Set(gear.SetLogger, log.New(gear.DefaultFilterWriter(), "", 0))

	app.Set(gear.SetParseError, func(err error) gear.HTTPError {
		msg := err.Error()

		if gorm.IsRecordNotFoundError(err) {
			return gear.ErrNotFound.WithMsg(msg)
		}
		if strings.Contains(msg, "Error 1062: Duplicate") {
			return gear.ErrConflict.WithMsg(msg)
		}

		return gear.ParseError(err)
	})

	// used for health check, so ingore logger
	app.Use(func(ctx *gear.Context) error {
		if ctx.Path == "/" || ctx.Path == "/version" {
			ctx.Set(gear.HeaderContentType, gear.MIMEApplicationJSONCharsetUTF8)
			return ctx.End(http.StatusOK, util.GetVersion())
		}
		return nil
	})
	app.Use(middleware.Logger)
	app.Use(requestid.New())

	app.Use(cors.New(cors.Options{
		AllowOrigins:  conf.Config.CorsWhiteList,
		Credentials:   true,
		ExposeHeaders: []string{gear.HeaderXRequestID},
	}))

	err := util.DigInvoke(func(routers []*gear.Router) error {
		for _, router := range routers {
			app.UseHandler(router)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return app
}
