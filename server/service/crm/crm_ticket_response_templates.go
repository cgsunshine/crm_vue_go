package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmTicketResponseTemplatesService struct {
}

// CreateCrmTicketResponseTemplates 创建快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) CreateCrmTicketResponseTemplates(crmTicketResponseTemplates *crm.CrmTicketResponseTemplates) (err error) {
	err = global.GVA_DB.Create(crmTicketResponseTemplates).Error
	return err
}

// DeleteCrmTicketResponseTemplates 删除快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService)DeleteCrmTicketResponseTemplates(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTicketResponseTemplates{},"id = ?",ID).Error
	return err
}

// DeleteCrmTicketResponseTemplatesByIds 批量删除快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService)DeleteCrmTicketResponseTemplatesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTicketResponseTemplates{},"id in ?",IDs).Error
	return err
}

// UpdateCrmTicketResponseTemplates 更新快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService)UpdateCrmTicketResponseTemplates(crmTicketResponseTemplates crm.CrmTicketResponseTemplates) (err error) {
	err = global.GVA_DB.Save(&crmTicketResponseTemplates).Error
	return err
}

// GetCrmTicketResponseTemplates 根据ID获取快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService)GetCrmTicketResponseTemplates(ID string) (crmTicketResponseTemplates crm.CrmTicketResponseTemplates, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmTicketResponseTemplates).Error
	return
}

// GetCrmTicketResponseTemplatesInfoList 分页获取快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService)GetCrmTicketResponseTemplatesInfoList(info crmReq.CrmTicketResponseTemplatesSearch) (list []crm.CrmTicketResponseTemplates, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmTicketResponseTemplates{})
    var crmTicketResponseTemplatess []crm.CrmTicketResponseTemplates
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
	
	err = db.Find(&crmTicketResponseTemplatess).Error
	return  crmTicketResponseTemplatess, total, err
}
