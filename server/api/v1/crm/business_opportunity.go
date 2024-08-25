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

// GetCrmPageBusinessOpportunityList 分页获取crmBusinessOpportunity表列表
// @Tags CrmBusinessOpportunity
// @Summary 分页获取crmBusinessOpportunity表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunitySearch true "分页获取crmBusinessOpportunity表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunity/getPageCrmBusinessOpportunityList [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) GetCrmPageBusinessOpportunityList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)
	if list, total, err := crmBusinessOpportunityService.GetPageCrmBusinessOpportunityInfoList(pageInfo); err != nil {
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

// FindCrmPageBusinessOpportunity 用id查询商机管理
// @Tags CrmBusinessOpportunity
// @Summary 用id查询商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunity true "用id查询商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunity/findCrmBusinessOpportunity [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) FindCrmPageBusinessOpportunity(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunity, err := crmBusinessOpportunityService.GetCrmPageBusinessOpportunity(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunity": recrmBusinessOpportunity}, c)
	}
}

// CreateCrmPageBusinessOpportunity 创建商机管理
// @Tags CrmBusinessOpportunity
// @Summary 创建商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunity true "创建商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunity/createCrmBusinessOpportunity [post]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) CreateCrmPageBusinessOpportunity(c *gin.Context) {

	var req crmReq.CrmReqBusinessOpportunity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	req.UserId = comm.GetHeaderUserId(c)

	//var products []crm.CrmProduct
	//for _, v := range req.ProductIds {
	//	products = append(products, crm.CrmProduct{
	//		GVA_MODEL: global.GVA_MODEL{
	//			ID: v,
	//		},
	//	})
	//}

	crmBusinessOpportunity := crm.CrmBusinessOpportunity{
		Amount:                  req.Amount,
		BusinessOpportunityName: req.BusinessOpportunityName,
		CustomerId:              req.CustomerId,
		Description:             req.Description,
		InputTime:               req.InputTime,
		OfferValidityPeriod:     req.OfferValidityPeriod,
		Status:                  req.Status,
		UserId:                  req.UserId,
		ReviewStatus:            req.ReviewStatus,
		Currency:                req.Currency,
		//Products:                products,
	}

	crmBusinessOpportunity.UserId = comm.GetHeaderUserId(c)

	crmBusinessOpportunity.ReviewStatus = comm.Approval_Status_Pending

	if err := crmBusinessOpportunityService.CreateCrmBusinessOpportunity(&crmBusinessOpportunity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}

	businessOpportunityId := int(crmBusinessOpportunity.ID)

	crmOrderProducts := make([]*crm.CrmOrderProduct, 0)
	for _, v := range req.OrderProduct {
		crmOrderProduct := &crm.CrmOrderProduct{
			OrderId:        &businessOpportunityId,
			ProductId:      v.ProductId,
			Quantity:       v.Quantity,
			Specifications: v.Specifications,
		}
		crmOrderProducts = append(crmOrderProducts, crmOrderProduct)
	}

	if err := crmOrderProductService.CreateCrmOrderProducts(crmOrderProducts); err != nil {
		global.GVA_LOG.Error("插入对应关系失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}

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
			AssociatedId:   &businessOpportunityId,
			Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
			StepId:         roleInfo.NodeId,
			ApprovalType:   utils.Pointer(comm.BusinessOpportunityApprovalType),
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}

	response.OkWithMessage("创建成功", c)
}
