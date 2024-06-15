// 自动生成模板CrmTicketCategories
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 工单类别 结构体  CrmTicketCategories
type CrmTicketCategories struct {
 global.GVA_MODEL 
      CategoryName  string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:工单类别名称;size:191;"`  //工单类别名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:类别描述;"`  //类别描述 
}


// TableName 工单类别 CrmTicketCategories自定义表名 crm_ticket_categories
func (CrmTicketCategories) TableName() string {
  return "crm_ticket_categories"
}

