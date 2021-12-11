package system

import "github.com/lliuhuan/arco-design-pro-gin/service"

type ApiGroup struct {
	BaseApi
	CasbinApi
}

var userService = service.ServiceGroupApp.System.UserService
var jwtService = service.ServiceGroupApp.System.JwtService
var casbinService = service.ServiceGroupApp.System.CasbinService
