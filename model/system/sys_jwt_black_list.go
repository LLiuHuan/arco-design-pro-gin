package system

import "github.com/lliuhuan/arco-design-pro-gin/global"

type JwtBlacklist struct {
	global.AdpModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
