import service from '@/utils/request'

// @Tags CrmPurchaseOrder
// @Summary 创建订购单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrder true "创建订购单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPurchaseOrder/createCrmPurchaseOrder [post]
export const createCrmPurchaseOrder = (data) => {
  return service({
    url: '/crmPurchaseOrder/createCrmPurchaseOrder',
    method: 'post',
    data
  })
}

// @Tags CrmPurchaseOrder
// @Summary 删除订购单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrder true "删除订购单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrder/deleteCrmPurchaseOrder [delete]
export const deleteCrmPurchaseOrder = (params) => {
  return service({
    url: '/crmPurchaseOrder/deleteCrmPurchaseOrder',
    method: 'delete',
    params
  })
}

// @Tags CrmPurchaseOrder
// @Summary 批量删除订购单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订购单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrder/deleteCrmPurchaseOrder [delete]
export const deleteCrmPurchaseOrderByIds = (params) => {
  return service({
    url: '/crmPurchaseOrder/deleteCrmPurchaseOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPurchaseOrder
// @Summary 更新订购单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrder true "更新订购单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPurchaseOrder/updateCrmPurchaseOrder [put]
export const updateCrmPurchaseOrder = (data) => {
  return service({
    url: '/crmPurchaseOrder/updateCrmPurchaseOrder',
    method: 'put',
    data
  })
}

// @Tags CrmPurchaseOrder
// @Summary 用id查询订购单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPurchaseOrder true "用id查询订购单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPurchaseOrder/findCrmPurchaseOrder [get]
export const findCrmPurchaseOrder = (params) => {
  return service({
    url: '/crmPurchaseOrder/findCrmPurchaseOrder',
    method: 'get',
    params
  })
}

// @Tags CrmPurchaseOrder
// @Summary 分页获取订购单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订购单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrder/getCrmPurchaseOrderList [get]
export const getCrmPurchaseOrderList = (params) => {
  return service({
    url: '/crmPurchaseOrder/getCrmPurchaseOrderList',
    method: 'get',
    params
  })
}
