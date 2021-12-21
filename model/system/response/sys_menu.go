// Package response
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 12:45
package response

import (
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
)

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}

type SysMenuTreeResponse struct {
	ID       uint                  `json:"key" gorm:"primarykey"`         // 主键ID
	Name     string                `json:"title" gorm:"comment:路由name"`   // 路由name
	ParentId string                `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Children []SysMenuTreeResponse `json:"children" gorm:"-"`
}
