package service

import (
	"context"

	"github.com/teambition/urbs-console/src/dto/urbssetting"
	"github.com/teambition/urbs-console/src/tpl"
)

// UrbsSettingInterface ....
type UrbsSettingInterface interface {
	// ***** product ******
	// 读取产品列表，支持条件筛选
	ProductList(ctx context.Context, args *tpl.Pagination) (*urbssetting.ProductsRes, error)
	// 创建产品
	ProductCreate(ctx context.Context, body *tpl.NameDescBody) (*urbssetting.ProductRes, error)
	// 更新产品
	ProductUpdate(ctx context.Context, product string, body *tpl.ProductUpdateBody) (*urbssetting.ProductRes, error)
	// 下线指定产品功能模块
	ProductOffline(ctx context.Context, product string) (*tpl.BoolRes, error)
	// 删除指定产品
	ProductDelete(ctx context.Context, product string) (*tpl.BoolRes, error)

	// ***** module ******
	// 读取指定产品的功能模块
	ModuleList(ctx context.Context, args *tpl.ProductPaginationURL) (*urbssetting.ModulesRes, error)
	// 指定产品创建功能模块
	ModuleCreate(ctx context.Context, product string, body *tpl.NameDescBody) (*urbssetting.ModuleRes, error)
	// 更新指定产品功能模块
	ModuleUpdate(ctx context.Context, product string, module string, body *tpl.ModuleUpdateBody) (*urbssetting.ModulesRes, error)
	// 下线指定产品功能模块
	ModuleOffline(ctx context.Context, product string, module string) (*tpl.BoolRes, error)

	// ***** setting ******
	// 读取指定产品功能模块的配置项
	SettingList(ctx context.Context, args *tpl.ProductModuleURL) (*urbssetting.SettingsInfoRes, error)
	// 创建指定产品功能模块配置项
	SettingCreate(ctx context.Context, args *tpl.ProductModuleURL, body *tpl.NameDescBody) (*urbssetting.SettingInfoRes, error)
	// 读取指定产品功能模块配置项
	SettingGet(ctx context.Context, args *tpl.ProductModuleSettingURL) (*urbssetting.SettingInfoRes, error)
	// 更新指定产品功能模块配置项
	SettingUpdate(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.SettingUpdateBody) (*urbssetting.SettingInfoRes, error)
	// 下线指定产品功能模块配置项
	SettingOffline(ctx context.Context, args *tpl.ProductModuleSettingURL) (*tpl.BoolRes, error)
	// 批量为用户或群组设置产品功能模块配置项
	SettingAssign(ctx context.Context, args *tpl.ProductModuleSettingURL, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error)

	// ***** label ******
	// 读取指定产品灰度标签
	LabelList(ctx context.Context, args *tpl.ProductPaginationURL) (*urbssetting.LabelsInfoRes, error)
	// 给指定产品创建灰度标签
	LabelCreate(ctx context.Context, product string, body *tpl.LabelBody) (*urbssetting.LabelInfoRes, error)
	// 更新指定产品灰度标签
	LabelUpdate(ctx context.Context, product string, label string, body *tpl.LabelUpdateBody) (*urbssetting.LabelInfoRes, error)
	// 删除指定产品灰度标签
	LabelDelete(ctx context.Context, product string, label string) (*tpl.BoolRes, error)
	// 下线指定产品灰度标签
	LabelOffline(ctx context.Context, product string, label string) (*tpl.BoolRes, error)
	// 批量为用户或群组设置产品灰度标签
	LabelAssign(ctx context.Context, product string, label string, body *tpl.UsersGroupsBody) (*tpl.BoolRes, error)

	// ***** user ******
	// 读取指定用户的灰度标签，支持条件筛选
	UserListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error)
	// 强制刷新指定用户的灰度标签列表缓存
	UserRefreshCached(ctx context.Context, uid string) (*tpl.BoolRes, error)
	// 读取指定用户的功能配置项，支持条件筛选
	UserListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error)
	// 读取指定用户的功能配置项，支持条件筛选，数据用于客户端
	UserListSettingsUnionAll(ctx context.Context, args *tpl.MySettingsQueryURL) (*urbssetting.MySettingsRes, error)
	// 查询指定用户是否存在
	UserCheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error)
	// 批量添加用户
	UserBatchAdd(ctx context.Context, users []string) (*tpl.BoolRes, error)
	// 删除指定用户的指定灰度标签
	UserRemoveLabled(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error)
	// 回滚指定用户的指定配置项
	UserRollbackSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error)
	// 删除指定用户的指定配置项
	UserRemoveSetting(ctx context.Context, uid string, hid string) (*tpl.BoolRes, error)

	// ***** group ******
	// 读取指定群组的灰度标签，支持条件筛选
	GroupListLables(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.LabelsInfoRes, error)
	// 读取指定群组的功能配置项，支持条件筛选
	GroupListSettings(ctx context.Context, args *tpl.UIDProductURL) (*urbssetting.MySettingsRes, error)
	// 读取群组列表，支持条件筛选
	GroupList(ctx context.Context, args *tpl.GroupsURL) (*urbssetting.GroupsRes, error)
	// 查询指定群组是否存在
	GroupCheckExists(ctx context.Context, uid string) (*tpl.BoolRes, error)
	// 批量添加群组
	GroupBatchAdd(ctx context.Context, groups []*tpl.GroupBody) (*tpl.BoolRes, error)
	// 更新指定群组
	GroupUpdate(ctx context.Context, uid string, body *tpl.GroupUpdateBody) (*urbssetting.GroupRes, error)
	// 删除指定群组
	GroupDelete(ctx context.Context, uid string) (*tpl.BoolRes, error)
	// 读取群组成员列表，支持条件筛选
	GroupListMembers(ctx context.Context, args *tpl.UIDPaginationURL) (*urbssetting.GroupMembersRes, error)
	// 指定群组批量添加成员
	GroupBatchAddMembers(ctx context.Context, groupId string, users []string) (*tpl.BoolRes, error)
	// 指定群组根据条件清理成员
	GroupRemoveMembers(ctx context.Context, args *tpl.GroupMembersURL) (*tpl.BoolRes, error)
	// 删除指定群组的指定灰度标签
	GroupRemoveLable(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error)
	// 回滚指定群组的指定配置项
	GroupRollbackSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error)
	// 删除指定群组的指定配置项
	GroupRemoveSetting(ctx context.Context, args *tpl.UIDHIDURL) (*tpl.BoolRes, error)
}
