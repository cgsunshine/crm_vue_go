import service from '@/utils/request'

// @Tags CrmTicketComments
// @Summary 创建共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketComments true "创建共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketComments/createCrmTicketComments [post]
export const createCrmTicketComments = (data) => {
  return service({
    url: '/crmTicketComments/createCrmTicketComments',
    method: 'post',
    data
  })
}

// @Tags CrmTicketComments
// @Summary 删除共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketComments true "删除共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketComments/deleteCrmTicketComments [delete]
export const deleteCrmTicketComments = (params) => {
  return service({
    url: '/crmTicketComments/deleteCrmTicketComments',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketComments
// @Summary 批量删除共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketComments/deleteCrmTicketComments [delete]
export const deleteCrmTicketCommentsByIds = (params) => {
  return service({
    url: '/crmTicketComments/deleteCrmTicketCommentsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketComments
// @Summary 更新共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketComments true "更新共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketComments/updateCrmTicketComments [put]
export const updateCrmTicketComments = (data) => {
  return service({
    url: '/crmTicketComments/updateCrmTicketComments',
    method: 'put',
    data
  })
}

// @Tags CrmTicketComments
// @Summary 用id查询共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmTicketComments true "用id查询共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketComments/findCrmTicketComments [get]
export const findCrmTicketComments = (params) => {
  return service({
    url: '/crmTicketComments/findCrmTicketComments',
    method: 'get',
    params
  })
}

// @Tags CrmTicketComments
// @Summary 分页获取共单回复列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取共单回复列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketComments/getCrmTicketCommentsList [get]
export const getCrmTicketCommentsList = (params) => {
  return service({
    url: '/crmTicketComments/getCrmTicketCommentsList',
    method: 'get',
    params
  })
}
