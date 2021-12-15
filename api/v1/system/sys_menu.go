// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 11:23
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
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
