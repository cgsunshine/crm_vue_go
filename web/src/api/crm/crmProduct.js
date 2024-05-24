import service from '@/utils/request'

// @Tags CrmProduct
// @Summary 创建crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProduct true "创建crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProduct/createCrmProduct [post]
export const createCrmProduct = (data) => {
  return service({
    url: '/crmProduct/createCrmProduct',
    method: 'post',
    data
  })
}

// @Tags CrmProduct
// @Summary 删除crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProduct true "删除crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProduct/deleteCrmProduct [delete]
export const deleteCrmProduct = (params) => {
  return service({
    url: '/crmProduct/deleteCrmProduct',
    method: 'delete',
    params
  })
}

// @Tags CrmProduct
// @Summary 批量删除crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProduct/deleteCrmProduct [delete]
export const deleteCrmProductByIds = (params) => {
  return service({
    url: '/crmProduct/deleteCrmProductByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmProduct
// @Summary 更新crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProduct true "更新crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProduct/updateCrmProduct [put]
export const updateCrmProduct = (data) => {
  return service({
    url: '/crmProduct/updateCrmProduct',
    method: 'put',
    data
  })
}

// @Tags CrmProduct
// @Summary 用id查询crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmProduct true "用id查询crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProduct/findCrmProduct [get]
export const findCrmProduct = (params) => {
  return service({
    url: '/crmProduct/findCrmProduct',
    method: 'get',
    params
  })
}

// @Tags CrmProduct
// @Summary 分页获取crmProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProduct/getCrmProductList [get]
export const getCrmProductList = (params) => {
  return service({
    url: '/crmProduct/getCrmProductList',
    method: 'get',
    params
  })
}
