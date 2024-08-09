package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmPageProcurementContract 根据ID获取订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) GetCrmPageProcurementContract(ID string) (crmProcurementContract crm.CrmProcurementContract, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmProcurementContract).Error
	return
}

// GetCrmPageProcurementContractInfoList 分页获取订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) GetCrmPageProcurementContractInfoList(info crmReq.CrmProcurementContractSearch) (list []crm.CrmPageProcurementContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmProcurementContract{})
	var crmProcurementContracts []crm.CrmPageProcurementContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ContractAmount != nil {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("contract_amount = ?"), info.ContractAmount)
	}
	if info.ContractName != "" {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("contract_name LIKE ?"), "%"+info.ContractName+"%")
	}
	if info.ContractStatus != "" {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("contract_status = ?"), info.ContractStatus)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("creation_time BETWEEN ? AND ? "), info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("expiration_time BETWEEN ? AND ? "), info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.UserId != nil {
		db = db.Where(crmProcurementContractService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Order("created_at DESC").Offset(offset)
	}

	err = db.Select("crm_procurement_contract.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_procurement_contract.user_id").
		Find(&crmProcurementContracts).Error
	return crmProcurementContracts, total, err
}

// SplicingQueryConditions 拼接条件
func (crmProcurementContractService *CrmProcurementContractService) SplicingQueryConditions(condition string) string {
	return "crm_procurement_contract." + condition
}
