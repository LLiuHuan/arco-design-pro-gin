package service

import (
	"github.com/lliuhuan/arco-design-pro-gin/service/example"
	"github.com/lliuhuan/arco-design-pro-gin/service/system"
)

type ServiceGroup struct {
	System  system.ServiceGroup
	Example example.ServerGroup
}

var ServiceGroupApp = new(ServiceGroup)
