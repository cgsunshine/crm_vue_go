<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="账单ID:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款金额:" prop="paymentAmount">
          <el-input-number v-model="formData.paymentAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="付款时间:" prop="paymentTime">
          <el-date-picker v-model="formData.paymentTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="付款凭证:" prop="paymentVoucher">
          <el-input v-model="formData.paymentVoucher" :clearable="true"  placeholder="请输入付款凭证" />
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
  createCrmPayment,
  updateCrmPayment,
  findCrmPayment
} from '@/api/crm/crmPayment'

defineOptions({
    name: 'CrmPaymentForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            orderId: 0,
            userId: 0,
            paymentAmount: 0,
            paymentTime: new Date(),
            paymentVoucher: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmPayment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmPayment
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmPayment(formData.value)
               break
             case 'update':
               res = await updateCrmPayment(formData.value)
               break
             default:
               res = await createCrmPayment(formData.value)
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
