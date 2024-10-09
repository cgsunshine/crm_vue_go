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
	"strconv"
)

type CrmDepositsApi struct {
}

var crmDepositsService = service.ServiceGroupApp.CrmServiceGroup.CrmDepositsService

// CreateCrmDeposits 创建押金管理
// @Tags CrmDeposits
// @Summary 创建押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmDeposits true "创建押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmDeposits/createCrmDeposits [post]
func (crmDepositsApi *CrmDepositsApi) CreateCrmDeposits(c *gin.Context) {
	var crmDeposits crm.CrmDeposits
	err := c.ShouldBindJSON(&crmDeposits)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//押金创建直接提交审核
	crmDeposits.DepositsStatus = comm.Deposits_Status_Under
	crmDeposits.ReviewStatus = comm.Approval_Status_Pending
	crmDeposits.RefundStatus = comm.RUnsubmitted_Refund_Status
	crmDeposits.RefundDate = nil
	crmDeposits.UserId = comm.GetHeaderUserId(c)
	//生成编号
	crmDeposits.DepositsNumber = comm.GetBusinessNumber(comm.DepositsNumberPrefix, crmDepositsService.GetCrmBillTodayCount())
	if err := crmDepositsService.CreateCrmDeposits(&crmDeposits); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}
	depositsId := int(crmDeposits.ID)

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.DepositsApproval)
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
			AssociatedId:   &depositsId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.DepositsApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)

}

// DeleteCrmDeposits 删除押金管理
// @Tags CrmDeposits
// @Summary 删除押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmDeposits true "删除押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmDeposits/deleteCrmDeposits [delete]
func (crmDepositsApi *CrmDepositsApi) DeleteCrmDeposits(c *gin.Context) {
	ID := c.Query("ID")
	recrmDeposits, err := crmDepositsService.GetCrmDeposits(ID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	if recrmDeposits.DepositsStatus != comm.Deposits_Status_Under {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("当前状态不能删除", c)
		return
	}
	if err := crmDepositsService.DeleteCrmDeposits(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}

	//关联删除审批
	err = crmApprovalTasksService.DelCrmAssociatedIdApprovalTasks(ID, comm.DepositsApprovalType)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}

	response.OkWithMessage("删除成功", c)
}

// DeleteCrmDepositsByIds 批量删除押金管理
// @Tags CrmDeposits
// @Summary 批量删除押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmDeposits/deleteCrmDepositsByIds [delete]
func (crmDepositsApi *CrmDepositsApi) DeleteCrmDepositsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmDepositsService.DeleteCrmDepositsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmDeposits 更新押金管理
// @Tags CrmDeposits
// @Summary 更新押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmDeposits true "更新押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmDeposits/updateCrmDeposits [put]
func (crmDepositsApi *CrmDepositsApi) UpdateCrmDeposits(c *gin.Context) {
	var crmDeposits crm.CrmDeposits
	err := c.ShouldBindJSON(&crmDeposits)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmDepositsService.UpdateCrmDeposits(crmDeposits); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmDeposits 用id查询押金管理
// @Tags CrmDeposits
// @Summary 用id查询押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmDeposits true "用id查询押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmDeposits/findCrmDeposits [get]
func (crmDepositsApi *CrmDepositsApi) FindCrmDeposits(c *gin.Context) {
	ID := c.Query("ID")
	if recrmDeposits, err := crmDepositsService.GetCrmDeposits(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {

		recrmDeposits.DepositsVoucher, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmDeposits.DepositsVoucher)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}

		recrmDeposits.RefundVoucher, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmDeposits.RefundVoucher)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}

		response.OkWithData(gin.H{"recrmDeposits": recrmDeposits}, c)
	}
}

// GetCrmDepositsList 分页获取押金管理列表
// @Tags CrmDeposits
// @Summary 分页获取押金管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmDepositsSearch true "分页获取押金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmDeposits/getCrmDepositsList [get]
func (crmDepositsApi *CrmDepositsApi) GetCrmDepositsList(c *gin.Context) {
	var pageInfo crmReq.CrmDepositsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)
	if list, total, err := crmDepositsService.GetCrmDepositsInfoList(pageInfo); err != nil {
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

// GetCrmDepositsPublic 不需要鉴权的押金管理接口
// @Tags CrmDeposits
// @Summary 不需要鉴权的押金管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmDepositsSearch true "分页获取押金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmDeposits/getCrmDepositsList [get]
func (crmDepositsApi *CrmDepositsApi) GetCrmDepositsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的押金管理接口信息",
	}, "获取成功", c)
}

// CreateCrmDepositsRefund 创建押金退款
// @Tags CrmDeposits
// @Summary 创建押金退款
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmDeposits true "创建押金退款"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmDeposits/createCrmDepositsRefund [post]
func (crmDepositsApi *CrmDepositsApi) CreateCrmDepositsRefund(c *gin.Context) {
	ID := c.Query("ID")
	//refundVoucher := c.Query("refundVoucher")
	recrmDeposits, err := crmDepositsService.GetCrmDeposits(ID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	if recrmDeposits.DepositsStatus != comm.Deposits_Status_Payment {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("当前状态不能退款，需要通过审批", c)
		return
	}

	if recrmDeposits.ReviewStatus != comm.Approval_Status_Pass {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("当前状态不能退款，审批通过可以退款", c)
		return
	}

	depositsId, _ := strconv.Atoi(ID)
	if err := crmRefundTasksService.CreateCrmRefundTasks(&crm.CrmRefundTasks{
		RefundStatus: comm.Approval_Status_Under,
		AssociatedId: &depositsId,
		RefundType:   utils.Pointer(comm.DepositsRefundType),
	}); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//提交完成以后，修改自己退款状态
	err = crmDepositsService.UpdDepositsInfo(&depositsId, map[string]interface{}{
		"refund_status": comm.Submitted_Refund_Status,
		//"refund_voucher": refundVoucher,
	})
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)

}
