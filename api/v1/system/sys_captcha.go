// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 13:48
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"
	systemRes "github.com/lliuhuan/arco-design-pro-gin/model/system/response"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct {
}

// Captcha 生成验证码
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.AdpConfig.Captcha.ImgHeight,
		global.AdpConfig.Captcha.ImgWidth,
		global.AdpConfig.Captcha.KeyLong,
		0.7,
		80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.AdpLog.Error("获取验证码失败！", zap.Error(err))
		response.FailWithMessage("获取验证码失败！", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.AdpConfig.Captcha.KeyLong,
		}, "获取验证码成功！", c)
	}
}
