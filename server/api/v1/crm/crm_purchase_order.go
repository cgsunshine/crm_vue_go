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
	"strings"
)

type CrmPurchaseOrderApi struct {
}

var crmPurchaseOrderService = service.ServiceGroupApp.CrmServiceGroup.CrmPurchaseOrderService

// CreateCrmPurchaseOrder 创建订购单管理
// @Tags CrmPurchaseOrder
// @Summary 创建订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrder true "创建订购单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPurchaseOrder/createCrmPurchaseOrder [post]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) CreateCrmPurchaseOrder(c *gin.Context) {
	var crmReqPurchaseOrder crmReq.CrmReqPurchaseOrder
	err := c.ShouldBindJSON(&crmReqPurchaseOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmReqPurchaseOrder.UserId = comm.GetHeaderUserId(c)
	crmReqPurchaseOrder.ReviewStatus = comm.Approval_Status_Pending

	crmPurchaseOrder := crm.CrmPurchaseOrder{
		Amount:            crmReqPurchaseOrder.Amount,
		ContractId:        crmReqPurchaseOrder.ContractId,
		CreationTime:      crmReqPurchaseOrder.CreationTime,
		Description:       crmReqPurchaseOrder.Description,
		ExpirationTime:    crmReqPurchaseOrder.ExpirationTime,
		ProductId:         crmReqPurchaseOrder.ProductId,
		Quantity:          crmReqPurchaseOrder.Quantity,
		UserId:            crmReqPurchaseOrder.UserId,
		PurchaseOrderName: crmReqPurchaseOrder.PurchaseOrderName,
		Currency:          crmReqPurchaseOrder.Currency,
		ReviewStatus:      crmReqPurchaseOrder.ReviewStatus,
	}

	//生成编号
	crmPurchaseOrder.PurchaseOrderNumber = comm.GetBusinessNumber(comm.PurchaseOrderNumberPrefix, crmPurchaseOrderService.GetCrmBillTodayCount())
	if err := crmPurchaseOrderService.CreateCrmPurchaseOrder(&crmPurchaseOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}

	purchaseOrderId := int(crmPurchaseOrder.ID)

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.PurchaseOrderApproval)
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

	crmPurchaseOrderProduct := make([]*crm.CrmPurchaseOrderProduct, 0)
	for _, v := range crmReqPurchaseOrder.ProductsInfo {
		crmOrderProduct := &crm.CrmPurchaseOrderProduct{
			PurchaseOrderId: &purchaseOrderId,
			ProductId:       v.ProductId,
			Quantity:        v.Quantity,
			Specifications:  v.Specifications,
			Price:           v.Price,
		}
		crmPurchaseOrderProduct = append(crmPurchaseOrderProduct, crmOrderProduct)
	}

	if err := crmPurchaseOrderProductService.CreateCrmPurchaseOrderProducts(crmPurchaseOrderProduct); err != nil {
		global.GVA_LOG.Error("插入对应关系失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//插入角色id对应的用户的审批记录
	//for _, userAuth := range ids {
	//	assigneeId := int(userAuth.SysUserId)
	//	if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
	//		AssigneeId:     &assigneeId,
	//		ApprovalStatus: comm.Approval_Status_Under,
	//		AssociatedId:   &purchaseOrderId,
	//		Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
	//		StepId:         roleInfo.NodeId,
	//		ApprovalType:   utils.Pointer(comm.PurchaseOrderApprovalType),
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
			AssociatedId:   &purchaseOrderId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.PurchaseOrderApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)

}

// DeleteCrmPurchaseOrder 删除订购单管理
// @Tags CrmPurchaseOrder
// @Summary 删除订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrder true "删除订购单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrder/deleteCrmPurchaseOrder [delete]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) DeleteCrmPurchaseOrder(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPurchaseOrderService.DeleteCrmPurchaseOrder(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {

		//关联删除审批
		err := crmApprovalTasksService.DelCrmAssociatedIdApprovalTasks(ID, comm.OrderApprovalType)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		}

	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCrmPurchaseOrderByIds 批量删除订购单管理
// @Tags CrmPurchaseOrder
// @Summary 批量删除订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPurchaseOrder/deleteCrmPurchaseOrderByIds [delete]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) DeleteCrmPurchaseOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPurchaseOrderService.DeleteCrmPurchaseOrderByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPurchaseOrder 更新订购单管理
// @Tags CrmPurchaseOrder
// @Summary 更新订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrder true "更新订购单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPurchaseOrder/updateCrmPurchaseOrder [put]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) UpdateCrmPurchaseOrder(c *gin.Context) {
	var crmPurchaseOrder crm.CrmPurchaseOrder
	err := c.ShouldBindJSON(&crmPurchaseOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPurchaseOrderService.UpdateCrmPurchaseOrder(crmPurchaseOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPurchaseOrder 用id查询订购单管理
// @Tags CrmPurchaseOrder
// @Summary 用id查询订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPurchaseOrder true "用id查询订购单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPurchaseOrder/findCrmPurchaseOrder [get]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) FindCrmPurchaseOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPurchaseOrder, err := crmPurchaseOrderService.GetCrmPurchaseOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPurchaseOrder": recrmPurchaseOrder}, c)
	}
}

// GetCrmPurchaseOrderList 分页获取订购单管理列表
// @Tags CrmPurchaseOrder
// @Summary 分页获取订购单管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPurchaseOrderSearch true "分页获取订购单管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrder/getCrmPurchaseOrderList [get]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) GetCrmPurchaseOrderList(c *gin.Context) {
	var pageInfo crmReq.CrmPurchaseOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmPurchaseOrderService.GetCrmPurchaseOrderInfoList(pageInfo); err != nil {
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

// GetCrmPurchaseOrderPublic 不需要鉴权的订购单管理接口
// @Tags CrmPurchaseOrder
// @Summary 不需要鉴权的订购单管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPurchaseOrderSearch true "分页获取订购单管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrder/getCrmPurchaseOrderList [get]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) GetCrmPurchaseOrderPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订购单管理接口信息",
	}, "获取成功", c)
}
