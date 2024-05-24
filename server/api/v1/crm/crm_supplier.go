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

type CrmSupplierApi struct {
}

var crmSupplierService = service.ServiceGroupApp.CrmServiceGroup.CrmSupplierService


// CreateCrmSupplier 创建crmSupplier表
// @Tags CrmSupplier
// @Summary 创建crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmSupplier true "创建crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmSupplier/createCrmSupplier [post]
func (crmSupplierApi *CrmSupplierApi) CreateCrmSupplier(c *gin.Context) {
	var crmSupplier crm.CrmSupplier
	err := c.ShouldBindJSON(&crmSupplier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmSupplierService.CreateCrmSupplier(&crmSupplier); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmSupplier 删除crmSupplier表
// @Tags CrmSupplier
// @Summary 删除crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmSupplier true "删除crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmSupplier/deleteCrmSupplier [delete]
func (crmSupplierApi *CrmSupplierApi) DeleteCrmSupplier(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmSupplierService.DeleteCrmSupplier(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmSupplierByIds 批量删除crmSupplier表
// @Tags CrmSupplier
// @Summary 批量删除crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmSupplier/deleteCrmSupplierByIds [delete]
func (crmSupplierApi *CrmSupplierApi) DeleteCrmSupplierByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmSupplierService.DeleteCrmSupplierByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmSupplier 更新crmSupplier表
// @Tags CrmSupplier
// @Summary 更新crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmSupplier true "更新crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmSupplier/updateCrmSupplier [put]
func (crmSupplierApi *CrmSupplierApi) UpdateCrmSupplier(c *gin.Context) {
	var crmSupplier crm.CrmSupplier
	err := c.ShouldBindJSON(&crmSupplier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmSupplierService.UpdateCrmSupplier(crmSupplier); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmSupplier 用id查询crmSupplier表
// @Tags CrmSupplier
// @Summary 用id查询crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmSupplier true "用id查询crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmSupplier/findCrmSupplier [get]
func (crmSupplierApi *CrmSupplierApi) FindCrmSupplier(c *gin.Context) {
	ID := c.Query("ID")
	if recrmSupplier, err := crmSupplierService.GetCrmSupplier(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmSupplier": recrmSupplier}, c)
	}
}

// GetCrmSupplierList 分页获取crmSupplier表列表
// @Tags CrmSupplier
// @Summary 分页获取crmSupplier表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmSupplierSearch true "分页获取crmSupplier表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmSupplier/getCrmSupplierList [get]
func (crmSupplierApi *CrmSupplierApi) GetCrmSupplierList(c *gin.Context) {
	var pageInfo crmReq.CrmSupplierSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmSupplierService.GetCrmSupplierInfoList(pageInfo); err != nil {
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

// GetCrmSupplierPublic 不需要鉴权的crmSupplier表接口
// @Tags CrmSupplier
// @Summary 不需要鉴权的crmSupplier表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmSupplierSearch true "分页获取crmSupplier表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmSupplier/getCrmSupplierList [get]
func (crmSupplierApi *CrmSupplierApi) GetCrmSupplierPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmSupplier表接口信息",
    }, "获取成功", c)
}
