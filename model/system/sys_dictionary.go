// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:47
// @desc: 字典表
package system

import "github.com/lliuhuan/arco-design-pro-gin/global"

// 如果含有time.Time 请自行import time包
type SysDictionary struct {
	global.AdpModel
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`   // 字典名（英）
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`       // 描述
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}
