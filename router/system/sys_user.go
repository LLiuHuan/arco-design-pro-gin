// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-12 14:04
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("users").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("users")
	var baseApi = v1.ApiV1GroupApp.System.BaseApi
	{
		userRouter.POST("register", baseApi.Register)              // 用户注册账号
		userRouter.POST("authority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.POST("authorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.PUT("", baseApi.SetUserInfo)                    // 设置用户信息
		userRouter.DELETE("/:id", baseApi.DeleteUser)              // 删除用户
	}
	{
		userRouterWithoutRecord.GET("", baseApi.GetUserList)
		userRouterWithoutRecord.GET("info", baseApi.GetUserInfo) // 获取自身信息
	}
}
