// 自动生成模板CrmCustomers
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 客户管理 结构体  CrmCustomers
type CrmCustomers struct {
 global.GVA_MODEL 
      CustomerName  string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"`  //客户名 
      CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" gorm:"column:customer_phone_data;comment:客户手机号;size:191;"`  //客户手机号 
      SysUserId  *int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:管理ID 销售代表;size:10;"`  //管理ID 销售代表 
      SysUserAuthorityId  *int `json:"sysUserAuthorityId" form:"sysUserAuthorityId" gorm:"column:sys_user_authority_id;comment:管理角色ID;size:10;"`  //管理角色ID 
      CustomerCompany  string `json:"customerCompany" form:"customerCompany" gorm:"column:customer_company;comment:客户公司;size:191;"`  //客户公司 
      CustomerAddress  string `json:"customerAddress" form:"customerAddress" gorm:"column:customer_address;comment:客户地址;size:191;"`  //客户地址 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:191;"`  //描述 
      CustomerStatus  string `json:"customerStatus" form:"customerStatus" gorm:"column:customer_status;comment:用户状态;size:191;"`  //用户状态 
      CustomerGroup  string `json:"customerGroup" form:"customerGroup" gorm:"column:customer_group;comment:客户分组;size:191;"`  //客户分组 
}


// TableName 客户管理 CrmCustomers自定义表名 crm_customers
func (CrmCustomers) TableName() string {
  return "crm_customers"
}

