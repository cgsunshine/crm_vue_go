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

type CrmCurrencyApi struct {
}

var crmCurrencyService = service.ServiceGroupApp.CrmServiceGroup.CrmCurrencyService


// CreateCrmCurrency 创建币种管理
// @Tags CrmCurrency
// @Summary 创建币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCurrency true "创建币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCurrency/createCrmCurrency [post]
func (crmCurrencyApi *CrmCurrencyApi) CreateCrmCurrency(c *gin.Context) {
	var crmCurrency crm.CrmCurrency
	err := c.ShouldBindJSON(&crmCurrency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCurrencyService.CreateCrmCurrency(&crmCurrency); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmCurrency 删除币种管理
// @Tags CrmCurrency
// @Summary 删除币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCurrency true "删除币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCurrency/deleteCrmCurrency [delete]
func (crmCurrencyApi *CrmCurrencyApi) DeleteCrmCurrency(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmCurrencyService.DeleteCrmCurrency(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmCurrencyByIds 批量删除币种管理
// @Tags CrmCurrency
// @Summary 批量删除币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmCurrency/deleteCrmCurrencyByIds [delete]
func (crmCurrencyApi *CrmCurrencyApi) DeleteCrmCurrencyByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmCurrencyService.DeleteCrmCurrencyByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmCurrency 更新币种管理
// @Tags CrmCurrency
// @Summary 更新币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCurrency true "更新币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCurrency/updateCrmCurrency [put]
func (crmCurrencyApi *CrmCurrencyApi) UpdateCrmCurrency(c *gin.Context) {
	var crmCurrency crm.CrmCurrency
	err := c.ShouldBindJSON(&crmCurrency)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCurrencyService.UpdateCrmCurrency(crmCurrency); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmCurrency 用id查询币种管理
// @Tags CrmCurrency
// @Summary 用id查询币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCurrency true "用id查询币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCurrency/findCrmCurrency [get]
func (crmCurrencyApi *CrmCurrencyApi) FindCrmCurrency(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCurrency, err := crmCurrencyService.GetCrmCurrency(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCurrency": recrmCurrency}, c)
	}
}

// GetCrmCurrencyList 分页获取币种管理列表
// @Tags CrmCurrency
// @Summary 分页获取币种管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCurrencySearch true "分页获取币种管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCurrency/getCrmCurrencyList [get]
func (crmCurrencyApi *CrmCurrencyApi) GetCrmCurrencyList(c *gin.Context) {
	var pageInfo crmReq.CrmCurrencySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmCurrencyService.GetCrmCurrencyInfoList(pageInfo); err != nil {
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

// GetCrmCurrencyPublic 不需要鉴权的币种管理接口
// @Tags CrmCurrency
// @Summary 不需要鉴权的币种管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCurrencySearch true "分页获取币种管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCurrency/getCrmCurrencyList [get]
func (crmCurrencyApi *CrmCurrencyApi) GetCrmCurrencyPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的币种管理接口信息",
    }, "获取成功", c)
}
