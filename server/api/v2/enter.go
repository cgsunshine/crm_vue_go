package v2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v2/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupAppv2 = new(ApiGroup)
