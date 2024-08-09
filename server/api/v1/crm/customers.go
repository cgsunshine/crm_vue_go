package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPageCrmCustomersList 分页获取客户管理列表
// @Tags CrmCustomers
// @Summary 分页获取客户管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomersSearch true "分页获取客户管理列表完整数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomers/getPageCrmCustomersList [get]
func (crmCustomersApi *CrmCustomersApi) GetCrmPageCustomersList(c *gin.Context) {
	var pageInfo crmReq.CrmCustomersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmCustomersService.GetPageCrmCustomersInfoList(pageInfo); err != nil {
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

// GetPageCrmCustomersList 分页获取客户管理列表所有用户
// @Tags CrmCustomers
// @Summary 分页获取客户管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomersSearch true "分页获取客户管理列表完整数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomers/getPageCrmCustomersList [get]
func (crmCustomersApi *CrmCustomersApi) GetCrmPageCustomersAllList(c *gin.Context) {
	var pageInfo crmReq.CrmCustomersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmCustomersService.GetPageCrmCustomersInfoList(pageInfo); err != nil {
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

// FindCrmPageCustomers 用id查询crmCustomers表
// @Tags CrmCustomers
// @Summary 用id查询crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCustomers true "用id查询crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomers/findCrmCustomers [get]
func (crmCustomersApi *CrmCustomersApi) FindCrmPageCustomers(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCustomers, err := crmCustomersService.GetCrmPageCustomers(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCustomers": recrmCustomers}, c)
	}
}

// CreateCrmCustomers 创建crmCustomers表
// @Tags CrmCustomers
// @Summary 创建crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomers true "创建crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCustomers/createCrmCustomers [post]
func (crmCustomersApi *CrmCustomersApi) CreateCrmPageCustomers(c *gin.Context) {
	var crmCustomers crm.CrmCustomers
	err := c.ShouldBindJSON(&crmCustomers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmCustomers.UserId = comm.GetHeaderUserId(c)

	if err := crmCustomersService.CreateCrmCustomers(&crmCustomers); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
