// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 11:22
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	//menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	var authorityMenuApi = v1.ApiV1GroupApp.System.AuthorityMenuApi
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu) // 获取菜单树
	}
	return menuRouterWithoutRecord
}
