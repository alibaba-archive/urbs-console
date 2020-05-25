package api

import (
	"github.com/teambition/gear"
	tracing "github.com/teambition/gear-tracing"

	"github.com/teambition/urbs-console/src/bll"
	"github.com/teambition/urbs-console/src/middleware"
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

func init() {
	util.DigProvide(newAPIs)
	util.DigProvide(newRouters)
}

var (
	services *service.Services
	blls     *bll.Blls
)

// APIs ..
type APIs struct {
	Canary *Canary
	User   *User
	Group  *Group

	Product *Product
	Label   *Label
	Module  *Module

	Setting    *Setting
	UrbsAcAcl  *UrbsAcAcl
	UrbsAcUser *UrbsAcUser
}

func newAPIs(b *bll.Blls, s *service.Services) *APIs {
	blls = b
	services = s
	apis := &APIs{
		Canary: &Canary{blls: blls},
		User:   &User{blls: blls},
		Group:  &Group{blls: blls},

		Product: &Product{blls: blls},
		Label:   &Label{blls: blls},
		Module:  &Module{blls: blls},

		Setting:    &Setting{blls: blls},
		UrbsAcAcl:  &UrbsAcAcl{blls: blls},
		UrbsAcUser: &UrbsAcUser{blls: blls},
	}
	return apis
}

func newRouters(apis *APIs) []*gear.Router {
	return []*gear.Router{newRouterV1(apis), newRouterAPIV1(apis)}
}

func newRouterV1(apis *APIs) *gear.Router {
	routerV1 := gear.NewRouter(gear.RouterOptions{
		Root: "/v1",
	})
	routerV1.Use(tracing.New())

	routerV1.Get("/canary", apis.Canary.Get)

	// ***** user ******
	// 读取指定用户的功能配置项，支持条件筛选，数据用于客户端
	routerV1.Get("/users/settings:unionAll", middleware.Verify(services), apis.User.ListSettingsUnionAllClient)

	return routerV1
}

func newRouterAPIV1(apis *APIs) *gear.Router {

	routerV1 := gear.NewRouter(gear.RouterOptions{
		Root: "/api/v1",
	})
	routerV1.Use(tracing.New())
	routerV1.Use(middleware.Verify(services))

	// ***** client ******
	// 读取指定用户的功能配置项，支持条件筛选，数据用于客户端
	routerV1.Get("/users/settings:unionAll", apis.User.ListSettingsUnionAllClient)

	checkSuperAdmin := middleware.CheckSuperAdmin(blls)

	checkViewer := middleware.CheckViewer(blls)
	// ***** UrbsAc ******
	// 添加用户
	routerV1.Post("/ac/users", checkSuperAdmin, apis.UrbsAcUser.Add)
	// 更新用户
	routerV1.Put("/ac/users/:uid", checkSuperAdmin, apis.UrbsAcUser.Update)
	// 删除用户
	routerV1.Delete("/ac/users/:uid", checkSuperAdmin, apis.UrbsAcUser.Delete)
	// 获取用户列表
	routerV1.Get("/ac/users", checkViewer, apis.UrbsAcUser.List)
	// 搜索用户
	routerV1.Get("/ac/users:search", checkViewer, apis.UrbsAcUser.Search)
	// 添加权限
	routerV1.Post("/ac/users/:uid/permissions", checkSuperAdmin, apis.UrbsAcAcl.Add)
	// 删除权限
	routerV1.Delete("/ac/users/:uid/permissions", checkSuperAdmin, apis.UrbsAcAcl.Delete)
	// 检查权限
	routerV1.Post("/ac/permission:check", checkViewer, apis.UrbsAcAcl.Check)

	// ***** product ******
	// 读取产品列表，支持条件筛选
	routerV1.Get("/products", checkViewer, apis.Product.List)
	// 读取指定产品的统计数据
	routerV1.Get("/products/:product/statistics", checkViewer, apis.Product.Statistics)
	// 创建产品
	routerV1.Post("/products", checkSuperAdmin, apis.Product.Create)
	// 更新指定产品
	routerV1.Put("/products/:product", apis.Product.Update)
	// 下线指定产品功能模块
	routerV1.Put("/products/:product+:offline", apis.Product.Offline)
	// 删除指定产品
	routerV1.Delete("/products/:product", checkSuperAdmin, apis.Product.Delete)

	// ***** label ******
	// 读取指定产品环境标签
	routerV1.Get("/products/:product/labels", checkViewer, apis.Label.List)
	// 读取指定产品下标签发布记录
	routerV1.Get("/products/:product/labels/:label/logs", checkViewer, apis.Label.Logs)
	// 获取标签下群组
	routerV1.Get("/products/:product/labels/:label/groups", checkViewer, apis.Label.ListGroups)
	// 移除指定群组的指定环境标签
	routerV1.Delete("/products/:product/labels/:label/groups/:uid", apis.Label.DeleteGroup)
	// 获取标签下用户
	routerV1.Get("/products/:product/labels/:label/users", checkViewer, apis.Label.ListUsers)
	// 移除指定用户的指定环境标签
	routerV1.Delete("/products/:product/labels/:label/users/:uid", apis.Label.DeleteUser)

	// 创建指定产品环境标签
	routerV1.Post("/products/:product/labels", apis.Label.Create)
	// 更新指定产品环境标签
	routerV1.Put("/products/:product/labels/:label", apis.Label.Update)
	// 删除指定产品环境标签
	routerV1.Delete("/products/:product/labels/:label", apis.Label.Delete)
	// 下线指定产品环境标签
	routerV1.Put("/products/:product/labels/:label+:offline", apis.Label.Offline)
	// 批量为用户或群组设置产品环境标签
	routerV1.Post("/products/:product/labels/:label+:assign", apis.Label.Assign)
	// 批量撤销对用户或群组设置的产品环境标签
	routerV1.Post("/products/:product/labels/:label+:recall", apis.Label.Recall)
	// 创建指定产品环境标签的灰度发布规则
	routerV1.Post("/products/:product/labels/:label/rules", apis.Label.CreateRule)
	// 读取指定产品环境标签的灰度发布规则列表
	routerV1.Get("/products/:product/labels/:label/rules", checkViewer, apis.Label.ListRules)
	// 更新指定产品环境标签的指定灰度发布规则
	routerV1.Put("/products/:product/labels/:label/rules/:hid", apis.Label.UpdateRule)
	// 删除指定产品环境标签的指定灰度发布规则
	routerV1.Delete("/products/:product/labels/:label/rules/:hid", apis.Label.DeleteRule)

	// ***** module ******
	// 读取指定产品的功能模块
	routerV1.Get("/products/:product/modules", checkViewer, apis.Module.List)
	// 指定产品创建功能模块
	routerV1.Post("/products/:product/modules", apis.Module.Create)
	// 更新指定产品功能模块
	routerV1.Put("/products/:product/modules/:module", apis.Module.Update)
	// 下线指定产品功能模块
	routerV1.Put("/products/:product/modules/:module+:offline", apis.Module.Offline)

	// ***** setting ******
	// 读取指定产品功能模块的配置项
	routerV1.Get("/products/:product/settings", apis.Setting.ListByProduct)
	// 读取指定产品功能模块的配置项
	routerV1.Get("/products/:product/modules/:module/settings", checkViewer, apis.Setting.List)
	// 读取指定具体配置项
	routerV1.Get("/products/:product/modules/:module/settings/:setting", apis.Setting.Get)
	// 读取指定产品功能模块配置项的发布记录
	routerV1.Get("/products/:product/modules/:module/settings/:setting/logs", checkViewer, apis.Setting.Logs)
	// 读取指定产品功能模块配置项的群组列表
	routerV1.Get("/products/:product/modules/:module/settings/:setting/groups", checkViewer, apis.Setting.ListGroups)
	// 移除指定群组的指定配置项
	routerV1.Delete("/products/:product/modules/:module/settings/:setting/groups/:uid", apis.Setting.DeleteGroup)
	// 回滚指定群组的指定配置项
	routerV1.Put("/products/:product/modules/:module/settings/:setting/groups/:uid+:rollback", apis.Setting.RollbackGroupSetting)
	// 读取指定产品功能模块配置项的用户列表
	routerV1.Get("/products/:product/modules/:module/settings/:setting/users", checkViewer, apis.Setting.ListUsers)
	// 移除指定用户的指定配置项
	routerV1.Delete("/products/:product/modules/:module/settings/:setting/users/:uid", apis.Setting.DeleteUser)
	// 回滚指定用户的指定配置项
	routerV1.Put("/products/:product/modules/:module/settings/:setting/users/:uid+:rollback", apis.Setting.RollbackUserSetting)
	// 创建指定产品功能模块配置项
	routerV1.Post("/products/:product/modules/:module/settings", apis.Setting.Create)
	// 更新指定产品功能模块配置项
	routerV1.Put("/products/:product/modules/:module/settings/:setting", apis.Setting.Update)
	// 下线指定产品功能模块配置项
	routerV1.Put("/products/:product/modules/:module/settings/:setting+:offline", apis.Setting.Offline)
	// 批量为用户或群组设置产品功能模块配置项
	routerV1.Post("/products/:product/modules/:module/settings/:setting+:assign", apis.Setting.Assign)
	// 批量撤销对用户或群组设置的产品灰度设置
	routerV1.Post("/products/:product/modules/:module/settings/:setting+:recall", apis.Setting.Recall)
	// 读取指定产品功能模块配置项的灰度发布规则列表
	routerV1.Get("/products/:product/modules/:module/settings/:setting/rules", checkViewer, apis.Setting.ListRules)
	// 创建指定产品功能模块配置项的灰度发布规则
	routerV1.Post("/products/:product/modules/:module/settings/:setting/rules", apis.Setting.CreateRule)
	// 更新指定产品功能模块配置项的指定灰度发布规则
	routerV1.Put("/products/:product/modules/:module/settings/:setting/rules/:hid", apis.Setting.UpdateRule)
	// 删除指定产品功能模块配置项的指定灰度发布规则
	routerV1.Delete("/products/:product/modules/:module/settings/:setting/rules/:hid", apis.Setting.DeleteRule)

	// ***** user ******
	// 读取用户列表，支持条件筛选
	routerV1.Get("/users", checkViewer, apis.User.List)
	// 读取指定用户的环境标签，支持条件筛选
	routerV1.Get("/users/:uid/labels", checkViewer, apis.User.ListLables)
	// 读取指定用户的功能配置项，支持条件筛选
	routerV1.Get("/users/:uid/settings", checkViewer, apis.User.ListSettings)
	// 强制刷新指定用户的环境标签列表缓存
	routerV1.Put("/users/:uid/labels:cache", checkViewer, apis.User.RefreshCachedLables)
	// 批量添加用户
	routerV1.Post("/users:batch", checkSuperAdmin, apis.User.BatchAdd)

	// ***** group ******
	// 读取群组列表，支持条件筛选
	routerV1.Get("/groups", checkViewer, apis.Group.List)
	// 读取指定群组的环境标签，支持条件筛选
	routerV1.Get("/groups/:uid/labels", checkViewer, apis.Group.ListLables)
	// 读取指定群组的功能配置项，支持条件筛选
	routerV1.Get("/groups/:uid/settings", checkViewer, apis.Group.ListSettings)
	// 读取群组成员列表，支持条件筛选
	routerV1.Get("/groups/:uid/members", checkViewer, apis.Group.ListMembers)
	// 批量添加群组
	routerV1.Post("/groups:batch", checkSuperAdmin, apis.Group.BatchAdd)
	// 更新指定群组
	routerV1.Put("/groups/:uid", checkSuperAdmin, apis.Group.Update)
	// 删除指定群组
	routerV1.Delete("/groups/:uid", checkSuperAdmin, apis.Group.Delete)
	// 指定群组批量添加成员
	routerV1.Post("/groups/:uid/members:batch", checkSuperAdmin, apis.Group.BatchAddMembers)
	// 指定群组根据条件清理成员
	routerV1.Delete("/groups/:uid/members", checkSuperAdmin, apis.Group.RemoveMembers)

	return routerV1
}
