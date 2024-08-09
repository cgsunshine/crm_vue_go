import service from '@/utils/request'

// @Tags CrmCurrency
// @Summary 创建币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCurrency true "创建币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCurrency/createCrmCurrency [post]
export const createCrmCurrency = (data) => {
  return service({
    url: '/crmCurrency/createCrmCurrency',
    method: 'post',
    data
  })
}

// @Tags CrmCurrency
// @Summary 删除币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCurrency true "删除币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCurrency/deleteCrmCurrency [delete]
export const deleteCrmCurrency = (params) => {
  return service({
    url: '/crmCurrency/deleteCrmCurrency',
    method: 'delete',
    params
  })
}

// @Tags CrmCurrency
// @Summary 批量删除币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCurrency/deleteCrmCurrency [delete]
export const deleteCrmCurrencyByIds = (params) => {
  return service({
    url: '/crmCurrency/deleteCrmCurrencyByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmCurrency
// @Summary 更新币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCurrency true "更新币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCurrency/updateCrmCurrency [put]
export const updateCrmCurrency = (data) => {
  return service({
    url: '/crmCurrency/updateCrmCurrency',
    method: 'put',
    data
  })
}

// @Tags CrmCurrency
// @Summary 用id查询币种管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmCurrency true "用id查询币种管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCurrency/findCrmCurrency [get]
export const findCrmCurrency = (params) => {
  return service({
    url: '/crmCurrency/findCrmCurrency',
    method: 'get',
    params
  })
}

// @Tags CrmCurrency
// @Summary 分页获取币种管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取币种管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCurrency/getCrmCurrencyList [get]
export const getCrmCurrencyList = (params) => {
  return service({
    url: '/crmCurrency/getCrmCurrencyList',
    method: 'get',
    params
  })
}
