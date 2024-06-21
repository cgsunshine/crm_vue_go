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
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmContractApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.ContractApprovalType)

	if list, total, err := crmApprovalTasksService.GetCrmContractApprovalTasksInfoList(pageInfo); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmBusinessOpportunityContractApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.BusinessOpportunityApprovalType)

	if list, total, err := crmApprovalTasksService.GetCrmBusinessOpportunityApprovalTasksInfoList(pageInfo); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmPaymentCollentionApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.PaymentCollentionApprovalType)

	if list, total, err := crmApprovalTasksService.GetCrmPaymentCollentionApprovalTasksInfoList(pageInfo); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmOrderApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	pageInfo.ApprovalType = utils.Pointer(comm.PaymentCollentionApprovalType)

	if list, total, err := crmApprovalTasksService.GetCrmOrderApprovalTasksInfoList(pageInfo); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) UpdateCrmMultipleApprovalTasks(c *gin.Context) {
	var CrmApprovalTasks crm.CrmApprovalTasks
	err := c.ShouldBindJSON(&CrmApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	//AssigneeId, _ := strconv.Atoi()

	id := strconv.Itoa(int(CrmApprovalTasks.ID))

	cats, err := crmApprovalTasksService.GetCrmApprovalTasks(id)
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

	if err := crmApprovalTasksService.UpdCrmApprovalTasks(CrmApprovalTasks.ID, map[string]interface{}{
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
			ok, err := crmBusinessOpportunityApprovalTasksService.GetCrmQueryStepApproved(*cats.AssociatedId, *node.NumberApprovedPersonnel, comm.Approval_Status_Pass)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			if ok {
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
						"status": comm.Approval_Status_Pass,
					})
					if err != nil {
						global.GVA_LOG.Error("更新失败!", zap.Error(err))
						response.FailWithMessage("更新失败", c)
						return
					}
				}

				if *cats.ApprovalType == comm.OrderApprovalType {
					//商机审批
					//所有人都同意，修改商机状态
					err = crmOrderService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
						"review_status": comm.Approval_Status_Pass,
					})
					if err != nil {
						global.GVA_LOG.Error("更新失败!", zap.Error(err))
						response.FailWithMessage("更新失败", c)
						return
					}
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
				//所有人都同意，修改合同状态
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
				//所有人都同意，修改商机状态
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
				//所有人都同意，修改商机状态
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
				//回款审批
				//所有人都同意，修改商机状态
				err = crmOrderService.UpdApprovalStatus(cats.AssociatedId, map[string]interface{}{
					"review_status": comm.Approval_Status_Rafuse,
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}

			//在将任务审批标记一下 改成失效 终止审批流程
			if err = crmApprovalTasksService.UpdCrmAssociatedIdApprovalTasks(*CrmApprovalTasks.AssociatedId, comm.PaymentCollentionApprovalType, map[string]interface{}{
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) FindCrmPageContactApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksService.GetCrmPageContractApprovalTasks(ID); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) FindCrmPageBusinessOpportunityApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksService.GetCrmPageBusinessOpportunityApprovalTasks(ID); err != nil {
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
func (crmApprovalTasksApi *CrmApprovalTasksApi) FindCrmPagePaymentCollentionApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmApprovalTasksService.GetCrmPagePaymentCollentionApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}
