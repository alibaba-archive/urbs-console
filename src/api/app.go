package api

import (
	"log"
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
	app.Set(gear.SetBodyParser, gear.DefaultBodyParser(2<<22))          // 8MB
	app.Set(gear.SetLogger, log.New(gear.DefaultFilterWriter(), "", 0)) // ignore TLS handshake error

	app.Set(gear.SetParseError, parseError)

	app.Use(middleware.Version)

	var staticServer gear.Middleware = nil
	if path := util.GetStaticFilePath(); path != "" {
		staticServer = static.New(static.Options{Root: path})
	}
	app.Use(middleware.StaticFile(staticServer))

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

func parseError(err error) gear.HTTPError {
	msg := err.Error()

	if gorm.IsRecordNotFoundError(err) {
		return gear.ErrNotFound.WithMsg(msg)
	}
	if strings.Contains(msg, "Error 1062: Duplicate") {
		return gear.ErrConflict.WithMsg(msg)
	}

	return gear.ParseError(err)
}
