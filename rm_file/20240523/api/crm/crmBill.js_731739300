import service from '@/utils/request'

// @Tags CrmBill
// @Summary 创建crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBill true "创建crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBill/createCrmBill [post]
export const createCrmBill = (data) => {
  return service({
    url: '/crmBill/createCrmBill',
    method: 'post',
    data
  })
}

// @Tags CrmBill
// @Summary 删除crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBill true "删除crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBill/deleteCrmBill [delete]
export const deleteCrmBill = (params) => {
  return service({
    url: '/crmBill/deleteCrmBill',
    method: 'delete',
    params
  })
}

// @Tags CrmBill
// @Summary 批量删除crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBill/deleteCrmBill [delete]
export const deleteCrmBillByIds = (params) => {
  return service({
    url: '/crmBill/deleteCrmBillByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBill
// @Summary 更新crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBill true "更新crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBill/updateCrmBill [put]
export const updateCrmBill = (data) => {
  return service({
    url: '/crmBill/updateCrmBill',
    method: 'put',
    data
  })
}

// @Tags CrmBill
// @Summary 用id查询crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/findCrmBill [get]
export const findCrmBill = (params) => {
  return service({
    url: '/crmBill/findCrmBill',
    method: 'get',
    params
  })
}

// @Tags CrmBill
// @Summary 分页获取crmBill表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmBill表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBill/getCrmBillList [get]
export const getCrmBillList = (params) => {
  return service({
    url: '/crmBill/getCrmBillList',
    method: 'get',
    params
  })
}
