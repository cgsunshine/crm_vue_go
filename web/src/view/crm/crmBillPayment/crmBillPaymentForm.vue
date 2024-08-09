<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="应付账单名称:" prop="billPaymentName">
          <el-input v-model="formData.billPaymentName" :clearable="true"  placeholder="请输入应付账单名称" />
       </el-form-item>
        <el-form-item label="币种:" prop="currency">
           <el-select v-model="formData.currency" placeholder="请选择币种" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in currencyOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="客户ID:" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款状态:" prop="paymentStatus">
           <el-select v-model="formData.paymentStatus" placeholder="请选择付款状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in payment_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="对账单ID:" prop="statementAccountId">
          <el-input v-model.number="formData.statementAccountId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCrmBillPayment,
  updateCrmBillPayment,
  findCrmBillPayment
} from '@/api/crm/crmBillPayment'

defineOptions({
    name: 'CrmBillPaymentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmBillPayment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmBillPayment
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    currencyOptions.value = await getDictFunc('currency')
    payment_statusOptions.value = await getDictFunc('payment_status')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
