package system

import "github.com/lliuhuan/arco-design-pro-gin/service"

type ApiGroup struct {
	JwtApi
	SysApi
	BaseApi
	InitDBApi
	CasbinApi
	AuthorityApi
	AuthorityMenuApi
}

var sysService = service.ServiceGroupApp.System.SysService
var jwtService = service.ServiceGroupApp.System.JwtService
var userService = service.ServiceGroupApp.System.UserService
var menuService = service.ServiceGroupApp.System.MenuService
var casbinService = service.ServiceGroupApp.System.CasbinService
var initDBService = service.ServiceGroupApp.System.InitDBService
var baseMenuService = service.ServiceGroupApp.System.BaseMenuService
var authorityService = service.ServiceGroupApp.System.AuthorityService
