package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmBusinessOpportunityApi struct {
}

var crmBusinessOpportunityService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityService

// CreateCrmBusinessOpportunity 创建商机管理
// @Tags CrmBusinessOpportunity
// @Summary 创建商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunity true "创建商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunity/createCrmBusinessOpportunity [post]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) CreateCrmBusinessOpportunity(c *gin.Context) {
	var req crmReq.CrmReqBusinessOpportunity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	req.UserId = comm.GetHeaderUserId(c)

	crmBusinessOpportunity := crm.CrmBusinessOpportunity{
		Amount:                  req.Amount,
		BusinessOpportunityName: req.BusinessOpportunityName,
		CustomerId:              req.CustomerId,
		Description:             req.Description,
		InputTime:               req.InputTime,
		OfferValidityPeriod:     req.OfferValidityPeriod,
		//ProductId:               req.CreatedAt,
		Status:       req.Status,
		UserId:       req.UserId,
		ReviewStatus: req.ReviewStatus,
		Currency:     req.Currency,
	}

	if err := crmBusinessOpportunityService.CreateCrmBusinessOpportunity(&crmBusinessOpportunity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunity 删除商机管理
// @Tags CrmBusinessOpportunity
// @Summary 删除商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunity true "删除商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunity/deleteCrmBusinessOpportunity [delete]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) DeleteCrmBusinessOpportunity(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityService.DeleteCrmBusinessOpportunity(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		//关联删除审批
		err := crmApprovalTasksService.DelCrmAssociatedIdApprovalTasks(ID, comm.BusinessOpportunityApprovalType)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		}
	}

	response.OkWithMessage("删除成功", c)
}

// DeleteCrmBusinessOpportunityByIds 批量删除商机管理
// @Tags CrmBusinessOpportunity
// @Summary 批量删除商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunity/deleteCrmBusinessOpportunityByIds [delete]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) DeleteCrmBusinessOpportunityByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityService.DeleteCrmBusinessOpportunityByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunity 更新商机管理
// @Tags CrmBusinessOpportunity
// @Summary 更新商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunity true "更新商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunity/updateCrmBusinessOpportunity [put]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) UpdateCrmBusinessOpportunity(c *gin.Context) {
	var req crm.CrmReqBusinessOpportunity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Products != nil {
		err = crmBusinessOpportunityService.SetBusinessOpportunityProducts(req.ID, req.ProductIds)
		if err != nil {
			global.GVA_LOG.Error("修改失败!", zap.Error(err))
			response.FailWithMessage("修改失败", c)
			return
		}
	}

	crmBusinessOpportunity := crm.CrmBusinessOpportunity{
		GVA_MODEL:               req.GVA_MODEL,
		Amount:                  req.Amount,
		BusinessOpportunityName: req.BusinessOpportunityName,
		CustomerId:              req.CustomerId,
		Description:             req.Description,
		InputTime:               req.InputTime,
		OfferValidityPeriod:     req.OfferValidityPeriod,
		//ProductId:               req.CreatedAt,
		Status:       req.Status,
		UserId:       req.UserId,
		ReviewStatus: req.ReviewStatus,
		Currency:     req.Currency,
	}

	if err := crmBusinessOpportunityService.UpdateCrmBusinessOpportunity(crmBusinessOpportunity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunity 用id查询商机管理
// @Tags CrmBusinessOpportunity
// @Summary 用id查询商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunity true "用id查询商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunity/findCrmBusinessOpportunity [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) FindCrmBusinessOpportunity(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunity, err := crmBusinessOpportunityService.GetCrmBusinessOpportunity(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunity": recrmBusinessOpportunity}, c)
	}
}

// GetCrmBusinessOpportunityList 分页获取商机管理列表
// @Tags CrmBusinessOpportunity
// @Summary 分页获取商机管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunitySearch true "分页获取商机管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunity/getCrmBusinessOpportunityList [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) GetCrmBusinessOpportunityList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityService.GetCrmBusinessOpportunityInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityPublic 不需要鉴权的商机管理接口
// @Tags CrmBusinessOpportunity
// @Summary 不需要鉴权的商机管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunitySearch true "分页获取商机管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunity/getCrmBusinessOpportunityList [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) GetCrmBusinessOpportunityPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商机管理接口信息",
	}, "获取成功", c)
}
