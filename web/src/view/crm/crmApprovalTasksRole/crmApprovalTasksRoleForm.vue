<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="审批状态:" prop="approvalStatus">
          <el-input v-model="formData.approvalStatus" :clearable="true"  placeholder="请输入审批状态" />
       </el-form-item>
        <el-form-item label="审批类型 1合同 2商机 3回款:" prop="approvalType">
          <el-input v-model.number="formData.approvalType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批人id:" prop="assigneeId">
          <el-input v-model.number="formData.assigneeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联id 合同 商机 回款:" prop="associatedId">
          <el-input v-model.number="formData.associatedId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批意见:" prop="comments">
          <el-input v-model="formData.comments" :clearable="true"  placeholder="请输入审批意见" />
       </el-form-item>
        <el-form-item label="关联的审批请求ID:" prop="requestId">
          <el-input v-model.number="formData.requestId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批角色id:" prop="roleId">
          <el-input v-model.number="formData.roleId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前审批步骤ID:" prop="stepId">
          <el-input v-model.number="formData.stepId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）:" prop="valid">
          <el-input v-model.number="formData.valid" :clearable="true" placeholder="请输入" />
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
  createCrmApprovalTasksRole,
  updateCrmApprovalTasksRole,
  findCrmApprovalTasksRole
} from '@/api/crm/crmApprovalTasksRole'

defineOptions({
    name: 'CrmApprovalTasksRoleForm'
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
            approvalStatus: '',
            approvalType: 0,
            assigneeId: 0,
            associatedId: 0,
            comments: '',
            requestId: 0,
            roleId: 0,
            stepId: 0,
            valid: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmApprovalTasksRole({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmApprovalTasksRole
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
               res = await createCrmApprovalTasksRole(formData.value)
               break
             case 'update':
               res = await updateCrmApprovalTasksRole(formData.value)
               break
             default:
               res = await createCrmApprovalTasksRole(formData.value)
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
