// 自动生成模板CrmProcurementContract
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 订购合同 结构体  CrmProcurementContract
type CrmProcurementContract struct {
 global.GVA_MODEL 
      ContractName  string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"`  //合同名称 
      CreationTime  *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`  //创建时间 
      ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`  //到期时间 
      ContractFile  string `json:"contractFile" form:"contractFile" gorm:"column:contract_file;comment:合同文件;size:191;"`  //合同文件 
      ContractAmount  *float64 `json:"contractAmount" form:"contractAmount" gorm:"column:contract_amount;comment:合同金额;size:10;"`  //合同金额 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人;size:10;"`  //负责人 
      ContractStatus  string `json:"contractStatus" form:"contractStatus" gorm:"column:contract_status;comment:合同状态;size:191;"`  //合同状态 
}


// TableName 订购合同 CrmProcurementContract自定义表名 crm_procurement_contract
func (CrmProcurementContract) TableName() string {
  return "crm_procurement_contract"
}

