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
	User   *User
	Group  *Group
}

func newAPIs(blls *bll.Blls) *APIs {
	return &APIs{
		Canary: &Canary{blls: blls},
		User:   &User{blls: blls},
		Group:  &Group{blls: blls},
	}
}

func newRouters(apis *APIs) []*gear.Router {

	routerV1 := gear.NewRouter(gear.RouterOptions{
		Root: "/v1",
	})
	routerV1.Use(tracing.New())

	routerV1.Get("/canary", apis.Canary.Get)

	// ***** user ******
	// 批量添加用户
	routerV1.Post("/users:batch", apis.User.BatchAdd)
	// ***** group ******
	// 批量添加群组
	routerV1.Post("/groups:batch", apis.Group.BatchAdd)

	return []*gear.Router{routerV1}
}
