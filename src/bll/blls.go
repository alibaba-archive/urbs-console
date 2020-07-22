package bll

import (
	"context"
	"time"

	"github.com/mushroomsir/tcc"
	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(NewBlls)
}

// Blls ...
type Blls struct {
	services *service.Services
	User     *User
	Group    *Group
	Product  *Product

	Label   *Label
	Module  *Module
	Setting *Setting

	OperationLog *OperationLog
	UrbsAcAcl    *UrbsAcAcl
	UrbsAcUser   *UrbsAcUser
}

// NewBlls ...
func NewBlls(services *service.Services, d *dao.Daos) *Blls {
	blls := &Blls{
		services:     services,
		UrbsAcAcl:    &UrbsAcAcl{daos: d},
		OperationLog: &OperationLog{daos: d},
		UrbsAcUser:   &UrbsAcUser{daos: d},

		User:  &User{services: services},
		Group: &Group{services: services, daos: d},
	}
	blls.Product = &Product{
		services:  services,
		daos:      d,
		urbsAcAcl: blls.UrbsAcAcl,
	}
	blls.Module = &Module{
		services:  services,
		urbsAcAcl: blls.UrbsAcAcl,
	}
	blls.Label = &Label{
		services:     services,
		daos:         d,
		operationLog: blls.OperationLog,
		urbsAcAcl:    blls.UrbsAcAcl,
		group:        blls.Group,
	}
	blls.Setting = &Setting{
		services:     services,
		daos:         d,
		operationLog: blls.OperationLog,
		urbsAcAcl:    blls.UrbsAcAcl,
		group:        blls.Group,
	}
	services.TCC.SetTryHandler(blls.Handler)
	services.TCC.SetConfirmHandler(blls.Handler)
	return blls
}

var (
	// TccSettingRecall ...
	TccSettingRecall = "setting.recall"
)

// Handler ...
func (a *Blls) Handler(task *tcc.Task) {
	if time.Now().Sub(task.CreatedAt) > time.Hour {
		err := task.Cancel()
		logger.Default.Warning("tccSuccess", task.Value, "cancel", err)
		return
	}
	if task.Name == TccSettingRecall {
		req := &settingRecallReq{}
		err := task.JSONToObj(req)
		if err != nil {
			logger.Default.Err("tccJsonToObj", err.Error())
			return
		}
		_, err = a.services.UrbsSetting.SettingRecall(context.Background(), req.Args, req.Body)
		if err != nil {
			logger.Default.Err("settingRecall", task.Value)
			return
		}
		task.Cancel()
		logger.Default.Info("tccSuccess", task.Value, "cancel", err)
	}
}
