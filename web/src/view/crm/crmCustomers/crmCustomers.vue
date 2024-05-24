<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="客户名" prop="customerName">
         <el-input v-model="searchInfo.customerName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户手机号" prop="customerPhoneData">
         <el-input v-model="searchInfo.customerPhoneData" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="管理ID 销售代表" prop="userId">
            
             <el-input v-model.number="searchInfo.userId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户公司" prop="customerCompany">
         <el-input v-model="searchInfo.customerCompany" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户地址" prop="customerAddress">
         <el-input v-model="searchInfo.customerAddress" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="用户状态" prop="customerStatus">
            <el-select v-model="searchInfo.customerStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.customerStatus=undefined}">
              <el-option v-for="(item,key) in customer_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="客户分组" prop="customerGroup">
         <el-input v-model="searchInfo.customerGroup" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="客户名" prop="customerName" width="120" />
        <el-table-column align="left" label="客户手机号" prop="customerPhoneData" width="120" />
        <el-table-column align="left" label="管理ID 销售代表" prop="userId" width="120" />
        <el-table-column align="left" label="管理角色ID" prop="userAuthorityId" width="120" />
        <el-table-column align="left" label="客户公司" prop="customerCompany" width="120" />
        <el-table-column align="left" label="客户地址" prop="customerAddress" width="120" />
        <el-table-column align="left" label="描述" prop="description" width="120" />
        <el-table-column align="left" label="用户状态" prop="customerStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.customerStatus,customer_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="客户分组id" prop="customerGroupId" width="120" />
        <el-table-column align="left" label="客户分组" prop="customerGroup" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCrmCustomersFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="客户名:"  prop="customerName" >
              <el-input v-model="formData.customerName" :clearable="true"  placeholder="请输入客户名" />
            </el-form-item>
            <el-form-item label="客户手机号:"  prop="customerPhoneData" >
              <el-input v-model="formData.customerPhoneData" :clearable="true"  placeholder="请输入客户手机号" />
            </el-form-item>
            <el-form-item label="管理ID 销售代表:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入管理ID 销售代表" />
            </el-form-item>
            <el-form-item label="管理角色ID:"  prop="userAuthorityId" >
              <el-input v-model.number="formData.userAuthorityId" :clearable="true" placeholder="请输入管理角色ID" />
            </el-form-item>
            <el-form-item label="客户公司:"  prop="customerCompany" >
              <el-input v-model="formData.customerCompany" :clearable="true"  placeholder="请输入客户公司" />
            </el-form-item>
            <el-form-item label="客户地址:"  prop="customerAddress" >
              <el-input v-model="formData.customerAddress" :clearable="true"  placeholder="请输入客户地址" />
            </el-form-item>
            <el-form-item label="描述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="用户状态:"  prop="customerStatus" >
              <el-select v-model="formData.customerStatus" placeholder="请选择用户状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in customer_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="客户分组id:"  prop="customerGroupId" >
              <el-input v-model.number="formData.customerGroupId" :clearable="true" placeholder="请输入客户分组id" />
            </el-form-item>
            <el-form-item label="客户分组:"  prop="customerGroup" >
              <el-input v-model="formData.customerGroup" :clearable="true"  placeholder="请输入客户分组" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="客户名">
                        {{ formData.customerName }}
                </el-descriptions-item>
                <el-descriptions-item label="客户手机号">
                        {{ formData.customerPhoneData }}
                </el-descriptions-item>
                <el-descriptions-item label="管理ID 销售代表">
                        {{ formData.userId }}
                </el-descriptions-item>
                <el-descriptions-item label="管理角色ID">
                        {{ formData.userAuthorityId }}
                </el-descriptions-item>
                <el-descriptions-item label="客户公司">
                        {{ formData.customerCompany }}
                </el-descriptions-item>
                <el-descriptions-item label="客户地址">
                        {{ formData.customerAddress }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.description }}
                </el-descriptions-item>
                <el-descriptions-item label="用户状态">
                        {{ filterDict(formData.customerStatus,customer_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="客户分组id">
                        {{ formData.customerGroupId }}
                </el-descriptions-item>
                <el-descriptions-item label="客户分组">
                        {{ formData.customerGroup }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCrmCustomers,
  deleteCrmCustomers,
  deleteCrmCustomersByIds,
  updateCrmCustomers,
  findCrmCustomers,
  getCrmCustomersList
} from '@/api/crm/crmCustomers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CrmCustomers'
})

// 自动化生成的字典（可能为空）以及字段
const customer_statusOptions = ref([])
const formData = ref({
        customerName: '',
        customerPhoneData: '',
        userId: 0,
        userAuthorityId: 0,
        customerCompany: '',
        customerAddress: '',
        description: '',
        customerStatus: '',
        customerGroupId: 0,
        customerGroup: '',
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCrmCustomersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    customer_statusOptions.value = await getDictFunc('customer_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCrmCustomersFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteCrmCustomersByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCrmCustomersFunc = async(row) => {
    const res = await findCrmCustomers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recrmCustomers
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCrmCustomersFunc = async (row) => {
    const res = await deleteCrmCustomers({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCrmCustomers({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recrmCustomers
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          customerName: '',
          customerPhoneData: '',
          userId: 0,
          userAuthorityId: 0,
          customerCompany: '',
          customerAddress: '',
          description: '',
          customerStatus: '',
          customerGroupId: 0,
          customerGroup: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        customerName: '',
        customerPhoneData: '',
        userId: 0,
        userAuthorityId: 0,
        customerCompany: '',
        customerAddress: '',
        description: '',
        customerStatus: '',
        customerGroupId: 0,
        customerGroup: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCrmCustomers(formData.value)
                  break
                case 'update':
                  res = await updateCrmCustomers(formData.value)
                  break
                default:
                  res = await createCrmCustomers(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
