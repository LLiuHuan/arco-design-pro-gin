package v1

import (
	"github.com/lliuhuan/arco-design-pro-gin/api/v1/example"
	"github.com/lliuhuan/arco-design-pro-gin/api/v1/system"
)

type ApiV1Group struct {
	System  system.ApiGroup
	Example example.ApiGroup
}

var ApiV1GroupApp = new(ApiV1Group)
