import service from '@/utils/request'

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 创建商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalTasks true "创建商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/createCrmBusinessOpportunityApprovalTasks [post]
export const createCrmBusinessOpportunityApprovalTasks = (data) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/createCrmBusinessOpportunityApprovalTasks',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 删除商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalTasks true "删除商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasks [delete]
export const deleteCrmBusinessOpportunityApprovalTasks = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasks',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 批量删除商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasks [delete]
export const deleteCrmBusinessOpportunityApprovalTasksByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 更新商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalTasks true "更新商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/updateCrmBusinessOpportunityApprovalTasks [put]
export const updateCrmBusinessOpportunityApprovalTasks = (data) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/updateCrmBusinessOpportunityApprovalTasks',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 用id查询商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunityApprovalTasks true "用id查询商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/findCrmBusinessOpportunityApprovalTasks [get]
export const findCrmBusinessOpportunityApprovalTasks = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/findCrmBusinessOpportunityApprovalTasks',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 分页获取商机审批列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商机审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/getCrmBusinessOpportunityApprovalTasksList [get]
export const getCrmBusinessOpportunityApprovalTasksList = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalTasks/getCrmBusinessOpportunityApprovalTasksList',
    method: 'get',
    params
  })
}
