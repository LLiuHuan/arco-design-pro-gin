// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:32
package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	systemReq "github.com/lliuhuan/arco-design-pro-gin/model/system/request"
	systemRes "github.com/lliuhuan/arco-design-pro-gin/model/system/response"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
	"go.uber.org/zap"
)

// Login 用户登录
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login

	if errStr, err := utils.BaseValidator(&l, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	fmt.Println(l)
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global.AdpLog.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			b.tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.AdpConfig.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.AdpLog.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.AdpConfig.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.AdpLog.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.AdpLog.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
