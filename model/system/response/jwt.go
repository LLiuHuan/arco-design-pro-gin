// Package response
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:38
package response

import uuid "github.com/satori/go.uuid"

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}
