package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmStatementAccountApi struct {
}

var crmStatementAccountService = service.ServiceGroupApp.CrmServiceGroup.CrmStatementAccountService

// CreateCrmStatementAccount 创建crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 创建crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccount true "创建crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccount/createCrmStatementAccount [post]
func (crmStatementAccountApi *CrmStatementAccountApi) CreateCrmStatementAccount(c *gin.Context) {
	var crmStatementAccount crm.CrmStatementAccount
	err := c.ShouldBindJSON(&crmStatementAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmStatementAccount.UserId = comm.GetHeaderUserId(c)
	crmStatementAccount.ReviewStatus = comm.Approval_Status_Pending

	if err := crmStatementAccountService.CreateCrmStatementAccount(&crmStatementAccount); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		statementAccountId := int(crmStatementAccount.ID)
		crmBillPayment := crm.CrmBillPayment{
			Amount:             crmStatementAccount.Amount,
			BillPaymentName:    "",
			Currency:           "", //缺个币种
			PaymentStatus:      comm.PaymentStatusUnpaid,
			StatementAccountId: &statementAccountId,
			UserId:             crmStatementAccount.UserId,
		}

		if err := crmBillPaymentService.CreateCrmBillPayment(&crmBillPayment); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		//在查出第一步对应的角色id
		roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.StatementAccountApproval)
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
		statementAccountId := int(crmStatementAccount.ID)
		//插入角色id对应的用户的审批记录
		for _, userAuth := range ids {
			assigneeId := int(userAuth.SysUserId)

			if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
				AssigneeId:     &assigneeId,
				ApprovalStatus: comm.Approval_Status_Under,
				AssociatedId:   &statementAccountId,
				Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
				StepId:         roleInfo.NodeId,
				ApprovalType:   utils.Pointer(comm.StatementAccountApprovalType),
			}); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败", c)
				return
			}
		}

		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmStatementAccount 删除crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 删除crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccount true "删除crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccount/deleteCrmStatementAccount [delete]
func (crmStatementAccountApi *CrmStatementAccountApi) DeleteCrmStatementAccount(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmStatementAccountService.DeleteCrmStatementAccount(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmStatementAccountByIds 批量删除crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 批量删除crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmStatementAccount/deleteCrmStatementAccountByIds [delete]
func (crmStatementAccountApi *CrmStatementAccountApi) DeleteCrmStatementAccountByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmStatementAccountService.DeleteCrmStatementAccountByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmStatementAccount 更新crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 更新crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccount true "更新crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccount/updateCrmStatementAccount [put]
func (crmStatementAccountApi *CrmStatementAccountApi) UpdateCrmStatementAccount(c *gin.Context) {
	var crmStatementAccount crm.CrmStatementAccount
	err := c.ShouldBindJSON(&crmStatementAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmStatementAccountService.UpdateCrmStatementAccount(crmStatementAccount); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmStatementAccount 用id查询crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 用id查询crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmStatementAccount true "用id查询crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccount/findCrmStatementAccount [get]
func (crmStatementAccountApi *CrmStatementAccountApi) FindCrmStatementAccount(c *gin.Context) {
	ID := c.Query("ID")
	if recrmStatementAccount, err := crmStatementAccountService.GetCrmStatementAccount(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmStatementAccount": recrmStatementAccount}, c)
	}
}

// GetCrmStatementAccountList 分页获取crmStatementAccount表列表
// @Tags CrmStatementAccount
// @Summary 分页获取crmStatementAccount表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountSearch true "分页获取crmStatementAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccount/getCrmStatementAccountList [get]
func (crmStatementAccountApi *CrmStatementAccountApi) GetCrmStatementAccountList(c *gin.Context) {
	var pageInfo crmReq.CrmStatementAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := crmStatementAccountService.GetCrmStatementAccountInfoList(pageInfo); err != nil {
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

// GetCrmStatementAccountPublic 不需要鉴权的crmStatementAccount表接口
// @Tags CrmStatementAccount
// @Summary 不需要鉴权的crmStatementAccount表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountSearch true "分页获取crmStatementAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccount/getCrmStatementAccountList [get]
func (crmStatementAccountApi *CrmStatementAccountApi) GetCrmStatementAccountPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmStatementAccount表接口信息",
	}, "获取成功", c)
}
