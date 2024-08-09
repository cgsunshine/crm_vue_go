import service from '@/utils/request'

// @Tags CrmOrderProduct
// @Summary 创建crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrderProduct true "创建crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrderProduct/createCrmOrderProduct [post]
export const createCrmOrderProduct = (data) => {
  return service({
    url: '/crmOrderProduct/createCrmOrderProduct',
    method: 'post',
    data
  })
}

// @Tags CrmOrderProduct
// @Summary 删除crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrderProduct true "删除crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrderProduct/deleteCrmOrderProduct [delete]
export const deleteCrmOrderProduct = (params) => {
  return service({
    url: '/crmOrderProduct/deleteCrmOrderProduct',
    method: 'delete',
    params
  })
}

// @Tags CrmOrderProduct
// @Summary 批量删除crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrderProduct/deleteCrmOrderProduct [delete]
export const deleteCrmOrderProductByIds = (params) => {
  return service({
    url: '/crmOrderProduct/deleteCrmOrderProductByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmOrderProduct
// @Summary 更新crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOrderProduct true "更新crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOrderProduct/updateCrmOrderProduct [put]
export const updateCrmOrderProduct = (data) => {
  return service({
    url: '/crmOrderProduct/updateCrmOrderProduct',
    method: 'put',
    data
  })
}

// @Tags CrmOrderProduct
// @Summary 用id查询crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmOrderProduct true "用id查询crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrderProduct/findCrmOrderProduct [get]
export const findCrmOrderProduct = (params) => {
  return service({
    url: '/crmOrderProduct/findCrmOrderProduct',
    method: 'get',
    params
  })
}

// @Tags CrmOrderProduct
// @Summary 分页获取crmOrderProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrderProduct/getCrmOrderProductList [get]
export const getCrmOrderProductList = (params) => {
  return service({
    url: '/crmOrderProduct/getCrmOrderProductList',
    method: 'get',
    params
  })
}
