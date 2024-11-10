package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm/custom"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdminHome struct{}

// CreateCrmApprovalNode 首页数据
// @Tags CrmApprovalNode
// @Summary 创建审批节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalNode true "创建审批节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalNode/createCrmApprovalNode [post]
func (adminHome *AdminHome) HomeData(c *gin.Context) {

	var search crmReq.CrmAdminHomeSearch
	err := c.ShouldBindQuery(&search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var roleId *int

	userId := GetSearchUserId(search.UserId, c)

	//AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	if userId != nil {
		roleId = userService.FindUserDataStatusByIdRole(userId)
	}

	businessOpportunityApproval, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.BusinessOpportunityApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	OrderApproval, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.OrderApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	salesContract, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.ContractApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	paymentCollentionApproval, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.PaymentCollentionApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	depositsApproval, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.DepositsApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	paymentApproval, err := crmApprovalTasksRoleService.ApprovalTasksCountRole(roleId, comm.Approval_Status_Under, comm.PaymentApprovalType, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	businessOpportunityCycleTimeAddRelation, err := crmBusinessOpportunityService.BusinessOpportunityCycleTimeAdd(userId, search.StartCreatedAt, search.EndCreatedAt, comm.BusinessOpportunityRelationType)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	businessOpportunityCycleTimeAddUnrelated, err := crmBusinessOpportunityService.BusinessOpportunityCycleTimeAdd(userId, search.StartCreatedAt, search.EndCreatedAt, comm.BusinessOpportunityUnrelatedType)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	orderCycleTimeAdd, err := crmOrderService.OrderCycleTimeAdd(userId, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	contractCycle, err := crmContractService.ContractCycleTimeAdd(userId, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	contractPendingRenewalCycle, err := crmContractService.ContractPendingRenewalCycleTime(userId, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	contractExpiredNotRenewed, err := crmContractService.ContractExpiredNotRenewedCycleTime(userId, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	paymentCollentionAmount, err := crmPaymentCollentionService.PaymentCollentionAmountCycleTime(userId, search.StartCreatedAt, search.EndCreatedAt)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(custom.CrmHomeData{
		PendingApprovalTask: &custom.PendingApprovalTask{
			BusinessOpportunity: businessOpportunityApproval,
			Order:               OrderApproval,
			SalesContract:       salesContract,
			PaymentCollection:   paymentCollentionApproval,
			Deposit:             depositsApproval,
			Payment:             paymentApproval,
		},
		AddBusiness: &custom.AddBusiness{
			BusinessOpportunity:         businessOpportunityCycleTimeAddRelation,
			BusinessOpportunityRelation: businessOpportunityCycleTimeAddUnrelated,
			Order:                       orderCycleTimeAdd,
		},
		CustomerInfo: &custom.CustomerInfo{
			TransactionCustomer:            10,
			PublicPoolCustomers:            11,
			PublicPoolTransactionCustomers: 12,
		},
		ContactInfo: &custom.ContactInfo{
			AddContact:        contractCycle,
			PendingRenewal:    contractPendingRenewalCycle,
			ExpiredNotRenewed: contractExpiredNotRenewed,
		},
		BusinessFollowUp: &custom.BusinessFollowUp{
			Customer:          15,
			PaymentCollection: paymentCollentionAmount,
		},
		PerformanceTarget: 60,
	}, "获取成功", c)

}
