// 自动生成模板CrmCurrency
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 币种管理 结构体  CrmCurrency
type CrmCurrency struct {
 global.GVA_MODEL 
      CurrencyName  string `json:"currencyName" form:"currencyName" gorm:"column:currency_name;comment:币种名称;size:10;"`  //币种名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
}


// TableName 币种管理 CrmCurrency自定义表名 crm_currency
func (CrmCurrency) TableName() string {
  return "crm_currency"
}

