import service from '@/utils/request'

// @Tags CrmTicketCategories
// @Summary 创建工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketCategories true "创建工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketCategories/createCrmTicketCategories [post]
export const createCrmTicketCategories = (data) => {
  return service({
    url: '/crmTicketCategories/createCrmTicketCategories',
    method: 'post',
    data
  })
}

// @Tags CrmTicketCategories
// @Summary 删除工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketCategories true "删除工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketCategories/deleteCrmTicketCategories [delete]
export const deleteCrmTicketCategories = (params) => {
  return service({
    url: '/crmTicketCategories/deleteCrmTicketCategories',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketCategories
// @Summary 批量删除工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketCategories/deleteCrmTicketCategories [delete]
export const deleteCrmTicketCategoriesByIds = (params) => {
  return service({
    url: '/crmTicketCategories/deleteCrmTicketCategoriesByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketCategories
// @Summary 更新工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketCategories true "更新工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketCategories/updateCrmTicketCategories [put]
export const updateCrmTicketCategories = (data) => {
  return service({
    url: '/crmTicketCategories/updateCrmTicketCategories',
    method: 'put',
    data
  })
}

// @Tags CrmTicketCategories
// @Summary 用id查询工单类别
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmTicketCategories true "用id查询工单类别"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketCategories/findCrmTicketCategories [get]
export const findCrmTicketCategories = (params) => {
  return service({
    url: '/crmTicketCategories/findCrmTicketCategories',
    method: 'get',
    params
  })
}

// @Tags CrmTicketCategories
// @Summary 分页获取工单类别列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工单类别列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketCategories/getCrmTicketCategoriesList [get]
export const getCrmTicketCategoriesList = (params) => {
  return service({
    url: '/crmTicketCategories/getCrmTicketCategoriesList',
    method: 'get',
    params
  })
}
