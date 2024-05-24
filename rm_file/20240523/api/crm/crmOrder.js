import service from '@/utils/request'

// @Tags CrmOrder
// @Summary 创建订单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrder true "创建订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrder/createCrmOrder [post]
export const createCrmOrder = (data) => {
  return service({
    url: '/crmOrder/createCrmOrder',
    method: 'post',
    data
  })
}

// @Tags CrmOrder
// @Summary 删除订单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrder true "删除订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrder/deleteCrmOrder [delete]
export const deleteCrmOrder = (params) => {
  return service({
    url: '/crmOrder/deleteCrmOrder',
    method: 'delete',
    params
  })
}

// @Tags CrmOrder
// @Summary 批量删除订单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrder/deleteCrmOrder [delete]
export const deleteCrmOrderByIds = (params) => {
  return service({
    url: '/crmOrder/deleteCrmOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmOrder
// @Summary 更新订单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrder true "更新订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOrder/updateCrmOrder [put]
export const updateCrmOrder = (data) => {
  return service({
    url: '/crmOrder/updateCrmOrder',
    method: 'put',
    data
  })
}

// @Tags CrmOrder
// @Summary 用id查询订单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmOrder true "用id查询订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrder/findCrmOrder [get]
export const findCrmOrder = (params) => {
  return service({
    url: '/crmOrder/findCrmOrder',
    method: 'get',
    params
  })
}

// @Tags CrmOrder
// @Summary 分页获取订单管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrder/getCrmOrderList [get]
export const getCrmOrderList = (params) => {
  return service({
    url: '/crmOrder/getCrmOrderList',
    method: 'get',
    params
  })
}
