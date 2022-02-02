// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-20 12:47
package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/errno"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/lliuhuan/arco-design-pro-gin/model/system/response"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BaseMenuService struct {
}

//GetBaseMenuById 返回当前选中menu
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: err error, menu model.SysBaseMenu
func (baseMenuService *BaseMenuService) GetBaseMenuById(id float64) (err error, menu system.SysBaseMenu) {
	err = global.AdpDb.Where("id = ?", id).First(&menu).Error
	return
}

//UpdateBaseMenu 更新路由
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return: err error
func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	upDateMap["redirect"] = menu.Redirect

	err = global.AdpDb.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.AdpLog.Debug("存在相同name修改失败")
				return errno.MenuIdenticalName
			}
		}
		txErr := tx.Unscoped().Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.AdpLog.Debug(txErr.Error())
			return txErr
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			global.AdpLog.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//DeleteBaseMenu 删除基础路由
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error
func (baseMenuService *BaseMenuService) DeleteBaseMenu(id float64) (err error) {
	err = global.AdpDb.Where("parent_id = ?", id).First(&system.SysBaseMenu{}).Error
	if err != nil {
		var menu system.SysBaseMenu
		db := global.AdpDb.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.AdpDb.Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthoritys) > 0 {
			err = global.AdpDb.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errno.MenuExistSubmenu
	}
	return err
}

//DeleteBaseMenus 删除基础路由
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error
func (baseMenuService *BaseMenuService) DeleteBaseMenus(ids []float64) (res response.SysBaseMenuDelete) {
	for _, id := range ids {
		err := baseMenuService.DeleteBaseMenu(id)
		if err != nil {
			res.Error += 1
		} else {
			res.Success += 1
		}
	}
	return res
}
