package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmApprovalRecordService struct {
}

// CreateCrmApprovalRecord 创建crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService) CreateCrmApprovalRecord(crmApprovalRecord *crm.CrmApprovalRecord) (err error) {
	err = global.GVA_DB.Create(crmApprovalRecord).Error
	return err
}

// DeleteCrmApprovalRecord 删除crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService)DeleteCrmApprovalRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmApprovalRecord{},"id = ?",ID).Error
	return err
}

// DeleteCrmApprovalRecordByIds 批量删除crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService)DeleteCrmApprovalRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmApprovalRecord{},"id in ?",IDs).Error
	return err
}

// UpdateCrmApprovalRecord 更新crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService)UpdateCrmApprovalRecord(crmApprovalRecord crm.CrmApprovalRecord) (err error) {
	err = global.GVA_DB.Save(&crmApprovalRecord).Error
	return err
}

// GetCrmApprovalRecord 根据ID获取crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService)GetCrmApprovalRecord(ID string) (crmApprovalRecord crm.CrmApprovalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalRecord).Error
	return
}

// GetCrmApprovalRecordInfoList 分页获取crmApprovalRecord表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalRecordService *CrmApprovalRecordService)GetCrmApprovalRecordInfoList(info crmReq.CrmApprovalRecordSearch) (list []crm.CrmApprovalRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalRecord{})
    var crmApprovalRecords []crm.CrmApprovalRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Code != "" {
        db = db.Where("code = ?",info.Code)
    }
    if info.ModuleId != "" {
        db = db.Where("moduleId LIKE ?","%"+ info.ModuleId+"%")
    }
        if info.StartCreateTime != nil && info.EndCreateTime != nil {
            db = db.Where("create_time BETWEEN ? AND ? ",info.StartCreateTime,info.EndCreateTime)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmApprovalRecords).Error
	return  crmApprovalRecords, total, err
}
