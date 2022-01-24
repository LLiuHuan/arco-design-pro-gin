// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-09 21:20
package system

import (
	"fmt"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

// CreateApi 新增基础api
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error
func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.AdpDb.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.AdpDb.Create(&api).Error
}

// DeleteApi 删除基础api
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	err = global.AdpDb.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetAPIInfoList 分页获取数据
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetAPIInfoList
//@description: 分页获取数据
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error
func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.AdpDb.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

// GetAllApis 获取所有的api
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi
func (apiService *ApiService) GetAllApis() (err error, apis []system.SysApi) {
	err = global.AdpDb.Find(&apis).Error
	for i, api := range apis {
		apis[i].Key = fmt.Sprintf("p:%sm:%s", api.Path, api.Method)
	}
	return
}

// GetApiById 根据id获取api
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi
func (apiService *ApiService) GetApiById(id float64) (err error, api system.SysApi) {
	err = global.AdpDb.Where("id = ?", id).First(&api).Error
	return
}

// UpdateApi 根据id更新api
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.AdpDb.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.AdpDb.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.AdpDb.Save(&api).Error
		}
	}
	return err
}

// DeleteApisByIds 删除选中API
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.AdpDb.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.AdpDb.Delete(&system.SysApi{}, "id in ?", ids).Error
}
