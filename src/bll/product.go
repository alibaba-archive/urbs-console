package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Product ...
type Product struct {
	services *service.Services
}

// Create 创建产品
func (a *Product) Create(ctx context.Context, req *tpl.NameDescBody) (*urbssetting.ProductRes, error) {
	return a.services.UrbsSetting.ProductCreate(ctx, req)
}

// List 返回产品列表
func (a *Product) List(ctx context.Context, req *tpl.Pagination) (*urbssetting.ProductsRes, error) {
	return a.services.UrbsSetting.ProductList(ctx, req)
}

// Update ...
func (a *Product) Update(ctx context.Context, productName string, req *tpl.ProductUpdateBody) (*urbssetting.ProductRes, error) {
	return a.services.UrbsSetting.ProductUpdate(ctx, productName, req)
}

// Offline 下线产品
func (a *Product) Offline(ctx context.Context, productName string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.ProductOffline(ctx, productName)
}

// Delete 逻辑删除产品
func (a *Product) Delete(ctx context.Context, productName string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.ProductDelete(ctx, productName)
}
