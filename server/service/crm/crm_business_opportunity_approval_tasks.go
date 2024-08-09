package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityApprovalTasksService struct {
}

// CreateCrmBusinessOpportunityApprovalTasks 创建商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) CreateCrmBusinessOpportunityApprovalTasks(crmBusinessOpportunityApprovalTasks *crm.CrmBusinessOpportunityApprovalTasks) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunityApprovalTasks).Error
	return err
}

// DeleteCrmBusinessOpportunityApprovalTasks 删除商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService)DeleteCrmBusinessOpportunityApprovalTasks(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunityApprovalTasks{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityApprovalTasksByIds 批量删除商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService)DeleteCrmBusinessOpportunityApprovalTasksByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunityApprovalTasks{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunityApprovalTasks 更新商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService)UpdateCrmBusinessOpportunityApprovalTasks(crmBusinessOpportunityApprovalTasks crm.CrmBusinessOpportunityApprovalTasks) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunityApprovalTasks).Error
	return err
}

// GetCrmBusinessOpportunityApprovalTasks 根据ID获取商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService)GetCrmBusinessOpportunityApprovalTasks(ID string) (crmBusinessOpportunityApprovalTasks crm.CrmBusinessOpportunityApprovalTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunityApprovalTasks).Error
	return
}

// GetCrmBusinessOpportunityApprovalTasksInfoList 分页获取商机审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService)GetCrmBusinessOpportunityApprovalTasksInfoList(info crmReq.CrmBusinessOpportunityApprovalTasksSearch) (list []crm.CrmBusinessOpportunityApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{})
    var crmBusinessOpportunityApprovalTaskss []crm.CrmBusinessOpportunityApprovalTasks
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
	
	err = db.Find(&crmBusinessOpportunityApprovalTaskss).Error
	return  crmBusinessOpportunityApprovalTaskss, total, err
}
