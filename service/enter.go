package service

import "github.com/lliuhuan/arco-design-pro-gin/service/system"

type ServiceGroup struct {
	System system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
