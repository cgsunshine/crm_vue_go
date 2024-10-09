package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmRefundTasksService struct {
}

// CreateCrmRefundTasks 创建退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) CreateCrmRefundTasks(crmRefundTasks *crm.CrmRefundTasks) (err error) {
	err = global.GVA_DB.Create(crmRefundTasks).Error
	return err
}

// DeleteCrmRefundTasks 删除退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) DeleteCrmRefundTasks(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmRefundTasks{}, "id = ?", ID).Error
	return err
}

// DeleteCrmRefundTasksByIds 批量删除退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) DeleteCrmRefundTasksByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmRefundTasks{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmRefundTasks 更新退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) UpdateCrmRefundTasks(crmRefundTasks crm.CrmRefundTasks) (err error) {
	err = global.GVA_DB.Save(&crmRefundTasks).Error
	return err
}

// GetCrmRefundTasks 根据ID获取退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) GetCrmRefundTasks(ID string) (crmRefundTasks crm.CrmRefundTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmRefundTasks).Error
	return
}

// GetCrmRefundTasks 根据ID获取退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) GetCrmRefundTasksInfo(ID uint) (crmRefundTasks crm.CrmRefundTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmRefundTasks).Error
	return
}

// UpdApprovalStatus 修改退款信息
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) UpdDepositsInfo(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmRefundTasks{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmRefundTasksInfoList 分页获取退款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmRefundTasksService *CrmRefundTasksService) GetCrmRefundTasksInfoList(info crmReq.CrmRefundTasksSearch) (list []crm.CrmPageRefundTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmRefundTasks{})
	var crmRefundTaskss []crm.CrmPageRefundTasks
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AssigneeId != nil {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("assignee_id = ?"), info.AssigneeId)
	}
	if info.RefundStatus != "" {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("refund_status = ?"), info.RefundStatus)
	}
	if info.Valid != nil {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("valid = ?"), info.Valid)
	}
	if info.AssociatedId != nil {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("associated_id = ?"), info.AssociatedId)
	}
	if info.RefundType != nil {
		db = db.Where(crmRefundTasksService.SplicingQueryConditions("refund_type = ?"), info.RefundType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	//Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
	//	Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
	//	Joins("LEFT JOIN crm_product ON crm_product.id = crm_order.product_id").
	//	Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_order.business_opportunity_id").
	//	Preload("Products").
	//	Order("crm_order.created_at DESC").
	err = db.Select("crm_refund_tasks.*,crm_deposits.deposits_name,crm_deposits.deposits_number").
		Joins("LEFT JOIN crm_deposits ON crm_deposits.id = crm_refund_tasks.associated_id").
		Order("crm_refund_tasks.created_at DESC").
		Find(&crmRefundTaskss).Error
	return crmRefundTaskss, total, err
}

// SplicingQueryConditions 拼接条件
func (crmRefundTasksService *CrmRefundTasksService) SplicingQueryConditions(condition string) string {
	return "crm_refund_tasks." + condition
}
