// Package middleware
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 14:20
package middleware

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/lliuhuan/arco-design-pro-gin/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"go.uber.org/zap"
)

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.AdpLog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					//c.Error(err.(error)) // nolint: errcheck
					_ = c.Error(errors.New("系统内部错误")) // nolint: errcheck
					response.ResponseAll(http.StatusInternalServerError, gin.H{}, "系统内部错误-1", c)
					c.Abort()
					return
				}

				if stack {
					global.AdpLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
					response.ResponseAll(http.StatusInternalServerError, gin.H{}, "系统内部错误-2", c)
				} else {
					global.AdpLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					response.ResponseAll(http.StatusInternalServerError, gin.H{}, "系统内部错误-3", c)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
