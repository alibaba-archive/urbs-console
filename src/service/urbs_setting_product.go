package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/util/request"
)

// ProductList ...
func (a *UrbsSetting) ProductList(ctx context.Context, args *urbssetting.Pagination) (*urbssetting.ProductsRes, error) {
	url := fmt.Sprintf("%s/v1/products?skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.Skip, args.PageSize)

	result := new(urbssetting.ProductsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductCreate ...
func (a *UrbsSetting) ProductCreate(ctx context.Context, body *urbssetting.NameDescBody) (*urbssetting.ProductsRes, error) {
	url := fmt.Sprintf("%s/v1/products", conf.Config.UrbsSetting.Addr)

	result := new(urbssetting.ProductsRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdate ...
func (a *UrbsSetting) ProductUpdate(ctx context.Context, product string, body *urbssetting.ProductUpdateBody) (*urbssetting.ProductsRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.ProductsRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductOffline ...
func (a *UrbsSetting) ProductOffline(ctx context.Context, product string) (*urbssetting.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s:offline", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDelete ...
func (a *UrbsSetting) ProductDelete(ctx context.Context, product string) (*urbssetting.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
