package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/tpl"
)

// ModuleList ...
func (a *UrbsSetting) ModuleList(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.ModulesInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.ModulesInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ModuleCreate ...
func (a *UrbsSetting) ModuleCreate(ctx context.Context, product string, body *tpl.NameDescBody) (*tpl.ModuleInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules", conf.Config.UrbsSetting.Addr, product)

	result := new(tpl.ModuleInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ModuleUpdate ...
func (a *UrbsSetting) ModuleUpdate(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*tpl.ModuleInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s", conf.Config.UrbsSetting.Addr, product, module)

	result := new(tpl.ModuleInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ModuleOffline ...
func (a *UrbsSetting) ModuleOffline(ctx context.Context, product string, module string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s:offline", conf.Config.UrbsSetting.Addr, product, module)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
