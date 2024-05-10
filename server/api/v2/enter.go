package v2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/crm"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CrmApiGroup     crm.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
