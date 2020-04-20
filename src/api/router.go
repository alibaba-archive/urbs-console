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

	checkSuperAdmin := middleware.CheckSuperAdmin(blls)

	checkViewer := middleware.CheckViewer(blls)
	// ***** UrbsAc ******
	// 添加用户
	routerV1.Post("/ac/users", checkSuperAdmin, apis.UrbsAcUser.Add)
	// 添加权限
	routerV1.Post("/ac/users/:uid/permissions", checkSuperAdmin, apis.UrbsAcAcl.Add)

	// ***** product ******
	// 读取产品列表，支持条件筛选
	routerV1.Get("/products", checkViewer, apis.Product.List)
	// 创建产品
	routerV1.Post("/products", checkSuperAdmin, apis.Product.Create)
	// 更新指定产品
	routerV1.Put("/products/:product", apis.Product.Update)
	// 下线指定产品功能模块
	routerV1.Put("/products/:product+:offline", apis.Product.Offline)
	// 删除指定产品
	routerV1.Delete("/products/:product", checkSuperAdmin, apis.Product.Delete)

	// ***** label ******
	// 读取指定产品灰度标签
	routerV1.Get("/products/:product/labels", checkViewer, apis.Label.List)
	// 读取指定产品下标签发布记录
	routerV1.Get("/products/:product/labels/:label/logs", checkViewer, apis.Label.Logs)
	// 获取标签下群组
	routerV1.Get("/products/:product/labels/:label/groups", checkViewer, apis.Label.GetGroups)
	// 获取标签下用户
	routerV1.Get("/products/:product/labels/:label/users", checkViewer, apis.Label.GetUsers)
	// 创建指定产品灰度标签
	routerV1.Post("/products/:product/labels", apis.Label.Create)
	// 更新指定产品灰度标签
	routerV1.Put("/products/:product/labels/:label", apis.Label.Update)
	// 删除指定产品灰度标签
	routerV1.Delete("/products/:product/labels/:label", apis.Label.Delete)
	// 下线指定产品灰度标签
	routerV1.Put("/products/:product/labels/:label+:offline", apis.Label.Offline)
	// 批量为用户或群组设置产品灰度标签
	routerV1.Post("/products/:product/labels/:label+:assign", apis.Label.Assign)

	// 批量撤销对用户或群组设置的产品灰度标签
	routerV1.Post("/products/:product/labels/:label+:recall", apis.Label.Recall)

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
	routerV1.Get("/products/:product/modules/:module/settings", checkViewer, apis.Setting.List)
	// 读取发布记录
	routerV1.Get("/products/:product/modules/:module/settings/:setting/logs", checkViewer, apis.Setting.Logs)
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

	// ***** user ******
	// 读取用户列表，支持条件筛选
	routerV1.Get("/users", checkViewer, apis.User.List)
	// 读取指定用户的灰度标签，支持条件筛选
	routerV1.Get("/users/:uid/labels", checkViewer, apis.User.ListLables)
	// 读取指定用户的功能配置项，支持条件筛选
	routerV1.Get("/users/:uid/settings", checkViewer, apis.User.ListSettings)
	// 强制刷新指定用户的灰度标签列表缓存
	routerV1.Put("/users/:uid/labels:cache", checkViewer, apis.User.RefreshCachedLables)
	// 删除指定用户的指定灰度标签
	routerV1.Delete("/users/:uid/labels/:hid", checkSuperAdmin, apis.User.RemoveLable)
	// 回滚指定用户的指定配置项
	routerV1.Put("/users/:uid/settings/:hid+:rollback", checkSuperAdmin, apis.User.RollbackSetting)
	// 删除指定用户的指定配置项
	routerV1.Delete("/users/:uid/settings/:hid", checkSuperAdmin, apis.User.RemoveSetting)
	// 批量添加用户
	routerV1.Post("/users:batch", checkSuperAdmin, apis.User.BatchAdd)

	// // 查询指定用户是否存在
	// routerV1.Get("/users/:uid+:exists", checkSuperAdmin, apis.User.CheckExists)

	// ***** group ******
	// 读取群组列表，支持条件筛选
	routerV1.Get("/groups", checkViewer, apis.Group.List)
	// 读取指定群组的灰度标签，支持条件筛选
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
	// 删除指定群组的指定灰度标签
	routerV1.Delete("/groups/:uid/labels/:hid", checkSuperAdmin, apis.Group.RemoveLable)
	// 回滚指定群组的指定配置项
	routerV1.Put("/groups/:uid/settings/:hid+:rollback", checkSuperAdmin, apis.Group.RollbackSetting)
	// 删除指定群组的指定配置项
	routerV1.Delete("/groups/:uid/settings/:hid", checkSuperAdmin, apis.Group.RemoveSetting)
	// 指定群组根据条件清理成员
	routerV1.Delete("/groups/:uid/members", checkSuperAdmin, apis.Group.RemoveMembers)

	// // 查询指定群组是否存在
	// routerV1.Get("/groups/:uid+:exists", checkSuperAdmin, apis.Group.CheckExists)
	return routerV1
}
