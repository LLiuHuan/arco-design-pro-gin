// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:33
package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

func (a *api) TableName() string {
	return "sys_apis"
}

func (a *api) Initialize() error {
	entities := []system.SysApi{
		{ApiGroup: "系统用户", Method: "GET", Path: "/v1/users", Description: "获取用户列表"},
		{ApiGroup: "系统用户", Method: "DELETE", Path: "/v1/users/:id", Description: "删除用户"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/v1/users", Description: "设置用户信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/api/resetPassword", Description: "重置用户密码"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/v1/users/info", Description: "获取自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/v1/users/authorities", Description: "设置权限组"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/v1/users/register", Description: "用户注册(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/v1/users/authority", Description: "修改用户角色(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/changePassword", Description: "修改密码（建(选择)"},

		{ApiGroup: "api", Method: "PUT", Path: "/v1/api", Description: "更新Api"},
		{ApiGroup: "api", Method: "POST", Path: "/v1/api", Description: "创建api"},
		{ApiGroup: "api", Method: "DELETE", Path: "/v1/api", Description: "删除Api"},
		{ApiGroup: "api", Method: "GET", Path: "/v1/api/all", Description: "获取所有api"},
		{ApiGroup: "api", Method: "POST", Path: "/v1/api/page", Description: "获取api列表"},
		{ApiGroup: "api", Method: "GET", Path: "/v1/api/:id", Description: "获取api详细信息"},
		{ApiGroup: "api", Method: "DELETE", Path: "/v1/api/byIds", Description: "批量删除api"},

		{ApiGroup: "角色", Method: "POST", Path: "/v1/authority", Description: "创建角色"},
		{ApiGroup: "角色", Method: "DELETE", Path: "/v1/authority", Description: "删除角色"},
		{ApiGroup: "角色", Method: "GET", Path: "/v1/authority", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "PUT", Path: "/v1/authority", Description: "更新角色信息"},
		{ApiGroup: "角色", Method: "POST", Path: "/v1/authority/copy", Description: "拷贝角色"},

		{ApiGroup: "菜单", Method: "PUT", Path: "/v1/menus", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/v1/menus", Description: "新增菜单"},
		{ApiGroup: "菜单", Method: "DELETE", Path: "/v1/menus/:id", Description: "删除菜单"},
		{ApiGroup: "菜单", Method: "GET", Path: "/v1/menus/:id", Description: "根据id获取菜单"},
		{ApiGroup: "菜单", Method: "GET", Path: "/v1/menus", Description: "分页获取基础menu列表"},
		{ApiGroup: "菜单", Method: "GET", Path: "/v1/menus/base", Description: "获取用户动态路由"},
		{ApiGroup: "菜单", Method: "GET", Path: "/v1/menus/user", Description: "获取菜单树(必选)"},
		{ApiGroup: "菜单", Method: "GET", Path: "/v1/menus/auth", Description: "获取指定角色menu"},
		{ApiGroup: "菜单", Method: "POST", Path: "/v1/menus/auth", Description: "增加menu和角色关联关系"},

		{ApiGroup: "操作记录", Method: "POST", Path: "/v1/operation/record", Description: "新增操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/v1/operation/record", Description: "删除操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/v1/operation/record", Description: "获取操作记录列表"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/v1/operation/record/:id", Description: "根据ID获取操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/v1/operation/record/byIds", Description: "批量删除操作历史"},

		{ApiGroup: "casbin", Method: "PUT", Path: "/v1/casbin", Description: "更改角色api权限"},
		{ApiGroup: "casbin", Method: "GET", Path: "/v1/casbin/:id", Description: "获取权限列表"},

		{ApiGroup: "系统服务", Method: "GET", Path: "/v1/system/info", Description: "获取服务器信息"},

		{ApiGroup: "base", Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/black", Description: "jwt加入黑名单(退出，必选)"},
	}
	if err := global.AdpDb.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *api) CheckDataExist() bool {
	if errors.Is(global.AdpDb.Where("path = ? AND method = ?", "/excel/downloadTemplate", "GET").First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
