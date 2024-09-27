package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CrmApprovalTasksRoleApi struct {
}

var crmApprovalTasksRoleService = service.ServiceGroupApp.CrmServiceGroup.CrmApprovalTasksRoleService


// CreateCrmApprovalTasksRole 创建角色审批表
// @Tags CrmApprovalTasksRole
// @Summary 创建角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasksRole true "创建角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalTasksRole/createCrmApprovalTasksRole [post]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) CreateCrmApprovalTasksRole(c *gin.Context) {
	var crmApprovalTasksRole crm.CrmApprovalTasksRole
	err := c.ShouldBindJSON(&crmApprovalTasksRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crmApprovalTasksRole); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmApprovalTasksRole 删除角色审批表
// @Tags CrmApprovalTasksRole
// @Summary 删除角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasksRole true "删除角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasksRole/deleteCrmApprovalTasksRole [delete]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) DeleteCrmApprovalTasksRole(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmApprovalTasksRoleService.DeleteCrmApprovalTasksRole(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmApprovalTasksRoleByIds 批量删除角色审批表
// @Tags CrmApprovalTasksRole
// @Summary 批量删除角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmApprovalTasksRole/deleteCrmApprovalTasksRoleByIds [delete]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) DeleteCrmApprovalTasksRoleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmApprovalTasksRoleService.DeleteCrmApprovalTasksRoleByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmApprovalTasksRole 更新角色审批表
// @Tags CrmApprovalTasksRole
// @Summary 更新角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasksRole true "更新角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalTasksRole/updateCrmApprovalTasksRole [put]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) UpdateCrmApprovalTasksRole(c *gin.Context) {
	var crmApprovalTasksRole crm.CrmApprovalTasksRole
	err := c.ShouldBindJSON(&crmApprovalTasksRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalTasksRoleService.UpdateCrmApprovalTasksRole(crmApprovalTasksRole); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmApprovalTasksRole 用id查询角色审批表
// @Tags CrmApprovalTasksRole
// @Summary 用id查询角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmApprovalTasksRole true "用id查询角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalTasksRole/findCrmApprovalTasksRole [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) FindCrmApprovalTasksRole(c *gin.Context) {
	ID := c.Query("ID")
	if recrmApprovalTasksRole, err := crmApprovalTasksRoleService.GetCrmApprovalTasksRole(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmApprovalTasksRole": recrmApprovalTasksRole}, c)
	}
}

// GetCrmApprovalTasksRoleList 分页获取角色审批表列表
// @Tags CrmApprovalTasksRole
// @Summary 分页获取角色审批表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksRoleSearch true "分页获取角色审批表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasksRole/getCrmApprovalTasksRoleList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmApprovalTasksRoleList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmApprovalTasksRoleService.GetCrmApprovalTasksRoleInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetCrmApprovalTasksRolePublic 不需要鉴权的角色审批表接口
// @Tags CrmApprovalTasksRole
// @Summary 不需要鉴权的角色审批表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksRoleSearch true "分页获取角色审批表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasksRole/getCrmApprovalTasksRoleList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmApprovalTasksRolePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的角色审批表接口信息",
    }, "获取成功", c)
}
