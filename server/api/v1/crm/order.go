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

// GetCrmOrderList 分页获取crmOrder表列表
// @Tags CrmOrder
// @Summary 分页获取crmOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderSearch true "分页获取crmOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrder/getCrmOrderList [get]
func (crmOrderApi *CrmOrderApi) GetCrmPageOrderList(c *gin.Context) {
	var pageInfo crmReq.CrmOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmOrderService.GetCrmPageOrderInfoList(pageInfo); err != nil {
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

// FindCrmPageOrder 用id查询crmOrder表
// @Tags CrmOrder
// @Summary 用id查询crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmOrder true "用id查询crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrder/findCrmOrder [get]
func (crmOrderApi *CrmOrderApi) FindCrmPageOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recrmOrder, err := crmOrderService.GetCrmPageOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmOrder": recrmOrder}, c)
	}
}

// CreateCrmPageOrder 创建crmOrder表 并且合同关联订单id
// @Tags CrmOrder
// @Summary 创建crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrder true "创建crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrder/createCrmOrder [post]
func (crmOrderApi *CrmOrderApi) CreateCrmPageOrder(c *gin.Context) {
	var crmOrder crm.CrmOrder
	err := c.ShouldBindJSON(&crmOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmOrder.UserId = comm.GetHeaderUserId(c)

	crmOrder.ReviewStatus = comm.Approval_Status_Pending

	if err := crmOrderService.CreateCrmOrder(&crmOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//修改合同中关联的订单id
	err = crmContractService.UpdApprovalStatus(crmOrder.ContactId, map[string]interface{}{
		"order_id": crmOrder.ID,
	})
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	orderId := int(crmOrder.ID)
	//自动生成账单信息
	crmBill := crm.CrmBill{
		Amount:            crmOrder.Price,
		Currency:          crmOrder.Currency,
		CustomerId:        crmOrder.CustomerId,
		Description:       "",
		ExpirationTime:    nil,
		OrderId:           &orderId,
		PaymentCollention: nil,
		PaymentStatus:     "",
		PaymentTime:       nil,
		PaymentType:       "",
		UserId:            crmOrder.UserId,
	}

	if err := crmBillService.CreateCrmBill(&crmBill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//插入审批逻辑
	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.BusinessOpportunityApproval)
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
			AssociatedId:   &orderId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.OrderApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)
}
