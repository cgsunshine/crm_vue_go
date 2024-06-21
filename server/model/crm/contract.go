// 自动生成模板CrmContract
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmContract表 结构体  CrmContract
type CrmPageContract struct {
	global.GVA_MODEL
	ApplicationTime *time.Time `json:"applicationTime" form:"applicationTime" gorm:"column:application_time;comment:申请时间;"`       //申请时间
	ContractFile    string     `json:"contractFile" form:"contractFile" gorm:"column:contract_file;comment:合同文件路径;size:191;"`     //合同文件路径
	ContractName    string     `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"`       //合同名称
	ContractStatus  string     `json:"contractStatus" form:"contractStatus" gorm:"column:contract_status;comment:合同状态;size:191;"` //合同状态
	ContractTypeId  *int       `json:"contractTypeId" form:"contractTypeId" gorm:"column:contract_type_id;comment:合同类型;"`         //合同类型
	CustomerId      *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                      //客户ID
	Description     string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`             //备注
	ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`          //到期时间
	OrderId         *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                               //订单ID
	ReviewResult    string     `json:"reviewResult" form:"reviewResult" gorm:"column:review_result;comment:审核结果;size:191;"`       //审核结果
	ReviewStatus    string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`       //审核状态
	UserId          *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                             //管理ID 销售代表
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"` //客户名
	//crm_order 表
	OrderName string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
}
