// 自动生成模板CrmProduct
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 产品管理 结构体  CrmProduct
type CrmProduct struct {
 global.GVA_MODEL 
      ProductName  string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:191;"`  //产品名称 
      CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" gorm:"column:customer_phone_data;comment:客户手机号;size:191;"`  //客户手机号 
      SysUserId  *int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:管理ID 销售代表;size:20;"`  //管理ID 销售代表 
      SysUserAuthorityId  *int `json:"sysUserAuthorityId" form:"sysUserAuthorityId" gorm:"column:sys_user_authority_id;comment:管理角色ID;size:20;"`  //管理角色ID 
      CustomerCompany  string `json:"customerCompany" form:"customerCompany" gorm:"column:customer_company;comment:客户公司;size:191;"`  //客户公司 
      CustomerAddress  string `json:"customerAddress" form:"customerAddress" gorm:"column:customer_address;comment:客户地址;size:191;"`  //客户地址 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      CustomerStatus  string `json:"customerStatus" form:"customerStatus" gorm:"column:customer_status;comment:用户状态;size:191;"`  //用户状态 
      ProductGroup  string `json:"productGroup" form:"productGroup" gorm:"column:product_group;comment:产品分组名称;size:191;"`  //产品分组名称 
      ProductGroupId  *int `json:"productGroupId" form:"productGroupId" gorm:"column:product_group_id;comment:产品分组id;size:10;"`  //产品分组id 
      ProductType  string `json:"productType" form:"productType" gorm:"column:product_type;comment:产品类型;size:255;"`  //产品类型 
      ProductTypeId  *int `json:"productTypeId" form:"productTypeId" gorm:"column:product_type_id;comment:产品类型id;size:10;"`  //产品类型id 
      Inventory  *int `json:"inventory" form:"inventory" gorm:"column:inventory;comment:库存;size:10;"`  //库存 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;size:10;"`  //产品原价 
      DiscountPrice  *float64 `json:"discountPrice" form:"discountPrice" gorm:"column:discount_price;comment:产品折扣价;size:10;"`  //产品折扣价 
      SalesPrice  *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品销售价格;size:10;"`  //产品销售价格 
}


// TableName 产品管理 CrmProduct自定义表名 crm_product
func (CrmProduct) TableName() string {
  return "crm_product"
}

