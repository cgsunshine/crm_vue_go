import service from '@/utils/request'

// @Tags CrmCustomers
// @Summary 创建crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomers true "创建crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCustomers/createCrmCustomers [post]
export const createCrmCustomers = (data) => {
  return service({
    url: '/crmCustomers/createCrmCustomers',
    method: 'post',
    data
  })
}

// @Tags CrmCustomers
// @Summary 删除crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomers true "删除crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomers/deleteCrmCustomers [delete]
export const deleteCrmCustomers = (params) => {
  return service({
    url: '/crmCustomers/deleteCrmCustomers',
    method: 'delete',
    params
  })
}

// @Tags CrmCustomers
// @Summary 批量删除crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomers/deleteCrmCustomers [delete]
export const deleteCrmCustomersByIds = (params) => {
  return service({
    url: '/crmCustomers/deleteCrmCustomersByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmCustomers
// @Summary 更新crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomers true "更新crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomers/updateCrmCustomers [put]
export const updateCrmCustomers = (data) => {
  return service({
    url: '/crmCustomers/updateCrmCustomers',
    method: 'put',
    data
  })
}

// @Tags CrmCustomers
// @Summary 用id查询crmCustomers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmCustomers true "用id查询crmCustomers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomers/findCrmCustomers [get]
export const findCrmCustomers = (params) => {
  return service({
    url: '/crmCustomers/findCrmCustomers',
    method: 'get',
    params
  })
}

// @Tags CrmCustomers
// @Summary 分页获取crmCustomers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmCustomers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomers/getCrmCustomersList [get]
export const getCrmCustomersList = (params) => {
  console.log(params,"passsss");
  console.log(new Date(),"pass");
  return service({
    url: '/crmCustomers/getCrmCustomersList',
    method: 'get',
    params
  })
}
