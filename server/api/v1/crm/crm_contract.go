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

type CrmContractApi struct {
}

var crmContractService = service.ServiceGroupApp.CrmServiceGroup.CrmContractService

// CreateCrmContract 创建crmContract表
// @Tags CrmContract
// @Summary 创建crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContract true "创建crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContract/createCrmContract [post]
func (crmContractApi *CrmContractApi) CreateCrmContract(c *gin.Context) {
	var crmContract crm.CrmContract
	err := c.ShouldBindJSON(&crmContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmContract.UserId = comm.GetHeaderUserId(c)

	if err := crmContractService.CreateCrmContract(&crmContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContract 删除crmContract表
// @Tags CrmContract
// @Summary 删除crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContract true "删除crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContract/deleteCrmContract [delete]
func (crmContractApi *CrmContractApi) DeleteCrmContract(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContractService.DeleteCrmContract(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContractByIds 批量删除crmContract表
// @Tags CrmContract
// @Summary 批量删除crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContract/deleteCrmContractByIds [delete]
func (crmContractApi *CrmContractApi) DeleteCrmContractByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContractService.DeleteCrmContractByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContract 更新crmContract表
// @Tags CrmContract
// @Summary 更新crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContract true "更新crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContract/updateCrmContract [put]
func (crmContractApi *CrmContractApi) UpdateCrmContract(c *gin.Context) {
	var crmContract crm.CrmContract
	err := c.ShouldBindJSON(&crmContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContractService.UpdateCrmContract(crmContract); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContract 用id查询crmContract表
// @Tags CrmContract
// @Summary 用id查询crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
func (crmContractApi *CrmContractApi) FindCrmContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContract": recrmContract}, c)
	}
}

// GetCrmContractList 分页获取crmContract表列表
// @Tags CrmContract
// @Summary 分页获取crmContract表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractSearch true "分页获取crmContract表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContract/getCrmContractList [get]
func (crmContractApi *CrmContractApi) GetCrmContractList(c *gin.Context) {
	var pageInfo crmReq.CrmContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContractService.GetCrmContractInfoList(pageInfo); err != nil {
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

// GetCrmContractPublic 不需要鉴权的crmContract表接口
// @Tags CrmContract
// @Summary 不需要鉴权的crmContract表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractSearch true "分页获取crmContract表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContract/getCrmContractList [get]
func (crmContractApi *CrmContractApi) GetCrmContractPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmContract表接口信息",
	}, "获取成功", c)
}
