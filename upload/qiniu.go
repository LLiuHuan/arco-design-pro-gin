// Package upload
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 15:50
package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

//UploadFile 上传文件
//@author: [lliuhuan](https://github.com/lliuhuan)
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.AdpConfig.Qiniu.Bucket}
	mac := qbox.NewMac(global.AdpConfig.Qiniu.AccessKey, global.AdpConfig.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.AdpLog.Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.AdpLog.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.AdpConfig.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//DeleteFile 删除文件
//@author: [lliuhuan](https://github.com/lliuhuan)
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error
func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.AdpConfig.Qiniu.AccessKey, global.AdpConfig.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.AdpConfig.Qiniu.Bucket, key); err != nil {
		global.AdpLog.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

//qiniuConfig 根据配置文件进行返回七牛云的配置
//@author: [lliuhuan](https://github.com/lliuhuan)
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config
func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.AdpConfig.Qiniu.UseHTTPS,
		UseCdnDomains: global.AdpConfig.Qiniu.UseCdnDomains,
	}
	switch global.AdpConfig.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
