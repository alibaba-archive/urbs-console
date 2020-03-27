package bll

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/tpl"
)

// Setting ...
type Setting struct {
	services *service.Services
}

// List 读取指定产品功能模块的配置项
func (a *Setting) List(ctx context.Context, args *tpl.ProductModuleURL) (*urbssetting.SettingsInfoRes, error) {

	return a.services.UrbsSetting.SettingList(ctx, args)
}

// Create 创建指定产品功能模块配置项
func (a *Setting) Create(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.NameDescBody) (*urbssetting.SettingInfoRes, error) {
	return a.services.UrbsSetting.SettingCreate(ctx, args, body)
}

// Get 读取指定产品功能模块配置项
func (a *Setting) Get(ctx context.Context, args *tpl.ProductModuleSettingURL) (*urbssetting.SettingInfoRes, error) {
	return a.services.UrbsSetting.SettingGet(ctx, args)
}

// Update 更新指定产品功能模块配置项
func (a *Setting) Update(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*urbssetting.SettingInfoRes, error) {
	return a.services.UrbsSetting.SettingUpdate(ctx, args, body)
}

// Offline 下线指定产品功能模块配置项
func (a *Setting) Offline(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingOffline(ctx, args)
}

// Assign 批量为用户或群组设置产品功能模块配置项
func (a *Setting) Assign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error) {
	return a.services.UrbsSetting.SettingAssign(ctx, args, body)
}
