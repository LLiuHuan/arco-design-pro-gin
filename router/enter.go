package router

import (
	"github.com/lliuhuan/arco-design-pro-gin/router/example"
	"github.com/lliuhuan/arco-design-pro-gin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
