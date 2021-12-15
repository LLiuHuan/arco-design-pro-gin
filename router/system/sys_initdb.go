// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 11:11
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
)

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	var dbApi = v1.ApiV1GroupApp.System.InitDBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 初始化用户数据库
		initRouter.POST("checkdb", dbApi.CheckDB) // 检测数据库是否初始化
	}
}
