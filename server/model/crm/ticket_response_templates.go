// 自动生成模板CrmTicketResponseTemplates
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 快捷回复模板 结构体  CrmTicketResponseTemplates
type CrmPageTicketResponseTemplates struct {
	global.GVA_MODEL
	CategoryId   *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:适用工单类别ID;"`              //适用工单类别ID
	Content      string `json:"content" form:"content" gorm:"column:content;comment:回复内容;size:191;"`                   //回复内容
	CreatedBy    *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者ID;"`                    //创建者ID
	TemplateName string `json:"templateName" form:"templateName" gorm:"column:template_name;comment:回复模板名称;size:255;"` //回复模板名称
	//以下是联表查询字段
	//crm_ticket_categories 工单分类
	CategoryName string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:工单类别名称;size:191;"` //工单类别名称
}
