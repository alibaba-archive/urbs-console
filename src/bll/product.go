package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dao"
	"github.com/teambition/urbs-console/src/logger"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Product ...
type Product struct {
	services  *service.Services
	daos      *dao.Daos
	urbsAcAcl *UrbsAcAcl
}

// Create 创建产品
func (a *Product) Create(ctx context.Context, args *tpl.NameDescBody) (*tpl.ProductRes, error) {
	err := a.urbsAcAcl.AddDefaultPermission(ctx, args.Uids, args.Name)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.ProductCreate(ctx, args)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, args.Name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// List 返回产品列表
func (a *Product) List(ctx context.Context, args *tpl.Pagination) (*tpl.ProductsRes, error) {
	products, err := a.services.UrbsSetting.ProductList(ctx, args)
	if err != nil {
		return nil, err
	}
	objects := make([]string, len(products.Result))
	for i, product := range products.Result {
		objects[i] = product.Name
	}
	subjects, err := a.urbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, product := range products.Result {
		product.Users = subjects[product.Name]
	}
	return products, nil
}

// Update ...
func (a *Product) Update(ctx context.Context, product string, body *tpl.ProductUpdateBody) (*tpl.ProductRes, error) {
	err := a.urbsAcAcl.Update(ctx, body.UidsBody, product)
	if err != nil {
		return nil, err
	}
	res, err := a.services.UrbsSetting.ProductUpdate(ctx, product, body)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = a.urbsAcAcl.FindUsersByObject(ctx, product)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Offline 下线产品
func (a *Product) Offline(ctx context.Context, product string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.ProductOffline(ctx, product)
}

// Delete 逻辑删除产品
func (a *Product) Delete(ctx context.Context, product string) (*tpl.BoolRes, error) {
	res, err := a.services.UrbsSetting.ProductDelete(ctx, product)
	if err != nil {
		return nil, err
	}
	err = a.daos.UrbsAcAcl.DeleteByObject(ctx, product)
	if err != nil {
		logger.Err(ctx, err.Error())
	}
	return res, nil
}

// Statistics 返回产品的统计数据
func (a *Product) Statistics(ctx context.Context, product string) (*tpl.ProductStatisticsRes, error) {
	return a.services.UrbsSetting.ProductStatistics(ctx, product)
}
