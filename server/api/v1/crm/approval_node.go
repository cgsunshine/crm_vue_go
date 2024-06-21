package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreateCrmPageApprovalNode 创建审批节点
// @Tags CrmApprovalNode
// @Summary 创建审批节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalNode true "创建审批节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalNode/createCrmPageApprovalNode [post]
func (crmApprovalNodeApi *CrmApprovalNodeApi) CreateCrmPageApprovalNode(c *gin.Context) {
	var crmPageApprovalNode crm.CrmPageApprovalNode
	err := c.ShouldBindJSON(&crmPageApprovalNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	nodes := make([]crm.CrmApprovalNode, 0)

	for index, roleId := range crmPageApprovalNode.RoleIds {
		crmApprovalNode := crm.CrmApprovalNode{
			NodeName:                crmPageApprovalNode.NodeName,
			NumberApprovedPersonnel: crmPageApprovalNode.NumberApprovedPersonnel,
			ProcessId:               crmPageApprovalNode.ProcessId,
			UserIds:                 crmPageApprovalNode.UserIds,
		}
		ind := index + 1
		crmApprovalNode.UserId = userService.FindUserDataStatusById(&userID)
		crmApprovalNode.NodeOrder = &ind
		crmApprovalNode.RoleIds = roleId
		nodes = append(nodes, crmApprovalNode)
	}

	if err := crmApprovalNodeService.CreateCrmBatApprovalNode(&nodes); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)

}
