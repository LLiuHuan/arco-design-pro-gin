package initialize

import (
	"os"

	"github.com/lliuhuan/arco-design-pro-gin/model/system"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.AdpConfig.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysOperationRecord{},
		system.SysBaseMenuParameter{},
	)

	if err != nil {
		global.AdpLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.AdpLog.Info("register table success")
}
