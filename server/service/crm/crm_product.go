package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"strconv"
	"strings"
)

type CrmProductService struct {
}

// CreateCrmProduct 创建crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) CreateCrmProduct(crmProduct *crm.CrmProduct) (err error) {
	err = global.GVA_DB.Create(crmProduct).Error
	return err
}

// DeleteCrmProduct 删除crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) DeleteCrmProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmProduct{}, "id = ?", ID).Error
	return err
}

// DeleteCrmProductByIds 批量删除crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) DeleteCrmProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmProduct{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmProduct 更新crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) UpdateCrmProduct(crmProduct crm.CrmProduct) (err error) {
	err = global.GVA_DB.Save(&crmProduct).Error
	return err
}

// GetCrmProduct 根据ID获取crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) GetCrmProduct(ID string) (crmProduct crm.CrmProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmProduct).Error
	return
}

// GetCrmProductInfoList 分页获取crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) GetCrmProductInfoList(info crmReq.CrmProductSearch) (list []crm.CrmProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmProduct{})
	var crmProducts []crm.CrmProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductName != "" {
		db = db.Where("product_name = ?", info.ProductName)
	}
	if info.ProductGroupId != nil {
		db = db.Where("product_group_id = ?", info.ProductGroupId)
	}
	if info.ProductTypeId != nil {
		db = db.Where("product_type_id = ?", info.ProductTypeId)
	}
	if info.Inventory != nil {
		db = db.Where("inventory <> ?", info.Inventory)
	}
	if info.Price != nil {
		db = db.Where("price <> ?", info.Price)
	}
	if info.DiscountPrice != nil {
		db = db.Where("discount_price <> ?", info.DiscountPrice)
	}
	if info.SalesPrice != nil {
		db = db.Where("sales_price <> ?", info.SalesPrice)
	}
	if info.ResourceId != nil {
		db = db.Where("resource_id = ?", info.ResourceId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmProducts).Error
	return crmProducts, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分通过ids 获取产品列表
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (crmProductService *CrmProductService) GetCrmProductInfoIdsList(ids string) (list []crm.CrmProduct, total int64, err error) {

	db := global.GVA_DB.Model(&crm.CrmProduct{})
	//var fileLists []example.ExaFileUploadAndDownload
	//if len(keyword) > 0 {
	//	db = db.Where("name LIKE ?", "%"+keyword+"%")
	//}

	// 将字符串转换为整数切片
	var idList []uint
	for _, idStr := range strings.Split(ids, ",") {
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Error parsing ID:", err)
			continue
		}
		idList = append(idList, uint(id))
	}
	db.Where("id in (?)", idList)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("updated_at desc").Debug().Find(&list).Error
	return list, total, err
}
