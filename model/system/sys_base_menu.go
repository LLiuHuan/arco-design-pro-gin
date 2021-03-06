// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:59
// @desc: 菜单表
package system

import "github.com/lliuhuan/arco-design-pro-gin/global"

type SysBaseMenu struct {
	global.AdpModel
	MenuLevel     uint                              `json:"-"`
	ParentId      string                            `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path          string                            `json:"path" gorm:"comment:路由path"`        // 路由path
	Name          string                            `json:"name" gorm:"comment:路由name"`        // 路由name
	Component     string                            `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int                               `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Redirect      string                            `json:"redirect" gorm:"comment:重定向"`       // 重定向
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
	MenuType      uint8                             `json:"menu_type" gorm:"comment:路由类型"` // 路由类型
	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                     `json:"children" gorm:"-"`
}

type Meta struct {
	Hidden      bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`           // 是否在列表隐藏
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
	Permissions string `json:"permissions" gorm:"comment:路由权限标识"`       // 路由权限标识
}

type SysBaseMenuParameter struct {
	global.AdpModel
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}
