// Package example
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-19 01:50
package example

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/example"
)

type FileService struct {
}

//GetFileRecordInfoList 分页获取数据
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func (e *FileService) GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.AdpDb.Model(&example.ExaFileUploadAndDownload{})
	var fileLists []example.ExaFileUploadAndDownload
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}
