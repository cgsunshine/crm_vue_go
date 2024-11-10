package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

// GetCrmContractApprovalTasksList 合同分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmContractApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//pageInfo.AssigneeId = GetSearchUserId(nil, c)
	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)
	pageInfo.ApprovalType = utils.Pointer(comm.ContractApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmContractApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityContractApprovalTasksList 商机分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmBusinessOpportunityContractApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.BusinessOpportunityApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmBusinessOpportunityApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityContractApprovalTasksList 商机分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmDepositsApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.DepositsApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmDepositsApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmPaymentCollentionApprovalTasksList 回款分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmPaymentCollentionApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.PaymentCollentionApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmPaymentCollentionApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmPaymentCollentionApprovalTasksList 订单分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmOrderApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.OrderApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmOrderApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmStatementAccountApprovalTasksList 对账单分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmStatementAccountApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.StatementAccountApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmStatementAccountApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmPaymentApprovalTasksList 付款分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmPaymentApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmPaymentApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.PaymentApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmPaymentApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmPurchaseOrderApprovalTasksList 订购单分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmPaymentApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmPurchaseOrderApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.PurchaseOrderApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmPurchaseOrderApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmProcurementContractApprovalTasksList 订购合同分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmPaymentApprovalTasksList [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) GetCrmProcurementContractApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)
	pageInfo.RoleId = userService.FindUserDataStatusByIdRole(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.ProcurementContractApprovalType)

	if list, total, err := crmApprovalTasksRoleService.GetCrmProcurementContractApprovalTasksInfoList(pageInfo); err != nil {
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

// UpdateCrmMultipleApprovalTasks 更新合同审批
// @Tags UpdateCrmMultipleApprovalTasks
// @Summary 更新合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "更新合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalTasks/updateCrmContactApprovalTasks [put]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) UpdateCrmMultipleApprovalTasks(c *gin.Context) {
	var CrmApprovalTasks crm.CrmApprovalTasks
	err := c.ShouldBindJSON(&CrmApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	//AssigneeId, _ := strconv.Atoi()

	id := strconv.Itoa(int(CrmApprovalTasks.ID))

	cats, err := crmApprovalTasksRoleService.GetCrmApprovalTasksRole(id)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	if cats.ApprovalStatus != comm.Approval_Status_Under {
		response.FailWithMessage("审核已操作，不能重复提交", c)
		return
	}

	if *cats.Valid == comm.Contact_Approval_Tasks_Valid_Invalid {
		response.FailWithMessage("审核流程结束，无需操作", c)
		return
	}

	if err := crmApprovalTasksRoleService.UpdCrmApprovalTasks(CrmApprovalTasks.ID, map[string]interface{}{
		"approval_status": CrmApprovalTasks.ApprovalStatus,
		"comments":        CrmApprovalTasks.Comments,
		"assignee_id":     AssigneeId,
	}); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {

		//审批更新成功以后，根据审批状态，更新合同审批状态
		if comm.Approval_Status_Pass == CrmApprovalTasks.ApprovalStatus {
			//审批通过检查所有人是否都同意
			//ok, err := crmContactApprovalTasksService.GetCrmQueryApproved(*crmContactApprovalTasks.ContactId, comm.Approval_Status_Pass)
			//if err != nil {
			//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
			//	response.FailWithMessage("更新失败", c)
			//	return
			//}

			//查询审批需要的人数
			node, err := crmApprovalNodeService.GetCrmApprovalNodeId(*cats.StepId)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
			}
			ok, err := crmApprovalTasksRoleService.GetCrmQueryStepApproved(*cats.AssociatedId, *node.NumberApprovedPersonnel, *cats.ApprovalType, comm.Approval_Status_Pass)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			if ok {

				nodeInfo, err := crmApprovalNodeService.GetCrmApprovalNodeId(*cats.StepId)
				if err != nil {
					global.GVA_LOG.Error("创建失败!", zap.Error(err))
					response.FailWithMessage("创建失败", c)
					return
				}

				//这里判断审批走的逻辑是不是最后一步如果是终止路程，如果不是要进入下一步审批
				if crmConfigService.GetCrmLastNameToConfig(comm.ApprovalConfigToType[*cats.ApprovalType], nodeInfo.NodeOrder) {
					err := ApprovalEntrance[*cats.ApprovalType].ApprovalProcessSuccess(cats.AssociatedId)
					if err != nil {
						global.GVA_LOG.Error("更新失败!", zap.Error(err))
						response.FailWithMessage("更新失败", c)
						return
					}

					//在将任务审批标记一下 通过审批也要终止流程
					if err = crmApprovalTasksRoleService.UpdCrmAssociatedIdApprovalTasks(*cats.AssociatedId, *cats.ApprovalType, map[string]interface{}{
						"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
					}); err != nil {
						global.GVA_LOG.Error("更新失败!", zap.Error(err))
						response.FailWithMessage("更新失败", c)
						return
					}
				} else {
					//不是最后一步，走下一步流程

					//在查出第一步对应的角色id
					roleInfo, err := crmConfigService.GetCrmNextNameToConfig(comm.ApprovalConfigToType[*cats.ApprovalType], nodeInfo.NodeOrder)
					if err != nil {
						global.GVA_LOG.Error("创建失败!", zap.Error(err))
						response.FailWithMessage("创建失败", c)
						return
					}

					//ids, err := userService.GetRoleUsers(roleInfo.RoleIds)
					//if err != nil {
					//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
					//	response.FailWithMessage("创建失败", c)
					//	return
					//}

					items := strings.Split(roleInfo.RoleIds, ",")
					for _, item := range items {
						roleId, _ := strconv.Atoi(item)
						if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crm.CrmApprovalTasksRole{
							RoleId:         &roleId,
							ApprovalStatus: comm.Approval_Status_Under,
							AssociatedId:   cats.AssociatedId,
							Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
							StepId:         roleInfo.NodeId,
							ApprovalType:   utils.Pointer(comm.ContractApprovalType),
						}); err != nil {
							global.GVA_LOG.Error("创建失败!", zap.Error(err))
							response.FailWithMessage("创建失败", c)
							return
						}
					}

					//插入角色id对应的用户的审批记录
					//for _, userAuth := range ids {
					//	assigneeId := int(userAuth.SysUserId)
					//
					//	if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crm.CrmApprovalTasksRole{
					//		AssigneeId:     &assigneeId,
					//		ApprovalStatus: comm.Approval_Status_Under,
					//		AssociatedId:   cats.AssociatedId,
					//		Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
					//		StepId:         roleInfo.NodeId,
					//		ApprovalType:   utils.Pointer(comm.ContractApprovalType),
					//	}); err != nil {
					//		global.GVA_LOG.Error("创建失败!", zap.Error(err))
					//		response.FailWithMessage("创建失败", c)
					//		return
					//	}
					//}
				}

			}
		}

		if comm.Approval_Status_Rafuse == CrmApprovalTasks.ApprovalStatus {
			//拒绝审批，中断审批流程
			////有人拒绝，修改合同状态
			//err = crmBusinessOpportunityService.UpdApprovalStatus(CrmApprovalTasks.BusinessOpportunityId, map[string]interface{}{
			//	"review_status": comm.Approval_Status_Rafuse,
			//	//"review_result": "newEmail@example.com",
			//})
			//if err != nil {
			//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
			//	response.FailWithMessage("更新失败", c)
			//	return
			//}

			err := ApprovalEntrance[comm.ContractApprovalType].ApprovalProcessFail(cats.AssociatedId)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			//在将任务审批标记一下 改成失效 终止审批流程
			if err = crmApprovalTasksRoleService.UpdCrmAssociatedIdApprovalTasks(*cats.AssociatedId, *cats.ApprovalType, map[string]interface{}{
				"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
			}); err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}
		}

		if err := crmApprovalRecordService.CreateCrmApprovalRecord(&crm.CrmApprovalRecord{
			ApproveOpinion: CrmApprovalTasks.Comments,
			//ApproverId:     utils.Pointer(int(crmContactApprovalTasks.ID)),
			ApproverId:   &AssigneeId,
			AssociatedId: CrmApprovalTasks.AssociatedId,
			Status:       CrmApprovalTasks.ApprovalStatus,
			ApprovalType: cats.ApprovalType,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		}
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateCrmMultipleApprovalTasksBack 更新合同审批备份
// @Tags UpdateCrmMultipleApprovalTasks
// @Summary 更新合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "更新合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalTasks/updateCrmContactApprovalTasks [put]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) UpdateCrmMultipleApprovalTasksBack(c *gin.Context) {
	var CrmApprovalTasks crm.CrmApprovalTasks
	err := c.ShouldBindJSON(&CrmApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	//AssigneeId, _ := strconv.Atoi()

	id := strconv.Itoa(int(CrmApprovalTasks.ID))

	cats, err := crmApprovalTasksRoleService.GetCrmApprovalTasksRole(id)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	if cats.ApprovalStatus != comm.Approval_Status_Under {
		response.FailWithMessage("审核已操作，不能重复提交", c)
		return
	}

	if *cats.Valid == comm.Contact_Approval_Tasks_Valid_Invalid {
		response.FailWithMessage("审核流程结束，无需操作", c)
		return
	}

	if err := crmApprovalTasksRoleService.UpdCrmApprovalTasks(CrmApprovalTasks.ID, map[string]interface{}{
		"approval_status": CrmApprovalTasks.ApprovalStatus,
		"comments":        CrmApprovalTasks.Comments,
	}); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {

		//审批更新成功以后，根据审批状态，更新合同审批状态
		if comm.Approval_Status_Pass == CrmApprovalTasks.ApprovalStatus {
			//审批通过检查所有人是否都同意
			//ok, err := crmContactApprovalTasksService.GetCrmQueryApproved(*crmContactApprovalTasks.ContactId, comm.Approval_Status_Pass)
			//if err != nil {
			//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
			//	response.FailWithMessage("更新失败", c)
			//	return
			//}

			//查询审批需要的人数
			node, err := crmApprovalNodeService.GetCrmApprovalNodeId(*cats.StepId)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
			}
			ok, err := crmApprovalTasksRoleService.GetCrmQueryStepApproved(*cats.AssociatedId, *node.NumberApprovedPersonnel, *cats.ApprovalType, comm.Approval_Status_Pass)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			if ok {

				nodeInfo, err := crmApprovalNodeService.GetCrmApprovalNodeId(*cats.StepId)
				if err != nil {
					global.GVA_LOG.Error("创建失败!", zap.Error(err))
					response.FailWithMessage("创建失败", c)
					return
				}

				//这里判断审批走的逻辑是不是最后一步如果是终止路程，如果不是要进入下一步审批
				if crmConfigService.GetCrmLastNameToConfig(comm.ApprovalConfigToType[*cats.ApprovalType], nodeInfo.NodeOrder) {
					if *cats.ApprovalType == comm.ContractApprovalType {
						//合同审批
						//所有人都同意，修改合同状态
						err = crmContractService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
							"review_status": comm.Approval_Status_Pass,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}
					}

					if *cats.ApprovalType == comm.BusinessOpportunityApprovalType {
						//商机审批
						//所有人都同意，修改商机状态
						err = crmBusinessOpportunityService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
							"review_status": comm.Approval_Status_Pass,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}
					}

					if *cats.ApprovalType == comm.PaymentCollentionApprovalType {
						//回款审批
						//所有人都同意，修改商机状态
						err = crmPaymentCollentionService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
							"review_status": comm.Approval_Status_Pass,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						pc, err := crmPaymentCollentionService.GetCrmPaymentIdCollention(*cats.AssociatedId)
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						//通过审批以后，商机更新为已关单
						err = crmBusinessOpportunityService.UpdApprovalStatus(pc.BusinessOpportunityId, map[string]interface{}{
							"status": comm.BusinessOpportunityStatus,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						order, err := crmOrderService.GetCrmOrderId(pc.OrderId)
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						//通过审批以后，商机更新为已关单
						err = crmBusinessOpportunityService.UpdApprovalStatus(order.BusinessOpportunityId, map[string]interface{}{
							"status": comm.BusinessOpportunityStatus,
						})

						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						//跟新账单状态 付款状态
						err = crmBillService.UpdApprovalStatus(pc.BillId, map[string]interface{}{
							"payment_status": comm.PaymentStatusPaid,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}

						if pc.PaymentCollentionType == comm.Pre_Deposit || pc.PaymentCollentionType == comm.Advance_charge {
							//更新用户余额 包含特定条件更新
							err = crmCustomersService.UpdateBalanceIncrease(*pc.CustomerId, *pc.Amount)
							if err != nil {
								global.GVA_LOG.Error("更新失败!", zap.Error(err))
								response.FailWithMessage("更新失败", c)
								return
							}
						}
					}

					if *cats.ApprovalType == comm.OrderApprovalType {
						//订单审批
						//所有人都同意，修改订单状态
						err = crmOrderService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
							"review_status": comm.Approval_Status_Pass,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}
					}

					if *cats.ApprovalType == comm.DepositsApprovalType {
						//押金审批
						//所有人都同意，修改押金状态
						err = crmDepositsService.UpdDepositsInfo(cats.AssociatedId, map[string]interface{}{
							"review_status":   comm.Approval_Status_Pass,
							"deposits_status": comm.Deposits_Status_Payment,
						})
						if err != nil {
							global.GVA_LOG.Error("更新失败!", zap.Error(err))
							response.FailWithMessage("更新失败", c)
							return
						}
					}

				} else {
					//不是最后一步，走下一步流程

					//在查出第一步对应的角色id
					roleInfo, err := crmConfigService.GetCrmNextNameToConfig(comm.ApprovalConfigToType[*cats.ApprovalType], nodeInfo.NodeOrder)
					if err != nil {
						global.GVA_LOG.Error("创建失败!", zap.Error(err))
						response.FailWithMessage("创建失败", c)
						return
					}

					//ids, err := userService.GetRoleUsers(roleInfo.RoleIds)
					//if err != nil {
					//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
					//	response.FailWithMessage("创建失败", c)
					//	return
					//}

					items := strings.Split(roleInfo.RoleIds, ",")
					for _, item := range items {
						roleId, _ := strconv.Atoi(item)
						if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crm.CrmApprovalTasksRole{
							RoleId:         &roleId,
							ApprovalStatus: comm.Approval_Status_Under,
							AssociatedId:   cats.AssociatedId,
							Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
							StepId:         roleInfo.NodeId,
							ApprovalType:   utils.Pointer(comm.ContractApprovalType),
						}); err != nil {
							global.GVA_LOG.Error("创建失败!", zap.Error(err))
							response.FailWithMessage("创建失败", c)
							return
						}
					}

					//插入角色id对应的用户的审批记录
					//for _, userAuth := range ids {
					//	assigneeId := int(userAuth.SysUserId)
					//	//这里需要修改，只插入一条
					//
					//}
				}

				//在将任务审批标记一下 通过审批也要终止流程
				if err = crmApprovalTasksRoleService.UpdCrmAssociatedIdApprovalTasks(*cats.AssociatedId, *cats.ApprovalType, map[string]interface{}{
					"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
				}); err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}

			}
		}

		if comm.Approval_Status_Rafuse == CrmApprovalTasks.ApprovalStatus {
			//拒绝审批，中断审批流程
			////有人拒绝，修改合同状态
			//err = crmBusinessOpportunityService.UpdApprovalStatus(CrmApprovalTasks.BusinessOpportunityId, map[string]interface{}{
			//	"review_status": comm.Approval_Status_Rafuse,
			//	//"review_result": "newEmail@example.com",
			//})
			//if err != nil {
			//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
			//	response.FailWithMessage("更新失败", c)
			//	return
			//}

			if *cats.ApprovalType == comm.ContractApprovalType {
				//合同审批
				//有人拒绝，修改合同状态
				err = crmContractService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			if *cats.ApprovalType == comm.BusinessOpportunityApprovalType {
				//商机审批
				//有人拒绝，修改商机状态
				err = crmBusinessOpportunityService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			if *cats.ApprovalType == comm.PaymentCollentionApprovalType {
				//回款审批
				//有人拒绝，修改审批状态
				err = crmPaymentCollentionService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			if *cats.ApprovalType == comm.OrderApprovalType {
				//订单审批
				//有人拒绝，修改订单状态
				err = crmOrderService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			if *cats.ApprovalType == comm.DepositsApprovalType {
				//订单审批
				//有人拒绝，修改订单状态
				err = crmDepositsService.UpdDepositsInfo(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			//在将任务审批标记一下 改成失效 终止审批流程
			if err = crmApprovalTasksRoleService.UpdCrmAssociatedIdApprovalTasks(*cats.AssociatedId, *cats.ApprovalType, map[string]interface{}{
				"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
			}); err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}
		}

		if err := crmApprovalRecordService.CreateCrmApprovalRecord(&crm.CrmApprovalRecord{
			ApproveOpinion: CrmApprovalTasks.Comments,
			//ApproverId:     utils.Pointer(int(crmContactApprovalTasks.ID)),
			ApproverId:   &AssigneeId,
			AssociatedId: CrmApprovalTasks.AssociatedId,
			Status:       CrmApprovalTasks.ApprovalStatus,
			ApprovalType: cats.ApprovalType,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPageContactApprovalTasks 用id查询合同审批
// @Tags FindCrmPageContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) FindCrmPageContactApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksRoleService.GetCrmPageContractApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}

// FindCrmPageBusinessOpportunityApprovalTasks 用id查询商机审批
// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) FindCrmPageBusinessOpportunityApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksRoleService.GetCrmPageBusinessOpportunityApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}

// FindCrmPageBusinessOpportunityApprovalTasks 用id查询商机审批
// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) FindCrmPagePaymentCollentionApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksRoleService.GetCrmPagePaymentCollentionApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}

// FindCrmPageDepositsApprovalTasks 用id查询押金审批
// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询押金审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmApprovalTasksRoleApi *CrmApprovalTasksRoleApi) FindCrmPageDepositsApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksRoleService.GetCrmPageDepositsApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}
