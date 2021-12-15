// Package example
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 11:06
package example

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
)

type ExaCustomer struct {
	global.AdpModel
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:客户名"`                // 客户名
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`    // 客户手机号
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                     // 管理ID
	SysUserAuthorityID string         `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"` // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`                         // 管理详情
}
