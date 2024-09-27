import service from '@/utils/request'

// @Tags CrmApprovalTasksRole
// @Summary 创建角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasksRole true "创建角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalTasksRole/createCrmApprovalTasksRole [post]
export const createCrmApprovalTasksRole = (data) => {
  return service({
    url: '/crmApprovalTasksRole/createCrmApprovalTasksRole',
    method: 'post',
    data
  })
}

// @Tags CrmApprovalTasksRole
// @Summary 删除角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasksRole true "删除角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasksRole/deleteCrmApprovalTasksRole [delete]
export const deleteCrmApprovalTasksRole = (params) => {
  return service({
    url: '/crmApprovalTasksRole/deleteCrmApprovalTasksRole',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalTasksRole
// @Summary 批量删除角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasksRole/deleteCrmApprovalTasksRole [delete]
export const deleteCrmApprovalTasksRoleByIds = (params) => {
  return service({
    url: '/crmApprovalTasksRole/deleteCrmApprovalTasksRoleByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalTasksRole
// @Summary 更新角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalTasksRole true "更新角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalTasksRole/updateCrmApprovalTasksRole [put]
export const updateCrmApprovalTasksRole = (data) => {
  return service({
    url: '/crmApprovalTasksRole/updateCrmApprovalTasksRole',
    method: 'put',
    data
  })
}

// @Tags CrmApprovalTasksRole
// @Summary 用id查询角色审批表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmApprovalTasksRole true "用id查询角色审批表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalTasksRole/findCrmApprovalTasksRole [get]
export const findCrmApprovalTasksRole = (params) => {
  return service({
    url: '/crmApprovalTasksRole/findCrmApprovalTasksRole',
    method: 'get',
    params
  })
}

// @Tags CrmApprovalTasksRole
// @Summary 分页获取角色审批表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取角色审批表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasksRole/getCrmApprovalTasksRoleList [get]
export const getCrmApprovalTasksRoleList = (params) => {
  return service({
    url: '/crmApprovalTasksRole/getCrmApprovalTasksRoleList',
    method: 'get',
    params
  })
}
