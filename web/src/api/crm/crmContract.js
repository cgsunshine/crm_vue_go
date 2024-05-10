import service from '@/utils/request'

// @Tags CrmContract
// @Summary 创建合同管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContract true "创建合同管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContract/createCrmContract [post]
export const createCrmContract = (data) => {
  return service({
    url: '/crmContract/createCrmContract',
    method: 'post',
    data
  })
}

// @Tags CrmContract
// @Summary 删除合同管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContract true "删除合同管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContract/deleteCrmContract [delete]
export const deleteCrmContract = (params) => {
  return service({
    url: '/crmContract/deleteCrmContract',
    method: 'delete',
    params
  })
}

// @Tags CrmContract
// @Summary 批量删除合同管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合同管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContract/deleteCrmContract [delete]
export const deleteCrmContractByIds = (params) => {
  return service({
    url: '/crmContract/deleteCrmContractByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContract
// @Summary 更新合同管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContract true "更新合同管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContract/updateCrmContract [put]
export const updateCrmContract = (data) => {
  return service({
    url: '/crmContract/updateCrmContract',
    method: 'put',
    data
  })
}

// @Tags CrmContract
// @Summary 用id查询合同管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContract true "用id查询合同管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
export const findCrmContract = (params) => {
  return service({
    url: '/crmContract/findCrmContract',
    method: 'get',
    params
  })
}

// @Tags CrmContract
// @Summary 分页获取合同管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合同管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContract/getCrmContractList [get]
export const getCrmContractList = (params) => {
  return service({
    url: '/crmContract/getCrmContractList',
    method: 'get',
    params
  })
}
