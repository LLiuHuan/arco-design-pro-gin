// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:39
package system

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c *casbin) Initialize() error {

	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: "/v1/users", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/users", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/v1/users/info", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/users/:id", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/v1/users/register", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/users/authority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/users/authorities", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/v1/api", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/v1/api", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/api", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/v1/api/all", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/api/:id", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/api/page", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/api/page", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/api/byIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/v1/authority", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/authority", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/v1/authority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/authority", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/v1/authority/copy", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/v1/casbin", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/v1/casbin/:id", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/base/login", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/resetPassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/jwt/black", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/v1/menus", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/menus", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/v1/menus", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/:id", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/user", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/base", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/auth", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/auth", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/menus/:id", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/v1/operation/record", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/operation/record", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/v1/operation/record", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/v1/operation/record/:id", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/v1/operation/record/byIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/v1/system/info", V2: "GET"},
	}
	if err := global.AdpDb.Create(&entities).Error; err != nil {
		return errors.Wrap(err, c.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (c *casbin) CheckDataExist() bool {
	if errors.Is(global.AdpDb.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "GET", V2: "/user/getUserInfo"}).First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
