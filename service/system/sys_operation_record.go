package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	systemReq "github.com/lliuhuan/arco-design-pro-gin/model/system/request"
)

type OperationRecordService struct {
}

// CreateSysOperationRecord 创建接口请求信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err erro
func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.AdpDb.Create(&sysOperationRecord).Error
	return err
}

// DeleteSysOperationRecordByIds 批量删除记录
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error
func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.AdpDb.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// DeleteSysOperationRecord 删除操作记录
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error
func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.AdpDb.Delete(&sysOperationRecord).Error
	return err
}

// GetSysOperationRecord 根据id获取单条操作记录
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord model.SysOperationRecord
func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (err error, sysOperationRecord system.SysOperationRecord) {
	err = global.AdpDb.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info systemReq.SysOperationRecordSearch
//@return: err error, list interface{}, total int64
func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info systemReq.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.AdpDb.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
