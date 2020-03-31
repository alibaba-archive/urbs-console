package bll

import (
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(NewBlls)
}

// Blls ...
type Blls struct {
	User    *User
	Group   *Group
	Product *Product

	Label   *Label
	Module  *Module
	Setting *Setting
}

// NewBlls ...
func NewBlls(services *service.Services) *Blls {
	return &Blls{
		User:    &User{services: services},
		Group:   &Group{services: services},
		Product: &Product{services: services},

		Label:   &Label{services: services},
		Module:  &Module{services: services},
		Setting: &Setting{services: services},
	}
}
