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

// CreateCrmContactApprovalTasks 创建合同审批多人
// @Tags CrmContactApprovalTasks
// @Summary 创建合同审批多人
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "创建合同审批多人"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactApprovalTasks/createCrmContactApprovalTasks [post]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) CreateCrmMultipleContactApprovalTasks(c *gin.Context) {
	var crmContactApprovalTasks crm.CrmReqContactApprovalTasks
	err := c.ShouldBindJSON(&crmContactApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	contact, err := crmContractService.GetCrmContract(strconv.Itoa(*crmContactApprovalTasks.ContactId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !(contact.ReviewStatus == comm.Approval_Status_Under || contact.ReviewStatus == comm.Approval_Status_Rafuse) {
		response.FailWithMessage("审批中或者已通过，请勿重复提交", c)
		return
	}

	for _, assigneeId := range crmContactApprovalTasks.AssigneeId {
		if err := crmContactApprovalTasksService.CreateCrmContactApprovalTasks(&crm.CrmContactApprovalTasks{
			AssigneeId:     assigneeId,
			ApprovalStatus: "1",
			ContactId:      crmContactApprovalTasks.ContactId,
			Valid:          utils.Pointer(1),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	//修改合同状态为审批中
	err = crmContractService.UpdApprovalStatus(crmContactApprovalTasks.ContactId, map[string]interface{}{
		"review_status": comm.Approval_Status_Pending,
	})
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// GetCrmContactApprovalTasksList 分页获取合同审批表列表
// @Tags CrmContactApprovalTasks
// @Summary 分页获取合同审批表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalTasksSearch true "分页获取合同审批表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalTasks/getCrmContactApprovalTasksList [get]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) GetCrmPageContactApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmContactApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = userService.FindUserDataStatusById(&AssigneeId)

	if list, total, err := crmContactApprovalTasksService.GetCrmPageContactApprovalTasksInfoList(pageInfo); err != nil {
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
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) UpdateCrmMultipleContactApprovalTasks(c *gin.Context) {
	var crmContactApprovalTasks crm.CrmContactApprovalTasks
	err := c.ShouldBindJSON(&crmContactApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	//AssigneeId, _ := strconv.Atoi()

	id := strconv.Itoa(int(crmContactApprovalTasks.ID))

	cats, err := crmContactApprovalTasksService.GetCrmContactApprovalTasks(id)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	if cats.ApprovalStatus != comm.Approval_Status_Under {
		response.FailWithMessage("审核已操作，不能重复提交", c)
		return
	}

	if err := crmContactApprovalTasksService.UpdCrmContactApprovalTasks(crmContactApprovalTasks.ID, map[string]interface{}{
		"approval_status": crmContactApprovalTasks.ApprovalStatus,
		"comments":        crmContactApprovalTasks.Comments,
	}); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {

		//审批更新成功以后，根据审批状态，更新合同审批状态
		if comm.Approval_Status_Pass == crmContactApprovalTasks.ApprovalStatus {
			//审批通过检查所有人是否都同意
			ok, err := crmContactApprovalTasksService.GetCrmQueryApproved(*crmContactApprovalTasks.ContactId, comm.Approval_Status_Pass)
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}
			if ok {
				//所有人都同意，修改合同状态
				err = crmContractService.UpdApprovalStatus(crmContactApprovalTasks.ContactId, map[string]interface{}{
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

		if comm.Approval_Status_Rafuse == crmContactApprovalTasks.ApprovalStatus {
			//拒绝审批，中断审批流程
			//有人拒绝，修改合同状态
			err = crmContractService.UpdApprovalStatus(crmContactApprovalTasks.ContactId, map[string]interface{}{
				"review_status": comm.Approval_Status_Rafuse,
				//"review_result": "newEmail@example.com",
			})
			if err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}

			//在将任务审批标记一下 改成失效

			if err = crmContactApprovalTasksService.UpdCrmContactIDContactApprovalTasks(*crmContactApprovalTasks.ContactId, map[string]interface{}{
				"valid": comm.Contact_Approval_Tasks_Valid_Invalid,
			}); err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				response.FailWithMessage("更新失败", c)
				return
			}
		}

		if err := crmContactApprovalRecordService.CreateCrmContactApprovalRecord(&crm.CrmContactApprovalRecord{
			ApproveOpinion: crmContactApprovalTasks.Comments,
			//ApproverId:     utils.Pointer(int(crmContactApprovalTasks.ID)),
			ApproverId: &AssigneeId,
			ContactId:  crmContactApprovalTasks.ContactId,
			Status:     crmContactApprovalTasks.ApprovalStatus,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		}
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPageContactApprovalTasks 用id查询合同审批
// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) FindCrmPageContactApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmContactApprovalTasksService.GetCrmPageContactApprovalTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}
