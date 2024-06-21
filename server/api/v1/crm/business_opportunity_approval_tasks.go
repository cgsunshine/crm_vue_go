package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmPageBusinessOpportunityApprovalTasksList 分页获取商机审批列表
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 分页获取商机审批列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityApprovalTasksSearch true "分页获取商机审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/getCrmBusinessOpportunityApprovalTasksList [get]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) GetCrmPageBusinessOpportunityApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	if list, total, err := crmBusinessOpportunityApprovalTasksService.GetCrmBusinessOpportunityApprovalTasksInfoList(pageInfo); err != nil {
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

// UpdateCrmMultipleContactApprovalTasks 更新合同审批
// @Tags CrmContactApprovalTasks
// @Summary 更新合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "更新合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalTasks/updateCrmContactApprovalTasks [put]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) UpdateCrmMultipleBusinessOpportunityApprovalTasks(c *gin.Context) {
	var crmBusinessOpportunityApprovalTasks crm.CrmBusinessOpportunityApprovalTasks
	err := c.ShouldBindJSON(&crmBusinessOpportunityApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	//AssigneeId, _ := strconv.Atoi()

	id := strconv.Itoa(int(crmBusinessOpportunityApprovalTasks.ID))

	cats, err := crmBusinessOpportunityApprovalTasksService.GetCrmBusinessOpportunityApprovalTasks(id)
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

	if err := crmBusinessOpportunityApprovalTasksService.UpdCrmBusinessOpportunityIdApprovalTasks(*crmBusinessOpportunityApprovalTasks.BusinessOpportunityId, map[string]interface{}{
		"approval_status": crmBusinessOpportunityApprovalTasks.ApprovalStatus,
		"comments":        crmBusinessOpportunityApprovalTasks.Comments,
	}); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {

		//审批更新成功以后，根据审批状态，更新合同审批状态
		if comm.Approval_Status_Pass == crmBusinessOpportunityApprovalTasks.ApprovalStatus {
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
			ok, err := crmBusinessOpportunityApprovalTasksService.GetCrmQueryStepApproved(*crmBusinessOpportunityApprovalTasks.BusinessOpportunityId, *node.NumberApprovedPersonnel, comm.Approval_Status_Pass)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			if ok {
				//所有人都同意，修改商机状态
				err = crmBusinessOpportunityService.UpdApprovalStatus(crmBusinessOpportunityApprovalTasks.BusinessOpportunityId, map[string]interface{}{
					"review_status": comm.Approval_Status_Pass,
					//"review_result": "newEmail@example.com",
				})
				if err != nil {
					global.GVA_LOG.Error("更新失败!", zap.Error(err))
					response.FailWithMessage("更新失败", c)
					return
				}
			}
		}

		if comm.Approval_Status_Rafuse == crmBusinessOpportunityApprovalTasks.ApprovalStatus {
			//拒绝审批，中断审批流程
			//有人拒绝，修改合同状态
			err = crmBusinessOpportunityService.UpdApprovalStatus(crmBusinessOpportunityApprovalTasks.BusinessOpportunityId, map[string]interface{}{
				"review_status": comm.Approval_Status_Rafuse,
				//"review_result": "newEmail@example.com",
			})
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			//在将任务审批标记一下 改成失效

			if err = crmBusinessOpportunityApprovalTasksService.UpdCrmBusinessOpportunityIdApprovalTasks(*crmBusinessOpportunityApprovalTasks.BusinessOpportunityId, map[string]interface{}{
				"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
			}); err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}
		}

		if err := crmBusinessOpportunityApprovalRecordService.CreateCrmBusinessOpportunityApprovalRecord(&crm.CrmBusinessOpportunityApprovalRecord{
			ApproveOpinion: crmBusinessOpportunityApprovalTasks.Comments,
			//ApproverId:     utils.Pointer(int(crmContactApprovalTasks.ID)),
			ApproverId:            &AssigneeId,
			BusinessOpportunityId: crmBusinessOpportunityApprovalTasks.BusinessOpportunityId,
			Status:                crmBusinessOpportunityApprovalTasks.ApprovalStatus,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		}
		response.OkWithMessage("更新成功", c)
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
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) FindCrmPageBusinessOpportunityApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmBusinessOpportunityApprovalTasksService.GetCrmPageBusinessOpportunityApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}
