import service from '@/utils/request'

// @Tags CrmSupplier
// @Summary 创建crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmSupplier true "创建crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmSupplier/createCrmSupplier [post]
export const createCrmSupplier = (data) => {
  return service({
    url: '/crmSupplier/createCrmSupplier',
    method: 'post',
    data
  })
}

// @Tags CrmSupplier
// @Summary 删除crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmSupplier true "删除crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmSupplier/deleteCrmSupplier [delete]
export const deleteCrmSupplier = (params) => {
  return service({
    url: '/crmSupplier/deleteCrmSupplier',
    method: 'delete',
    params
  })
}

// @Tags CrmSupplier
// @Summary 批量删除crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmSupplier/deleteCrmSupplier [delete]
export const deleteCrmSupplierByIds = (params) => {
  return service({
    url: '/crmSupplier/deleteCrmSupplierByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmSupplier
// @Summary 更新crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmSupplier true "更新crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmSupplier/updateCrmSupplier [put]
export const updateCrmSupplier = (data) => {
  return service({
    url: '/crmSupplier/updateCrmSupplier',
    method: 'put',
    data
  })
}

// @Tags CrmSupplier
// @Summary 用id查询crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmSupplier true "用id查询crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmSupplier/findCrmSupplier [get]
export const findCrmSupplier = (params) => {
  return service({
    url: '/crmSupplier/findCrmSupplier',
    method: 'get',
    params
  })
}

// @Tags CrmSupplier
// @Summary 分页获取crmSupplier表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmSupplier表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmSupplier/getCrmSupplierList [get]
export const getCrmSupplierList = (params) => {
  return service({
    url: '/crmSupplier/getCrmSupplierList',
    method: 'get',
    params
  })
}
