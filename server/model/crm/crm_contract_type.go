// 自动生成模板CrmContractType
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 合同类型 结构体  CrmContractType
type CrmContractType struct {
	global.GVA_MODEL
	ContractTypeName string `json:"contractTypeName" form:"contractTypeName" gorm:"column:contract_type_name;comment:合同类型名称;size:191;"` //合同类型名称
	Description      string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                      //备注
	Unit             string `json:"unit" form:"unit" gorm:"column:unit;comment:单位;size:191;"`                                           //单位
}

// TableName 合同类型 CrmContractType自定义表名 crm_contract_type
func (CrmContractType) TableName() string {
	return "crm_contract_type"
}
