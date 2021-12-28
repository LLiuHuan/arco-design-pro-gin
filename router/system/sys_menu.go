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
	menuRouter := Router.Group("menus").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menus")
	var authorityMenuApi = v1.ApiV1GroupApp.System.AuthorityMenuApi
	{
		menuRouter.POST("", authorityMenuApi.AddBaseMenu)          // 新增菜单
		menuRouter.PUT("", authorityMenuApi.UpdateBaseMenu)        // 更新菜单
		menuRouter.POST("auth", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系
		menuRouter.DELETE("/:id", authorityMenuApi.DeleteBaseMenu) // 删除菜单
	}
	{
		menuRouterWithoutRecord.GET("user", authorityMenuApi.GetMenu)          // 获取菜单树
		menuRouterWithoutRecord.GET("", authorityMenuApi.GetMenuList)          // 分页获取基础menu列表
		menuRouterWithoutRecord.GET("base", authorityMenuApi.GetBaseMenuTree)  // 获取用户动态路由
		menuRouterWithoutRecord.GET("auth", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouterWithoutRecord.GET("/:id", authorityMenuApi.GetBaseMenuById)  // 根据id获取菜单
	}
}
