package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmTicketCommentsService struct {
}

// CreateCrmTicketComments 创建crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService) CreateCrmTicketComments(crmTicketComments *crm.CrmTicketComments) (err error) {
	err = global.GVA_DB.Create(crmTicketComments).Error
	return err
}

// DeleteCrmTicketComments 删除crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService)DeleteCrmTicketComments(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTicketComments{},"id = ?",ID).Error
	return err
}

// DeleteCrmTicketCommentsByIds 批量删除crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService)DeleteCrmTicketCommentsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTicketComments{},"id in ?",IDs).Error
	return err
}

// UpdateCrmTicketComments 更新crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService)UpdateCrmTicketComments(crmTicketComments crm.CrmTicketComments) (err error) {
	err = global.GVA_DB.Save(&crmTicketComments).Error
	return err
}

// GetCrmTicketComments 根据ID获取crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService)GetCrmTicketComments(ID string) (crmTicketComments crm.CrmTicketComments, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmTicketComments).Error
	return
}

// GetCrmTicketCommentsInfoList 分页获取crmTicketComments表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCommentsService *CrmTicketCommentsService)GetCrmTicketCommentsInfoList(info crmReq.CrmTicketCommentsSearch) (list []crm.CrmTicketComments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmTicketComments{})
    var crmTicketCommentss []crm.CrmTicketComments
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
	
	err = db.Find(&crmTicketCommentss).Error
	return  crmTicketCommentss, total, err
}
