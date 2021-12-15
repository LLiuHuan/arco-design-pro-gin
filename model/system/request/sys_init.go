// Package request
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:28
package request

import (
	"fmt"

	"github.com/lliuhuan/arco-design-pro-gin/config"
)

type InitDB struct {
	DBType   string `json:"dbType"`                      // 数据库类型
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 数据库连接端口
	UserName string `json:"userName" binding:"required"` // 数据库用户名
	Password string `json:"password"`                    // 数据库密码
	DBName   string `json:"dbName" binding:"required"`   // 数据库名
}

// MysqlEmptyDsn msyql 空数据库 建库链接
// Author SliverHorn
func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

// ToMysqlConfig 转换 config.Mysql
// @author: [lliuhuan](https://github.com/lliuhuan)
func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		Path:     i.Host,
		Port:     i.Port,
		Dbname:   i.DBName,
		Username: i.UserName,
		Password: i.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
}
