import service from '@/utils/request'

// @Tags CrmStatementAccountFile
// @Summary 创建对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFile true "创建对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccountFile/createCrmStatementAccountFile [post]
export const createCrmStatementAccountFile = (data) => {
  return service({
    url: '/crmStatementAccountFile/createCrmStatementAccountFile',
    method: 'post',
    data
  })
}

// @Tags CrmStatementAccountFile
// @Summary 删除对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFile true "删除对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFile/deleteCrmStatementAccountFile [delete]
export const deleteCrmStatementAccountFile = (params) => {
  return service({
    url: '/crmStatementAccountFile/deleteCrmStatementAccountFile',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccountFile
// @Summary 批量删除对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFile/deleteCrmStatementAccountFile [delete]
export const deleteCrmStatementAccountFileByIds = (params) => {
  return service({
    url: '/crmStatementAccountFile/deleteCrmStatementAccountFileByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccountFile
// @Summary 更新对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFile true "更新对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccountFile/updateCrmStatementAccountFile [put]
export const updateCrmStatementAccountFile = (data) => {
  return service({
    url: '/crmStatementAccountFile/updateCrmStatementAccountFile',
    method: 'put',
    data
  })
}

// @Tags CrmStatementAccountFile
// @Summary 用id查询对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmStatementAccountFile true "用id查询对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccountFile/findCrmStatementAccountFile [get]
export const findCrmStatementAccountFile = (params) => {
  return service({
    url: '/crmStatementAccountFile/findCrmStatementAccountFile',
    method: 'get',
    params
  })
}

// @Tags CrmStatementAccountFile
// @Summary 分页获取对账单文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取对账单文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFile/getCrmStatementAccountFileList [get]
export const getCrmStatementAccountFileList = (params) => {
  return service({
    url: '/crmStatementAccountFile/getCrmStatementAccountFileList',
    method: 'get',
    params
  })
}
