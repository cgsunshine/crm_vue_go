package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmApprovalNodeService struct {
}

// CreateCrmApprovalNode 创建crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService) CreateCrmApprovalNode(crmApprovalNode *crm.CrmApprovalNode) (err error) {
	err = global.GVA_DB.Create(crmApprovalNode).Error
	return err
}

// DeleteCrmApprovalNode 删除crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService)DeleteCrmApprovalNode(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmApprovalNode{},"id = ?",ID).Error
	return err
}

// DeleteCrmApprovalNodeByIds 批量删除crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService)DeleteCrmApprovalNodeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmApprovalNode{},"id in ?",IDs).Error
	return err
}

// UpdateCrmApprovalNode 更新crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService)UpdateCrmApprovalNode(crmApprovalNode crm.CrmApprovalNode) (err error) {
	err = global.GVA_DB.Save(&crmApprovalNode).Error
	return err
}

// GetCrmApprovalNode 根据ID获取crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService)GetCrmApprovalNode(ID string) (crmApprovalNode crm.CrmApprovalNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalNode).Error
	return
}

// GetCrmApprovalNodeInfoList 分页获取crmApprovalNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalNodeService *CrmApprovalNodeService)GetCrmApprovalNodeInfoList(info crmReq.CrmApprovalNodeSearch) (list []crm.CrmApprovalNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalNode{})
    var crmApprovalNodes []crm.CrmApprovalNode
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
	
	err = db.Find(&crmApprovalNodes).Error
	return  crmApprovalNodes, total, err
}
