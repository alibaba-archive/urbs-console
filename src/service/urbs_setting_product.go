package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// ProductCreate ...
func (a *UrbsSetting) ProductCreate(ctx context.Context, body *tpl.NameDescBody) (*urbssetting.ProductRes, error) {
	url := fmt.Sprintf("%s/v1/products", conf.Config.UrbsSetting.Addr)

	result := new(urbssetting.ProductRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductList ...
func (a *UrbsSetting) ProductList(ctx context.Context, args *tpl.Pagination) (*urbssetting.ProductsRes, error) {
	url := fmt.Sprintf("%s/v1/products?skip=%d&pageSize=%d&pageToken=%s", conf.Config.UrbsSetting.Addr, args.Skip, args.PageSize, args.PageToken)

	result := new(urbssetting.ProductsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdate ...
func (a *UrbsSetting) ProductUpdate(ctx context.Context, product string, body *tpl.ProductUpdateBody) (*urbssetting.ProductRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.ProductRes)

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
