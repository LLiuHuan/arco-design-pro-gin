// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 11:23
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	systemRes "github.com/lliuhuan/arco-design-pro-gin/model/system/response"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct {
}

// GetMenu 获取用户动态路由
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	if err, menus := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// GetBaseMenuById 根据id获取菜单
// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	if errStr, err := utils.BaseValidator(&idInfo, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, menu := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// GetMenuList 分页获取基础menu列表
// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	if errStr, err := utils.BaseValidator(&pageInfo, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, menuList, total := menuService.GetInfoList(); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// UpdateBaseMenu 更新菜单
// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	if errStr, err := utils.BaseValidator(&menu, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		global.AdpLog.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
