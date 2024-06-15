// 自动生成模板CrmContactFile
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 合同文件 结构体  CrmContactFile
type CrmContactFile struct {
 global.GVA_MODEL 
      ContactId  *int `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:19;"`  //合同id 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:19;"`  //排序 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName 合同文件 CrmContactFile自定义表名 crm_contact_file
func (CrmContactFile) TableName() string {
  return "crm_contact_file"
}

