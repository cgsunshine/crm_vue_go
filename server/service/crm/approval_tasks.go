package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"gorm.io/gorm"
)

// SearchCriteria 搜索条件
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) SearchCriteria(info crmReq.CrmApprovalTasksSearch, db *gorm.DB) {

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ApprovalStatus != "" {
		db = db.Where("approval_status = ?", info.ApprovalStatus)
	}
	if info.ApprovalType != nil {
		db = db.Where("approval_type = ?", info.ApprovalType)
	}
	if info.AssigneeId != nil {
		db = db.Where("assignee_id = ?", info.AssigneeId)
	}
	if info.AssociatedId != nil {
		db = db.Where("associated_id = ?", info.AssociatedId)
	}
	if info.RequestId != nil {
		db = db.Where("request_id = ?", info.RequestId)
	}
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 合同
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmContractApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmContactInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmContactInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 商机
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmBusinessOpportunityApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmBusinessOpportunityInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmBusinessOpportunityInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_business_opportunity.business_opportunity_name").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页押金审批任务记录 商机
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmDepositsApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmDepositsInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmDepositsInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_deposits.deposits_name").
		Joins("LEFT JOIN crm_deposits ON crm_deposits.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 回款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmPaymentCollentionApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmPaymentCollentionInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmPaymentCollentionInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_payment_collention.payment_collention_name").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 回款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmOrderApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmOrderInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmOrderInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_order.order_name").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 回款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmStatementAccountApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmStatementAccountInfoApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
	var crmApprovalTaskss []crm.CrmStatementAccountInfoApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks.*,crm_statement_account.statement_account_name").
		Joins("LEFT JOIN crm_statement_account ON crm_statement_account.id = crm_approval_tasks.associated_id").
		Order("crm_approval_tasks.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmPageContractApprovalTasks 根据ID获合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmPageContractApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("id = ?", ID).
		Select("crm_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_approval_tasks.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPageBusinessOpportunityApprovalTasks 根据ID获商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmPageBusinessOpportunityApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("id = ?", ID).
		Select("crm_approval_tasks.*,crm_business_opportunity.business_opportunity_name").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_approval_tasks.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPagePaymentCollentionApprovalTasks 根据ID获回款审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmPagePaymentCollentionApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("id = ?", ID).
		Select("crm_approval_tasks.*,crm_contract.payment_collention_name").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.id = crm_approval_tasks.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPageDepositsApprovalTasks 根据ID获押金审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmPageDepositsApprovalTasks(ID string) (crmDepositsInfoApprovalTasks crm.CrmDepositsInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("id = ?", ID).
		Select("crm_approval_tasks.*,crm_deposits.deposits_name").
		Joins("LEFT JOIN crm_deposits ON crm_deposits.id = crm_approval_tasks.associated_id").
		First(&crmDepositsInfoApprovalTasks).Error
	return
}

//以下都是通用方法

// GetCrmQueryApproved 查询审核是否通过（所有人都要同意）
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmQueryApproved(associatedId int, approvalStatus string) (bool, error) {
	total := int64(0)
	count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})

	db = db.Where("associated_id = ? AND valid = ?", associatedId, comm.Contact_Approval_Tasks_valid_Effective)

	err := db.Count(&total).Error
	if err != nil {
		return false, err
	}

	db = db.Where("approval_status = ?", approvalStatus)

	err = db.Count(&count).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}

	if total == count {
		return true, nil
	}
	return false, nil
}

// GetCrmQueryStepApproved 查询审核是否通过 再此步骤中，可能需要部分人他同意即可通过审批，具体看配置
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) GetCrmQueryStepApproved(associatedId, numberApprovedPersonnel int, approvalType int, approvalStatus string) (bool, error) {
	total := int64(0)
	//count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})

	db = db.Where("associated_id = ? AND valid = ? AND approval_type = ?", associatedId, comm.Contact_Approval_Tasks_valid_Effective, approvalType)
	db = db.Where("approval_status = ?", approvalStatus)

	err := db.Count(&total).Error
	if err != nil {
		return false, err
	}

	if total == 0 {
		return false, nil
	}

	//if total == count {
	//	return true, nil
	//}

	if total >= int64(numberApprovedPersonnel) {
		return true, nil
	}
	return false, nil
}

// UpdCrmApprovalTasks 更新
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) UpdCrmApprovalTasks(ID uint, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("id = ?", ID).Updates(data).Error
	return
}

// UpdCrmAssociatedIdApprovalTasks 更新 通过associatedId关联
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) UpdCrmAssociatedIdApprovalTasks(associatedId, approvalType int, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("associated_id = ? AND approval_type = ?", associatedId, approvalType).Updates(data).Error
	return
}

// DelCrmAssociatedIdApprovalTasks 删除 通过associatedId关联
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) DelCrmAssociatedIdApprovalTasks(associatedId string, approvalType int) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasks{}).Where("associated_id = ? AND approval_type = ?", associatedId, approvalType).Delete(&crm.CrmApprovalTasks{}).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmApprovalTasksService *CrmApprovalTasksService) SplicingQueryConditions(condition string) string {
	return "crm_approval_tasks." + condition
}
