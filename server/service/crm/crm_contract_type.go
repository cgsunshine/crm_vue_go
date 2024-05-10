package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmContractTypeService struct {
}

// CreateCrmContractType 创建合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService) CreateCrmContractType(crmContractType *crm.CrmContractType) (err error) {
	err = global.GVA_DB.Create(crmContractType).Error
	return err
}

// DeleteCrmContractType 删除合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService)DeleteCrmContractType(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmContractType{},"id = ?",ID).Error
	return err
}

// DeleteCrmContractTypeByIds 批量删除合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService)DeleteCrmContractTypeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmContractType{},"id in ?",IDs).Error
	return err
}

// UpdateCrmContractType 更新合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService)UpdateCrmContractType(crmContractType crm.CrmContractType) (err error) {
	err = global.GVA_DB.Save(&crmContractType).Error
	return err
}

// GetCrmContractType 根据ID获取合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService)GetCrmContractType(ID string) (crmContractType crm.CrmContractType, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmContractType).Error
	return
}

// GetCrmContractTypeInfoList 分页获取合同类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContractTypeService *CrmContractTypeService)GetCrmContractTypeInfoList(info crmReq.CrmContractTypeSearch) (list []crm.CrmContractType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmContractType{})
    var crmContractTypes []crm.CrmContractType
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
	
	err = db.Find(&crmContractTypes).Error
	return  crmContractTypes, total, err
}
