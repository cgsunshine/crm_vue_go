package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
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

// GetCrmApprovalNodeInfoList 分页获取审批节点记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService) GetCrmPageApprovalNodeInfoList(info crmReq.CrmApprovalNodeSearch) (list []crm.CrmPageApprovalNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalNode{})
	var crmApprovalNodes []crm.CrmPageApprovalNode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProcessId != nil {
		db = db.Where("processId = ?", info.ProcessId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Debug()

	err = db.Select("crm_approval_node.*,sys_authorities.authority_name").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_node.roleIds").
		Find(&crmApprovalNodes).
		Error

	return crmApprovalNodes, total, err
}
