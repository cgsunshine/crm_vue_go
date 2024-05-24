package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CrmCustomersApi struct {
}

var crmCustomersService = service.ServiceGroupApp.CrmServiceGroup.CrmCustomersService


// CreateCrmCustomers 创建客户管理
// @Tags CrmCustomers
// @Summary 创建客户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomers true "创建客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCustomers/createCrmCustomers [post]
func (crmCustomersApi *CrmCustomersApi) CreateCrmCustomers(c *gin.Context) {
	var crmCustomers crm.CrmCustomers
	err := c.ShouldBindJSON(&crmCustomers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCustomersService.CreateCrmCustomers(&crmCustomers); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmCustomers 删除客户管理
// @Tags CrmCustomers
// @Summary 删除客户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomers true "删除客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomers/deleteCrmCustomers [delete]
func (crmCustomersApi *CrmCustomersApi) DeleteCrmCustomers(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmCustomersService.DeleteCrmCustomers(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmCustomersByIds 批量删除客户管理
// @Tags CrmCustomers
// @Summary 批量删除客户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmCustomers/deleteCrmCustomersByIds [delete]
func (crmCustomersApi *CrmCustomersApi) DeleteCrmCustomersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmCustomersService.DeleteCrmCustomersByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmCustomers 更新客户管理
// @Tags CrmCustomers
// @Summary 更新客户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomers true "更新客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomers/updateCrmCustomers [put]
func (crmCustomersApi *CrmCustomersApi) UpdateCrmCustomers(c *gin.Context) {
	var crmCustomers crm.CrmCustomers
	err := c.ShouldBindJSON(&crmCustomers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCustomersService.UpdateCrmCustomers(crmCustomers); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmCustomers 用id查询客户管理
// @Tags CrmCustomers
// @Summary 用id查询客户管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCustomers true "用id查询客户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomers/findCrmCustomers [get]
func (crmCustomersApi *CrmCustomersApi) FindCrmCustomers(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCustomers, err := crmCustomersService.GetCrmCustomers(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCustomers": recrmCustomers}, c)
	}
}

// GetCrmCustomersList 分页获取客户管理列表
// @Tags CrmCustomers
// @Summary 分页获取客户管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomersSearch true "分页获取客户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomers/getCrmCustomersList [get]
func (crmCustomersApi *CrmCustomersApi) GetCrmCustomersList(c *gin.Context) {
	var pageInfo crmReq.CrmCustomersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmCustomersService.GetCrmCustomersInfoList(pageInfo); err != nil {
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

// GetCrmCustomersPublic 不需要鉴权的客户管理接口
// @Tags CrmCustomers
// @Summary 不需要鉴权的客户管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomersSearch true "分页获取客户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomers/getCrmCustomersList [get]
func (crmCustomersApi *CrmCustomersApi) GetCrmCustomersPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的客户管理接口信息",
    }, "获取成功", c)
}
