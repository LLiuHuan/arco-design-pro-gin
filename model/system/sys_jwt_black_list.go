// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:22
// @desc: token 黑名单表
package system

import "github.com/lliuhuan/arco-design-pro-gin/global"

type JwtBlacklist struct {
	global.AdpModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
