package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// SettingList ...
func (a *UrbsSetting) SettingList(ctx context.Context, args *tpl.ProductModuleURL) (*urbssetting.SettingsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings?skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Skip, args.PageSize)

	result := new(urbssetting.SettingsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingCreate ...
func (a *UrbsSetting) SettingCreate(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.NameDescBody) (*urbssetting.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings", conf.Config.UrbsSetting.Addr, args.Product, args.Module)

	result := new(urbssetting.SettingInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingGet ...
func (a *UrbsSetting) SettingGet(ctx context.Context, args *tpl.ProductModuleSettingURL) (*urbssetting.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(urbssetting.SettingInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingUpdate ...
func (a *UrbsSetting) SettingUpdate(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*urbssetting.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(urbssetting.SettingInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingOffline ...
func (a *UrbsSetting) SettingOffline(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s:offline", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingAssign ...
func (a *UrbsSetting) SettingAssign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s:assign", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
