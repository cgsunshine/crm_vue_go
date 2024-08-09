package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmQueryApproved 查询审核是否通过（所有人都要同意）
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) GetCrmQueryApproved(businessOpportunityId int, approvalStatus string) (bool, error) {
	total := int64(0)
	count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{})

	db = db.Where("business_opportunity_id = ? AND valid = ?", businessOpportunityId, comm.Contact_Approval_Tasks_valid_Effective)

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

// GetCrmQueryStepApproved 查询审核是否通过 再此步骤中，可能需要部分人他同意即可通过审批，具体看配置
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) GetCrmQueryStepApproved(businessOpportunityId, numberApprovedPersonnel int, approvalStatus string) (bool, error) {
	total := int64(0)
	//count := int64(0)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{})

	db = db.Where("business_opportunity_id = ? AND valid = ?", businessOpportunityId, comm.Contact_Approval_Tasks_valid_Effective)

	err := db.Count(&total).Error
	if err != nil {
		return false, err
	}

	db = db.Where("approval_status = ?", approvalStatus)

	if total == 0 {
		return false, nil
	}

	//if total == count {
	//	return true, nil
	//}

	if total >= int64(numberApprovedPersonnel) {
		return true, nil
	}
	return false, nil
}

// GetCrmBusinessOpportunityApprovalTasksInfoList 分页获取合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) GetCrmPageContactApprovalTasksInfoList(info crmReq.CrmBusinessOpportunityApprovalTasksSearch) (list []crm.CrmPageBusinessOpportunityApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{})
	var CrmBusinessOpportunityApprovalTaskss []crm.CrmPageBusinessOpportunityApprovalTasks
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmBusinessOpportunityApprovalTasksService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ApprovalStatus != "" {
		db = db.Where(crmBusinessOpportunityApprovalTasksService.SplicingQueryConditions("approval_status = ?"), info.ApprovalStatus)
	}
	if info.AssigneeId != nil {
		db = db.Where(crmBusinessOpportunityApprovalTasksService.SplicingQueryConditions("assignee_id = ?"), info.AssigneeId)
	}
	if info.BusinessOpportunityId != nil {
		db = db.Where(crmBusinessOpportunityApprovalTasksService.SplicingQueryConditions("business_opportunity_id = ?"), info.BusinessOpportunityId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_business_opportunity_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_business_opportunity_approval_tasks.business_opportunity_id").
		Find(&CrmBusinessOpportunityApprovalTaskss).Error
	return CrmBusinessOpportunityApprovalTaskss, total, err
}

// GetCrmBusinessOpportunityApprovalTasks 更新
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) UpdCrmBusinessOpportunityApprovalTasks(ID uint, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{}).Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmBusinessOpportunityApprovalTasks 更新 通合同id
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) UpdCrmBusinessOpportunityIdApprovalTasks(businessOpportunityId int, data map[string]interface{}) (err error) {
	err = global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{}).Where("business_opportunity_id = ?", businessOpportunityId).Updates(data).Error
	return
}

// GetCrmBusinessOpportunityApprovalTasks 根据ID获取合同审批记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) GetCrmPageBusinessOpportunityApprovalTasks(ID string) (CrmBusinessOpportunityApprovalTasks crm.CrmPageBusinessOpportunityApprovalTasks, err error) {
	err = global.GVA_DB.Model(&crm.CrmBusinessOpportunityApprovalTasks{}).Where("id = ?", ID).
		Select("crm_business_opportunity_approval_tasks.*,crm_contract.contract_name").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_business_opportunity_approval_tasks.business_opportunity_id").
		First(&CrmBusinessOpportunityApprovalTasks).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmBusinessOpportunityApprovalTasksService *CrmBusinessOpportunityApprovalTasksService) SplicingQueryConditions(condition string) string {
	return "crm_business_opportunity_approval_tasks." + condition
}
