// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-15 17:34
package system

import (
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

type AuthorityApi struct {
}

// GetAuthorityList 分页获取角色列表
// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [get]
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	if errStr, err := utils.BaseValidatorQuery(&pageInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateAuthority 创建角色
// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority system.SysAuthority
	if errStr, err := utils.BaseValidator(&authority, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		if errors.Cause(err) == errno.AuthExist {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// CopyAuthority 拷贝角色
// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	if errStr, err := utils.BaseValidator(&copyInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
		if errors.Cause(err) == errno.AuthExist {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// DeleteAuthority 删除角色
// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	if errStr, err := utils.BaseValidator(&authority, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		if errors.Cause(err) == errno.AuthInUse || errors.Cause(err) == errno.AuthExistSubRole {
			global.AdpLog.Error(err.Error(), zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		global.AdpLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateAuthority 更新角色信息
// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	if errStr, err := utils.BaseValidator(&auth, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		global.AdpLog.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}
