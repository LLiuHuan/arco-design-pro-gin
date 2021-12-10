// Package upload
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 16:07
package upload

import (
	"mime/multipart"

	"github.com/lliuhuan/arco-design-pro-gin/global"
)

//OSS OSS接口
//@author: [lliuhuan](https://github.com/lliuhuan)
//@interface_name: OSS
//@description: OSS接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

//NewOss 创建oss对象
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS
func NewOss() OSS {
	switch global.AdpConfig.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &Local{}
	}
}
