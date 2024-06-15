package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmTicketsService struct {
}

// CreateCrmTickets 创建crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService) CreateCrmTickets(crmTickets *crm.CrmTickets) (err error) {
	err = global.GVA_DB.Create(crmTickets).Error
	return err
}

// DeleteCrmTickets 删除crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService)DeleteCrmTickets(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTickets{},"id = ?",ID).Error
	return err
}

// DeleteCrmTicketsByIds 批量删除crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService)DeleteCrmTicketsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTickets{},"id in ?",IDs).Error
	return err
}

// UpdateCrmTickets 更新crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService)UpdateCrmTickets(crmTickets crm.CrmTickets) (err error) {
	err = global.GVA_DB.Save(&crmTickets).Error
	return err
}

// GetCrmTickets 根据ID获取crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService)GetCrmTickets(ID string) (crmTickets crm.CrmTickets, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmTickets).Error
	return
}

// GetCrmTicketsInfoList 分页获取crmTickets表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketsService *CrmTicketsService)GetCrmTicketsInfoList(info crmReq.CrmTicketsSearch) (list []crm.CrmTickets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmTickets{})
    var crmTicketss []crm.CrmTickets
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
	
	err = db.Find(&crmTicketss).Error
	return  crmTicketss, total, err
}
