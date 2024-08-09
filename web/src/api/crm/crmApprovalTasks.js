import service from '@/utils/request'

// @Tags CrmApprovalTasks
// @Summary 创建审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasks true "创建审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalTasks/createCrmApprovalTasks [post]
export const createCrmApprovalTasks = (data) => {
  return service({
    url: '/crmApprovalTasks/createCrmApprovalTasks',
    method: 'post',
    data
  })
}

// @Tags CrmApprovalTasks
// @Summary 删除审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasks true "删除审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasks/deleteCrmApprovalTasks [delete]
export const deleteCrmApprovalTasks = (params) => {
  return service({
    url: '/crmApprovalTasks/deleteCrmApprovalTasks',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalTasks
// @Summary 批量删除审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasks/deleteCrmApprovalTasks [delete]
export const deleteCrmApprovalTasksByIds = (params) => {
  return service({
    url: '/crmApprovalTasks/deleteCrmApprovalTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalTasks
// @Summary 更新审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasks true "更新审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalTasks/updateCrmApprovalTasks [put]
export const updateCrmApprovalTasks = (data) => {
  return service({
    url: '/crmApprovalTasks/updateCrmApprovalTasks',
    method: 'put',
    data
  })
}

// @Tags CrmApprovalTasks
// @Summary 用id查询审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmApprovalTasks true "用id查询审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalTasks/findCrmApprovalTasks [get]
export const findCrmApprovalTasks = (params) => {
  return service({
    url: '/crmApprovalTasks/findCrmApprovalTasks',
    method: 'get',
    params
  })
}

// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
export const getCrmApprovalTasksList = (params) => {
  return service({
    url: '/crmApprovalTasks/getCrmApprovalTasksList',
    method: 'get',
    params
  })
}
