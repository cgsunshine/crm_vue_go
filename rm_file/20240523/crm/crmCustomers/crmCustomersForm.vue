<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户名:" prop="customerName">
          <el-input v-model="formData.customerName" :clearable="true"  placeholder="请输入客户名" />
       </el-form-item>
        <el-form-item label="客户手机号:" prop="customerPhoneData">
          <el-input v-model="formData.customerPhoneData" :clearable="true"  placeholder="请输入客户手机号" />
       </el-form-item>
        <el-form-item label="管理ID 销售代表:" prop="sysUserId">
          <el-input v-model.number="formData.sysUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理角色ID:" prop="sysUserAuthorityId">
          <el-input v-model.number="formData.sysUserAuthorityId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户公司:" prop="customerCompany">
          <el-input v-model="formData.customerCompany" :clearable="true"  placeholder="请输入客户公司" />
       </el-form-item>
        <el-form-item label="客户地址:" prop="customerAddress">
          <el-input v-model="formData.customerAddress" :clearable="true"  placeholder="请输入客户地址" />
       </el-form-item>
        <el-form-item label="描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入描述" />
       </el-form-item>
        <el-form-item label="用户状态:" prop="customerStatus">
          <el-input v-model="formData.customerStatus" :clearable="true"  placeholder="请输入用户状态" />
       </el-form-item>
        <el-form-item label="客户分组:" prop="customerGroup">
          <el-input v-model="formData.customerGroup" :clearable="true"  placeholder="请输入客户分组" />
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
  createCrmCustomers,
  updateCrmCustomers,
  findCrmCustomers
} from '@/api/crm/crmCustomers'

defineOptions({
    name: 'CrmCustomersForm'
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
            customerName: '',
            customerPhoneData: '',
            sysUserId: 0,
            sysUserAuthorityId: 0,
            customerCompany: '',
            customerAddress: '',
            description: '',
            customerStatus: '',
            customerGroup: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmCustomers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmCustomers
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
