// 自动生成模板CrmCustomers
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户管理 结构体  CrmCustomers
type CrmPageCustomers struct {
	global.GVA_MODEL
	CustomerName      string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"`                   //客户名
	CustomerPhoneData string `json:"customerPhoneData" form:"customerPhoneData" gorm:"column:customer_phone_data;comment:客户手机号;size:191;"` //客户手机号
	UserId            *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:19;"`                                //管理ID 销售代表
	UserAuthorityId   *int   `json:"userAuthorityId" form:"userAuthorityId" gorm:"column:user_authority_id;comment:管理角色ID;size:19;"`       //管理角色ID
	CustomerCompany   string `json:"customerCompany" form:"customerCompany" gorm:"column:customer_company;comment:客户公司;size:191;"`         //客户公司
	CustomerAddress   string `json:"customerAddress" form:"customerAddress" gorm:"column:customer_address;comment:客户地址;size:191;"`         //客户地址
	Description       string `json:"description" form:"description" gorm:"column:description;comment:描述;size:191;"`                        //描述
	CustomerStatus    string `json:"customerStatus" form:"customerStatus" gorm:"column:customer_status;comment:用户状态;size:191;"`            //用户状态
	CustomerGroupId   *int   `json:"customerGroupId" form:"customerGroupId" gorm:"column:customer_group_id;comment:客户分组id;"`               //客户分组id
	CustomerGroup     string `json:"customerGroup" form:"customerGroup" gorm:"column:customer_group;comment:客户分组;size:191;"`               //客户分组
	//以下都是联表字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customer_group 表
	GroupName string   `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:90;"` //分组名称
	Balance   *float64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额;default:0;"`        //余额 （预付款金额）
}
