// Package response
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-13 12:45
package response

import "github.com/lliuhuan/arco-design-pro-gin/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}
