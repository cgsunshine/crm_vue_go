import service from '@/utils/request'

// @Tags CrmProductGroup
// @Summary 创建产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductGroup true "创建产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProductGroup/createCrmProductGroup [post]
export const createCrmProductGroup = (data) => {
  return service({
    url: '/crmProductGroup/createCrmProductGroup',
    method: 'post',
    data
  })
}

// @Tags CrmProductGroup
// @Summary 删除产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductGroup true "删除产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductGroup/deleteCrmProductGroup [delete]
export const deleteCrmProductGroup = (params) => {
  return service({
    url: '/crmProductGroup/deleteCrmProductGroup',
    method: 'delete',
    params
  })
}

// @Tags CrmProductGroup
// @Summary 批量删除产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductGroup/deleteCrmProductGroup [delete]
export const deleteCrmProductGroupByIds = (params) => {
  return service({
    url: '/crmProductGroup/deleteCrmProductGroupByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmProductGroup
// @Summary 更新产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmProductGroup true "更新产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProductGroup/updateCrmProductGroup [put]
export const updateCrmProductGroup = (data) => {
  return service({
    url: '/crmProductGroup/updateCrmProductGroup',
    method: 'put',
    data
  })
}

// @Tags CrmProductGroup
// @Summary 用id查询产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmProductGroup true "用id查询产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProductGroup/findCrmProductGroup [get]
export const findCrmProductGroup = (params) => {
  return service({
    url: '/crmProductGroup/findCrmProductGroup',
    method: 'get',
    params
  })
}

// @Tags CrmProductGroup
// @Summary 分页获取产品分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取产品分组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductGroup/getCrmProductGroupList [get]
export const getCrmProductGroupList = (params) => {
  return service({
    url: '/crmProductGroup/getCrmProductGroupList',
    method: 'get',
    params
  })
}
