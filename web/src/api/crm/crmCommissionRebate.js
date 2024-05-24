import service from '@/utils/request'

// @Tags CrmCommissionRebate
// @Summary 创建crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCommissionRebate true "创建crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCommissionRebate/createCrmCommissionRebate [post]
export const createCrmCommissionRebate = (data) => {
  return service({
    url: '/crmCommissionRebate/createCrmCommissionRebate',
    method: 'post',
    data
  })
}

// @Tags CrmCommissionRebate
// @Summary 删除crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCommissionRebate true "删除crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCommissionRebate/deleteCrmCommissionRebate [delete]
export const deleteCrmCommissionRebate = (params) => {
  return service({
    url: '/crmCommissionRebate/deleteCrmCommissionRebate',
    method: 'delete',
    params
  })
}

// @Tags CrmCommissionRebate
// @Summary 批量删除crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCommissionRebate/deleteCrmCommissionRebate [delete]
export const deleteCrmCommissionRebateByIds = (params) => {
  return service({
    url: '/crmCommissionRebate/deleteCrmCommissionRebateByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmCommissionRebate
// @Summary 更新crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCommissionRebate true "更新crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCommissionRebate/updateCrmCommissionRebate [put]
export const updateCrmCommissionRebate = (data) => {
  return service({
    url: '/crmCommissionRebate/updateCrmCommissionRebate',
    method: 'put',
    data
  })
}

// @Tags CrmCommissionRebate
// @Summary 用id查询crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmCommissionRebate true "用id查询crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCommissionRebate/findCrmCommissionRebate [get]
export const findCrmCommissionRebate = (params) => {
  return service({
    url: '/crmCommissionRebate/findCrmCommissionRebate',
    method: 'get',
    params
  })
}

// @Tags CrmCommissionRebate
// @Summary 分页获取crmCommissionRebate表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmCommissionRebate表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCommissionRebate/getCrmCommissionRebateList [get]
export const getCrmCommissionRebateList = (params) => {
  return service({
    url: '/crmCommissionRebate/getCrmCommissionRebateList',
    method: 'get',
    params
  })
}
