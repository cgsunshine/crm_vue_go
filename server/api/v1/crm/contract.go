package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCrmContractList 分页获取crmContract表列表
// @Tags CrmContract
// @Summary 分页获取crmContract表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractSearch true "分页获取crmContract表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContract/getCrmContractList [get]
func (crmContractApi *CrmContractApi) GetCrmPageContractList(c *gin.Context) {
	var pageInfo crmReq.CrmContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmContractService.GetCrmPageContractInfoList(pageInfo); err != nil {
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

// CreateCrmContract 创建crmContract表
// @Tags CrmContract
// @Summary 创建crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContract true "创建crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContract/createCrmContract [post]
func (crmContractApi *CrmContractApi) CreatePageCrmContract(c *gin.Context) {
	var crmContract crm.CrmContract
	err := c.ShouldBindJSON(&crmContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmContract.UserId = comm.GetHeaderUserId(c)

	crmContract.ContractStatus = "1"
	crmContract.ReviewStatus = comm.Approval_Status_Pending
	crmContract.ReviewResult = "1"

	//先查询合同审批对应的流程

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.ContractApproval)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	ids, err := userService.GetRoleUsers(roleInfo.RoleIds)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	if err := crmContractService.CreateCrmContract(&crmContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		contactId := int(crmContract.ID)
		//插入角色id对应的用户的审批记录
		for _, userAuth := range ids {
			assigneeId := int(userAuth.SysUserId)

			if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
				AssigneeId:     &assigneeId,
				ApprovalStatus: comm.Approval_Status_Under,
				AssociatedId:   &contactId,
				Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
				StepId:         roleInfo.NodeId,
				ApprovalType:   utils.Pointer(comm.ContractApprovalType),
			}); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败", c)
				return
			}
		}
		response.OkWithMessage("创建成功", c)
	}
}

// FindPageCrmContract 用id查询crmContract表
// @Tags CrmContract
// @Summary 用id查询crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
func (crmContractApi *CrmContractApi) FindCrmPageContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContract": recrmContract}, c)
	}
}

// FindCrmPageFileContract 用id查询crmContract表 获取相关文件图片
// @Tags CrmContract
// @Summary 用id查询crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
func (crmContractApi *CrmContractApi) FindCrmPageFileContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		//查询关联文件
		list, _, err := fileUploadAndDownloadService.GetFileRecordInfoIdsList(recrmContract.ContractFile)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmContract": list}, c)
	}
}

// DownloadCrmPageFileContractExcel 用id查询crmContract表 下载合同Excel
// @Tags CrmContract
// @Summary 下载合同Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/downloadCrmPageFileContractExcel [get]
func (crmContractApi *CrmContractApi) DownloadCrmPageFileContractExcel(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		//查询关联文件
		list, _, err := fileUploadAndDownloadService.GetFileRecordInfoIdsList(recrmContract.ContractFile)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmContract": list}, c)
	}
}
