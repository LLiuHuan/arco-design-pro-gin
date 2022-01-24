// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-18 15:53
package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
)

type OperationRecordRouter struct {
}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("operation/record")
	var authorityMenuApi = v1.ApiV1GroupApp.System.OperationRecordApi
	{
		operationRecordRouter.POST("", authorityMenuApi.CreateSysOperationRecord)              // 新建SysOperationRecord
		operationRecordRouter.DELETE("", authorityMenuApi.DeleteSysOperationRecord)            // 删除SysOperationRecord
		operationRecordRouter.DELETE("/byIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("/:id", authorityMenuApi.FindSysOperationRecord)             // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("", authorityMenuApi.GetSysOperationRecordList)              // 获取SysOperationRecord列表

	}
}
