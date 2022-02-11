// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:59
// @desc: 角色菜单关联视图
package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string    `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string    `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu `json:"children" gorm:"-"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
