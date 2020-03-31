package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// ModuleList ...
func (a *UrbsSetting) ModuleList(ctx context.Context, args *tpl.ProductPaginationURL) (*urbssetting.ModulesRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules?skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.Product, args.Skip, args.PageSize)

	result := new(urbssetting.ModulesRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ModuleCreate ...
func (a *UrbsSetting) ModuleCreate(ctx context.Context, product string, body *tpl.NameDescBody) (*urbssetting.ModuleRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.ModuleRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// ModuleUpdate ...
func (a *UrbsSetting) ModuleUpdate(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*urbssetting.ModuleRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s", conf.Config.UrbsSetting.Addr, product, module)

	result := new(urbssetting.ModuleRes)

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
