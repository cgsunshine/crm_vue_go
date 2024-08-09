// 自动生成模板CrmStatementAccountFile
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 对账单文件 结构体  CrmStatementAccountFile
type CrmStatementAccountFile struct {
 global.GVA_MODEL 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:19;"`  //排序 
      StatementAccountId  *int `json:"statementAccountId" form:"statementAccountId" gorm:"column:statement_account_id;comment:对账单id;size:19;"`  //对账单id 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName 对账单文件 CrmStatementAccountFile自定义表名 crm_statement_account_file
func (CrmStatementAccountFile) TableName() string {
  return "crm_statement_account_file"
}

