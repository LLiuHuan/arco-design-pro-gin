package internal

import (
	"fmt"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// @author: [lliuhuan](https://github.com/lliuhuan)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// @author: [lliuhuan](https://github.com/lliuhuan)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.AdpConfig.System.DbType {
	case "mysql":
		logZap = global.AdpConfig.Mysql.LogZap
		if logZap {
			global.AdpLog.Info(fmt.Sprintf(message+"\n", data...))
		} else {
			w.Writer.Printf(message, data...)
		}
	}
}
