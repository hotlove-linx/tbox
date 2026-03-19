<template>
  <div class="page-container">
    <!-- Status Tabs -->
    <el-card shadow="never">
      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <el-tab-pane name="pending">
          <template #label>待审核 <el-badge :value="12" :max="99" class="tab-badge" /></template>
        </el-tab-pane>
        <el-tab-pane name="reviewing">
          <template #label>审核中 <el-badge :value="3" :max="99" class="tab-badge" /></template>
        </el-tab-pane>
        <el-tab-pane name="approved">
          <template #label>已通过 <el-badge :value="256" :max="99" class="tab-badge" /></template>
        </el-tab-pane>
        <el-tab-pane name="rejected">
          <template #label>已拒绝 <el-badge :value="18" :max="99" class="tab-badge" /></template>
        </el-tab-pane>
      </el-tabs>

      <!-- Filters -->
      <div class="filter-bar">
        <el-input v-model="searchText" placeholder="搜索软件名称" prefix-icon="Search" clearable style="width: 200px" />
        <el-select v-model="categoryFilter" placeholder="分类" clearable style="width: 140px">
          <el-option label="工具软件" value="tools" />
          <el-option label="游戏娱乐" value="games" />
          <el-option label="效率办公" value="productivity" />
        </el-select>
        <el-date-picker v-model="dateRange" type="daterange" range-separator="至" start-placeholder="开始" end-placeholder="结束" style="width: 240px" />
      </div>

      <!-- Table -->
      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="软件信息" min-width="200">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 10px;">
              <el-avatar :size="36" shape="square">{{ row.name[0] }}</el-avatar>
              <div>
                <div style="font-weight: 500;">{{ row.name }}</div>
                <div style="font-size: 12px; color: #909399;">v{{ row.version }} | {{ row.category }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="submitter" label="提交人" width="100" />
        <el-table-column prop="submitTime" label="提交时间" width="170" />
        <el-table-column prop="fileSize" label="文件大小" width="90" />
        <el-table-column label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)" size="small">{{ statusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="viewDetail(row)">详情</el-button>
            <el-button v-if="row.status === 'pending'" link type="success" size="small" @click="handleApprove(row)">通过</el-button>
            <el-button v-if="row.status === 'pending'" link type="danger" size="small" @click="handleReject(row)">拒绝</el-button>
            <el-button v-if="row.status === 'pending'" link type="warning" size="small" @click="handleTransfer(row)">转交</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" :page-size="20" :total="total" layout="total, prev, pager, next" background />
      </div>
    </el-card>

    <!-- Detail Dialog -->
    <el-dialog v-model="detailVisible" title="审核详情" width="800px" top="5vh">
      <template v-if="currentItem">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="软件名称">{{ currentItem.name }}</el-descriptions-item>
          <el-descriptions-item label="版本">v{{ currentItem.version }}</el-descriptions-item>
          <el-descriptions-item label="分类">{{ currentItem.category }}</el-descriptions-item>
          <el-descriptions-item label="文件大小">{{ currentItem.fileSize }}</el-descriptions-item>
          <el-descriptions-item label="提交人">{{ currentItem.submitter }}</el-descriptions-item>
          <el-descriptions-item label="提交时间">{{ currentItem.submitTime }}</el-descriptions-item>
          <el-descriptions-item label="描述" :span="2">
            一款功能强大的应用软件，提供丰富的功能和优质的用户体验。
          </el-descriptions-item>
        </el-descriptions>

        <div style="margin-top: 20px;">
          <h4>安全扫描报告</h4>
          <el-result icon="success" title="扫描通过" sub-title="未发现恶意代码、病毒或隐私违规问题" />
        </div>

        <div v-if="currentItem.status === 'pending'" style="margin-top: 20px;">
          <h4>审核操作</h4>
          <div style="display: flex; gap: 12px; margin-top: 12px;">
            <el-button type="success" @click="handleApprove(currentItem)"><el-icon><Select /></el-icon>通过</el-button>
            <el-button type="danger" @click="handleReject(currentItem)"><el-icon><CloseBold /></el-icon>拒绝</el-button>
            <el-button type="warning" @click="handleTransfer(currentItem)"><el-icon><Right /></el-icon>转交</el-button>
          </div>
        </div>
      </template>
    </el-dialog>

    <!-- Reject Dialog -->
    <el-dialog v-model="rejectVisible" title="拒绝审核" width="480px">
      <el-form :model="rejectForm" :rules="rejectRules" ref="rejectFormRef" label-width="80px">
        <el-form-item label="拒绝原因" prop="reason">
          <el-input v-model="rejectForm.reason" type="textarea" :rows="4" placeholder="请输入拒绝原因" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="rejectVisible = false">取消</el-button>
        <el-button type="danger" @click="confirmReject">确定拒绝</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const activeTab = ref('pending')
