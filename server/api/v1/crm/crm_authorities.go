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

type CrmAuthoritiesApi struct {
}

var crmAuthoritiesService = service.ServiceGroupApp.CrmServiceGroup.CrmAuthoritiesService


// CreateCrmAuthorities 创建数据权限
// @Tags CrmAuthorities
// @Summary 创建数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmAuthorities true "创建数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmAuthorities/createCrmAuthorities [post]
func (crmAuthoritiesApi *CrmAuthoritiesApi) CreateCrmAuthorities(c *gin.Context) {
	var crmAuthorities crm.CrmAuthorities
	err := c.ShouldBindJSON(&crmAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmAuthoritiesService.CreateCrmAuthorities(&crmAuthorities); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmAuthorities 删除数据权限
// @Tags CrmAuthorities
// @Summary 删除数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmAuthorities true "删除数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmAuthorities/deleteCrmAuthorities [delete]
func (crmAuthoritiesApi *CrmAuthoritiesApi) DeleteCrmAuthorities(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmAuthoritiesService.DeleteCrmAuthorities(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmAuthoritiesByIds 批量删除数据权限
// @Tags CrmAuthorities
// @Summary 批量删除数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmAuthorities/deleteCrmAuthoritiesByIds [delete]
func (crmAuthoritiesApi *CrmAuthoritiesApi) DeleteCrmAuthoritiesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmAuthoritiesService.DeleteCrmAuthoritiesByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmAuthorities 更新数据权限
// @Tags CrmAuthorities
// @Summary 更新数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmAuthorities true "更新数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmAuthorities/updateCrmAuthorities [put]
func (crmAuthoritiesApi *CrmAuthoritiesApi) UpdateCrmAuthorities(c *gin.Context) {
	var crmAuthorities crm.CrmAuthorities
	err := c.ShouldBindJSON(&crmAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmAuthoritiesService.UpdateCrmAuthorities(crmAuthorities); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmAuthorities 用id查询数据权限
// @Tags CrmAuthorities
// @Summary 用id查询数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmAuthorities true "用id查询数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmAuthorities/findCrmAuthorities [get]
func (crmAuthoritiesApi *CrmAuthoritiesApi) FindCrmAuthorities(c *gin.Context) {
	ID := c.Query("ID")
	if recrmAuthorities, err := crmAuthoritiesService.GetCrmAuthorities(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmAuthorities": recrmAuthorities}, c)
	}
}

// GetCrmAuthoritiesList 分页获取数据权限列表
// @Tags CrmAuthorities
// @Summary 分页获取数据权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmAuthoritiesSearch true "分页获取数据权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmAuthorities/getCrmAuthoritiesList [get]
func (crmAuthoritiesApi *CrmAuthoritiesApi) GetCrmAuthoritiesList(c *gin.Context) {
	var pageInfo crmReq.CrmAuthoritiesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmAuthoritiesService.GetCrmAuthoritiesInfoList(pageInfo); err != nil {
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

// GetCrmAuthoritiesPublic 不需要鉴权的数据权限接口
// @Tags CrmAuthorities
// @Summary 不需要鉴权的数据权限接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmAuthoritiesSearch true "分页获取数据权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmAuthorities/getCrmAuthoritiesList [get]
func (crmAuthoritiesApi *CrmAuthoritiesApi) GetCrmAuthoritiesPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的数据权限接口信息",
    }, "获取成功", c)
}
