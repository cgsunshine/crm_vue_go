// 自动生成模板CrmStatementAccountFileUploadAndDownloads
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 对账单上传文件 结构体  CrmStatementAccountFileUploadAndDownloads
type CrmStatementAccountFileUploadAndDownloads struct {
 global.GVA_MODEL 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`  //排序 
      StatementAccountId  *int `json:"statementAccountId" form:"statementAccountId" gorm:"column:statement_account_id;comment:对账打死你id;"`  //对账打死你id 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName 对账单上传文件 CrmStatementAccountFileUploadAndDownloads自定义表名 crm_statement_account_file_upload_and_downloads
func (CrmStatementAccountFileUploadAndDownloads) TableName() string {
  return "crm_statement_account_file_upload_and_downloads"
}

