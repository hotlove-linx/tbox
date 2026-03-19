<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="filter-bar">
        <el-input v-model="searchText" placeholder="搜索反馈标题/内容" prefix-icon="Search" clearable style="width: 220px" />
        <el-select v-model="typeFilter" placeholder="反馈类型" clearable style="width: 130px">
          <el-option label="Bug反馈" value="bug" />
          <el-option label="功能建议" value="feature" />
          <el-option label="使用咨询" value="question" />
          <el-option label="其他" value="other" />
        </el-select>
        <el-select v-model="statusFilter" placeholder="处理状态" clearable style="width: 120px">
          <el-option label="待处理" value="pending" />
          <el-option label="处理中" value="processing" />
          <el-option label="已回复" value="replied" />
          <el-option label="已关闭" value="closed" />
        </el-select>
        <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="反馈类型" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="typeTagMap[row.type]" size="small">{{ typeTextMap[row.type] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题" width="180" show-overflow-tooltip />
        <el-table-column prop="content" label="内容摘要" min-width="200" show-overflow-tooltip />
        <el-table-column prop="user" label="提交用户" width="100" />
        <el-table-column prop="submitTime" label="提交时间" width="170" />
        <el-table-column label="处理状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="statusTagMap[row.status]" size="small">{{ statusTextMap[row.status] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="handler" label="处理人" width="100" />
        <el-table-column label="操作" width="220">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="viewFeedback(row)">查看</el-button>
            <el-button link type="success" size="small" @click="replyFeedback(row)">回复</el-button>
            <el-button link type="warning" size="small" @click="assignFeedback(row)">分配</el-button>
            <el-button v-if="row.status !== 'closed'" link type="info" size="small" @click="closeFeedback(row)">关闭</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" :page-size="20" :total="total" layout="total, prev, pager, next" background />
      </div>
    </el-card>

    <!-- Detail / Reply Dialog -->
    <el-dialog v-model="dialogVisible" :title="dialogMode === 'view' ? '反馈详情' : '回复反馈'" width="600px">
      <template v-if="currentFeedback">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="类型"><el-tag :type="typeTagMap[currentFeedback.type]" size="small">{{ typeTextMap[currentFeedback.type] }}</el-tag></el-descriptions-item>
          <el-descriptions-item label="提交用户">{{ currentFeedback.user }}</el-descriptions-item>
          <el-descriptions-item label="标题" :span="2">{{ currentFeedback.title }}</el-descriptions-item>
          <el-descriptions-item label="内容" :span="2">{{ currentFeedback.content }}</el-descriptions-item>
          <el-descriptions-item label="提交时间">{{ currentFeedback.submitTime }}</el-descriptions-item>
          <el-descriptions-item label="状态"><el-tag :type="statusTagMap[currentFeedback.status]" size="small">{{ statusTextMap[currentFeedback.status] }}</el-tag></el-descriptions-item>
        </el-descriptions>

        <div v-if="dialogMode === 'reply'" style="margin-top: 20px;">
          <el-form :model="replyForm" :rules="replyRules" ref="replyFormRef" label-width="80px">
            <el-form-item label="回复内容" prop="content">
              <el-input v-model="replyForm.content" type="textarea" :rows="4" placeholder="请输入回复内容" />
            </el-form-item>
          </el-form>
        </div>
      </template>
      <template #footer>
        <el-button @click="dialogVisible = false">关闭</el-button>
        <el-button v-if="dialogMode === 'reply'" type="primary" @click="submitReply">发送回复</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const searchText = ref('')
const typeFilter = ref('')
const statusFilter = ref('')
const currentPage = ref(1)
const total = ref(56)
const dialogVisible = ref(false)
const dialogMode = ref<'view' | 'reply'>('view')
const currentFeedback = ref<any>(null)
const replyFormRef = ref<FormInstance>()
const replyForm = reactive({ content: '' })
const replyRules: FormRules = { content: [{ required: true, message: '请输入回复内容', trigger: 'blur' }] }

const typeTagMap: Record<string, any> = { bug: 'danger', feature: 'primary', question: 'warning', other: 'info' }
const typeTextMap: Record<string, string> = { bug: 'Bug反馈', feature: '功能建议', question: '使用咨询', other: '其他' }
const statusTagMap: Record<string, any> = { pending: 'warning', processing: 'primary', replied: 'success', closed: 'info' }
const statusTextMap: Record<string, string> = { pending: '待处理', processing: '处理中', replied: '已回复', closed: '已关闭' }

const tableData = ref([
  { id: 1, type: 'bug', title: '下载功能异常', content: '在Wi-Fi环境下下载经常中断，需要重新开始', user: '张三', submitTime: '2024-03-15 10:00:00', status: 'pending', handler: '' },
  { id: 2, type: 'feature', title: '建议增加深色模式', content: '希望应用能支持深色模式，晚上使用更舒适', user: '李四', submitTime: '2024-03-14 18:30:00', status: 'processing', handler: '管理员A' },
  { id: 3, type: 'question', title: '如何批量下载', content: '请问有没有批量下载的功能？在哪里可以找到？', user: '王五', submitTime: '2024-03-14 14:00:00', status: 'replied', handler: '管理员B' },
  { id: 4, type: 'bug', title: '搜索结果不准确', content: '搜索时输入关键词，结果不相关的软件也会出现', user: '赵六', submitTime: '2024-03-13 09:00:00', status: 'closed', handler: '管理员A' },
  { id: 5, type: 'feature', title: '支持分享到社交平台', content: '希望可以直接分享软件到微信、QQ等社交平台', user: '孙七', submitTime: '2024-03-12 11:20:00', status: 'pending', handler: '' },
  { id: 6, type: 'other', title: '软件推荐', content: '推荐一些好用的办公软件', user: '周八', submitTime: '2024-03-11 16:00:00', status: 'replied', handler: '管理员C' }
])

function handleSearch() { ElMessage.info('搜索功能已触发') }

function viewFeedback(row: any) {
  currentFeedback.value = row
  dialogMode.value = 'view'
  dialogVisible.value = true
}

function replyFeedback(row: any) {
  currentFeedback.value = row
  dialogMode.value = 'reply'
  replyForm.content = ''
  dialogVisible.value = true
}

function assignFeedback(row: any) {
  ElMessage.info(`分配反馈 #${row.id}`)
}

function closeFeedback(row: any) {
  row.status = 'closed'
  ElMessage.success('已关闭')
}

async function submitReply() {
  await replyFormRef.value?.validate()
  if (currentFeedback.value) {
    currentFeedback.value.status = 'replied'
    currentFeedback.value.handler = '当前管理员'
  }
  dialogVisible.value = false
  ElMessage.success('回复成功')
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; flex-wrap: wrap; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
</style>
