<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商机名称:" prop="businessOpportunityName">
          <el-input v-model="formData.businessOpportunityName" :clearable="true"  placeholder="请输入商机名称" />
       </el-form-item>
        <el-form-item label="备注:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="客户ID:" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品id:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商机金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商机状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入商机状态" />
       </el-form-item>
        <el-form-item label="商机录入时间:" prop="inputTime">
          <el-date-picker v-model="formData.inputTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="报价有效期:" prop="offerValidityPeriod">
          <el-date-picker v-model="formData.offerValidityPeriod" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createCrmBusinessOpportunity,
  updateCrmBusinessOpportunity,
  findCrmBusinessOpportunity
} from '@/api/crm/crmBusinessOpportunity'

defineOptions({
    name: 'CrmBusinessOpportunityForm'
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
            businessOpportunityName: '',
            description: '',
            customerId: 0,
            userId: 0,
            productId: 0,
            amount: 0,
            status: '',
            inputTime: new Date(),
            offerValidityPeriod: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmBusinessOpportunity({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmBusinessOpportunity
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
               res = await createCrmBusinessOpportunity(formData.value)
               break
             case 'update':
               res = await updateCrmBusinessOpportunity(formData.value)
               break
             default:
               res = await createCrmBusinessOpportunity(formData.value)
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
