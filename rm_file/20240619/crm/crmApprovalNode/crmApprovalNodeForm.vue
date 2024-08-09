<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="预留条件:" prop="conditionExpression">
          <el-input v-model="formData.conditionExpression" :clearable="true"  placeholder="请输入预留条件" />
       </el-form-item>
        <el-form-item label="流程名称:" prop="nodeName">
          <el-input v-model="formData.nodeName" :clearable="true"  placeholder="请输入流程名称" />
       </el-form-item>
        <el-form-item label="节点顺序:" prop="nodeOrder">
          <el-input v-model.number="formData.nodeOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批通过的人数，需要几人同意才能通过审批:" prop="numberApprovedPersonnel">
          <el-input v-model.number="formData.numberApprovedPersonnel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属审批流程ID:" prop="processId">
          <el-input v-model="formData.processId" :clearable="true"  placeholder="请输入所属审批流程ID" />
       </el-form-item>
        <el-form-item label="角色ID:" prop="roleIds">
          <el-input v-model="formData.roleIds" :clearable="true"  placeholder="请输入角色ID" />
       </el-form-item>
        <el-form-item label="添加记录的用户:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="特定审批人ID:" prop="userIds">
          <el-input v-model="formData.userIds" :clearable="true"  placeholder="请输入特定审批人ID" />
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
  createCrmApprovalNode,
  updateCrmApprovalNode,
  findCrmApprovalNode
} from '@/api/crm/crmApprovalNode'

defineOptions({
    name: 'CrmApprovalNodeForm'
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
            conditionExpression: '',
            nodeName: '',
            nodeOrder: 0,
            numberApprovedPersonnel: 0,
            processId: '',
            roleIds: '',
            userId: 0,
            userIds: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmApprovalNode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmApprovalNode
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
               res = await createCrmApprovalNode(formData.value)
               break
             case 'update':
               res = await updateCrmApprovalNode(formData.value)
               break
             default:
               res = await createCrmApprovalNode(formData.value)
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
