package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/tpl"
)

// SettingListByProduct ...
func (a *UrbsSetting) SettingListByProduct(ctx context.Context, args *tpl.ProductPaginationURL) (*tpl.SettingsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/settings?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.SettingsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingList ...
func (a *UrbsSetting) SettingList(ctx context.Context, args *tpl.ProductModuleURL) (*tpl.SettingsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.SettingsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingCreate ...
func (a *UrbsSetting) SettingCreate(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.NameDescBody) (*tpl.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings", conf.Config.UrbsSetting.Addr, args.Product, args.Module)

	result := new(tpl.SettingInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingGet ...
func (a *UrbsSetting) SettingGet(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.SettingInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingUpdate ...
func (a *UrbsSetting) SettingUpdate(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*tpl.SettingInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.SettingInfoRes)

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
func (a *UrbsSetting) SettingAssign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) (*tpl.SettingReleaseInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s:assign", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.SettingReleaseInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingRecall 撤销指定批次的用户或群组的配置项
func (a *UrbsSetting) SettingRecall(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.RecallBody) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s:recall", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingListUsers ...
func (a *UrbsSetting) SettingListUsers(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingUsersInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s/users?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.SettingUsersInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingListGroups ...
func (a *UrbsSetting) SettingListGroups(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingGroupsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s/groups?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.SettingGroupsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingCreateRule ...
func (a *UrbsSetting) SettingCreateRule(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error) {

	url := fmt.Sprintf("%s/products/%s/modules/%s/settings/%s/rules", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.SettingRuleInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingUpdateRule ...
func (a *UrbsSetting) SettingUpdateRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL, body *tpl.SettingRuleBody) (*tpl.SettingRuleInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s/rules/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting, args.HID)

	result := new(tpl.SettingRuleInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingDeleteRule ...
func (a *UrbsSetting) SettingDeleteRule(ctx context.Context, args *tpl.ProductModuleSettingHIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s/rules/%s", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting, args.HID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// SettingListRule ...
func (a *UrbsSetting) SettingListRule(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.SettingRulesInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/modules/%s/settings/%s/rules", conf.Config.UrbsSetting.Addr, args.Product, args.Module, args.Setting)

	result := new(tpl.SettingRulesInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
