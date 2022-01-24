// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-09 21:18
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	var apiRouterApi = v1.ApiV1GroupApp.System.SystemApiApi
	{
		apiRouter.POST("", apiRouterApi.CreateApi)              // 创建Api
		apiRouter.DELETE("", apiRouterApi.DeleteApi)            // 删除Api
		apiRouter.GET("/:id", apiRouterApi.GetApiById)          // 获取单条Api消息
		apiRouter.PUT("", apiRouterApi.UpdateApi)               // 更新api
		apiRouter.DELETE("byIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.GET("all", apiRouterApi.GetAllApis)   // 获取所有api
		apiRouterWithoutRecord.POST("page", apiRouterApi.GetApiList) // 获取Api列表
	}
}