const searchText = ref('')
const categoryFilter = ref('')
const dateRange = ref(null)
const currentPage = ref(1)
const total = ref(12)
const detailVisible = ref(false)
const rejectVisible = ref(false)
const currentItem = ref<any>(null)
const rejectFormRef = ref<FormInstance>()
const rejectForm = reactive({ reason: '' })
const rejectRules: FormRules = { reason: [{ required: true, message: '请输入拒绝原因', trigger: 'blur' }] }

const allData: Record<string, any[]> = {
  pending: [
    { id: 2001, name: '智能笔记', version: '2.1.0', category: '效率办公', submitter: '王五', submitTime: '2024-03-15 08:30:00', fileSize: '15.8MB', status: 'pending' },
    { id: 2002, name: '密码管理器', version: '1.5.3', category: '工具软件', submitter: '郑十', submitTime: '2024-03-14 14:20:00', fileSize: '8.9MB', status: 'pending' },
    { id: 2003, name: '日程管理', version: '1.0.0', category: '效率办公', submitter: '陈一', submitTime: '2024-03-14 11:00:00', fileSize: '12.3MB', status: 'pending' },
    { id: 2004, name: '计算器Plus', version: '2.0.0', category: '工具软件', submitter: '林二', submitTime: '2024-03-13 16:45:00', fileSize: '5.2MB', status: 'pending' }
  ],
  reviewing: [
    { id: 2010, name: 'PDF阅读器', version: '3.0.0', category: '效率办公', submitter: '黄三', submitTime: '2024-03-12 09:00:00', fileSize: '22.1MB', status: 'reviewing' }
  ],
  approved: [
    { id: 1001, name: '超级文件管理器', version: '3.2.1', category: '工具软件', submitter: '张三', submitTime: '2024-03-10 10:30:00', fileSize: '28.5MB', status: 'approved' },
    { id: 1002, name: '极速浏览器', version: '5.0.0', category: '工具软件', submitter: '李四', submitTime: '2024-03-09 09:00:00', fileSize: '45.2MB', status: 'approved' }
  ],
  rejected: [
    { id: 3001, name: '违规应用', version: '1.0.0', category: '其他', submitter: '匿名', submitTime: '2024-03-08 15:00:00', fileSize: '3.1MB', status: 'rejected' }
  ]
}

const tableData = ref(allData.pending)

function handleTabChange(tab: string) {
  tableData.value = allData[tab] || []
  total.value = tableData.value.length
}

function statusType(s: string) {
  const map: Record<string, any> = { pending: 'warning', reviewing: 'primary', approved: 'success', rejected: 'danger' }
  return map[s] || 'info'
}

function statusText(s: string) {
  const map: Record<string, string> = { pending: '待审核', reviewing: '审核中', approved: '已通过', rejected: '已拒绝' }
  return map[s] || s
}

function viewDetail(row: any) {
  currentItem.value = row
  detailVisible.value = true
}

function handleApprove(row: any) {
  row.status = 'approved'
  ElMessage.success(`${row.name} 审核通过`)
  detailVisible.value = false
}

function handleReject(row: any) {
  currentItem.value = row
  rejectVisible.value = true
}

async function confirmReject() {
  await rejectFormRef.value?.validate()
  if (currentItem.value) {
    currentItem.value.status = 'rejected'
    ElMessage.success('已拒绝')
  }
  rejectVisible.value = false
  detailVisible.value = false
  rejectForm.reason = ''
}

function handleTransfer(row: any) {
  ElMessage.info(`转交 ${row.name} 的审核任务`)
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
.tab-badge { margin-left: 4px; }
:deep(.el-tabs__header) { margin-bottom: 16px; }
</style>
