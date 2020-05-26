package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Blls ...
type Blls struct {
	services *service.Services

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
func NewBlls(services *service.Services, d *dao.Daos) *Blls {
	blls = &Blls{
		services: services,

		User:    &User{services: services},
		Group:   &Group{services: services, daos: d},
		Product: &Product{services: services},

		Label:   &Label{services: services, daos: d},
		Module:  &Module{services: services},
		Setting: &Setting{services: services, daos: d},

		UrbsAcAcl:    &UrbsAcAcl{daos: d},
		OperationLog: &OperationLog{daos: d},
		UrbsAcUser:   &UrbsAcUser{daos: d},
	}
	return blls
}

// AddUserAndOrg ...
func (a *Blls) AddUserAndOrg(ctx context.Context, users []string, groups []string) {
	if len(users) > 0 {
		_, err := a.services.UrbsSetting.UserBatchAdd(ctx, users)
		if err != nil {
			logger.Err(ctx, "userBatchAdd", "error", err.Error())
		} else {
			logger.Info(ctx, "userBatchAdd", "users", users)
		}
	}
	if len(groups) > 0 {
		groupBody := []tpl.GroupBody{}
		for _, g := range groups {
			groupBody = append(groupBody, tpl.GroupBody{
				UID:  g,
				Kind: "organization",
			})
		}
		err := blls.Group.BatchAdd(ctx, groupBody)
		if err != nil {
			logger.Err(ctx, "groupBatchAdd", "error", err.Error())
		} else {
			logger.Info(ctx, "groupBatchAdd", "groups", groupBody)
		}
	}
}
