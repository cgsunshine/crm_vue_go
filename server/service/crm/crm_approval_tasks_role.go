package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmApprovalTasksRoleService struct {
}

// CreateCrmApprovalTasksRole 创建角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) CreateCrmApprovalTasksRole(crmApprovalTasksRole *crm.CrmApprovalTasksRole) (err error) {
	err = global.GVA_DB.Create(crmApprovalTasksRole).Error
	return err
}

// DeleteCrmApprovalTasksRole 删除角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) DeleteCrmApprovalTasksRole(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmApprovalTasksRole{}, "id = ?", ID).Error
	return err
}

// DeleteCrmApprovalTasksRoleByIds 批量删除角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) DeleteCrmApprovalTasksRoleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmApprovalTasksRole{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmApprovalTasksRole 更新角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) UpdateCrmApprovalTasksRole(crmApprovalTasksRole crm.CrmApprovalTasksRole) (err error) {
	err = global.GVA_DB.Save(&crmApprovalTasksRole).Error
	return err
}

// GetCrmApprovalTasksRole 根据ID获取角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmApprovalTasksRole(ID string) (crmApprovalTasksRole crm.CrmApprovalTasksRole, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmApprovalTasksRole).Error
	return
}

// GetCrmApprovalTasksRoleInfoList 分页获取角色审批表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) GetCrmApprovalTasksRoleInfoList(info crmReq.CrmApprovalTasksRoleSearch) (list []crm.CrmApprovalTasksRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	var crmApprovalTasksRoles []crm.CrmApprovalTasksRole
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ApprovalStatus != "" {
		db = db.Where("approval_status = ?", info.ApprovalStatus)
	}
	if info.ApprovalType != nil {
		db = db.Where("approval_type = ?", info.ApprovalType)
	}
	if info.AssigneeId != nil {
		db = db.Where("assignee_id = ?", info.AssigneeId)
	}
	if info.AssociatedId != nil {
		db = db.Where("associated_id = ?", info.AssociatedId)
	}
	if info.Comments != "" {
		db = db.Where("comments LIKE ?", "%"+info.Comments+"%")
	}
	if info.RequestId != nil {
		db = db.Where("request_id = ?", info.RequestId)
	}
	if info.RoleId != nil {
		db = db.Where("role_id = ?", info.RoleId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmApprovalTasksRoles).Error
	return crmApprovalTasksRoles, total, err
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmApprovalTasksRoleService *CrmApprovalTasksRoleService) ApprovalTasksCountRole(roleId *int, approvalStatus string, approvalType int, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmApprovalTasksRole{})
	SearchConditionRole(db, roleId, startDate, endDate)
	err = db.Where("approval_status = ? AND approval_type = ?", approvalStatus, approvalType).Debug().Count(&total).Error
	return
}
