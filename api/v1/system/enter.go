package system

import "github.com/lliuhuan/arco-design-pro-gin/service"

type ApiGroup struct {
	JwtApi
	BaseApi
	InitDBApi
	CasbinApi
	AuthorityMenuApi
}

var userService = service.ServiceGroupApp.System.UserService
var jwtService = service.ServiceGroupApp.System.JwtService
var casbinService = service.ServiceGroupApp.System.CasbinService
var menuService = service.ServiceGroupApp.System.MenuService
var initDBService = service.ServiceGroupApp.System.InitDBService
