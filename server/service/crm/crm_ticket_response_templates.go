package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
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
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) DeleteCrmTicketResponseTemplates(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTicketResponseTemplates{}, "id = ?", ID).Error
	return err
}

// DeleteCrmTicketResponseTemplatesByIds 批量删除快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) DeleteCrmTicketResponseTemplatesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTicketResponseTemplates{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmTicketResponseTemplates 更新快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) UpdateCrmTicketResponseTemplates(crmTicketResponseTemplates crm.CrmTicketResponseTemplates) (err error) {
	err = global.GVA_DB.Save(&crmTicketResponseTemplates).Error
	return err
}

// GetCrmTicketResponseTemplates 根据ID获取快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) GetCrmTicketResponseTemplates(ID string) (crmTicketResponseTemplates crm.CrmTicketResponseTemplates, err error) {
	err = global.GVA_DB.Where(crmTicketResponseTemplatesService.SplicingQueryConditions("id = ?"), ID).
		Select("crm_ticket_response_templates.*,crm_ticket_categories.category_name").
		Joins("LEFT JOIN crm_ticket_categories ON crm_ticket_categories.id = crm_ticket_response_templates.category_id").
		First(&crmTicketResponseTemplates).Error
	return
}

// GetCrmTicketResponseTemplatesInfoList 分页获取快捷回复模板记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) GetCrmTicketResponseTemplatesInfoList(info crmReq.CrmTicketResponseTemplatesSearch) (list []crm.CrmPageTicketResponseTemplates, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmTicketResponseTemplates{})
	var crmTicketResponseTemplatess []crm.CrmPageTicketResponseTemplates
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmTicketResponseTemplatesService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_ticket_response_templates.*,crm_ticket_categories.category_name").
		Joins("LEFT JOIN crm_ticket_categories ON crm_ticket_categories.id = crm_ticket_response_templates.category_id").
		Order(fmt.Sprintf("crm_ticket_response_templates.created_at %s", comm.OrderHandle(info.CreatedAtOrder))).
		Find(&crmTicketResponseTemplatess).Error
	return crmTicketResponseTemplatess, total, err
}

// SplicingQueryConditions 拼接条件
func (crmTicketResponseTemplatesService *CrmTicketResponseTemplatesService) SplicingQueryConditions(condition string) string {
	return "crm_ticket_response_templates." + condition
}
