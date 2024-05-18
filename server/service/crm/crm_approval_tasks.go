package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmApprovalTasksService struct {
}

// CreateCrmApprovalTasks 创建crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService) CreateCrmApprovalTasks(crmApprovalTasks *crm.CrmApprovalTasks) (err error) {
	err = global.GVA_DB.Create(crmApprovalTasks).Error
	return err
}

// DeleteCrmApprovalTasks 删除crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService)DeleteCrmApprovalTasks(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmApprovalTasks{},"id = ?",ID).Error
	return err
}

// DeleteCrmApprovalTasksByIds 批量删除crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService)DeleteCrmApprovalTasksByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmApprovalTasks{},"id in ?",IDs).Error
	return err
}

// UpdateCrmApprovalTasks 更新crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService)UpdateCrmApprovalTasks(crmApprovalTasks crm.CrmApprovalTasks) (err error) {
	err = global.GVA_DB.Save(&crmApprovalTasks).Error
	return err
}

// GetCrmApprovalTasks 根据ID获取crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService)GetCrmApprovalTasks(ID string) (crmApprovalTasks crm.CrmApprovalTasks, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalTasks).Error
	return
}

// GetCrmApprovalTasksInfoList 分页获取crmApprovalTasks表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksService *CrmApprovalTasksService)GetCrmApprovalTasksInfoList(info crmReq.CrmApprovalTasksSearch) (list []crm.CrmApprovalTasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasks{})
    var crmApprovalTaskss []crm.CrmApprovalTasks
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
	
	err = db.Find(&crmApprovalTaskss).Error
	return  crmApprovalTaskss, total, err
}
