// 自动生成模板CrmSupplier
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmSupplier表 结构体  CrmSupplier
type CrmSupplier struct {
	global.GVA_MODEL
	CommpanyName  string     `json:"commpanyName" form:"commpanyName" gorm:"column:commpany_name;comment:公司名称;size:191;" binding:"required"`   //公司名称
	ContactPerson string     `json:"contactPerson" form:"contactPerson" gorm:"column:contact_person;comment:联系人;size:191;" binding:"required"` //联系人
	Email         string     `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:191;"`                                              //邮箱
	Telephone     string     `json:"telephone" form:"telephone" gorm:"column:telephone;comment:电话;size:191;"`                                  //电话
	NoteAddTime   *time.Time `json:"noteAddTime" form:"noteAddTime" gorm:"column:note_add_time;comment:备注添加时间;"`                               //备注添加时间
	UserId        *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人id;"`                                                //负责人id
	Description   string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                            //备注
}

// TableName crmSupplier表 CrmSupplier自定义表名 crm_supplier
func (CrmSupplier) TableName() string {
	return "crm_supplier"
}
