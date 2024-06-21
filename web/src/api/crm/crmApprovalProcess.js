import service from '@/utils/request'

// @Tags CrmApprovalProcess
// @Summary 创建审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalProcess true "创建审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalProcess/createCrmApprovalProcess [post]
export const createCrmApprovalProcess = (data) => {
  return service({
    url: '/crmApprovalProcess/createCrmApprovalProcess',
    method: 'post',
    data
  })
}

// @Tags CrmApprovalProcess
// @Summary 删除审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalProcess true "删除审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalProcess/deleteCrmApprovalProcess [delete]
export const deleteCrmApprovalProcess = (params) => {
  return service({
    url: '/crmApprovalProcess/deleteCrmApprovalProcess',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalProcess
// @Summary 批量删除审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalProcess/deleteCrmApprovalProcess [delete]
export const deleteCrmApprovalProcessByIds = (params) => {
  return service({
    url: '/crmApprovalProcess/deleteCrmApprovalProcessByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalProcess
// @Summary 更新审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalProcess true "更新审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalProcess/updateCrmApprovalProcess [put]
export const updateCrmApprovalProcess = (data) => {
  return service({
    url: '/crmApprovalProcess/updateCrmApprovalProcess',
    method: 'put',
    data
  })
}

// @Tags CrmApprovalProcess
// @Summary 用id查询审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmApprovalProcess true "用id查询审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalProcess/findCrmApprovalProcess [get]
export const findCrmApprovalProcess = (params) => {
  return service({
    url: '/crmApprovalProcess/findCrmApprovalProcess',
    method: 'get',
    params
  })
}

// @Tags CrmApprovalProcess
// @Summary 分页获取审批流程列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取审批流程列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalProcess/getCrmApprovalProcessList [get]
export const getCrmApprovalProcessList = (params) => {
  return service({
    url: '/crmApprovalProcess/getCrmApprovalProcessList',
    method: 'get',
    params
  })
}
