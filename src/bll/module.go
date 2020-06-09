package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Module ...
type Module struct {
	services *service.Services

	urbsAcAcl *UrbsAcAcl
}

// List 读取指定产品的功能模块
func (a *Module) List(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.ModulesInfoRes, error) {
	ress, err := a.services.UrbsSetting.ModuleList(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(ress.Result))
	for i, module := range ress.Result {
		objects[i] = args.Product + module.Name
	}
	subjects, err := a.urbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, module := range ress.Result {
		module.Users = subjects[args.Product+module.Name]
	}
	return ress, nil
}

// Create 指定产品创建功能模块
func (a *Module) Create(ctx context.Context, product string, body *tpl.NameDescBody) (*tpl.ModuleInfoRes, error) {
	err := a.urbsAcAcl.AddDefaultPermission(ctx, body.Uids, product+body.Name)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.ModuleCreate(ctx, product, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, product+body.Name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update 更新指定产品功能模块
func (a *Module) Update(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*tpl.ModuleInfoRes, error) {
	err := a.urbsAcAcl.Update(ctx, body.UidsBody, product+module)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.ModuleUpdate(ctx, product, module, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, product+module)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Offline 下线指定产品功能模块
func (a *Module) Offline(ctx context.Context, product string, module string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.ModuleOffline(ctx, product, module)
}
