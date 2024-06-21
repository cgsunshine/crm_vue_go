package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

// GetCrmNameToConfig 获取审批对应的第一步角色id
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService) GetCrmNameToConfig(config_name string) (crmConfig crm.CrmPageConfig, err error) {
	err = global.GVA_DB.Model(&crm.CrmConfig{}).Where("crm_config.config_name = ? AND nodeOrder = 1", config_name).
		Select("crm_config.*,crm_approval_node.roleIds,crm_approval_node.id as node").
		Joins("LEFT JOIN crm_approval_node ON crm_approval_node.processId = crm_config.associated_id").
		First(&crmConfig).Error
	return
}
