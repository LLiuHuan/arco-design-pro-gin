// Package errno
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 14:39
package errno

import (
	"github.com/pkg/errors"
)

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

	MenuExistSubmenu  = errors.New("此菜单存在子菜单不可删除")
	MenuIdenticalName = errors.New("存在相同name")

	AuthExist        = errors.New("存在相同角色id")
	AuthExistSubRole = errors.New("此角色存在子角色不允许删除")
	AuthInUse        = errors.New("此角色有用户正在使用禁止删除")
)

const (
	ERROR   = -1
	SUCCESS = 0
	TIMEOUT = 10001
)
