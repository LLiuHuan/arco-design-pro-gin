// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:52
package system

import (
	"fmt"
	"strings"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
)

var ViewAuthorityMenuMysql = new(viewAuthorityMenuMysql)

type viewAuthorityMenuMysql struct{}

func (v *viewAuthorityMenuMysql) TableName() string {
	var entity system.SysMenu
	return entity.TableName()
}

func (v *viewAuthorityMenuMysql) Initialize() error {
	var entity AuthorityMenus
	sql := `
	CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW @table_name AS
	select @menus.id                AS id,
		   @menus.path              AS path,
		   @menus.icon              AS icon,
		   @menus.name              AS name,
		   @menus.sort              AS sort,
		   @menus.title             AS title,
		   @menus.hidden            AS hidden,
		   @menus.permissions       AS permissions,
		   @menus.menu_type         AS menu_type,
		   @menus.component         AS component,
		   @menus.parent_id         AS parent_id,
		   @menus.created_at        AS created_at,
		   @menus.updated_at        AS updated_at,
		   @menus.deleted_at        AS deleted_at,
		   @menus.keep_alive        AS keep_alive,
		   @menus.menu_level        AS menu_level,
		   @menus.default_menu      AS default_menu,
           @menus.redirect          AS redirect,
		   @authorities_menus.sys_base_menu_id      AS menu_id,
		   @authorities_menus.sys_authority_authority_id AS authority_id
	from (@authorities_menus
			 join @menus on ((@authorities_menus.sys_base_menu_id = @menus.id)));
	`
	sql = strings.ReplaceAll(sql, "@table_name", v.TableName())
	sql = strings.ReplaceAll(sql, "@menus", "sys_base_menus")
	sql = strings.ReplaceAll(sql, "@authorities_menus", entity.TableName())
	if err := global.AdpDb.Exec(sql).Error; err != nil {
		return errors.Wrap(err, v.TableName()+"视图创建失败!")
	}
	return nil
}

func (v *viewAuthorityMenuMysql) CheckDataExist() bool {
	err1 := global.AdpDb.Find(&[]system.SysMenu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist", global.AdpConfig.Mysql.Dbname, v.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
