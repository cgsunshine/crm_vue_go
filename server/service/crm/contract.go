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
	db := global.GVA_DB.Model(&crm.CrmContract{})
	var crmContracts []crm.CrmPageContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartApplicationTime != nil && info.EndApplicationTime != nil {
		db = db.Where("application_time BETWEEN ? AND ? ", info.StartApplicationTime, info.EndApplicationTime)
	}
	if info.ContractName != "" {
		db = db.Where("contract_name LIKE ?", "%"+info.ContractName+"%")
	}
	if info.ContractStatus != "" {
		db = db.Where("contract_status = ?", info.ContractStatus)
	}
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where("expiration_time BETWEEN ? AND ? ", info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.ReviewResult != "" {
		db = db.Where("review_result = ?", info.ReviewResult)
	}
	if info.ReviewStatus != "" {
		db = db.Where("review_status = ?", info.ReviewStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_contract.*,crm_customers.customer_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_contract.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_contract.customer_id").
		Find(&crmContracts).Error
	return crmContracts, total, err
}
