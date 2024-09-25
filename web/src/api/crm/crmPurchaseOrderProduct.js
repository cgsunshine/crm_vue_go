import service from '@/utils/request'

// @Tags CrmPurchaseOrderProduct
// @Summary 创建crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrderProduct true "创建crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPurchaseOrderProduct/createCrmPurchaseOrderProduct [post]
export const createCrmPurchaseOrderProduct = (data) => {
  return service({
    url: '/crmPurchaseOrderProduct/createCrmPurchaseOrderProduct',
    method: 'post',
    data
  })
}

// @Tags CrmPurchaseOrderProduct
// @Summary 删除crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrderProduct true "删除crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrderProduct/deleteCrmPurchaseOrderProduct [delete]
export const deleteCrmPurchaseOrderProduct = (params) => {
  return service({
    url: '/crmPurchaseOrderProduct/deleteCrmPurchaseOrderProduct',
    method: 'delete',
    params
  })
}

// @Tags CrmPurchaseOrderProduct
// @Summary 批量删除crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrderProduct/deleteCrmPurchaseOrderProduct [delete]
export const deleteCrmPurchaseOrderProductByIds = (params) => {
  return service({
    url: '/crmPurchaseOrderProduct/deleteCrmPurchaseOrderProductByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPurchaseOrderProduct
// @Summary 更新crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPurchaseOrderProduct true "更新crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPurchaseOrderProduct/updateCrmPurchaseOrderProduct [put]
export const updateCrmPurchaseOrderProduct = (data) => {
  return service({
    url: '/crmPurchaseOrderProduct/updateCrmPurchaseOrderProduct',
    method: 'put',
    data
  })
}

// @Tags CrmPurchaseOrderProduct
// @Summary 用id查询crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPurchaseOrderProduct true "用id查询crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPurchaseOrderProduct/findCrmPurchaseOrderProduct [get]
export const findCrmPurchaseOrderProduct = (params) => {
  return service({
    url: '/crmPurchaseOrderProduct/findCrmPurchaseOrderProduct',
    method: 'get',
    params
  })
}

// @Tags CrmPurchaseOrderProduct
// @Summary 分页获取crmPurchaseOrderProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmPurchaseOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrderProduct/getCrmPurchaseOrderProductList [get]
export const getCrmPurchaseOrderProductList = (params) => {
  return service({
    url: '/crmPurchaseOrderProduct/getCrmPurchaseOrderProductList',
    method: 'get',
    params
  })
}
