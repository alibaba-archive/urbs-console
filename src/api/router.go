package api

import (
	"github.com/teambition/gear"
	tracing "github.com/teambition/gear-tracing"

	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(newAPIs)
	util.DigProvide(newRouters)
}

// APIs ..
type APIs struct {
	Canary *Canary
}

func newAPIs(blls *bll.Blls) *APIs {
	return &APIs{
		Canary: &Canary{blls: blls},
	}
}

func newRouters(apis *APIs) []*gear.Router {

	routerV1 := gear.NewRouter(gear.RouterOptions{
		Root: "/v1",
	})
	routerV1.Use(tracing.New())

	routerV1.Get("/canary", apis.Canary.Get)

	return []*gear.Router{routerV1}
}
