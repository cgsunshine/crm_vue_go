import service from '@/utils/request'

// @Tags CrmProcurementContract
// @Summary 创建订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProcurementContract true "创建订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProcurementContract/createCrmProcurementContract [post]
export const createCrmProcurementContract = (data) => {
  return service({
    url: '/crmProcurementContract/createCrmProcurementContract',
    method: 'post',
    data
  })
}

// @Tags CrmProcurementContract
// @Summary 删除订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProcurementContract true "删除订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProcurementContract/deleteCrmProcurementContract [delete]
export const deleteCrmProcurementContract = (params) => {
  return service({
    url: '/crmProcurementContract/deleteCrmProcurementContract',
    method: 'delete',
    params
  })
}

// @Tags CrmProcurementContract
// @Summary 批量删除订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProcurementContract/deleteCrmProcurementContract [delete]
export const deleteCrmProcurementContractByIds = (params) => {
  return service({
    url: '/crmProcurementContract/deleteCrmProcurementContractByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmProcurementContract
// @Summary 更新订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProcurementContract true "更新订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProcurementContract/updateCrmProcurementContract [put]
export const updateCrmProcurementContract = (data) => {
  return service({
    url: '/crmProcurementContract/updateCrmProcurementContract',
    method: 'put',
    data
  })
}

// @Tags CrmProcurementContract
// @Summary 用id查询订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmProcurementContract true "用id查询订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProcurementContract/findCrmProcurementContract [get]
export const findCrmProcurementContract = (params) => {
  return service({
    url: '/crmProcurementContract/findCrmProcurementContract',
    method: 'get',
    params
  })
}

// @Tags CrmProcurementContract
// @Summary 分页获取订购合同列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订购合同列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProcurementContract/getCrmProcurementContractList [get]
export const getCrmProcurementContractList = (params) => {
  return service({
    url: '/crmProcurementContract/getCrmProcurementContractList',
    method: 'get',
    params
  })
}
