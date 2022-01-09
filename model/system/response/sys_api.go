// Package response
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-01-09 21:33
package response

import "github.com/lliuhuan/arco-design-pro-gin/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
