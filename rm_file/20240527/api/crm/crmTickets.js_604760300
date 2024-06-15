import service from '@/utils/request'

// @Tags CrmTickets
// @Summary 创建工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTickets true "创建工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTickets/createCrmTickets [post]
export const createCrmTickets = (data) => {
  return service({
    url: '/crmTickets/createCrmTickets',
    method: 'post',
    data
  })
}

// @Tags CrmTickets
// @Summary 删除工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTickets true "删除工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTickets/deleteCrmTickets [delete]
export const deleteCrmTickets = (params) => {
  return service({
    url: '/crmTickets/deleteCrmTickets',
    method: 'delete',
    params
  })
}

// @Tags CrmTickets
// @Summary 批量删除工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTickets/deleteCrmTickets [delete]
export const deleteCrmTicketsByIds = (params) => {
  return service({
    url: '/crmTickets/deleteCrmTicketsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmTickets
// @Summary 更新工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTickets true "更新工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTickets/updateCrmTickets [put]
export const updateCrmTickets = (data) => {
  return service({
    url: '/crmTickets/updateCrmTickets',
    method: 'put',
    data
  })
}

// @Tags CrmTickets
// @Summary 用id查询工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmTickets true "用id查询工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTickets/findCrmTickets [get]
export const findCrmTickets = (params) => {
  return service({
    url: '/crmTickets/findCrmTickets',
    method: 'get',
    params
  })
}

// @Tags CrmTickets
// @Summary 分页获取工单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTickets/getCrmTicketsList [get]
export const getCrmTicketsList = (params) => {
  return service({
    url: '/crmTickets/getCrmTicketsList',
    method: 'get',
    params
  })
}
