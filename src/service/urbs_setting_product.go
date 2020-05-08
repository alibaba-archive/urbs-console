package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/tpl"
)

// ProductList ...
func (a *UrbsSetting) ProductList(ctx context.Context, args *tpl.Pagination) (*tpl.ProductsRes, error) {
	url := fmt.Sprintf("%s/v1/products?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.ProductsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductStatistics ...
func (a *UrbsSetting) ProductStatistics(ctx context.Context, product string) (*tpl.ProductStatisticsRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/statistics", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.ProductStatisticsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductCreate ...
func (a *UrbsSetting) ProductCreate(ctx context.Context, body *tpl.NameDescBody) (*tpl.ProductRes, error) {
	url := fmt.Sprintf("%s/v1/products", conf.Config.UrbsSetting.Addr)

	result := new(tpl.ProductRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdate ...
func (a *UrbsSetting) ProductUpdate(ctx context.Context, product string, body *tpl.ProductUpdateBody) (*tpl.ProductRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.ProductRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductOffline ...
func (a *UrbsSetting) ProductOffline(ctx context.Context, product string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s:offline", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDelete ...
func (a *UrbsSetting) ProductDelete(ctx context.Context, product string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
