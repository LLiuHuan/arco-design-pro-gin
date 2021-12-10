package system

import "github.com/lliuhuan/arco-design-pro-gin/service"

type ApiGroup struct {
	BaseApi
	CasbinApi
}

var casbinService = service.ServiceGroupApp.System.CasbinService
