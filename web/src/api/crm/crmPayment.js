import service from '@/utils/request'

// @Tags CrmPayment
// @Summary 创建crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPayment true "创建crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPayment/createCrmPayment [post]
export const createCrmPayment = (data) => {
  return service({
    url: '/crmPayment/createCrmPayment',
    method: 'post',
    data
  })
}

// @Tags CrmPayment
// @Summary 删除crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPayment true "删除crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPayment/deleteCrmPayment [delete]
export const deleteCrmPayment = (params) => {
  return service({
    url: '/crmPayment/deleteCrmPayment',
    method: 'delete',
    params
  })
}

// @Tags CrmPayment
// @Summary 批量删除crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPayment/deleteCrmPayment [delete]
export const deleteCrmPaymentByIds = (params) => {
  return service({
    url: '/crmPayment/deleteCrmPaymentByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPayment
// @Summary 更新crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPayment true "更新crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPayment/updateCrmPayment [put]
export const updateCrmPayment = (data) => {
  return service({
    url: '/crmPayment/updateCrmPayment',
    method: 'put',
    data
  })
}

// @Tags CrmPayment
// @Summary 用id查询crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPayment true "用id查询crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPayment/findCrmPayment [get]
export const findCrmPayment = (params) => {
  return service({
    url: '/crmPayment/findCrmPayment',
    method: 'get',
    params
  })
}

// @Tags CrmPayment
// @Summary 分页获取crmPayment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmPayment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPayment/getCrmPaymentList [get]
export const getCrmPaymentList = (params) => {
  return service({
    url: '/crmPayment/getCrmPaymentList',
    method: 'get',
    params
  })
}
