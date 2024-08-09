import service from '@/utils/request'

// @Tags CrmProductType
// @Summary 创建产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductType true "创建产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProductType/createCrmProductType [post]
export const createCrmProductType = (data) => {
  return service({
    url: '/crmProductType/createCrmProductType',
    method: 'post',
    data
  })
}

// @Tags CrmProductType
// @Summary 删除产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductType true "删除产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductType/deleteCrmProductType [delete]
export const deleteCrmProductType = (params) => {
  return service({
    url: '/crmProductType/deleteCrmProductType',
    method: 'delete',
    params
  })
}

// @Tags CrmProductType
// @Summary 批量删除产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductType/deleteCrmProductType [delete]
export const deleteCrmProductTypeByIds = (params) => {
  return service({
    url: '/crmProductType/deleteCrmProductTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmProductType
// @Summary 更新产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductType true "更新产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProductType/updateCrmProductType [put]
export const updateCrmProductType = (data) => {
  return service({
    url: '/crmProductType/updateCrmProductType',
    method: 'put',
    data
  })
}

// @Tags CrmProductType
// @Summary 用id查询产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmProductType true "用id查询产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProductType/findCrmProductType [get]
export const findCrmProductType = (params) => {
  return service({
    url: '/crmProductType/findCrmProductType',
    method: 'get',
    params
  })
}

// @Tags CrmProductType
// @Summary 分页获取产品类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取产品类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductType/getCrmProductTypeList [get]
export const getCrmProductTypeList = (params) => {
  return service({
    url: '/crmProductType/getCrmProductTypeList',
    method: 'get',
    params
  })
}
