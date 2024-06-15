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

type CrmTicketCategoriesApi struct {
}

var crmTicketCategoriesService = service.ServiceGroupApp.CrmServiceGroup.CrmTicketCategoriesService


// CreateCrmTicketCategories 创建工单类别
// @Tags CrmTicketCategories
// @Summary 创建工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketCategories true "创建工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketCategories/createCrmTicketCategories [post]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) CreateCrmTicketCategories(c *gin.Context) {
	var crmTicketCategories crm.CrmTicketCategories
	err := c.ShouldBindJSON(&crmTicketCategories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketCategoriesService.CreateCrmTicketCategories(&crmTicketCategories); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmTicketCategories 删除工单类别
// @Tags CrmTicketCategories
// @Summary 删除工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketCategories true "删除工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketCategories/deleteCrmTicketCategories [delete]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) DeleteCrmTicketCategories(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmTicketCategoriesService.DeleteCrmTicketCategories(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmTicketCategoriesByIds 批量删除工单类别
// @Tags CrmTicketCategories
// @Summary 批量删除工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmTicketCategories/deleteCrmTicketCategoriesByIds [delete]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) DeleteCrmTicketCategoriesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmTicketCategoriesService.DeleteCrmTicketCategoriesByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmTicketCategories 更新工单类别
// @Tags CrmTicketCategories
// @Summary 更新工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketCategories true "更新工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketCategories/updateCrmTicketCategories [put]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) UpdateCrmTicketCategories(c *gin.Context) {
	var crmTicketCategories crm.CrmTicketCategories
	err := c.ShouldBindJSON(&crmTicketCategories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketCategoriesService.UpdateCrmTicketCategories(crmTicketCategories); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmTicketCategories 用id查询工单类别
// @Tags CrmTicketCategories
// @Summary 用id查询工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTicketCategories true "用id查询工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketCategories/findCrmTicketCategories [get]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) FindCrmTicketCategories(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTicketCategories, err := crmTicketCategoriesService.GetCrmTicketCategories(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTicketCategories": recrmTicketCategories}, c)
	}
}

// GetCrmTicketCategoriesList 分页获取工单类别列表
// @Tags CrmTicketCategories
// @Summary 分页获取工单类别列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketCategoriesSearch true "分页获取工单类别列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketCategories/getCrmTicketCategoriesList [get]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) GetCrmTicketCategoriesList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketCategoriesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTicketCategoriesService.GetCrmTicketCategoriesInfoList(pageInfo); err != nil {
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

// GetCrmTicketCategoriesPublic 不需要鉴权的工单类别接口
// @Tags CrmTicketCategories
// @Summary 不需要鉴权的工单类别接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketCategoriesSearch true "分页获取工单类别列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketCategories/getCrmTicketCategoriesList [get]
func (crmTicketCategoriesApi *CrmTicketCategoriesApi) GetCrmTicketCategoriesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的工单类别接口信息",
    }, "获取成功", c)
}
