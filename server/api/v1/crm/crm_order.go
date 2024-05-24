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

type CrmOrderApi struct {
}

var crmOrderService = service.ServiceGroupApp.CrmServiceGroup.CrmOrderService


// CreateCrmOrder 创建crmOrder表
// @Tags CrmOrder
// @Summary 创建crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrder true "创建crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrder/createCrmOrder [post]
func (crmOrderApi *CrmOrderApi) CreateCrmOrder(c *gin.Context) {
	var crmOrder crm.CrmOrder
	err := c.ShouldBindJSON(&crmOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOrderService.CreateCrmOrder(&crmOrder); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmOrder 删除crmOrder表
// @Tags CrmOrder
// @Summary 删除crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrder true "删除crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrder/deleteCrmOrder [delete]
func (crmOrderApi *CrmOrderApi) DeleteCrmOrder(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmOrderService.DeleteCrmOrder(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmOrderByIds 批量删除crmOrder表
// @Tags CrmOrder
// @Summary 批量删除crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmOrder/deleteCrmOrderByIds [delete]
func (crmOrderApi *CrmOrderApi) DeleteCrmOrderByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmOrderService.DeleteCrmOrderByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmOrder 更新crmOrder表
// @Tags CrmOrder
// @Summary 更新crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrder true "更新crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOrder/updateCrmOrder [put]
func (crmOrderApi *CrmOrderApi) UpdateCrmOrder(c *gin.Context) {
	var crmOrder crm.CrmOrder
	err := c.ShouldBindJSON(&crmOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOrderService.UpdateCrmOrder(crmOrder); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmOrder 用id查询crmOrder表
// @Tags CrmOrder
// @Summary 用id查询crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmOrder true "用id查询crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrder/findCrmOrder [get]
func (crmOrderApi *CrmOrderApi) FindCrmOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recrmOrder, err := crmOrderService.GetCrmOrder(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmOrder": recrmOrder}, c)
	}
}

// GetCrmOrderList 分页获取crmOrder表列表
// @Tags CrmOrder
// @Summary 分页获取crmOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderSearch true "分页获取crmOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrder/getCrmOrderList [get]
func (crmOrderApi *CrmOrderApi) GetCrmOrderList(c *gin.Context) {
	var pageInfo crmReq.CrmOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmOrderService.GetCrmOrderInfoList(pageInfo); err != nil {
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

// GetCrmOrderPublic 不需要鉴权的crmOrder表接口
// @Tags CrmOrder
// @Summary 不需要鉴权的crmOrder表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderSearch true "分页获取crmOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrder/getCrmOrderList [get]
func (crmOrderApi *CrmOrderApi) GetCrmOrderPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmOrder表接口信息",
    }, "获取成功", c)
}
