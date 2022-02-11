// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:25
package system

import (
	"database/sql"
	"fmt"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system/request"

	"github.com/lliuhuan/arco-design-pro-gin/model/system"
)

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
// @author: [lliuhuan](https://github.com/lliuhuan)
func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMsqlDB(conf)
	default:
		return initDBService.initMsqlDB(conf)
	}
}

// initTables 初始化表
// @author: [lliuhuan](https://github.com/lliuhuan)
func (initDBService *InitDBService) initTables() error {
	return global.AdpDb.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.JwtBlacklist{},
		//system.SysDictionary{},
		system.SysOperationRecord{},
		//system.SysDictionaryDetail{},
		//system.SysBaseMenuParameter{},

		adapter.CasbinRule{},

		//example.ExaFile{},
		//example.ExaCustomer{},
		//example.ExaFileChunk{},
		//example.ExaFileUploadAndDownload{},
	)

}

// createDatabase 创建数据库(mysql)
// @author: [lliuhuan](https://github.com/lliuhuan)
func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
