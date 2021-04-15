package bll

import (
	"context"
	"time"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// User ...
type User struct {
	services *service.Services
}

// List ...
func (a *User) List(ctx context.Context, args *tpl.Pagination) (*tpl.UsersRes, error) {
	return a.services.UrbsSetting.UserList(ctx, args)
}

// ListLables ...
func (a *User) ListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MyLabelsRes, error) {
	return a.services.UrbsSetting.UserListLables(ctx, args)
}

// RefreshCachedLables ...
func (a *User) RefreshCachedLables(ctx context.Context, uid string) (*tpl.UserRes, error) {
	return a.services.UrbsSetting.UserRefreshCached(ctx, uid)
}

// ListSettings ...
func (a *User) ListSettings(ctx context.Context, args *tpl.UIDPaginationURL) (*tpl.MySettingsRes, error) {
	return a.services.UrbsSetting.UserListSettings(ctx, args)
}

// ListSettingsUnionAll ...
func (a *User) ListSettingsUnionAll(ctx *gear.Context, args *tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error) {
	mySettingsRes, mySettingsErr := a.services.UrbsSetting.UserListSettingsUnionAll(ctx, args)
	if mySettingsErr != nil {
		return nil, mySettingsErr
	}
	if args.Setting != "" {
		return mySettingsRes, nil
	}
	if !args.WithLabel {
		return mySettingsRes, nil
	}
	res, err := a.services.UrbsSetting.LabelsCache(ctx, args.Product, args.UID)
	if err != nil {
		return nil, err
	}
	var label string
	for _, val := range res.Result {
		if !MatchClient(val.Clients, args.Client) {
			continue
		}
		if !MatchChannel(val.Clients, args.Channel) {
			continue
		}
		label = val.Label
		break
	}
	if label != "" {
		temp := []*tpl.MySetting{{
			Product:    args.Product,
			Module:     "urbs",
			Name:       label,
			Value:      "true",
			AssignedAt: time.Now().UTC(),
		}}
		mySettingsRes.Result = append(temp, mySettingsRes.Result...)
	}
	return mySettingsRes, nil
}

// CheckExists ...
func (a *User) CheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserCheckExists(ctx, uid)
}

// BatchAdd 批量添加用户
func (a *User) BatchAdd(ctx context.Context, users []string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserBatchAdd(ctx, users)
}

// ApplyRules ...
func (a *User) ApplyRules(ctx context.Context, product string, body *tpl.ApplyRulesBody) (*tpl.BoolRes, error) {
	res, err := a.services.UrbsSetting.UserBatchAdd(ctx, body.Users)
	if err != nil {
		return nil, err
	}
	settingBody := &tpl.ApplyRulesBody{Kind: body.Kind}
	settingBody.Users = body.Users
	res, err = a.services.UrbsSetting.ProductApplyRule(ctx, product, settingBody)
	return res, err
}
