package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// UserListLables ...
func (a *UrbsSetting) UserListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/labels?skip=%d&pageSize=%d&pageToken=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken)

	result := new(urbssetting.LabelsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRefreshCached ...
func (a *UrbsSetting) UserRefreshCached(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/labels:cache", conf.Config.UrbsSetting.Addr, uid)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserListSettings ...
func (a *UrbsSetting) UserListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/settings?skip=%d&pageSize=%d&pageToken=%s&product=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Product)

	result := new(urbssetting.MySettingsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserListSettingsUnionAll ...
func (a *UrbsSetting) UserListSettingsUnionAll(ctx context.Context, args *tpl.MySettingsQueryURL) (*urbssetting.MySettingsRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/settings:unionAll?skip=%d&pageSize=%d&pageToken=%s&product=%s&client=%s&channel=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Product, args.Client, args.Channel)

	result := new(urbssetting.MySettingsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserCheckExists ...
func (a *UrbsSetting) UserCheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s:exists", conf.Config.UrbsSetting.Addr, uid)

	result := new(tpl.BoolRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserBatchAdd ...
func (a *UrbsSetting) UserBatchAdd(ctx context.Context, users []string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users:batch", conf.Config.UrbsSetting.Addr)

	body := new(urbssetting.UsersBody)
	body.Users = users

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRemoveLabled ...
func (a *UrbsSetting) UserRemoveLabled(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/labels/%s", conf.Config.UrbsSetting.Addr, uid, hid)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRollbackSetting ...
func (a *UrbsSetting) UserRollbackSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/settings/%s:rollback", conf.Config.UrbsSetting.Addr, uid, hid)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRemoveSetting ...
func (a *UrbsSetting) UserRemoveSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/settings/%s", conf.Config.UrbsSetting.Addr, uid, hid)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
