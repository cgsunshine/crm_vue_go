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

type CrmTestApi struct {
}

var crmTestService = service.ServiceGroupApp.CrmServiceGroup.CrmTestService


// CreateCrmTest 创建crmTest表
// @Tags CrmTest
// @Summary 创建crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTest true "创建crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTest/createCrmTest [post]
func (crmTestApi *CrmTestApi) CreateCrmTest(c *gin.Context) {
	var crmTest crm.CrmTest
	err := c.ShouldBindJSON(&crmTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTestService.CreateCrmTest(&crmTest); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmTest 删除crmTest表
// @Tags CrmTest
// @Summary 删除crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTest true "删除crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTest/deleteCrmTest [delete]
func (crmTestApi *CrmTestApi) DeleteCrmTest(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmTestService.DeleteCrmTest(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmTestByIds 批量删除crmTest表
// @Tags CrmTest
// @Summary 批量删除crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmTest/deleteCrmTestByIds [delete]
func (crmTestApi *CrmTestApi) DeleteCrmTestByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmTestService.DeleteCrmTestByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmTest 更新crmTest表
// @Tags CrmTest
// @Summary 更新crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTest true "更新crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTest/updateCrmTest [put]
func (crmTestApi *CrmTestApi) UpdateCrmTest(c *gin.Context) {
	var crmTest crm.CrmTest
	err := c.ShouldBindJSON(&crmTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTestService.UpdateCrmTest(crmTest); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmTest 用id查询crmTest表
// @Tags CrmTest
// @Summary 用id查询crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTest true "用id查询crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTest/findCrmTest [get]
func (crmTestApi *CrmTestApi) FindCrmTest(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTest, err := crmTestService.GetCrmTest(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTest": recrmTest}, c)
	}
}

// GetCrmTestList 分页获取crmTest表列表
// @Tags CrmTest
// @Summary 分页获取crmTest表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTestSearch true "分页获取crmTest表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTest/getCrmTestList [get]
func (crmTestApi *CrmTestApi) GetCrmTestList(c *gin.Context) {
	var pageInfo crmReq.CrmTestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTestService.GetCrmTestInfoList(pageInfo); err != nil {
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

// GetCrmTestPublic 不需要鉴权的crmTest表接口
// @Tags CrmTest
// @Summary 不需要鉴权的crmTest表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTestSearch true "分页获取crmTest表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTest/getCrmTestList [get]
func (crmTestApi *CrmTestApi) GetCrmTestPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmTest表接口信息",
    }, "获取成功", c)
}
