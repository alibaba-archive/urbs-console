package service

import (
	"context"
	"fmt"

	"github.com/mushroomsir/request"
	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/tpl"
)

// UserList ...
func (a *UrbsSetting) UserList(ctx context.Context, args *tpl.Pagination) (*tpl.UsersRes, error) {
	url := fmt.Sprintf("%s/v1/users?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.UsersRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserListLables ...
func (a *UrbsSetting) UserListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MyLabelsRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/labels?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.MyLabelsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRefreshCached ...
func (a *UrbsSetting) UserRefreshCached(ctx context.Context, uid string) (*tpl.UserRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/labels:cache", conf.Config.UrbsSetting.Addr, uid)

	result := new(tpl.UserRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserListSettings ...
func (a *UrbsSetting) UserListSettings(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MySettingsRes, error) {
	url := fmt.Sprintf("%s/v1/users/%s/settings?skip=%d&pageSize=%d&pageToken=%s&q=%s", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Q)

	result := new(tpl.MySettingsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// UserListSettingsUnionAll ...
func (a *UrbsSetting) UserListSettingsUnionAll(ctx context.Context, args *tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error) {
	path := "%s/v1/users/%s/settings:unionAll?skip=%d&pageSize=%d&pageToken=%s"
	path += "&product=%s"
	path += "&module=%s"
	path += "&setting=%s"
	path += "&client=%s"
	path += "&channel=%s"
	url := fmt.Sprintf(path, conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize, args.PageToken, args.Product, args.Module, args.Setting, args.Client, args.Channel)

	result := new(tpl.MySettingsRes)

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

	body := new(tpl.UsersBody)
	body.Users = users

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
