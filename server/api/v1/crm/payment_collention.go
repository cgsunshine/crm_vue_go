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
	"time"
)

// GetCrmPaymentCollentionList 分页获取crmPaymentCollention表列表
// @Tags CrmPaymentCollention
// @Summary 分页获取crmPaymentCollention表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionSearch true "分页获取crmPaymentCollention表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollention/getCrmPaymentCollentionList [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) GetCrmPagePaymentCollentionInfoList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentCollentionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmPaymentCollentionService.GetCrmPagePaymentCollentionInfoList(pageInfo); err != nil {
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

// FindCrmPagePaymentCollention 用id查询crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 用id查询crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPaymentCollention true "用id查询crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollention/findCrmPaymentCollention [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) FindCrmPagePaymentCollention(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPaymentCollention, err := crmPaymentCollentionService.GetCrmPagePaymentCollention(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		recrmPaymentCollention.Proof, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmPaymentCollention.Proof)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmPaymentCollention": recrmPaymentCollention}, c)
	}
}

// CreateCrmPagePaymentCollention 创建crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 创建crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollention true "创建crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollention/createCrmPaymentCollention [post]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) CreateCrmPagePaymentCollention(c *gin.Context) {
	var crmPaymentCollention crm.CrmPaymentCollention
	err := c.ShouldBindJSON(&crmPaymentCollention)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	t := time.Now()
	crmPaymentCollention.UserId = comm.GetHeaderUserId(c)
	crmPaymentCollention.PaymentTime = &t

	crmPaymentCollention.UserId = comm.GetHeaderUserId(c)

	crmPaymentCollention.ReviewStatus = comm.Approval_Status_Pending

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.PaymentCollentionApproval)
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
	crmPaymentCollention.PaymentCollentionNumber = comm.GetBusinessNumber(comm.PaymentCollentionNumberPrefix, crmPaymentCollentionService.GetCrmPaymentCollentionTodayCount())
	if err := crmPaymentCollentionService.CreateCrmPaymentCollention(&crmPaymentCollention); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}

	err = crmBillService.UpdAssOrderID(crmPaymentCollention.OrderId, map[string]interface{}{
		"payment_collention_id": crmPaymentCollention.ID,
	})

	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	paymentCollentionId := int(crmPaymentCollention.ID)

	//插入角色id对应的用户的审批记录
	//for _, userAuth := range ids {
	//	assigneeId := int(userAuth.SysUserId)
	//	if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
	//		AssigneeId:     &assigneeId,
	//		ApprovalStatus: comm.Approval_Status_Under,
	//		AssociatedId:   &paymentCollentionId,
	//		Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
	//		StepId:         roleInfo.NodeId,
	//		ApprovalType:   utils.Pointer(comm.PaymentCollentionApprovalType),
	//	}); err != nil {
	//		global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//		response.FailWithMessage("创建失败", c)
	//		return
	//	}
	//}

	items := strings.Split(roleInfo.RoleIds, ",")
	for _, item := range items {
		roleId, _ := strconv.Atoi(item)
		if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crm.CrmApprovalTasksRole{
			RoleId:         &roleId,
			ApprovalStatus: comm.Approval_Status_Under,
			AssociatedId:   &paymentCollentionId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.PaymentCollentionApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)
}
