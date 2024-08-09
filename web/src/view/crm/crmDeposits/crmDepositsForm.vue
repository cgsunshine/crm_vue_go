<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="押金金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="货币类型:" prop="currency">
           <el-select v-model="formData.currency" placeholder="请选择货币类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in currencyOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="客户id:" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入描述" />
       </el-form-item>
        <el-form-item label="订单id:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付日期:" prop="paymentDate">
          <el-date-picker v-model="formData.paymentDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="退款日期:" prop="refundDate">
          <el-date-picker v-model="formData.refundDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="押金状态 已支付 已退款 待处理:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入押金状态 已支付 已退款 待处理" />
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
  createCrmDeposits,
  updateCrmDeposits,
  findCrmDeposits
} from '@/api/crm/crmDeposits'

defineOptions({
    name: 'CrmDepositsForm'
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
const formData = ref({
            amount: 0,
            currency: '',
            customerId: 0,
            description: '',
            orderId: 0,
            paymentDate: new Date(),
            refundDate: new Date(),
            status: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmDeposits({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmDeposits
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    currencyOptions.value = await getDictFunc('currency')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmDeposits(formData.value)
               break
             case 'update':
               res = await updateCrmDeposits(formData.value)
               break
             default:
               res = await createCrmDeposits(formData.value)
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
