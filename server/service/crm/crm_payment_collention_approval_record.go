package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPaymentCollentionApprovalRecordService struct {
}

// CreateCrmPaymentCollentionApprovalRecord 创建回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService) CreateCrmPaymentCollentionApprovalRecord(crmPaymentCollentionApprovalRecord *crm.CrmPaymentCollentionApprovalRecord) (err error) {
	err = global.GVA_DB.Create(crmPaymentCollentionApprovalRecord).Error
	return err
}

// DeleteCrmPaymentCollentionApprovalRecord 删除回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService)DeleteCrmPaymentCollentionApprovalRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPaymentCollentionApprovalRecord{},"id = ?",ID).Error
	return err
}

// DeleteCrmPaymentCollentionApprovalRecordByIds 批量删除回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService)DeleteCrmPaymentCollentionApprovalRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPaymentCollentionApprovalRecord{},"id in ?",IDs).Error
	return err
}

// UpdateCrmPaymentCollentionApprovalRecord 更新回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService)UpdateCrmPaymentCollentionApprovalRecord(crmPaymentCollentionApprovalRecord crm.CrmPaymentCollentionApprovalRecord) (err error) {
	err = global.GVA_DB.Save(&crmPaymentCollentionApprovalRecord).Error
	return err
}

// GetCrmPaymentCollentionApprovalRecord 根据ID获取回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService)GetCrmPaymentCollentionApprovalRecord(ID string) (crmPaymentCollentionApprovalRecord crm.CrmPaymentCollentionApprovalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPaymentCollentionApprovalRecord).Error
	return
}

// GetCrmPaymentCollentionApprovalRecordInfoList 分页获取回款审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalRecordService *CrmPaymentCollentionApprovalRecordService)GetCrmPaymentCollentionApprovalRecordInfoList(info crmReq.CrmPaymentCollentionApprovalRecordSearch) (list []crm.CrmPaymentCollentionApprovalRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmPaymentCollentionApprovalRecord{})
    var crmPaymentCollentionApprovalRecords []crm.CrmPaymentCollentionApprovalRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ApproverId != nil {
        db = db.Where("approver_id = ?",info.ApproverId)
    }
    if info.PaymentCollentionId != nil {
        db = db.Where("payment_collention_id = ?",info.PaymentCollentionId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmPaymentCollentionApprovalRecords).Error
	return  crmPaymentCollentionApprovalRecords, total, err
}
