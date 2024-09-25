// 自动生成模板CrmProcurementContract
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 订购合同 结构体  CrmProcurementContract
type CrmPageProcurementContract struct {
	global.GVA_MODEL
	ContractAmount *float64   `json:"contractAmount" form:"contractAmount" gorm:"column:contract_amount;comment:合同金额;" binding:"required"`    //合同金额
	ContractFile   string     `json:"contractFile" form:"contractFile" gorm:"column:contract_file;comment:合同文件;size:191;"`                    //合同文件
	ContractName   string     `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;" binding:"required"` //合同名称
	ContractStatus string     `json:"contractStatus" form:"contractStatus" gorm:"column:contract_status;comment:合同状态;size:191;"`              //合同状态
	CreationTime   *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`                             //创建时间
	ExpirationTime *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`                       //到期时间
	UserId         *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人;"`                                                //负责人
	ReviewStatus   string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                    //审核状态
	Description    string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                          //备注
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}
