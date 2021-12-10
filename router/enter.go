package router

import "github.com/lliuhuan/arco-design-pro-gin/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
