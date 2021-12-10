package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/core"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/initialize"
)

func init() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath
func main() {
	global.AdpVp = core.Viper()      // 初始化Viper
	global.AdpLog = core.Zap()       // 初始化zap日志
	global.AdpDb = initialize.Gorm() // gorm 链接数据库
	if global.AdpDb != nil {
		initialize.RegisterTables(global.AdpDb) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.AdpDb.DB()
		defer db.Close()
	}
	if err := initialize.InitTrans("zh"); err != nil {
		er := fmt.Sprintf("init trans failed, err:%v\n", err)
		panic(er)
	}

	fmt.Println(gin.Mode())
	core.RunServer()
}
