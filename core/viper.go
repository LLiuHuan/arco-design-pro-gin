package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				if gin.Mode() == "debug" {
					config = "config_dev.yaml"
				} else if gin.Mode() == "release" {
					config = "config_pro.yaml"
				}
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用APG_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.AdpConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.AdpConfig); err != nil {
		fmt.Println(err)
	}
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	fmt.Println(filepath.Abs(".."))
	//global.AdpConfig.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.AdpConfig.JWT.ExpiresTime)),
	)
	return v
}
