<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="filter-bar">
        <el-select v-model="typeFilter" placeholder="举报类型" clearable style="width: 140px">
          <el-option label="软件举报" value="software" />
          <el-option label="评论举报" value="review" />
          <el-option label="用户举报" value="user" />
        </el-select>
        <el-select v-model="statusFilter" placeholder="处理状态" clearable style="width: 120px">
          <el-option label="待处理" value="pending" />
          <el-option label="已处理" value="handled" />
          <el-option label="已忽略" value="ignored" />
        </el-select>
        <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="举报类型" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="reportTypeTag(row.type)" size="small">{{ reportTypeText(row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="target" label="举报对象" width="150" show-overflow-tooltip />
        <el-table-column prop="reason" label="举报原因" width="120" />
        <el-table-column prop="description" label="举报描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="reporter" label="举报人" width="100" />
        <el-table-column prop="reportTime" label="举报时间" width="170" />
        <el-table-column label="处理状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="handleStatusType(row.status)" size="small">{{ handleStatusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="result" label="处理结果" width="120" show-overflow-tooltip />
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-button v-if="row.status === 'pending'" link type="primary" size="small" @click="openHandleDialog(row)">处理</el-button>
            <el-button v-else link type="info" size="small" @click="viewDetail(row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" :page-size="20" :total="total" layout="total, prev, pager, next" background />
      </div>
    </el-card>

    <!-- Handle Dialog -->
    <el-dialog v-model="handleDialogVisible" title="处理举报" width="600px">
      <template v-if="currentReport">
        <el-descriptions :column="1" border>
          <el-descriptions-item label="举报类型">{{ reportTypeText(currentReport.type) }}</el-descriptions-item>
          <el-descriptions-item label="举报对象">{{ currentReport.target }}</el-descriptions-item>
          <el-descriptions-item label="举报原因">{{ currentReport.reason }}</el-descriptions-item>
          <el-descriptions-item label="举报描述">{{ currentReport.description }}</el-descriptions-item>
          <el-descriptions-item label="举报人">{{ currentReport.reporter }}</el-descriptions-item>
          <el-descriptions-item label="举报时间">{{ currentReport.reportTime }}</el-descriptions-item>
        </el-descriptions>

        <el-divider />

        <el-form :model="handleForm" :rules="handleRules" ref="handleFormRef" label-width="80px">
          <el-form-item label="处理判定" prop="valid">
            <el-radio-group v-model="handleForm.valid">
              <el-radio :value="true">举报有效</el-radio>
              <el-radio :value="false">举报无效</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="处理说明" prop="note">
            <el-input v-model="handleForm.note" type="textarea" :rows="3" placeholder="请输入处理说明" />
          </el-form-item>
        </el-form>
      </template>
      <template #footer>
        <el-button @click="handleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmHandle">确定处理</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const typeFilter = ref('')
const statusFilter = ref('')
const currentPage = ref(1)
const total = ref(28)
const handleDialogVisible = ref(false)
const currentReport = ref<any>(null)
const handleFormRef = ref<FormInstance>()

const handleForm = reactive({ valid: true, note: '' })
const handleRules: FormRules = {
  valid: [{ required: true, message: '请选择判定结果' }],
  note: [{ required: true, message: '请输入处理说明', trigger: 'blur' }]
}

const tableData = ref([
  { id: 1, type: 'software', target: '违规工具 v1.0', reason: '恶意软件', description: '该软件包含恶意代码，窃取用户信息', reporter: '张三', reportTime: '2024-03-15 10:00:00', status: 'pending', result: '' },
  { id: 2, type: 'review', target: '评论#3892', reason: '广告内容', description: '评论中包含广告推广信息和联系方式', reporter: '李四', reportTime: '2024-03-14 15:30:00', status: 'pending', result: '' },
  { id: 3, type: 'user', target: '用户 广告号', reason: '恶意行为', description: '该用户多次发布广告评论', reporter: '王五', reportTime: '2024-03-14 11:20:00', status: 'handled', result: '已封禁用户' },
  { id: 4, type: 'software', target: '仿冒微信', reason: '侵权', description: '仿冒知名软件，存在侵权行为', reporter: '赵六', reportTime: '2024-03-13 09:00:00', status: 'handled', result: '已下架软件' },
  { id: 5, type: 'review', target: '评论#4210', reason: '辱骂他人', description: '评论内容包含人身攻击和辱骂', reporter: '孙七', reportTime: '2024-03-12 18:00:00', status: 'ignored', result: '举报无效' }
])

function reportTypeTag(t: string) {
  const map: Record<string, any> = { software: 'danger', review: 'warning', user: 'primary' }
  return map[t] || 'info'
}

function reportTypeText(t: string) {
  const map: Record<string, string> = { software: '软件举报', review: '评论举报', user: '用户举报' }
  return map[t] || t
}

function handleStatusType(s: string) {
  const map: Record<string, any> = { pending: 'warning', handled: 'success', ignored: 'info' }
  return map[s] || 'info'
}

function handleStatusText(s: string) {
  const map: Record<string, string> = { pending: '待处理', handled: '已处理', ignored: '已忽略' }
  return map[s] || s
}

function handleSearch() {
  ElMessage.info('搜索功能已触发')
}

function openHandleDialog(row: any) {
  currentReport.value = row
  handleForm.valid = true
  handleForm.note = ''
  handleDialogVisible.value = true
}

function viewDetail(row: any) {
  currentReport.value = row
  handleDialogVisible.value = true
}

async function confirmHandle() {
  await handleFormRef.value?.validate()
  if (currentReport.value) {
    currentReport.value.status = handleForm.valid ? 'handled' : 'ignored'
    currentReport.value.result = handleForm.valid ? '举报有效，已处理' : '举报无效'
  }
  handleDialogVisible.value = false
  ElMessage.success('处理完成')
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
</style>
