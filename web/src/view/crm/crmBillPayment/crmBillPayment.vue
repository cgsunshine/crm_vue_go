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
      
        <el-form-item label="金额" prop="amount">
            
             <el-input v-model.number="searchInfo.amount" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="应付账单名称" prop="billPaymentName">
         <el-input v-model="searchInfo.billPaymentName" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="币种" prop="currency">
            <el-select v-model="searchInfo.currency" clearable placeholder="请选择" @clear="()=>{searchInfo.currency=undefined}">
              <el-option v-for="(item,key) in currencyOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="客户ID" prop="customerId">
            
             <el-input v-model.number="searchInfo.customerId" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="付款状态" prop="paymentStatus">
            <el-select v-model="searchInfo.paymentStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.paymentStatus=undefined}">
              <el-option v-for="(item,key) in payment_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="对账单ID" prop="statementAccountId">
            
             <el-input v-model.number="searchInfo.statementAccountId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="用户id" prop="userId">
            
             <el-input v-model.number="searchInfo.userId" placeholder="搜索条件" />

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
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column sortable align="left" label="金额" prop="amount" width="120" />
        <el-table-column align="left" label="应付账单名称" prop="billPaymentName" width="120" />
        <el-table-column align="left" label="币种" prop="currency" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.currency,currencyOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="客户ID" prop="customerId" width="120" />
        <el-table-column align="left" label="付款状态" prop="paymentStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.paymentStatus,payment_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="对账单ID" prop="statementAccountId" width="120" />
        <el-table-column align="left" label="用户id" prop="userId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCrmBillPaymentFunc(scope.row)">变更</el-button>
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
            <el-form-item label="金额:"  prop="amount" >
              <el-input-number v-model="formData.amount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="应付账单名称:"  prop="billPaymentName" >
              <el-input v-model="formData.billPaymentName" :clearable="true"  placeholder="请输入应付账单名称" />
            </el-form-item>
            <el-form-item label="币种:"  prop="currency" >
              <el-select v-model="formData.currency" placeholder="请选择币种" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in currencyOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="客户ID:"  prop="customerId" >
              <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入客户ID" />
            </el-form-item>
            <el-form-item label="付款状态:"  prop="paymentStatus" >
              <el-select v-model="formData.paymentStatus" placeholder="请选择付款状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in payment_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="对账单ID:"  prop="statementAccountId" >
              <el-input v-model.number="formData.statementAccountId" :clearable="true" placeholder="请输入对账单ID" />
            </el-form-item>
            <el-form-item label="用户id:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户id" />
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
                <el-descriptions-item label="金额">
                        {{ formData.amount }}
                </el-descriptions-item>
                <el-descriptions-item label="应付账单名称">
                        {{ formData.billPaymentName }}
                </el-descriptions-item>
                <el-descriptions-item label="币种">
                        {{ filterDict(formData.currency,currencyOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="客户ID">
                        {{ formData.customerId }}
                </el-descriptions-item>
                <el-descriptions-item label="付款状态">
                        {{ filterDict(formData.paymentStatus,payment_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="对账单ID">
                        {{ formData.statementAccountId }}
                </el-descriptions-item>
                <el-descriptions-item label="用户id">
                        {{ formData.userId }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCrmBillPayment,
  deleteCrmBillPayment,
  deleteCrmBillPaymentByIds,
  updateCrmBillPayment,
  findCrmBillPayment,
  getCrmBillPaymentList
} from '@/api/crm/crmBillPayment'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CrmBillPayment'
})

// 自动化生成的字典（可能为空）以及字段
const currencyOptions = ref([])
const payment_statusOptions = ref([])
const formData = ref({
        amount: 0,
        billPaymentName: '',
        currency: '',
        customerId: 0,
        paymentStatus: '',
        statementAccountId: 0,
        userId: 0,
        })


// 验证规则
const rule = reactive({
               amount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            amount: 'amount',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
  const table = await getCrmBillPaymentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    currencyOptions.value = await getDictFunc('currency')
    payment_statusOptions.value = await getDictFunc('payment_status')
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
            deleteCrmBillPaymentFunc(row)
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
      const res = await deleteCrmBillPaymentByIds({ IDs })
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
const updateCrmBillPaymentFunc = async(row) => {
    const res = await findCrmBillPayment({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recrmBillPayment
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCrmBillPaymentFunc = async (row) => {
    const res = await deleteCrmBillPayment({ ID: row.ID })
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
  const res = await findCrmBillPayment({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recrmBillPayment
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          amount: 0,
          billPaymentName: '',
          currency: '',
          customerId: 0,
          paymentStatus: '',
          statementAccountId: 0,
          userId: 0,
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
        amount: 0,
        billPaymentName: '',
        currency: '',
        customerId: 0,
        paymentStatus: '',
        statementAccountId: 0,
        userId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCrmBillPayment(formData.value)
                  break
                case 'update':
                  res = await updateCrmBillPayment(formData.value)
                  break
                default:
                  res = await createCrmBillPayment(formData.value)
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
