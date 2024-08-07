package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm/custom"
	"github.com/gin-gonic/gin"
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

	response.OkWithDetailed(custom.CrmHomeData{
		PendingApprovalTask: &custom.PendingApprovalTask{
			BusinessOpportunity: 2,
			Order:               2,
			SalesContract:       3,
			PaymentCollection:   4,
			Deposit:             5,
			Payment:             6,
		},
		AddBusiness: &custom.AddBusiness{
			BusinessOpportunity:         7,
			BusinessOpportunityRelation: 8,
			Order:                       9,
		},
		CustomerInfo: &custom.CustomerInfo{
			TransactionCustomer:            10,
			PublicPoolCustomers:            11,
			PublicPoolTransactionCustomers: 12,
		},
		ContactInfo: &custom.ContactInfo{
			AddContact:        12,
			PendingRenewal:    13,
			ExpiredNotRenewed: 14,
		},
		BusinessFollowUp: &custom.BusinessFollowUp{
			Customer:          15,
			PaymentCollection: 16,
		},
		PerformanceTarget: 60,
	}, "获取成功", c)

}
