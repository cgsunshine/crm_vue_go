package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmContactApprovalRecordService struct {
}

// CreateCrmContactApprovalRecord 创建合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService) CreateCrmContactApprovalRecord(crmContactApprovalRecord *crm.CrmContactApprovalRecord) (err error) {
	err = global.GVA_DB.Create(crmContactApprovalRecord).Error
	return err
}

// DeleteCrmContactApprovalRecord 删除合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService)DeleteCrmContactApprovalRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmContactApprovalRecord{},"id = ?",ID).Error
	return err
}

// DeleteCrmContactApprovalRecordByIds 批量删除合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService)DeleteCrmContactApprovalRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmContactApprovalRecord{},"id in ?",IDs).Error
	return err
}

// UpdateCrmContactApprovalRecord 更新合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService)UpdateCrmContactApprovalRecord(crmContactApprovalRecord crm.CrmContactApprovalRecord) (err error) {
	err = global.GVA_DB.Save(&crmContactApprovalRecord).Error
	return err
}

// GetCrmContactApprovalRecord 根据ID获取合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService)GetCrmContactApprovalRecord(ID string) (crmContactApprovalRecord crm.CrmContactApprovalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmContactApprovalRecord).Error
	return
}

// GetCrmContactApprovalRecordInfoList 分页获取合同审批记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalRecordService *CrmContactApprovalRecordService)GetCrmContactApprovalRecordInfoList(info crmReq.CrmContactApprovalRecordSearch) (list []crm.CrmContactApprovalRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmContactApprovalRecord{})
    var crmContactApprovalRecords []crm.CrmContactApprovalRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ApproveOpinion != "" {
        db = db.Where("approve_opinion = ?",info.ApproveOpinion)
    }
    if info.ApproverId != nil {
        db = db.Where("approver_id = ?",info.ApproverId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmContactApprovalRecords).Error
	return  crmContactApprovalRecords, total, err
}
