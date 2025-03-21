package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"gorm.io/gorm"
	"time"
)

// SearchCriteria 搜索条件
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) SearchCriteria(info crmReq.CrmApprovalTasksRoleSearch, db *gorm.DB) {

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ApprovalStatus != "" {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("approval_status = ?"), info.ApprovalStatus)
	}
	if info.ApprovalType != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("approval_type = ?"), info.ApprovalType)
	}
	if info.AssigneeId != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("assignee_id = ?"), info.AssigneeId)
	}
	if info.AssociatedId != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("associated_id = ?"), info.AssociatedId)
	}
	if info.RequestId != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("request_id = ?"), info.RequestId)
	}
	if info.RoleId != nil {
		db = db.Where(crmApprovalTasksRoleService.SplicingQueryConditions("role_id = ?"), info.RoleId)
	}
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 合同
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmContractApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmContactInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmContactInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_contract.contract_name,sys_authorities.authority_name,crm_contract.contract_number,sys_users.username").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 商机
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmBusinessOpportunityApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmBusinessOpportunityInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmBusinessOpportunityInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_business_opportunity.business_opportunity_name,crm_business_opportunity.business_opportunity_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页押金审批任务记录 商机
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmDepositsApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmDepositsInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmDepositsInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_deposits.deposits_name,crm_deposits.deposits_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_deposits ON crm_deposits.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 回款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPaymentCollentionApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmPaymentCollentionInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmPaymentCollentionInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_payment_collention.payment_collention_name,crm_payment_collention.payment_collention_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 回款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmOrderApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmOrderInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmOrderInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_order.order_name,crm_order.order_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 对账单
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmStatementAccountApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmStatementAccountInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmStatementAccountInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_statement_account.statement_account_name,crm_statement_account.statement_account_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_statement_account ON crm_statement_account.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 付款
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPaymentApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmPaymentInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmPaymentInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_payment.payment_name,crm_payment.payment_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_payment ON crm_payment.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 订购单
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPurchaseOrderApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmPurchaseOrderInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmPurchaseOrderInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_purchase_order.purchase_order_name,crm_purchase_order.purchase_order_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_purchase_order ON crm_purchase_order.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmApprovalTasksInfoList 分页获取审批任务记录 订购合同
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmProcurementContractApprovalTasksInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmProcurementContractInfoApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTaskss []crm.CrmProcurementContractInfoApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	crmApprovalTasksRoleService.SearchCriteria(info, db)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_approval_tasks_role.*,crm_procurement_contract.contract_name,crm_procurement_contract.procurement_contract_number,sys_authorities.authority_name,sys_users.username").
		Joins("LEFT JOIN crm_procurement_contract ON crm_procurement_contract.id = crm_approval_tasks_role.associated_id").
		Joins("LEFT JOIN sys_authorities ON sys_authorities.authority_id = crm_approval_tasks_role.role_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_approval_tasks_role.assignee_id").
		Order("crm_approval_tasks_role.created_at DESC").
		Find(&crmApprovalTaskss).Error
	return crmApprovalTaskss, total, err
}

// GetCrmPageContractApprovalTasks 根据ID获合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPageContractApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("id = ?", ID).
		Select("crm_approval_tasks_role.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_approval_tasks_role.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPageBusinessOpportunityApprovalTasks 根据ID获商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPageBusinessOpportunityApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("id = ?", ID).
		Select("crm_approval_tasks_role.*,crm_business_opportunity.business_opportunity_name").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_approval_tasks_role.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPagePaymentCollentionApprovalTasks 根据ID获回款审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPagePaymentCollentionApprovalTasks(ID string) (CrmContactInfoApprovalTasks crm.CrmContactInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("id = ?", ID).
		Select("crm_approval_tasks_role.*,crm_contract.payment_collention_name").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.id = crm_approval_tasks_role.associated_id").
		First(&CrmContactInfoApprovalTasks).Error
	return
}

// GetCrmPageDepositsApprovalTasks 根据ID获押金审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmPageDepositsApprovalTasks(ID string) (crmDepositsInfoApprovalTasks crm.CrmDepositsInfoApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("id = ?", ID).
		Select("crm_approval_tasks_role.*,crm_deposits.deposits_name").
		Joins("LEFT JOIN crm_deposits ON crm_deposits.id = crm_approval_tasks_role.associated_id").
		First(&crmDepositsInfoApprovalTasks).Error
	return
}

//以下都是通用方法

// GetCrmQueryApproved 查询审核是否通过（所有人都要同意）
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmQueryApproved(associatedId int, approvalStatus string) (bool, error) {
	total := int64(0)
	count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})

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
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmQueryStepApproved(associatedId, numberApprovedPersonnel int, approvalType int, approvalStatus string) (bool, error) {
	total := int64(0)
	//count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})

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
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) UpdCrmApprovalTasks(ID uint, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("id = ?", ID).Updates(data).Error
	return
}

// UpdCrmAssociatedIdApprovalTasks 更新 通过associatedId关联
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) UpdCrmAssociatedIdApprovalTasks(associatedId, approvalType int, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("associated_id = ? AND approval_type = ?", associatedId, approvalType).Updates(data).Error
	return
}

// DelCrmAssociatedIdApprovalTasks 删除 通过associatedId关联
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) DelCrmAssociatedIdApprovalTasks(associatedId string, approvalType int) (err error) {
	err = global.GVA_DB.Model(&crm.CrmApprovalTasksRole{}).Where("associated_id = ? AND approval_type = ?", associatedId, approvalType).Delete(&crm.CrmApprovalTasksRole{}).Error
	return
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) ApprovalTasksCount(assigneeId *int, approvalType int, approvalStatus string, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	SearchCondition(db, nil, startDate, endDate)
	err = db.Where("assignee_id = ? AND  approval_type = ? AND valid = 1 AND approval_status = ? ", assigneeId, approvalType, approvalStatus).Count(&total).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) SplicingQueryConditions(condition string) string {
	return "crm_approval_tasks_role." + condition
}
