package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// GroupListLables ...
func (a *UrbsSetting) GroupListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/labels?skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize)

	result := new(urbssetting.LabelsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupListSettings ...
func (a *UrbsSetting) GroupListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/settings?product=%s&skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.UID, args.Product, args.Skip, args.PageSize)

	result := new(urbssetting.MySettingsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupList ...
func (a *UrbsSetting) GroupList(ctx context.Context, args *tpl.GroupsURL) (*urbssetting.GroupsRes, error) {
	url := fmt.Sprintf("%s/v1/groups?kind=%s&skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.Kind, args.Skip, args.PageSize)

	result := new(urbssetting.GroupsRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupCheckExists ...
func (a *UrbsSetting) GroupCheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s:exists", conf.Config.UrbsSetting.Addr, uid)

	result := new(tpl.BoolRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupBatchAdd ...
func (a *UrbsSetting) GroupBatchAdd(ctx context.Context, groups []*tpl.GroupBody) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups:batch", conf.Config.UrbsSetting.Addr)

	body := new(tpl.GroupsBody)
	body.Groups = groups

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupUpdate ...
func (a *UrbsSetting) GroupUpdate(ctx context.Context, uid string, body *tpl.GroupUpdateBody) (*urbssetting.GroupRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s", conf.Config.UrbsSetting.Addr, uid)

	result := new(urbssetting.GroupRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupDelete ...
func (a *UrbsSetting) GroupDelete(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s", conf.Config.UrbsSetting.Addr, uid)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupListMembers ...
func (a *UrbsSetting) GroupListMembers(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.GroupMembersRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members?skip=%d&pageSize=%d", conf.Config.UrbsSetting.Addr, args.UID, args.Skip, args.PageSize)

	result := new(urbssetting.GroupMembersRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupBatchAddMembers ...
func (a *UrbsSetting) GroupBatchAddMembers(ctx context.Context, groupId string, users []string) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members:batch", conf.Config.UrbsSetting.Addr, groupId)

	body := new(urbssetting.UsersBody)
	body.Users = users

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupRemoveMembers ...
func (a *UrbsSetting) GroupRemoveMembers(ctx context.Context, args *tpl.GroupMembersURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/members", conf.Config.UrbsSetting.Addr, args.UID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupRemoveLable ...
func (a *UrbsSetting) GroupRemoveLable(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/labels/%s", conf.Config.UrbsSetting.Addr, args.UID, args.HID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupRollbackSetting ...
func (a *UrbsSetting) GroupRollbackSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/settings/%s:rollback", conf.Config.UrbsSetting.Addr, args.UID, args.HID)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// GroupRemoveSetting ...
func (a *UrbsSetting) GroupRemoveSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error) {
	url := fmt.Sprintf("%s/v1/groups/%s/settings/%s", conf.Config.UrbsSetting.Addr, args.UID, args.HID)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()
	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}
