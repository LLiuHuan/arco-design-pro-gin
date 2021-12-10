package response

import "github.com/lliuhuan/arco-design-pro-gin/model/system/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
