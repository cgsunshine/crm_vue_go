package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmContactFileUploadAndDownloadsService struct {
}

// CreateCrmContactFileUploadAndDownloads 创建crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService) CreateCrmContactFileUploadAndDownloads(crmContactFileUploadAndDownloads *crm.CrmContactFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Create(crmContactFileUploadAndDownloads).Error
	return err
}

// DeleteCrmContactFileUploadAndDownloads 删除crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService)DeleteCrmContactFileUploadAndDownloads(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmContactFileUploadAndDownloads{},"id = ?",ID).Error
	return err
}

// DeleteCrmContactFileUploadAndDownloadsByIds 批量删除crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService)DeleteCrmContactFileUploadAndDownloadsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmContactFileUploadAndDownloads{},"id in ?",IDs).Error
	return err
}

// UpdateCrmContactFileUploadAndDownloads 更新crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService)UpdateCrmContactFileUploadAndDownloads(crmContactFileUploadAndDownloads crm.CrmContactFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Save(&crmContactFileUploadAndDownloads).Error
	return err
}

// GetCrmContactFileUploadAndDownloads 根据ID获取crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService)GetCrmContactFileUploadAndDownloads(ID string) (crmContactFileUploadAndDownloads crm.CrmContactFileUploadAndDownloads, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmContactFileUploadAndDownloads).Error
	return
}

// GetCrmContactFileUploadAndDownloadsInfoList 分页获取crmContactFileUploadAndDownloads表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileUploadAndDownloadsService *CrmContactFileUploadAndDownloadsService)GetCrmContactFileUploadAndDownloadsInfoList(info crmReq.CrmContactFileUploadAndDownloadsSearch) (list []crm.CrmContactFileUploadAndDownloads, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmContactFileUploadAndDownloads{})
    var crmContactFileUploadAndDownloadss []crm.CrmContactFileUploadAndDownloads
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ContactId != nil {
        db = db.Where("contact_id = ?",info.ContactId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["sort"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmContactFileUploadAndDownloadss).Error
	return  crmContactFileUploadAndDownloadss, total, err
}
