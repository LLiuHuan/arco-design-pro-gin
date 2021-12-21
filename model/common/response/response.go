package response

import (
	"net/http"

	"github.com/lliuhuan/arco-design-pro-gin/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultAll(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		message,
	})
}

func Ok(c *gin.Context) {
	Result(errno.SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(errno.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(errno.SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errno.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(errno.ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(errno.ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errno.ERROR, data, message, c)
}

func FailTokenWithMessage(message string, c *gin.Context) {
	Result(errno.TIMEOUT, map[string]interface{}{}, message, c)
}
