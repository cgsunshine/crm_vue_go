package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmApprovalProcessService struct {
}

// CreateCrmApprovalProcess 创建审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService) CreateCrmApprovalProcess(crmApprovalProcess *crm.CrmApprovalProcess) (err error) {
	err = global.GVA_DB.Create(crmApprovalProcess).Error
	return err
}

// DeleteCrmApprovalProcess 删除审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService)DeleteCrmApprovalProcess(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmApprovalProcess{},"id = ?",ID).Error
	return err
}

// DeleteCrmApprovalProcessByIds 批量删除审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService)DeleteCrmApprovalProcessByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmApprovalProcess{},"id in ?",IDs).Error
	return err
}

// UpdateCrmApprovalProcess 更新审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService)UpdateCrmApprovalProcess(crmApprovalProcess crm.CrmApprovalProcess) (err error) {
	err = global.GVA_DB.Save(&crmApprovalProcess).Error
	return err
}

// GetCrmApprovalProcess 根据ID获取审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService)GetCrmApprovalProcess(ID string) (crmApprovalProcess crm.CrmApprovalProcess, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalProcess).Error
	return
}

// GetCrmApprovalProcessInfoList 分页获取审批流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalProcessService *CrmApprovalProcessService)GetCrmApprovalProcessInfoList(info crmReq.CrmApprovalProcessSearch) (list []crm.CrmApprovalProcess, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalProcess{})
    var crmApprovalProcesss []crm.CrmApprovalProcess
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
	
	err = db.Find(&crmApprovalProcesss).Error
	return  crmApprovalProcesss, total, err
}
