// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-12 02:29
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	jwtRouter := Router.Group("jwt")
	var jwtApi = v1.ApiV1GroupApp.System.JwtApi
	{
		jwtRouter.POST("black", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
	return jwtRouter
}
