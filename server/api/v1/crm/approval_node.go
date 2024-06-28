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
	var crmPageApprovalNode crm.CrmPageInfoApprovalNode
	err := c.ShouldBindJSON(&crmPageApprovalNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := comm.GetHeaderUserId(c)

	nodes := make([]crm.CrmApprovalNode, 0)

	for index, roleId := range crmPageApprovalNode.RoleIds {
		crmApprovalNode := crm.CrmApprovalNode{
			NodeName: crmPageApprovalNode.NodeName,
			//NumberApprovedPersonnel: crmPageApprovalNode.NumberApprovedPersonnel,
			NumberApprovedPersonnel: utils.Pointer(comm.NumberApprovedPersonnel),
			ProcessId:               crmPageApprovalNode.ProcessId,
			UserIds:                 crmPageApprovalNode.UserIds,
		}
		ind := index + 1
		crmApprovalNode.UserId = userService.FindUserDataStatusById(userID)
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

// GetCrmPageApprovalNodeList 分页获取审批节点列表
// @Tags CrmApprovalNode
// @Summary 分页获取审批节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalNodeSearch true "分页获取审批节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalNode/getCrmApprovalNodeList [get]
func (crmApprovalNodeApi *CrmApprovalNodeApi) GetCrmPageApprovalNodeList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := crmApprovalNodeService.GetCrmPageApprovalNodeInfoList(pageInfo); err != nil {
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
