package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmTicketCategoriesService struct {
}

// CreateCrmTicketCategories 创建工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService) CreateCrmTicketCategories(crmTicketCategories *crm.CrmTicketCategories) (err error) {
	err = global.GVA_DB.Create(crmTicketCategories).Error
	return err
}

// DeleteCrmTicketCategories 删除工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService)DeleteCrmTicketCategories(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTicketCategories{},"id = ?",ID).Error
	return err
}

// DeleteCrmTicketCategoriesByIds 批量删除工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService)DeleteCrmTicketCategoriesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTicketCategories{},"id in ?",IDs).Error
	return err
}

// UpdateCrmTicketCategories 更新工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService)UpdateCrmTicketCategories(crmTicketCategories crm.CrmTicketCategories) (err error) {
	err = global.GVA_DB.Save(&crmTicketCategories).Error
	return err
}

// GetCrmTicketCategories 根据ID获取工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService)GetCrmTicketCategories(ID string) (crmTicketCategories crm.CrmTicketCategories, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmTicketCategories).Error
	return
}

// GetCrmTicketCategoriesInfoList 分页获取工单类别记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTicketCategoriesService *CrmTicketCategoriesService)GetCrmTicketCategoriesInfoList(info crmReq.CrmTicketCategoriesSearch) (list []crm.CrmTicketCategories, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmTicketCategories{})
    var crmTicketCategoriess []crm.CrmTicketCategories
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
	
	err = db.Find(&crmTicketCategoriess).Error
	return  crmTicketCategoriess, total, err
}
