// Package upload
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 16:04
package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

//UploadFile 上传文件
//@author: [lliuhuan](https://github.com/lliuhuan)
//@object: *TencentCOS
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		global.AdpLog.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.AdpConfig.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.AdpConfig.TencentCOS.BaseURL + "/" + global.AdpConfig.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

//DeleteFile 删除文件
//@author: [lliuhuan](https://github.com/lliuhuan)
//@object: *TencentCOS
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := global.AdpConfig.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.AdpLog.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.AdpConfig.TencentCOS.Bucket + ".cos." + global.AdpConfig.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.AdpConfig.TencentCOS.SecretID,
			SecretKey: global.AdpConfig.TencentCOS.SecretKey,
		},
	})
	return client
}
