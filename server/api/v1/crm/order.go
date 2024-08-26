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

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

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
	var crmPageOrder crm.CrmReqOrder
	err := c.ShouldBindJSON(&crmPageOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmPageOrder.UserId = comm.GetHeaderUserId(c)

	crmPageOrder.ReviewStatus = comm.Approval_Status_Pending

	//var products []crm.CrmProduct
	//for _, v := range crmPageOrder.ProductIds {
	//	products = append(products, crm.CrmProduct{
	//		GVA_MODEL: global.GVA_MODEL{
	//			ID: v,
	//		},
	//	})
	//}

	crmOrder := crm.CrmOrder{
		Currency:                 crmPageOrder.Currency,
		CustomerId:               crmPageOrder.CustomerId,
		Description:              crmPageOrder.Description,
		DiscountRate:             crmPageOrder.DiscountRate,
		Amount:                   crmPageOrder.Amount,
		SalesPrice:               crmPageOrder.SalesPrice,
		UserId:                   crmPageOrder.UserId,
		ContactId:                crmPageOrder.ContactId,
		SupplementaryInformation: crmPageOrder.SupplementaryInformation,
		OrderName:                crmPageOrder.OrderName,
		ReviewStatus:             crmPageOrder.ReviewStatus,
		CurrencyId:               crmPageOrder.CurrencyId,
		OrderNumber:              crmPageOrder.OrderNumber,
		//Products:                 products,
		BusinessOpportunityId: crmPageOrder.BusinessOpportunityId,
	}

	if err := crmOrderService.CreateCrmOrder(&crmOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//添加多个关联产品
	//
	//for _, idStr := range strings.Split(crmOrder.ProductIds, ",") {
	//	id, err := strconv.ParseInt(idStr, 10, 32)
	//	if err != nil {
	//		fmt.Println("Error parsing ID:", err)
	//		continue
	//	}
	//	orderId := int(crmOrder.ID)
	//	productId := int(id)
	//
	//	if err := crmOrderProductService.CreateCrmOrderProduct(&crm.CrmOrderProduct{
	//		OrderId:   &orderId,
	//		ProductId: &productId,
	//	}); err != nil {
	//		global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//		response.FailWithMessage("创建失败", c)
	//		return
	//	}
	//}

	orderId := int(crmOrder.ID)
	orderProducts := make([]*crm.CrmOrderProduct, 0)

	for _, productInfo := range crmPageOrder.ProductsInfo {
		orderProducts = append(orderProducts, &crm.CrmOrderProduct{
			OrderId:        &orderId,
			ProductId:      productInfo.ProductId,
			Quantity:       productInfo.Quantity,
			Specifications: productInfo.Specifications,
		})
	}

	if err := crmOrderProductService.CreateCrmOrderProductInc(orderProducts); err != nil {
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

	////自动生成账单信息
	//crmBill := crm.CrmBill{
	//	OrderId: &orderId,
	//	UserId:  crmOrder.UserId,
	//}
	//
	//if err := crmBillService.CreateCrmBill(&crmBill); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//	return
	//}

	//插入审批逻辑
	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.OrderApproval)
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

// CreateCrmPageOrder 创建crmOrder表 并且合同关联订单id
// @Tags CrmOrder
// @Summary 创建crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrder true "创建crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrder/createCrmOrder [post]
func (crmOrderApi *CrmOrderApi) SetOrderProducts(c *gin.Context) {
	var crmPageOrder crmReq.SetOrderProduct
	err := c.ShouldBindJSON(&crmPageOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	orderId := int(crmPageOrder.ID)
	err = crmOrderProductService.DeleteCrmOrderProductBy(&orderId)
	if err != nil {
		global.GVA_LOG.Error("删除异常!", zap.Error(err))
		return
	}
	orderProducts := make([]*crm.CrmOrderProduct, 0)
	for _, productInfo := range crmPageOrder.ProductsInfo {
		orderProducts = append(orderProducts, &crm.CrmOrderProduct{
			OrderId:        &orderId,
			ProductId:      productInfo.ProductId,
			Quantity:       productInfo.Quantity,
			Specifications: productInfo.Specifications,
		})

	}

	if err := crmOrderProductService.CreateCrmOrderProductInc(orderProducts); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//err = crmOrderService.SetOrderProducts(crmPageOrder.ID, crmPageOrder.ProductIds)
	//if err != nil {
	//	global.GVA_LOG.Error("修改失败!", zap.Error(err))
	//	response.FailWithMessage("修改失败", c)
	//	return
	//}
	response.OkWithMessage("修改成功", c)

}
