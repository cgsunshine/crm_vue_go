// 自动生成模板CrmAuthorities
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 数据权限 结构体  CrmAuthorities
type CrmAuthorities struct {
	global.GVA_MODEL
	AuthorityId *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:角色id;size:10;"` //角色id
	Status      string `json:"status" form:"status" gorm:"column:status;comment:数据权限;size:10;"`                 //数据权限  1 个人 2 所有
}

// TableName 数据权限 CrmAuthorities自定义表名 crm_authorities
func (CrmAuthorities) TableName() string {
	return "crm_authorities"
}
