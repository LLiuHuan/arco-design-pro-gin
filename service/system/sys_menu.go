// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:57
package system

import (
	"strconv"

	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
)

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

//getMenuTreeMap 获取路由总树map
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu
func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]system.SysMenu) {
	var allMenus []system.SysMenu
	treeMap = make(map[string][]system.SysMenu)
	err = global.AdpDb.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//getChildrenList 获取子菜单
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//getBaseChildrenList 获取菜单的子菜单
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error
func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//getBaseMenuTreeMap 获取路由总树map
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu
func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]system.SysBaseMenu) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.AdpDb.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//GetMenuTree 获取动态菜单树
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []system.SysMenu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//GetInfoList 获取路由分页
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64
func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []system.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//GetBaseMenuTree 获取基础路由树
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu
func (menuService *MenuService) GetBaseMenuTree() (err error, menus []system.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//AddMenuAuthority 为角色增加menu树
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error
func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId string) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//GetMenuAuthority 查看当前角色树
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []system.SysMenu) {
	err = global.AdpDb.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
