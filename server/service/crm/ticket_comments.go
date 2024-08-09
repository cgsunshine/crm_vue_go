package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmTicketCommentsInfoList 分页获取共单回复记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService) GetCrmPageTicketCommentsInfoList(info crmReq.CrmTicketCommentsSearch) (list []crm.CrmPageTicketComments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmTicketComments{})
	var crmTicketCommentss []crm.CrmPageTicketComments
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmTicketCommentsService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TicketId != nil {
		db = db.Where("ticket_id = ?", info.TicketId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_ticket_comments.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_ticket_comments.user_id").
		Order("crm_ticket_comments.created_at DESC").
		Find(&crmTicketCommentss).Error
	return crmTicketCommentss, total, err
}

// GetCrmTicketComments 根据ID获取共单回复记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService) GetCrmPageTicketComments(ID string) (crmTicketComments crm.CrmPageTicketComments, err error) {
	err = global.GVA_DB.Model(&crm.CrmTicketComments{}).Where("crm_ticket_comments.id = ?", ID).
		Select("crm_ticket_comments.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_ticket_comments.user_id").
		First(&crmTicketComments).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmTicketCommentsService *CrmTicketCommentsService) SplicingQueryConditions(condition string) string {
	return "crm_ticket_comments." + condition
}
