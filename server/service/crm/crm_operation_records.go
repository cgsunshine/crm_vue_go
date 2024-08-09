package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmOperationRecordsService struct {
}

// CreateCrmOperationRecords 创建操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService) CreateCrmOperationRecords(crmOperationRecords *crm.CrmOperationRecords) (err error) {
	err = global.GVA_DB.Create(crmOperationRecords).Error
	return err
}

// DeleteCrmOperationRecords 删除操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService)DeleteCrmOperationRecords(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmOperationRecords{},"id = ?",ID).Error
	return err
}

// DeleteCrmOperationRecordsByIds 批量删除操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService)DeleteCrmOperationRecordsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmOperationRecords{},"id in ?",IDs).Error
	return err
}

// UpdateCrmOperationRecords 更新操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService)UpdateCrmOperationRecords(crmOperationRecords crm.CrmOperationRecords) (err error) {
	err = global.GVA_DB.Save(&crmOperationRecords).Error
	return err
}

// GetCrmOperationRecords 根据ID获取操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService)GetCrmOperationRecords(ID string) (crmOperationRecords crm.CrmOperationRecords, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmOperationRecords).Error
	return
}

// GetCrmOperationRecordsInfoList 分页获取操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOperationRecordsService *CrmOperationRecordsService)GetCrmOperationRecordsInfoList(info crmReq.CrmOperationRecordsSearch) (list []crm.CrmOperationRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmOperationRecords{})
    var crmOperationRecordss []crm.CrmOperationRecords
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
	
	err = db.Find(&crmOperationRecordss).Error
	return  crmOperationRecordss, total, err
}
