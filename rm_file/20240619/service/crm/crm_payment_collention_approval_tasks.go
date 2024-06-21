package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPaymentCollentionApprovalTasksService struct {
}

// CreateCrmPaymentCollentionApprovalTasks 创建crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService) CreateCrmPaymentCollentionApprovalTasks(crmPaymentCollentionApprovalTasks *crm.CrmPaymentCollentionApprovalTasks) (err error) {
	err = global.GVA_DB.Create(crmPaymentCollentionApprovalTasks).Error
	return err
}

// DeleteCrmPaymentCollentionApprovalTasks 删除crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService)DeleteCrmPaymentCollentionApprovalTasks(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPaymentCollentionApprovalTasks{},"id = ?",ID).Error
	return err
}

// DeleteCrmPaymentCollentionApprovalTasksByIds 批量删除crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService)DeleteCrmPaymentCollentionApprovalTasksByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPaymentCollentionApprovalTasks{},"id in ?",IDs).Error
	return err
}

// UpdateCrmPaymentCollentionApprovalTasks 更新crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService)UpdateCrmPaymentCollentionApprovalTasks(crmPaymentCollentionApprovalTasks crm.CrmPaymentCollentionApprovalTasks) (err error) {
	err = global.GVA_DB.Save(&crmPaymentCollentionApprovalTasks).Error
	return err
}

// GetCrmPaymentCollentionApprovalTasks 根据ID获取crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService)GetCrmPaymentCollentionApprovalTasks(ID string) (crmPaymentCollentionApprovalTasks crm.CrmPaymentCollentionApprovalTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPaymentCollentionApprovalTasks).Error
	return
}

// GetCrmPaymentCollentionApprovalTasksInfoList 分页获取crmPaymentCollentionApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionApprovalTasksService *CrmPaymentCollentionApprovalTasksService)GetCrmPaymentCollentionApprovalTasksInfoList(info crmReq.CrmPaymentCollentionApprovalTasksSearch) (list []crm.CrmPaymentCollentionApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmPaymentCollentionApprovalTasks{})
    var crmPaymentCollentionApprovalTaskss []crm.CrmPaymentCollentionApprovalTasks
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ApprovalStatus != "" {
        db = db.Where("approval_status = ?",info.ApprovalStatus)
    }
    if info.AssigneeId != nil {
        db = db.Where("assignee_id = ?",info.AssigneeId)
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
	
	err = db.Find(&crmPaymentCollentionApprovalTaskss).Error
	return  crmPaymentCollentionApprovalTaskss, total, err
}
