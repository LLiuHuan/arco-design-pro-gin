// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:42
package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "LAYOUT", Sort: 1, Redirect: "/dashboard/workplace", Meta: system.Meta{Title: "menu.dashboard.index", Icon: "IconDashboard", Hidden: false, Permissions: "dashboard"}, MenuType: 1},
		{MenuLevel: 0, ParentId: "1", Path: "workplace", Name: "workplace", Component: "/dashboard/workplace/index.vue", Sort: 1, Meta: system.Meta{Title: "menu.dashboard.workplace", Icon: "IconDashboard", Hidden: false, Permissions: "dashboard.workplace"}, MenuType: 2},

		{MenuLevel: 0, ParentId: "0", Path: "about", Name: "about", Component: "LAYOUT", Sort: 7, Redirect: "/about/index", Meta: system.Meta{Title: "menu.about.index", Icon: "IconExclamationCircleFill", Hidden: false, Permissions: "about.index"}, MenuType: 1},
		{MenuLevel: 0, ParentId: "3", Path: "index", Name: "about_index", Component: "/about/index.vue", Sort: 7, Meta: system.Meta{Title: "menu.about.index", Icon: "IconExclamationCircleFill", Hidden: false, Permissions: "about.about_index"}, MenuType: 2},

		{MenuLevel: 0, ParentId: "0", Path: "system", Name: "系统管理", Component: "LAYOUT", Sort: 6, Redirect: "/system/menu", Meta: system.Meta{Title: "menu.system.index", Icon: "IconSettings", Hidden: false, KeepAlive: true, Permissions: "system.index"}, MenuType: 1},
		{MenuLevel: 0, ParentId: "5", Path: "menu", Name: "菜单管理", Component: "/system/menu/index.vue", Sort: 1, Meta: system.Meta{Title: "menu.system.menu", Hidden: false, Permissions: "system.menu"}, MenuType: 2},
		{MenuLevel: 0, ParentId: "5", Path: "user", Name: "用户管理", Component: "/system/user/index.vue", Sort: 2, Meta: system.Meta{Title: "menu.system.user", Hidden: false, Permissions: "system.user"}, MenuType: 2},
		{MenuLevel: 0, ParentId: "5", Name: "system.user.add", Sort: 1, Meta: system.Meta{Title: "添加", Hidden: false, Permissions: "system.user.add"}, MenuType: 3},
		{MenuLevel: 0, ParentId: "5", Name: "system.user.del", Sort: 2, Meta: system.Meta{Title: "删除", Hidden: false, Permissions: "system.user.del"}, MenuType: 3},
		{MenuLevel: 0, ParentId: "5", Name: "system.user.auth", Sort: 3, Meta: system.Meta{Title: "角色", Hidden: false, Permissions: "system.user.auth"}, MenuType: 3},
		{MenuLevel: 0, ParentId: "5", Path: "auth", Name: "角色管理", Component: "/system/auth/index.vue", Sort: 3, Meta: system.Meta{Title: "menu.system.auth", Hidden: false, Permissions: "system.auth"}, MenuType: 2},
		{MenuLevel: 0, ParentId: "5", Path: "api", Name: "API管理", Component: "/system/api/index.vue", Sort: 4, Meta: system.Meta{Title: "menu.system.api", Hidden: false, Permissions: "system.api"}, MenuType: 2},
		{MenuLevel: 0, ParentId: "5", Path: "operationRecord", Name: "操作历史", Component: "/system/operationRecord/index.vue", Sort: 5, Meta: system.Meta{Title: "menu.system.operationRecord", Hidden: false, Permissions: "system.operationRecord"}, MenuType: 2},

		{MenuLevel: 0, ParentId: "0", Path: "state", Name: "服务器状态首页", Component: "LAYOUT", Sort: 8, Redirect: "/state/index", Meta: system.Meta{Title: "menu.state.index", Icon: "IconLoading", Hidden: false, KeepAlive: true, Permissions: "state.index"}, MenuType: 1},
		{MenuLevel: 0, ParentId: "14", Path: "index", Name: "服务器状态", Component: "/state/index.vue", Sort: 1, Meta: system.Meta{Title: "menu.state.index", Icon: "IconLoading", Hidden: false, Permissions: "state.index"}, MenuType: 2},

		{MenuLevel: 0, ParentId: "0", Path: "https://www.baidu.com", Name: "百度", Component: "/", Sort: 0, Meta: system.Meta{Title: "百度", Icon: "IconDashboard", Hidden: false}, MenuType: 4},
	}
	if err := global.AdpDb.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.AdpDb.Where("path = ?", "autoCodeEdit/:id").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
