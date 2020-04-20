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
	blls     *Blls
	daos     *dao.Daos
	services *service.Services
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
	UrbsAcUser   *UrbsAcUser
}

// NewBlls ...
func NewBlls(s *service.Services, d *dao.Daos) *Blls {
	daos = d
	services = s
	blls = &Blls{
		User:    &User{services: services},
		Group:   &Group{services: services},
		Product: &Product{services: services},

		Label:   &Label{services: services},
		Module:  &Module{services: services},
		Setting: &Setting{services: services},

		UrbsAcAcl:    &UrbsAcAcl{daos: d},
		OperationLog: &OperationLog{daos: d},
		UrbsAcUser:   &UrbsAcUser{daos: d},
	}
	return blls
}
