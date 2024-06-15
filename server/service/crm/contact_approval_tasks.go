package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmQueryApproved 查询审核是否通过
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) GetCrmQueryApproved(contactId int, approvalStatus string) (bool, error) {
	total := int64(0)
	count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmContactApprovalTasks{})

	db = db.Where("contact_id = ? AND valid = ?", contactId, comm.Contact_Approval_Tasks_valid_Effective)

	err := db.Count(&total).Error
	if err != nil {
		return false, err
	}

	db = db.Where("approval_status = ?", approvalStatus)

	err = db.Count(&count).Error
	if err != nil {
		return false, err
	}
	if total == 0 {
		return false, nil
	}

	if total == count {
		return true, nil
	}
	return false, nil
}

// GetCrmContactApprovalTasksInfoList 分页获取合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) GetCrmPageContactApprovalTasksInfoList(info crmReq.CrmContactApprovalTasksSearch) (list []crm.CrmRespContactApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmContactApprovalTasks{})
	var crmContactApprovalTaskss []crm.CrmRespContactApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmContactApprovalTasksService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AssigneeId != nil {
		db = db.Where(crmContactApprovalTasksService.SplicingQueryConditions("assignee_id = ?"), info.AssigneeId)
	}
	if info.ApprovalStatus != "" {
		db = db.Where(crmContactApprovalTasksService.SplicingQueryConditions("approval_status = ?"), info.ApprovalStatus)
	}
	if info.ContactId != nil {
		db = db.Where(crmContactApprovalTasksService.SplicingQueryConditions("contact_id = ?"), info.ContactId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_contact_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_contact_approval_tasks.contact_id").
		Find(&crmContactApprovalTaskss).Error
	return crmContactApprovalTaskss, total, err
}

// GetCrmContactApprovalTasks 更新
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) UpdCrmContactApprovalTasks(ID uint, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmContactApprovalTasks{}).Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmContactApprovalTasks 更新 通合同id
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) UpdCrmContactIDContactApprovalTasks(ContactID int, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmContactApprovalTasks{}).Where("contact_id = ?", ContactID).Updates(data).Error
	return
}

// GetCrmContactApprovalTasks 根据ID获取合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) GetCrmPageContactApprovalTasks(ID string) (crmContactApprovalTasks crm.CrmRespContactApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmContactApprovalTasks{}).Where("id = ?", ID).
		Select("crm_contact_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_contact_approval_tasks.contact_id").
		First(&crmContactApprovalTasks).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmContactApprovalTasksService *CrmContactApprovalTasksService) SplicingQueryConditions(condition string) string {
	return "crm_contact_approval_tasks." + condition
}
