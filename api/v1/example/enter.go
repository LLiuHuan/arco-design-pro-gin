// Package example
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-19 01:44
package example

import "github.com/lliuhuan/arco-design-pro-gin/service"

type ApiGroup struct {
	FileApi
}

var fileServer = service.ServiceGroupApp.Example.FileService
