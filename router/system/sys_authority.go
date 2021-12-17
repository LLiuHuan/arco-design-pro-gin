// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-15 17:36
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	var authorityApi = v1.ApiV1GroupApp.System.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
	}
	{
		authorityRouterWithoutRecord.GET("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
