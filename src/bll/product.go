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
	services *service.Services
	daos     *dao.Daos
}

// Create 创建产品
func (a *Product) Create(ctx context.Context, args *tpl.NameDescBody) (*tpl.ProductRes, error) {
	for _, uid := range args.Uids {
		err := blls.UrbsAcAcl.AddDefaultPermission(ctx, uid, args.Name)
		if err != nil {
			return nil, err
		}
	}
	res, err := a.services.UrbsSetting.ProductCreate(ctx, args)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, args.Name)
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
	subjects, err := blls.UrbsAcAcl.FindUsersByObjects(ctx, objects)
	if err != nil {
		return nil, err
	}
	for _, product := range products.Result {
		product.Users = subjects[product.Name]
	}
	return products, nil
}

// Update ...
func (a *Product) Update(ctx context.Context, product string, args *tpl.ProductUpdateBody) (*tpl.ProductRes, error) {
	if len(args.Uids) > 0 {
		err := blls.UrbsAcAcl.Update(ctx, args.Uids, product)
		if err != nil {
			return nil, err
		}
	}
	res, err := a.services.UrbsSetting.ProductUpdate(ctx, product, args)
	if err != nil {
		return nil, err
	}
	res.Result.Users, err = blls.UrbsAcAcl.FindUsersByObject(ctx, product)
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
	err := daos.UrbsAcAcl.DeleteByObject(ctx, product)
	if err != nil {
		logger.Err(ctx, err.Error())
	}
	return a.services.UrbsSetting.ProductDelete(ctx, product)
}

// Statistics 返回产品的统计数据
func (a *Product) Statistics(ctx context.Context, product string) (*tpl.ProductStatisticsRes, error) {
	return a.services.UrbsSetting.ProductStatistics(ctx, product)
}
