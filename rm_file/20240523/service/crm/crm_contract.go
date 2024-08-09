package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmContractService struct {
}

// CreateCrmContract 创建合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService) CreateCrmContract(crmContract *crm.CrmContract) (err error) {
	err = global.GVA_DB.Create(crmContract).Error
	return err
}

// DeleteCrmContract 删除合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService)DeleteCrmContract(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmContract{},"id = ?",ID).Error
	return err
}

// DeleteCrmContractByIds 批量删除合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService)DeleteCrmContractByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmContract{},"id in ?",IDs).Error
	return err
}

// UpdateCrmContract 更新合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService)UpdateCrmContract(crmContract crm.CrmContract) (err error) {
	err = global.GVA_DB.Save(&crmContract).Error
	return err
}

// GetCrmContract 根据ID获取合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService)GetCrmContract(ID string) (crmContract crm.CrmContract, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmContract).Error
	return
}

// GetCrmContractInfoList 分页获取合同管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractService *CrmContractService)GetCrmContractInfoList(info crmReq.CrmContractSearch) (list []crm.CrmContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmContract{})
    var crmContracts []crm.CrmContract
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmContracts).Error
	return  crmContracts, total, err
}
