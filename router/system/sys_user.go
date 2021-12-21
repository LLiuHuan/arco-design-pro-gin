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
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	var baseApi = v1.ApiV1GroupApp.System.BaseApi
	{
		userRouter.POST("register", baseApi.Register)                     // 用户注册账号
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
	}
	{
		userRouterWithoutRecord.GET("getUserList", baseApi.GetUserList)
	}
}