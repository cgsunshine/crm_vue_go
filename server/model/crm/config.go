// 自动生成模板CrmConfig
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 系统配置 结构体  CrmConfig
type CrmPageConfig struct {
	global.GVA_MODEL
	ConfigName    string `json:"configName" form:"configName" gorm:"column:config_name;comment:配置名称;size:10;"`               //配置名称
	Description   string `json:"description" form:"description" gorm:"column:description;comment:描述;size:199;"`              //描述
	AssociatedId  *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 可为空;size:10;"`     //关联id 可为空
	ConfigContent string `json:"configContent" form:"configContent" gorm:"column:config_content;comment:配置内容 可为空;size:255;"` //配置内容 可为空
	//以下是联表查询字段
	//crm_approval_node 表
	RoleIds   string `json:"roleIds" form:"roleIds" gorm:"column:roleIds;comment:角色ID;size:255;"`      //角色ID
	NodeId    *int   `json:"node" form:"node" gorm:"column:node;comment:角色ID;size:10;"`                //节点id
	NodeOrder *int   `json:"nodeOrder" form:"nodeOrder" gorm:"column:nodeOrder;comment:节点顺序;size:19;"` //节点顺序
}
