// Package example
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-19 01:42
package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lliuhuan/arco-design-pro-gin/api/v1"
)

type FileRouter struct {
}

func (e *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	var exaFileApi = v1.ApiV1GroupApp.Example.FileApi
	{
		fileRouter.POST("getFileList", exaFileApi.GetFileList) // 获取上传文件列表
	}
}
