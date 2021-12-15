// Package errno
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 14:39
package errno

import "github.com/pkg/errors"

var (
	// OK = &Errno{Code: 0, Message: "OK"}

	// TokenExpired token 已过期
	TokenExpired = errors.New("token is expired")
	// TokenNotValidYet token 未激活
	TokenNotValidYet = errors.New("token not active yet")
	// TokenMalformed 不是token
	TokenMalformed = errors.New("that's not even a token")
	// TokenInvalid token 无法处理
	TokenInvalid = errors.New("couldn't handle this token: ")
)

const (
	ERROR   = -1
	SUCCESS = 0
	TIMEOUT = 10001
)
