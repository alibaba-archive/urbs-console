package service

import (
	"context"
	"fmt"

	"github.com/teambition/urbs-console/src/conf"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
	"github.com/teambition/urbs-console/src/util/request"
)

// LabelList ...
func (a *UrbsSetting) LabelList(ctx context.Context, args *tpl.ProductPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels?skip=%d&pageSize=%d&pageToken=%s", conf.Config.UrbsSetting.Addr, args.Product, args.Skip, args.PageSize, args.PageToken)

	result := new(urbssetting.LabelsInfoRes)

	resp, err := request.Get(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelCreate ...
func (a *UrbsSetting) LabelCreate(ctx context.Context, product string, args *tpl.LabelBody) (*urbssetting.LabelInfoRes, error) {
	url := fmt.Sprintf("%s/v1/products/%s/labels", conf.Config.UrbsSetting.Addr, product)

	result := new(urbssetting.LabelInfoRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(args).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelUpdate ...
func (a *UrbsSetting) LabelUpdate(ctx context.Context, product string, label string, body *tpl.LabelUpdateBody) (*urbssetting.LabelInfoRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s", conf.Config.UrbsSetting.Addr, product, label)

	result := new(urbssetting.LabelInfoRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelDelete ...
func (a *UrbsSetting) LabelDelete(ctx context.Context, product string, label string) (*tpl.BoolRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.BoolRes)

	resp, err := request.Delete(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelOffline ...
func (a *UrbsSetting) LabelOffline(ctx context.Context, product string, label string) (*tpl.BoolRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s:offline", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.BoolRes)

	resp, err := request.Put(url).Header(UrbsSettingHeader(ctx)).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}

// LabelAssign ...
func (a *UrbsSetting) LabelAssign(ctx context.Context, product string, label string, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error) {

	url := fmt.Sprintf("%s/v1/products/%s/labels/%s:assign", conf.Config.UrbsSetting.Addr, product, label)

	result := new(tpl.BoolRes)

	resp, err := request.Post(url).Header(UrbsSettingHeader(ctx)).Body(body).Result(result).Do()

	if err := HanderResponse(resp, err); err != nil {
		return nil, err
	}
	return result, nil
}