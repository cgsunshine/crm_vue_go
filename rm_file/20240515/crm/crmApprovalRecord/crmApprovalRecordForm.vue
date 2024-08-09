<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="审批流程代码:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入审批流程代码" />
       </el-form-item>
        <el-form-item label="审批事项ID或模块ID:" prop="moduleId">
          <el-input v-model="formData.moduleId" :clearable="true"  placeholder="请输入审批事项ID或模块ID" />
       </el-form-item>
        <el-form-item label="审批相关信息或备注:" prop="message">
          <el-input v-model="formData.message" :clearable="true"  placeholder="请输入审批相关信息或备注" />
       </el-form-item>
        <el-form-item label="发起人ID:" prop="applicantId">
          <el-input v-model="formData.applicantId" :clearable="true"  placeholder="请输入发起人ID" />
       </el-form-item>
        <el-form-item label="申请时间:" prop="applyTime">
          <el-date-picker v-model="formData.applyTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="审批状态:" prop="流程状态">
       </el-form-item>
        <el-form-item label="当前审批节点ID:" prop="currentNodeId">
          <el-input v-model="formData.currentNodeId" :clearable="true"  placeholder="请输入当前审批节点ID" />
       </el-form-item>
        <el-form-item label="当前审批人ID:" prop="approverId">
          <el-input v-model="formData.approverId" :clearable="true"  placeholder="请输入当前审批人ID" />
       </el-form-item>
        <el-form-item label="审批时间:" prop="approveTime">
          <el-date-picker v-model="formData.approveTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="审批意见:" prop="approveOpinion">
          <el-input v-model="formData.approveOpinion" :clearable="true"  placeholder="请输入审批意见" />
       </el-form-item>
        <el-form-item label="最终审批结果:" prop="finalResult">
        <el-select v-model="formData.finalResult" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in []" :key="item" :label="item" :value="item" />
        </el-select>
       </el-form-item>
        <el-form-item label="流程关闭时间:" prop="closeTime">
          <el-date-picker v-model="formData.closeTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="记录创建者:" prop="creator">
          <el-input v-model="formData.creator" :clearable="true"  placeholder="请输入记录创建者" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="记录最后更新者:" prop="updator">
          <el-input v-model="formData.updator" :clearable="true"  placeholder="请输入记录最后更新者" />
       </el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="逻辑删除标志:" prop="isDeleted">
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
  createCrmApprovalRecord,
  updateCrmApprovalRecord,
  findCrmApprovalRecord
} from '@/api/crm/crmApprovalRecord'

defineOptions({
    name: 'CrmApprovalRecordForm'
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
            code: '',
            moduleId: '',
            message: '',
            applicantId: '',
            applyTime: new Date(),
            currentNodeId: '',
            approverId: '',
            approveTime: new Date(),
            approveOpinion: '',
            closeTime: new Date(),
            creator: '',
            createTime: new Date(),
            updator: '',
            updateTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmApprovalRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmApprovalRecord
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
               res = await createCrmApprovalRecord(formData.value)
               break
             case 'update':
               res = await updateCrmApprovalRecord(formData.value)
               break
             default:
               res = await createCrmApprovalRecord(formData.value)
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
