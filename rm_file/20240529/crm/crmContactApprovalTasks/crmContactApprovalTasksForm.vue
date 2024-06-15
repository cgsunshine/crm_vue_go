<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联的审批请求ID:" prop="requestId">
          <el-input v-model.number="formData.requestId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前审批步骤ID:" prop="stepId">
          <el-input v-model.number="formData.stepId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="指派审批人:" prop="assignee">
          <el-input v-model="formData.assignee" :clearable="true"  placeholder="请输入指派审批人" />
       </el-form-item>
        <el-form-item label="审批状态:" prop="approvalStatus">
           <el-select v-model="formData.approvalStatus" placeholder="请选择审批状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in approval_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="审批意见:" prop="comments">
          <el-input v-model="formData.comments" :clearable="true"  placeholder="请输入审批意见" />
       </el-form-item>
        <el-form-item label="合同id:" prop="contactId">
          <el-input v-model.number="formData.contactId" :clearable="true" placeholder="请输入" />
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
  createCrmContactApprovalTasks,
  updateCrmContactApprovalTasks,
  findCrmContactApprovalTasks
} from '@/api/crm/crmContactApprovalTasks'

defineOptions({
    name: 'CrmContactApprovalTasksForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const approval_statusOptions = ref([])
const formData = ref({
            requestId: 0,
            stepId: 0,
            assignee: '',
            approvalStatus: '',
            comments: '',
            contactId: 0,
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
      const res = await findCrmContactApprovalTasks({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmContactApprovalTasks
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    approval_statusOptions.value = await getDictFunc('approval_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmContactApprovalTasks(formData.value)
               break
             case 'update':
               res = await updateCrmContactApprovalTasks(formData.value)
               break
             default:
               res = await createCrmContactApprovalTasks(formData.value)
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
