// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:22
// @desc: 初始化数据表
package system

import (
	"github.com/gookit/color"
)

type InitDBFunc interface {
	Init() (err error)
}

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	AuthorityMenu   = "\n[%v] --> %v 视图已存在!\n"
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 表初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!\n"
)

// MysqlDataInitialize Mysql 初始化接口使用封装
// @author: [lliuhuan](https://github.com/lliuhuan)
func MysqlDataInitialize(inits ...InitData) error {
	var entity SysMenu
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == entity.TableName() {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, Mysql, inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, Mysql, inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, err)
			return err
		} else {
			color.Info.Printf(InitDataSuccess, Mysql, inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
