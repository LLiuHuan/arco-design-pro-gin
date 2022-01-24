// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-09 21:25
package system

import (
	"net/http"

	systemReq "github.com/lliuhuan/arco-design-pro-gin/model/system/request"

	systemRes "github.com/lliuhuan/arco-design-pro-gin/model/system/response"

	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
	"go.uber.org/zap"
)

type SystemApiApi struct {
}

// CreateApi 创建基础api
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/api [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		global.AdpLog.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApi 删除api
// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/api [delete]
func (s *SystemApiApi) DeleteApi(c *gin.Context) {
	var api system.SysApi
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := apiService.DeleteApi(api); err != nil {
		global.AdpLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetApiList 分页获取API列表
// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (s *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	if errStr, err := utils.BaseValidator(&pageInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err, list, total := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetApiById 根据id获取api
// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/:id [get]
func (s *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	if errStr, err := utils.BaseValidatorUri(&idInfo, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	err, api := apiService.GetApiById(idInfo.ID)
	if err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(systemRes.SysAPIResponse{Api: api}, c)
	}
}

// UpdateApi 创建基础api
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/api [put]
func (s *SystemApiApi) UpdateApi(c *gin.Context) {
	var api system.SysApi
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := apiService.UpdateApi(api); err != nil {
		global.AdpLog.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// GetAllApis 获取所有的Api 不分页
// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/api/all [get]
func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	if err, apis := apiService.GetAllApis(); err != nil {
		global.AdpLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "获取成功", c)
	}
}

// DeleteApisByIds 删除选中Api
// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/byIds [delete]
func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	if errStr, err := utils.BaseValidator(&ids, c); err != nil {
		response.FailCodeMessage(http.StatusBadRequest, errStr, c)
		return
	}
	if err := apiService.DeleteApisByIds(ids); err != nil {
		global.AdpLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
