import service from '@/utils/request'

// @Tags CrmBusinessOpportunityProduct
// @Summary 创建crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityProduct true "创建crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityProduct/createCrmBusinessOpportunityProduct [post]
export const createCrmBusinessOpportunityProduct = (data) => {
  return service({
    url: '/crmBusinessOpportunityProduct/createCrmBusinessOpportunityProduct',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunityProduct
// @Summary 删除crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityProduct true "删除crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProduct [delete]
export const deleteCrmBusinessOpportunityProduct = (params) => {
  return service({
    url: '/crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProduct',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityProduct
// @Summary 批量删除crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProduct [delete]
export const deleteCrmBusinessOpportunityProductByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProductByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityProduct
// @Summary 更新crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityProduct true "更新crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityProduct/updateCrmBusinessOpportunityProduct [put]
export const updateCrmBusinessOpportunityProduct = (data) => {
  return service({
    url: '/crmBusinessOpportunityProduct/updateCrmBusinessOpportunityProduct',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunityProduct
// @Summary 用id查询crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunityProduct true "用id查询crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityProduct/findCrmBusinessOpportunityProduct [get]
export const findCrmBusinessOpportunityProduct = (params) => {
  return service({
    url: '/crmBusinessOpportunityProduct/findCrmBusinessOpportunityProduct',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunityProduct
// @Summary 分页获取crmBusinessOpportunityProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmBusinessOpportunityProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityProduct/getCrmBusinessOpportunityProductList [get]
export const getCrmBusinessOpportunityProductList = (params) => {
  return service({
    url: '/crmBusinessOpportunityProduct/getCrmBusinessOpportunityProductList',
    method: 'get',
    params
  })
}
