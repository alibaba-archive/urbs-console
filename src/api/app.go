package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/teambition/gear/middleware/cors"
	"github.com/teambition/gear/middleware/static"

	"github.com/teambition/gear"
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
	staticServer := static.New(static.Options{Root: util.GetStaticFilePath()})

	app.Use(func(ctx *gear.Context) error {
		if ctx.Path == "/version" { // used for health check, so ingore logger
			ctx.Set(gear.HeaderContentType, gear.MIMEApplicationJSONCharsetUTF8)
			return ctx.End(http.StatusOK, util.GetVersion())
		}
		if !strings.HasPrefix(ctx.Path, "/api") && !strings.HasPrefix(ctx.Path, "/v1") {
			if !strings.HasPrefix(ctx.Path, "/umi.css") && !strings.HasPrefix(ctx.Path, "/umi.js") {
				ctx.Path = "/"
			}
			return staticServer(ctx)
		}
		return nil
	})
	app.Use(middleware.Logger)

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
