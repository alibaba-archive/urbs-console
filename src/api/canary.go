package api

import (
	"github.com/go-http-utils/cookie"
	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/bll"
)

// Canary ...
type Canary struct {
	blls *bll.Blls
}

// Get ..
func (a *Canary) Get(ctx *gear.Context) error {
	label := ctx.Query("label")
	if label == "" {
		label = "beta"
	}
	option := &cookie.Options{
		Domain:   ctx.Query("domain"),
		Path:     "/",
		HTTPOnly: true,
	}
	ctx.Cookies.Set("X-Canary", "label="+label, option)
	ctx.Cookies.Set("tb_gateway", "traefik", option)
	return ctx.OkJSON(struct{}{})
}
