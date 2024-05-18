package routerv2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/routerv2/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
