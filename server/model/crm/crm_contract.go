// 自动生成模板CrmContract
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 合同管理 结构体  CrmContract
type CrmContract struct {
 global.GVA_MODEL 
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:产品名称;size:10;"`  //产品名称 
      CustomerId  *int `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:产品ID;size:10;"`  //产品ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:10;"`  //管理ID 销售代表 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;size:10;"`  //产品原价 
      SalesPrice  *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品折扣价;size:10;"`  //产品折扣价 
      Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`  //币种 
      DiscountRate  *int `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;size:10;"`  //折扣率 
      ContractTypeId  *int `json:"contractTypeId" form:"contractTypeId" gorm:"column:contract_type_id;comment:合同类型;size:10;"`  //合同类型 
      ApplicationTime  *time.Time `json:"applicationTime" form:"applicationTime" gorm:"column:application_time;comment:申请时间;"`  //申请时间 
      ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`  //到期时间 
      ContractStatus  string `json:"contractStatus" form:"contractStatus" gorm:"column:contract_status;comment:合同状态;size:255;"`  //合同状态 
      ReviewStatus  string `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:255;"`  //审核状态 
      ReviewResult  string `json:"reviewResult" form:"reviewResult" gorm:"column:review_result;comment:审核结果;size:255;"`  //审核结果 
      ContractDocument  string `json:"contractDocument" form:"contractDocument" gorm:"column:contract_document;comment:合同文件;size:255;"`  //合同文件 
}


// TableName 合同管理 CrmContract自定义表名 crm_contract
func (CrmContract) TableName() string {
  return "crm_contract"
}

