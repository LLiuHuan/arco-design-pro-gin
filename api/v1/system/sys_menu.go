// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 11:23
package system

import (
	"fmt"
	"net/http"

	"github.com/lliuhuan/arco-design-pro-gin/errno"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	systemReq "github.com/lliuhuan/arco-design-pro-gin/model/system/request"
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
// @Router /menus/user [get]
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
// @Router /menus/:id [get]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	if errStr, err := utils.BaseValidatorUri(&idInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	fmt.Println("idInfo", idInfo)
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
// @Router /menus [get]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	if errStr, err := utils.BaseValidatorQuery(&pageInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
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
// @Router /menus [put]
func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	if errStr, err := utils.BaseValidator(&menu, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		if errors.Cause(err) == errno.MenuIdenticalName {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AddBaseMenu 新增菜单
// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menus [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	if errStr, err := utils.BaseValidator(&menu, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := menuService.AddBaseMenu(menu); err != nil {
		if errors.Cause(err) == errno.MenuIdenticalName {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// DeleteBaseMenu 删除菜单
// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data uri request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menus/:id [delete]
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	if errStr, err := utils.BaseValidatorUri(&menu, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		if errors.Cause(err) == errno.MenuExistSubmenu {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetMenuAuthority 获取指定角色menu
// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menus/auth [get]
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	if errStr, err := utils.BaseValidatorQuery(&param, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// GetBaseMenuTree 获取用户动态路由
// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/base [get]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	if err, menus := menuService.GetBaseMenuTree(); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// AddMenuAuthority 增加menu和角色关联关系
// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/auth [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	if errStr, err := utils.BaseValidator(&authorityMenu, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.AdpLog.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}
