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

type CrmCommissionRebateApi struct {
}

var crmCommissionRebateService = service.ServiceGroupApp.CrmServiceGroup.CrmCommissionRebateService


// CreateCrmCommissionRebate 创建返佣管理
// @Tags CrmCommissionRebate
// @Summary 创建返佣管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCommissionRebate true "创建返佣管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCommissionRebate/createCrmCommissionRebate [post]
func (crmCommissionRebateApi *CrmCommissionRebateApi) CreateCrmCommissionRebate(c *gin.Context) {
	var crmCommissionRebate crm.CrmCommissionRebate
	err := c.ShouldBindJSON(&crmCommissionRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCommissionRebateService.CreateCrmCommissionRebate(&crmCommissionRebate); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmCommissionRebate 删除返佣管理
// @Tags CrmCommissionRebate
// @Summary 删除返佣管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCommissionRebate true "删除返佣管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCommissionRebate/deleteCrmCommissionRebate [delete]
func (crmCommissionRebateApi *CrmCommissionRebateApi) DeleteCrmCommissionRebate(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmCommissionRebateService.DeleteCrmCommissionRebate(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmCommissionRebateByIds 批量删除返佣管理
// @Tags CrmCommissionRebate
// @Summary 批量删除返佣管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmCommissionRebate/deleteCrmCommissionRebateByIds [delete]
func (crmCommissionRebateApi *CrmCommissionRebateApi) DeleteCrmCommissionRebateByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmCommissionRebateService.DeleteCrmCommissionRebateByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmCommissionRebate 更新返佣管理
// @Tags CrmCommissionRebate
// @Summary 更新返佣管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCommissionRebate true "更新返佣管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCommissionRebate/updateCrmCommissionRebate [put]
func (crmCommissionRebateApi *CrmCommissionRebateApi) UpdateCrmCommissionRebate(c *gin.Context) {
	var crmCommissionRebate crm.CrmCommissionRebate
	err := c.ShouldBindJSON(&crmCommissionRebate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCommissionRebateService.UpdateCrmCommissionRebate(crmCommissionRebate); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmCommissionRebate 用id查询返佣管理
// @Tags CrmCommissionRebate
// @Summary 用id查询返佣管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCommissionRebate true "用id查询返佣管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCommissionRebate/findCrmCommissionRebate [get]
func (crmCommissionRebateApi *CrmCommissionRebateApi) FindCrmCommissionRebate(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCommissionRebate, err := crmCommissionRebateService.GetCrmCommissionRebate(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCommissionRebate": recrmCommissionRebate}, c)
	}
}

// GetCrmCommissionRebateList 分页获取返佣管理列表
// @Tags CrmCommissionRebate
// @Summary 分页获取返佣管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCommissionRebateSearch true "分页获取返佣管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCommissionRebate/getCrmCommissionRebateList [get]
func (crmCommissionRebateApi *CrmCommissionRebateApi) GetCrmCommissionRebateList(c *gin.Context) {
	var pageInfo crmReq.CrmCommissionRebateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmCommissionRebateService.GetCrmCommissionRebateInfoList(pageInfo); err != nil {
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

// GetCrmCommissionRebatePublic 不需要鉴权的返佣管理接口
// @Tags CrmCommissionRebate
// @Summary 不需要鉴权的返佣管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCommissionRebateSearch true "分页获取返佣管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCommissionRebate/getCrmCommissionRebateList [get]
func (crmCommissionRebateApi *CrmCommissionRebateApi) GetCrmCommissionRebatePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的返佣管理接口信息",
    }, "获取成功", c)
}
