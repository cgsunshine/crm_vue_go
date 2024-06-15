import service from '@/utils/request'

// @Tags CrmContactApprovalTasks
// @Summary 创建合同审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalTasks true "创建合同审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactApprovalTasks/createCrmContactApprovalTasks [post]
export const createCrmContactApprovalTasks = (data) => {
  return service({
    url: '/crmContactApprovalTasks/createCrmContactApprovalTasks',
    method: 'post',
    data
  })
}

// @Tags CrmContactApprovalTasks
// @Summary 删除合同审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalTasks true "删除合同审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalTasks/deleteCrmContactApprovalTasks [delete]
export const deleteCrmContactApprovalTasks = (params) => {
  return service({
    url: '/crmContactApprovalTasks/deleteCrmContactApprovalTasks',
    method: 'delete',
    params
  })
}

// @Tags CrmContactApprovalTasks
// @Summary 批量删除合同审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合同审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalTasks/deleteCrmContactApprovalTasks [delete]
export const deleteCrmContactApprovalTasksByIds = (params) => {
  return service({
    url: '/crmContactApprovalTasks/deleteCrmContactApprovalTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContactApprovalTasks
// @Summary 更新合同审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalTasks true "更新合同审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalTasks/updateCrmContactApprovalTasks [put]
export const updateCrmContactApprovalTasks = (data) => {
  return service({
    url: '/crmContactApprovalTasks/updateCrmContactApprovalTasks',
    method: 'put',
    data
  })
}

// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContactApprovalTasks true "用id查询合同审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
export const findCrmContactApprovalTasks = (params) => {
  return service({
    url: '/crmContactApprovalTasks/findCrmContactApprovalTasks',
    method: 'get',
    params
  })
}

// @Tags CrmContactApprovalTasks
// @Summary 分页获取合同审批表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合同审批表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalTasks/getCrmContactApprovalTasksList [get]
export const getCrmContactApprovalTasksList = (params) => {
  return service({
    url: '/crmContactApprovalTasks/getCrmContactApprovalTasksList',
    method: 'get',
    params
  })
}
