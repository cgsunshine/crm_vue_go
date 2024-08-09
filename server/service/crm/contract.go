package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmContractInfoList 分页获取crmContract表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService) GetCrmPageContractInfoList(info crmReq.CrmContractSearch) (list []crm.CrmPageContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	//db := global.GVA_DB.Model(&crm.CrmContract{})
	db := global.GVA_DB.Model(&crm.CrmContract{})
	var crmContracts []crm.CrmPageContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartApplicationTime != nil && info.EndApplicationTime != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("application_time BETWEEN ? AND ? "), info.StartApplicationTime, info.EndApplicationTime)
	}
	if info.ContractName != "" {
		db = db.Where(crmContractService.SplicingQueryConditions("contract_name LIKE ?"), "%"+info.ContractName+"%")
	}
	if info.ContractStatus != "" {
		db = db.Where(crmContractService.SplicingQueryConditions("contract_status = ?"), info.ContractStatus)
	}
	if info.CustomerId != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("expiration_time BETWEEN ? AND ? "), info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.OrderId != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("order_id = ?"), info.OrderId)
	}
	if info.ReviewResult != "" {
		db = db.Where(crmContractService.SplicingQueryConditions("review_result = ?"), info.ReviewResult)
	}
	if info.ReviewStatus != "" {
		db = db.Where(crmContractService.SplicingQueryConditions("review_status = ?"), info.ReviewStatus)
	}
	if info.UserId != nil {
		db = db.Where(crmContractService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_contract.*,crm_customers.customer_name,sys_users.username,crm_order.order_name,crm_contract_type.contract_type_name,crm_order.amount,crm_order.currency").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_contract.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_contract.customer_id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_contract.order_id").
		Joins("LEFT JOIN crm_contract_type ON crm_contract_type.id = crm_contract.contract_type_id").
		Order("crm_contract.created_at DESC").
		Find(&crmContracts).Error
	return crmContracts, total, err
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmContract{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmPageContract 根据ID获取crmContract表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService) GetCrmPageContract(ID string) (crmContract crm.CrmPageContract, err error) {
	err = global.GVA_DB.Model(&crm.CrmContract{}).
		Select("crm_contract.*,crm_customers.customer_name,sys_users.username,crm_order.order_name,crm_contract_type.contract_type_name,crm_order.amount,crm_order.currency").
		Where("crm_contract.id = ?", ID).
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_contract.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_contract.customer_id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_contract.order_id").
		Joins("LEFT JOIN crm_contract_type ON crm_contract_type.id = crm_contract.contract_type_id").
		First(&crmContract).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmContractService *CrmContractService) SplicingQueryConditions(condition string) string {
	return "crm_contract." + condition
}
