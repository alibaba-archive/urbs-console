package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Module ...
type Module struct {
	services *service.Services
}

// List 读取指定产品的功能模块
func (a *Module) List(ctx context.Context, args *tpl.ProductPaginationURL) (*urbssetting.ModulesRes, error) {
	return a.services.UrbsSetting.ModuleList(ctx, args)
}

// Create 指定产品创建功能模块
func (a *Module) Create(ctx context.Context, product string, body *tpl.NameDescBody) (*urbssetting.ModuleRes, error) {
	return a.services.UrbsSetting.ModuleCreate(ctx, product, body)
}

// Update 更新指定产品功能模块
func (a *Module) Update(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*urbssetting.ModuleRes, error) {
	return a.services.UrbsSetting.ModuleUpdate(ctx, product, module, body)
}

// Offline 下线指定产品功能模块
func (a *Module) Offline(ctx context.Context, product string, module string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.ModuleOffline(ctx, product, module)
}
