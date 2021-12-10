//go:build !windows
// +build !windows

// Package core
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 13:09
package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 30 * time.Second
	s.WriteTimeout = 30 * time.Second
	s.MaxHeaderBytes = 1 << 23
	return s
}
