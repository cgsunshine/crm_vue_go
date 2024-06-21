import service from '@/utils/request'

// @Tags CrmApprovalNode
// @Summary 创建crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalNode true "创建crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalNode/createCrmApprovalNode [post]
export const createCrmApprovalNode = (data) => {
  return service({
    url: '/crmApprovalNode/createCrmApprovalNode',
    method: 'post',
    data
  })
}

// @Tags CrmApprovalNode
// @Summary 删除crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalNode true "删除crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalNode/deleteCrmApprovalNode [delete]
export const deleteCrmApprovalNode = (params) => {
  return service({
    url: '/crmApprovalNode/deleteCrmApprovalNode',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalNode
// @Summary 批量删除crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalNode/deleteCrmApprovalNode [delete]
export const deleteCrmApprovalNodeByIds = (params) => {
  return service({
    url: '/crmApprovalNode/deleteCrmApprovalNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalNode
// @Summary 更新crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalNode true "更新crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalNode/updateCrmApprovalNode [put]
export const updateCrmApprovalNode = (data) => {
  return service({
    url: '/crmApprovalNode/updateCrmApprovalNode',
    method: 'put',
    data
  })
}

// @Tags CrmApprovalNode
// @Summary 用id查询crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmApprovalNode true "用id查询crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalNode/findCrmApprovalNode [get]
export const findCrmApprovalNode = (params) => {
  return service({
    url: '/crmApprovalNode/findCrmApprovalNode',
    method: 'get',
    params
  })
}

// @Tags CrmApprovalNode
// @Summary 分页获取crmApprovalNode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmApprovalNode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalNode/getCrmApprovalNodeList [get]
export const getCrmApprovalNodeList = (params) => {
  return service({
    url: '/crmApprovalNode/getCrmApprovalNodeList',
    method: 'get',
    params
  })
}
