package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

// GetCrmNameToConfig 获取审批对应的第一步角色id
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService) GetCrmNameToConfig(config_name string) (crmConfig crm.CrmPageConfig, err error) {
	err = global.GVA_DB.Model(&crm.CrmConfig{}).Where("crm_config.config_name = ?", config_name).
		Select("crm_config.*,crm_approval_node.roleIds,crm_approval_node.id as node").
		Joins("LEFT JOIN crm_approval_node ON crm_approval_node.processId = crm_config.associated_id").
		Order("crm_approval_node.nodeOrder ASC").
		First(&crmConfig).Error
	return
}

// GetCrmNameToConfig 获取审批对应的下一步角色id
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService) GetCrmNextNameToConfig(config_name string, step *int) (crmConfig crm.CrmPageConfig, err error) {
	err = global.GVA_DB.Model(&crm.CrmConfig{}).Where("crm_config.config_name = ? AND crm_approval_node.nodeOrder > ?", config_name, step).
		Select("crm_config.*,crm_approval_node.roleIds,crm_approval_node.id as node").
		Joins("LEFT JOIN crm_approval_node ON crm_approval_node.processId = crm_config.associated_id").
		Order("crm_approval_node.nodeOrder ASC").
		First(&crmConfig).Error
	return
}

// GetCrmNameToConfig 查询审批节点是否是最后一步
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService) GetCrmLastNameToConfig(config_name string, step *int) bool {
	var crmConfig crm.CrmPageConfig
	err := global.GVA_DB.Model(&crm.CrmConfig{}).Where("crm_config.config_name = ?", config_name).
		Select("crm_config.*,crm_approval_node.roleIds,crm_approval_node.id as node,crm_approval_node.nodeOrder").
		Joins("LEFT JOIN crm_approval_node ON crm_approval_node.processId = crm_config.associated_id").
		Order("crm_approval_node.nodeOrder DESC").
		First(&crmConfig).Error

	if err != nil {
		fmt.Println(err)
		return false
	}

	//fmt.Println(*crmConfig.NodeOrder == *step)

	if *crmConfig.NodeOrder == *step {
		return true
	}

	return false
}
