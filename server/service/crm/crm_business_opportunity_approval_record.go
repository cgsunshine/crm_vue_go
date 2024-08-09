package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityApprovalRecordService struct {
}

// CreateCrmBusinessOpportunityApprovalRecord 创建商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService) CreateCrmBusinessOpportunityApprovalRecord(crmBusinessOpportunityApprovalRecord *crm.CrmBusinessOpportunityApprovalRecord) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunityApprovalRecord).Error
	return err
}

// DeleteCrmBusinessOpportunityApprovalRecord 删除商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService)DeleteCrmBusinessOpportunityApprovalRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunityApprovalRecord{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityApprovalRecordByIds 批量删除商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService)DeleteCrmBusinessOpportunityApprovalRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunityApprovalRecord{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunityApprovalRecord 更新商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService)UpdateCrmBusinessOpportunityApprovalRecord(crmBusinessOpportunityApprovalRecord crm.CrmBusinessOpportunityApprovalRecord) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunityApprovalRecord).Error
	return err
}

// GetCrmBusinessOpportunityApprovalRecord 根据ID获取商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService)GetCrmBusinessOpportunityApprovalRecord(ID string) (crmBusinessOpportunityApprovalRecord crm.CrmBusinessOpportunityApprovalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunityApprovalRecord).Error
	return
}

// GetCrmBusinessOpportunityApprovalRecordInfoList 分页获取商机审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalRecordService *CrmBusinessOpportunityApprovalRecordService)GetCrmBusinessOpportunityApprovalRecordInfoList(info crmReq.CrmBusinessOpportunityApprovalRecordSearch) (list []crm.CrmBusinessOpportunityApprovalRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalRecord{})
    var crmBusinessOpportunityApprovalRecords []crm.CrmBusinessOpportunityApprovalRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.BusinessOpportunityId != nil {
        db = db.Where("business_opportunity_id = ?",info.BusinessOpportunityId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmBusinessOpportunityApprovalRecords).Error
	return  crmBusinessOpportunityApprovalRecords, total, err
}
