import service from '@/utils/request'

// @Tags CrmBillPayment
// @Summary 创建应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBillPayment true "创建应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBillPayment/createCrmBillPayment [post]
export const createCrmBillPayment = (data) => {
  return service({
    url: '/crmBillPayment/createCrmBillPayment',
    method: 'post',
    data
  })
}

// @Tags CrmBillPayment
// @Summary 删除应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBillPayment true "删除应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBillPayment/deleteCrmBillPayment [delete]
export const deleteCrmBillPayment = (params) => {
  return service({
    url: '/crmBillPayment/deleteCrmBillPayment',
    method: 'delete',
    params
  })
}

// @Tags CrmBillPayment
// @Summary 批量删除应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBillPayment/deleteCrmBillPayment [delete]
export const deleteCrmBillPaymentByIds = (params) => {
  return service({
    url: '/crmBillPayment/deleteCrmBillPaymentByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBillPayment
// @Summary 更新应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBillPayment true "更新应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBillPayment/updateCrmBillPayment [put]
export const updateCrmBillPayment = (data) => {
  return service({
    url: '/crmBillPayment/updateCrmBillPayment',
    method: 'put',
    data
  })
}

// @Tags CrmBillPayment
// @Summary 用id查询应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBillPayment true "用id查询应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBillPayment/findCrmBillPayment [get]
export const findCrmBillPayment = (params) => {
  return service({
    url: '/crmBillPayment/findCrmBillPayment',
    method: 'get',
    params
  })
}

// @Tags CrmBillPayment
// @Summary 分页获取应付账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取应付账单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBillPayment/getCrmBillPaymentList [get]
export const getCrmBillPaymentList = (params) => {
  return service({
    url: '/crmBillPayment/getCrmBillPaymentList',
    method: 'get',
    params
  })
}
