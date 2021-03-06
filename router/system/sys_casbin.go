package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
	"github.com/lliuhuan/arco-design-pro-gin/middleware"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	var casbinApi = v1.ApiV1GroupApp.System.CasbinApi
	{
		casbinRouter.PUT("", casbinApi.UpdateCasbin)
	}
	{
		// TODO: 这个用到的时候再修改
		casbinRouterWithoutRecord.GET("/:id", casbinApi.GetPolicyPathByAuthorityId)
	}
}
