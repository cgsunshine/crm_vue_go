import service from '@/utils/request'

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 创建crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalTasks true "创建crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollentionApprovalTasks/createCrmPaymentCollentionApprovalTasks [post]
export const createCrmPaymentCollentionApprovalTasks = (data) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/createCrmPaymentCollentionApprovalTasks',
    method: 'post',
    data
  })
}

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 删除crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalTasks true "删除crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasks [delete]
export const deleteCrmPaymentCollentionApprovalTasks = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasks',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 批量删除crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasks [delete]
export const deleteCrmPaymentCollentionApprovalTasksByIds = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 更新crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalTasks true "更新crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollentionApprovalTasks/updateCrmPaymentCollentionApprovalTasks [put]
export const updateCrmPaymentCollentionApprovalTasks = (data) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/updateCrmPaymentCollentionApprovalTasks',
    method: 'put',
    data
  })
}

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 用id查询crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPaymentCollentionApprovalTasks true "用id查询crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollentionApprovalTasks/findCrmPaymentCollentionApprovalTasks [get]
export const findCrmPaymentCollentionApprovalTasks = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/findCrmPaymentCollentionApprovalTasks',
    method: 'get',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 分页获取crmPaymentCollentionApprovalTasks表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmPaymentCollentionApprovalTasks表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalTasks/getCrmPaymentCollentionApprovalTasksList [get]
export const getCrmPaymentCollentionApprovalTasksList = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalTasks/getCrmPaymentCollentionApprovalTasksList',
    method: 'get',
    params
  })
}
