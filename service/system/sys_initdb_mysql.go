// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:26
package system

import (
	"fmt"

	"github.com/lliuhuan/arco-design-pro-gin/source/example"

	"github.com/lliuhuan/arco-design-pro-gin/source/system"

	model "github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/lliuhuan/arco-design-pro-gin/model/system/request"

	"github.com/lliuhuan/arco-design-pro-gin/config"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// writeMysqlConfig mysql回写配置
// @author: [lliuhuan](https://github.com/lliuhuan))
func (initDBService *InitDBService) writeMysqlConfig(mysql config.Mysql) error {
	global.AdpConfig.Mysql = mysql
	cs := utils.StructToMap(global.AdpConfig)
	for k, v := range cs {
		global.AdpVp.Set(k, v)
	}
	global.AdpVp.Set("jwt.signing-key", uuid.NewV4().String())
	return global.AdpVp.WriteConfig()
}

// initMsqlDB 创建数据库并初始化 mysql
// @author: [lliuhuan](https://github.com/lliuhuan)
func (initDBService *InitDBService) initMsqlDB(conf request.InitDB) error {
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	} // 创建数据库

	mysqlConfig := conf.ToMysqlConfig()
	if mysqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名, 则跳出初始化数据

	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(), // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: true,              // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.AdpDb = db
	}

	if err := initDBService.initTables(); err != nil {
		global.AdpDb = nil
		return err
	}

	if err := initDBService.initMysqlData(); err != nil {
		global.AdpDb = nil
		return err
	}

	if err := initDBService.writeMysqlConfig(mysqlConfig); err != nil {
		return err
	}

	return nil
}

// initData mysql 初始化数据
// @author: [lliuhuan](https://github.com/lliuhuan)
func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system.Api,
		system.User,
		system.Casbin,
		system.BaseMenu,
		system.Authority,
		system.Dictionary,
		system.UserAuthority,
		system.DataAuthorities,
		system.AuthoritiesMenus,
		system.DictionaryDetail,
		system.ViewAuthorityMenuMysql,
		example.FileMysql,
	)
}
