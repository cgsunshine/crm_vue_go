package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmTicketsInfoList 分页获取工单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService) GetCrmPageTicketsInfoList(info crmReq.CrmTicketsSearch) (list []crm.CrmPageTickets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmTickets{})
	var crmTicketss []crm.CrmPageTickets
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartActualResolutionTime != nil && info.EndActualResolutionTime != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("actual_resolution_time BETWEEN ? AND ?"), info.StartActualResolutionTime, info.EndActualResolutionTime)
	}
	if info.StartLastReplyTime != nil && info.EndLastReplyTime != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("last_reply_time BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SubmitterId != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("submitter_id = ?"), info.SubmitterId)
	}
	if info.AssigneeId != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("assignee_id = ?"), info.AssigneeId)
	}

	if info.TicketStatus != "" {
		db = db.Where(crmTicketsService.SplicingQueryConditions("ticket_status = ?"), info.TicketStatus)
	}

	if info.Title != "" {
		db = db.Where(crmTicketsService.SplicingQueryConditions("title LIKE ?"), "%"+info.Title+"%")
	}

	if info.Priority != "" {
		db = db.Where(crmTicketsService.SplicingQueryConditions("priority = ?"), info.Priority)
	}

	if info.CategoryId != nil {
		db = db.Where(crmTicketsService.SplicingQueryConditions("category_id = ?"), info.CategoryId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_tickets.*,assignee_user.username as assignee_user_name,submitter_user.username as submitter_user_name,crm_ticket_categories.category_name").
		Joins("LEFT JOIN sys_users as  submitter_user ON submitter_user.id = crm_tickets.submitter_id").
		Joins("LEFT JOIN sys_users as assignee_user ON assignee_user.id = crm_tickets.assignee_id").
		Joins("LEFT JOIN crm_ticket_categories ON crm_ticket_categories.id = crm_tickets.category_id").
		Order(fmt.Sprintf("crm_tickets.created_at %s", comm.OrderHandle(info.CreatedAtOrder))).
		Find(&crmTicketss).Error
	return crmTicketss, total, err
}

// GetCrmTickets 根据ID获取工单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService) GetCrmPageTickets(ID string) (crmTickets crm.CrmPageTickets, err error) {
	err = global.GVA_DB.Model(&crm.CrmTickets{}).Where("crm_tickets.id = ?", ID).
		Select("crm_tickets.*,assignee_user.username as assignee_user_name,submitter_user.username as submitter_user_name,crm_ticket_categories.category_name").
		Joins("LEFT JOIN sys_users as  submitter_user ON submitter_user.id = crm_tickets.submitter_id").
		Joins("LEFT JOIN sys_users as assignee_user ON assignee_user.id = crm_tickets.assignee_id").
		Joins("LEFT JOIN crm_ticket_categories ON crm_ticket_categories.id = crm_tickets.category_id").
		First(&crmTickets).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmTicketsService *CrmTicketsService) SplicingQueryConditions(condition string) string {
	return "crm_tickets." + condition
}

// UpdApprovalStatus 修改工单信息
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService) UpdTicketsInfo(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmTickets{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}
