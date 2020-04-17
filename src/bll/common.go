package bll

import (
	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(NewBlls)
}

var (
	blls *Blls
)

// Blls ...
type Blls struct {
	User    *User
	Group   *Group
	Product *Product

	Label   *Label
	Module  *Module
	Setting *Setting

	OperationLog *OperationLog
	UrbsAcAcl    *UrbsAcAcl
}

// NewBlls ...
func NewBlls(services *service.Services, daos *dao.Daos) *Blls {
	blls = &Blls{
		User:    &User{services: services},
		Group:   &Group{services: services},
		Product: &Product{services: services},

		Label:   &Label{services: services},
		Module:  &Module{services: services},
		Setting: &Setting{services: services},

		OperationLog: &OperationLog{daos: daos},
	}
	return blls
}