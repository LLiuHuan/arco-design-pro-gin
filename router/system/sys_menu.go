// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 11:22
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	var authorityMenuApi = v1.ApiV1GroupApp.System.AuthorityMenuApi
	{
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu) // 更新菜单
	}
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)                 // 获取菜单树
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)         // 分页获取基础menu列表
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById) // 根据id获取菜单
	}
}
