package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmProcurementContractService struct {
}

// CreateCrmProcurementContract 创建订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) CreateCrmProcurementContract(crmProcurementContract *crm.CrmProcurementContract) (err error) {
	err = global.GVA_DB.Create(crmProcurementContract).Error
	return err
}

// DeleteCrmProcurementContract 删除订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) DeleteCrmProcurementContract(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmProcurementContract{}, "id = ?", ID).Error
	return err
}

// GetCrmBillTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) GetCrmBillTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmProcurementContract{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// DeleteCrmProcurementContractByIds 批量删除订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) DeleteCrmProcurementContractByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmProcurementContract{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmProcurementContract 更新订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) UpdateCrmProcurementContract(crmProcurementContract crm.CrmProcurementContract) (err error) {
	err = global.GVA_DB.Save(&crmProcurementContract).Error
	return err
}

// GetCrmProcurementContract 根据ID获取订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) GetCrmProcurementContract(ID string) (crmProcurementContract crm.CrmPageProcurementContract, err error) {
	err = global.GVA_DB.Model(&crm.CrmProcurementContract{}).Where("crm_procurement_contract.id = ?", ID).Select("crm_procurement_contract.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_procurement_contract.user_id").First(&crmProcurementContract).Error
	return
}

// GetCrmProcurementContractInfoList 分页获取订购合同记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) GetCrmProcurementContractInfoList(info crmReq.CrmProcurementContractSearch) (list []crm.CrmProcurementContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmProcurementContract{})
	var crmProcurementContracts []crm.CrmProcurementContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ContractAmount != nil {
		db = db.Where("contract_amount = ?", info.ContractAmount)
	}
	if info.ContractName != "" {
		db = db.Where("contract_name LIKE ?", "%"+info.ContractName+"%")
	}
	if info.ContractStatus != "" {
		db = db.Where("contract_status = ?", info.ContractStatus)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where("creation_time BETWEEN ? AND ? ", info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where("expiration_time BETWEEN ? AND ? ", info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.ProcurementContractNumber != "" {
		db = db.Where("procurement_contract_number = ?", info.ProcurementContractNumber)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Order("created_at DESC").Offset(offset)
	}

	err = db.Find(&crmProcurementContracts).Error
	return crmProcurementContracts, total, err
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmProcurementContractService *CrmProcurementContractService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmProcurementContract{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}
