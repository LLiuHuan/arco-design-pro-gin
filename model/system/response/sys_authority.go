// Package response
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:55
package response

import "github.com/lliuhuan/arco-design-pro-gin/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"oldAuthorityId"` // 旧角色ID
}
