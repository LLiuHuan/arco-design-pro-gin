// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:51
package system

import (
	"reflect"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	var entity AuthorityMenus
	return entity.TableName()
}

func (a *authoritiesMenus) Initialize() error {
	entities := []AuthorityMenus{
		{BaseMenuId: 1, AuthorityId: "888"},
		{BaseMenuId: 2, AuthorityId: "888"},
		{BaseMenuId: 3, AuthorityId: "888"},
		{BaseMenuId: 4, AuthorityId: "888"},
		{BaseMenuId: 5, AuthorityId: "888"},
		{BaseMenuId: 6, AuthorityId: "888"},
		{BaseMenuId: 7, AuthorityId: "888"},
		{BaseMenuId: 8, AuthorityId: "888"},
		{BaseMenuId: 9, AuthorityId: "888"},
		{BaseMenuId: 10, AuthorityId: "888"},
		{BaseMenuId: 11, AuthorityId: "888"},
		{BaseMenuId: 12, AuthorityId: "888"},
		{BaseMenuId: 13, AuthorityId: "888"},
		{BaseMenuId: 14, AuthorityId: "888"},
		{BaseMenuId: 15, AuthorityId: "888"},
		{BaseMenuId: 16, AuthorityId: "888"},
		{BaseMenuId: 17, AuthorityId: "888"},
	}
	if err := global.AdpDb.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.AdpDb.Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?", 17, "9528").First(&AuthorityMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

func (a *AuthorityMenus) TableName() string {
	var entity system.SysAuthority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("SysBaseMenus"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}
