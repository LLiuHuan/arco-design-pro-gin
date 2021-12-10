package core

import (
	"fmt"
	"time"

	"github.com/lliuhuan/arco-design-pro-gin/service/system"

	"go.uber.org/zap"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.AdpConfig.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.AdpDb != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%s", global.AdpConfig.System.Port)

	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.AdpLog.Info("server run success on ", zap.String("address: ", address))

	fmt.Printf(`
                                          _               _                                         
                                         | |             (_)                                        
      ____   ____   ____   ___   ___   _ | |  ____   ___  _   ____  ____   ___  ____    ____   ___  
     / _  | / ___) / ___) / _ \ (___) / || | / _  ) /___)| | / _  ||  _ \ (___)|  _ \  / ___) / _ \ 
    ( ( | || |    ( (___ | |_| |     ( (_| |( (/ / |___ || |( ( | || | | |     | | | || |    | |_| |
     \_||_||_|     \____) \___/       \____| \____)(___/ |_| \_|| ||_| |_|     | ||_/ |_|     \___/ 
                                                            (_____|            |_|

	欢迎使用 github.com/lliuhuan/arco-design-pro
	当前版本:V0.1.1 Beta
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认后端接口运行地址:http://127.0.0.1%s
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address, address)
	global.AdpLog.Error(s.ListenAndServe().Error())
}
