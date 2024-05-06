package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmTestService struct {
}

// CreateCrmTest 创建crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService) CreateCrmTest(crmTest *crm.CrmTest) (err error) {
	err = global.GVA_DB.Create(crmTest).Error
	return err
}

// DeleteCrmTest 删除crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService)DeleteCrmTest(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmTest{},"id = ?",ID).Error
	return err
}

// DeleteCrmTestByIds 批量删除crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService)DeleteCrmTestByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmTest{},"id in ?",IDs).Error
	return err
}

// UpdateCrmTest 更新crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService)UpdateCrmTest(crmTest crm.CrmTest) (err error) {
	err = global.GVA_DB.Save(&crmTest).Error
	return err
}

// GetCrmTest 根据ID获取crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService)GetCrmTest(ID string) (crmTest crm.CrmTest, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmTest).Error
	return
}

// GetCrmTestInfoList 分页获取crmTest表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmTestService *CrmTestService)GetCrmTestInfoList(info crmReq.CrmTestSearch) (list []crm.CrmTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmTest{})
    var crmTests []crm.CrmTest
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
	
	err = db.Find(&crmTests).Error
	return  crmTests, total, err
}
