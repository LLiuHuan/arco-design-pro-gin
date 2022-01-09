// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-07 16:50
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type SysRouter struct {
}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	var systemApi = v1.ApiV1GroupApp.System.SysApi
	{
		sysRouter.GET("info", systemApi.GetServerInfo)   // 获取服务器信息
		sysRouter.POST("reload", systemApi.ReloadSystem) // 重启服务
	}
}
