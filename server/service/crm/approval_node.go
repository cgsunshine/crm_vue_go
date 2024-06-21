package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

// GetCrmApprovalNodeId 根据ID获取节点审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService) GetCrmApprovalNodeId(ID int) (crmApprovalNode crm.CrmApprovalNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalNode).Error
	return
}

// CreateCrmBatApprovalNode 批量插入
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService) CreateCrmBatApprovalNode(crmApprovalNode *[]crm.CrmApprovalNode) (err error) {
	err = global.GVA_DB.Create(crmApprovalNode).Error
	return err
}
