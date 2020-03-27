package bll

import (
	"context"

	"github.com/teambition/gear"
	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// User ...
type User struct {
	services *service.Services
}

// ListLables ...
func (a *User) ListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error) {
	return a.services.UrbsSetting.UserListLables(ctx, args)
}

// RefreshCachedLables ...
func (a *User) RefreshCachedLables(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserRefreshCached(ctx, uid)
}

// ListSettings ...
func (a *User) ListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error) {
	return a.services.UrbsSetting.UserListSettings(ctx, args)
}

// ListSettingsUnionAll ...
func (a *User) ListSettingsUnionAll(ctx *gear.Context, args *tpl.MySettingsQueryURL) (*urbssetting.MySettingsRes, error) {
	return a.services.UrbsSetting.UserListSettingsUnionAll(ctx, args)
}

// CheckExists ...
func (a *User) CheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserCheckExists(ctx, uid)
}

// BatchAdd 批量添加用户
func (a *User) BatchAdd(ctx context.Context, users []string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserBatchAdd(ctx, users)
}

// RemoveLable ...
func (a *User) RemoveLable(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserRemoveLabled(ctx, uid, hid)
}

// RollbackSetting ...
func (a *User) RollbackSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserRollbackSetting(ctx, uid, hid)
}

// RemoveSetting ...
func (a *User) RemoveSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.UserRemoveSetting(ctx, uid, hid)
}
