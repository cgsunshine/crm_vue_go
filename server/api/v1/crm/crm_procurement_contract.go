package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmProcurementContractApi struct {
}

var crmProcurementContractService = service.ServiceGroupApp.CrmServiceGroup.CrmProcurementContractService

// CreateCrmProcurementContract 创建订购合同
// @Tags CrmProcurementContract
// @Summary 创建订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProcurementContract true "创建订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProcurementContract/createCrmProcurementContract [post]
func (crmProcurementContractApi *CrmProcurementContractApi) CreateCrmProcurementContract(c *gin.Context) {
	var crmProcurementContract crm.CrmProcurementContract
	err := c.ShouldBindJSON(&crmProcurementContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmProcurementContract.UserId = comm.GetHeaderUserId(c)
	crmProcurementContract.ReviewStatus = comm.Approval_Status_Pending
	crmProcurementContract.ContractStatus = comm.ProcurementContractStatusEffectiveType
	//生成编号
	crmProcurementContract.ProcurementContractNumber = comm.GetBusinessNumber(comm.ProcurementContractNumberPrefix, crmProcurementContractService.GetCrmBillTodayCount())

	if err := crmProcurementContractService.CreateCrmProcurementContract(&crmProcurementContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	procurementContractOrderId := int(crmProcurementContract.ID)

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.ProcurementContractApproval)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	ids, err := userService.GetRoleUsers(roleInfo.RoleIds)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//插入角色id对应的用户的审批记录
	for _, userAuth := range ids {
		assigneeId := int(userAuth.SysUserId)
		if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
			AssigneeId:     &assigneeId,
			ApprovalStatus: comm.Approval_Status_Under,
			AssociatedId:   &procurementContractOrderId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.ProcurementContractApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteCrmProcurementContract 删除订购合同
// @Tags CrmProcurementContract
// @Summary 删除订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProcurementContract true "删除订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProcurementContract/deleteCrmProcurementContract [delete]
func (crmProcurementContractApi *CrmProcurementContractApi) DeleteCrmProcurementContract(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmProcurementContractService.DeleteCrmProcurementContract(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmProcurementContractByIds 批量删除订购合同
// @Tags CrmProcurementContract
// @Summary 批量删除订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmProcurementContract/deleteCrmProcurementContractByIds [delete]
func (crmProcurementContractApi *CrmProcurementContractApi) DeleteCrmProcurementContractByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmProcurementContractService.DeleteCrmProcurementContractByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmProcurementContract 更新订购合同
// @Tags CrmProcurementContract
// @Summary 更新订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProcurementContract true "更新订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProcurementContract/updateCrmProcurementContract [put]
func (crmProcurementContractApi *CrmProcurementContractApi) UpdateCrmProcurementContract(c *gin.Context) {
	var crmProcurementContract crm.CrmProcurementContract
	err := c.ShouldBindJSON(&crmProcurementContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProcurementContractService.UpdateCrmProcurementContract(crmProcurementContract); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmProcurementContract 用id查询订购合同
// @Tags CrmProcurementContract
// @Summary 用id查询订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProcurementContract true "用id查询订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProcurementContract/findCrmProcurementContract [get]
func (crmProcurementContractApi *CrmProcurementContractApi) FindCrmProcurementContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProcurementContract, err := crmProcurementContractService.GetCrmProcurementContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		//recrmProcurementContract.ContractFile, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmProcurementContract.ContractFile)
		//if err != nil {
		//	global.GVA_LOG.Error("查询失败!", zap.Error(err))
		//	response.FailWithMessage("查询失败", c)
		//}
		response.OkWithData(gin.H{"recrmProcurementContract": recrmProcurementContract}, c)
	}
}

// GetCrmProcurementContractList 分页获取订购合同列表
// @Tags CrmProcurementContract
// @Summary 分页获取订购合同列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProcurementContractSearch true "分页获取订购合同列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProcurementContract/getCrmProcurementContractList [get]
func (crmProcurementContractApi *CrmProcurementContractApi) GetCrmProcurementContractList(c *gin.Context) {
	var pageInfo crmReq.CrmProcurementContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmProcurementContractService.GetCrmPageProcurementContractInfoList(pageInfo); err != nil {
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

// GetCrmProcurementContractPublic 不需要鉴权的订购合同接口
// @Tags CrmProcurementContract
// @Summary 不需要鉴权的订购合同接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProcurementContractSearch true "分页获取订购合同列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProcurementContract/getCrmProcurementContractList [get]
func (crmProcurementContractApi *CrmProcurementContractApi) GetCrmProcurementContractPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订购合同接口信息",
	}, "获取成功", c)
}
