<template>
  <div class="space-page">
    <div class="page-header">
      <h2 class="page-title">个人空间</h2>
      <el-button type="primary" @click="showUpload = true">
        <el-icon><Upload /></el-icon> 上传软件
      </el-button>
    </div>

    <!-- Stats Cards -->
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
          <el-icon :size="24" color="#fff"><Box /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ stats.softwareCount }}</div>
          <div class="stat-label">已上传软件</div>
          <div class="stat-detail">{{ stats.publishedCount }} 个已上架 · {{ stats.reviewingCount }} 个审核中 · {{ stats.privateCount }} 个私有</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%)">
          <el-icon :size="24" color="#fff"><Download /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ formatCount(stats.totalDownloads) }}</div>
          <div class="stat-label">总下载量</div>
          <div class="stat-detail">本周增量 +{{ formatCount(stats.weeklyDownloads) }}</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
          <el-icon :size="24" color="#fff"><Coin /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ stats.usedSpace }} / {{ stats.totalSpace }}</div>
          <div class="stat-label">空间容量</div>
          <el-progress :percentage="stats.usedPercent" :stroke-width="6" :show-text="false" style="margin-top: 6px" />
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <el-tabs v-model="activeTab" class="space-tabs">
      <!-- My Software -->
      <el-tab-pane label="我的软件" name="software">
        <el-table :data="mySoftware" stripe style="width: 100%">
          <el-table-column label="软件" min-width="220">
            <template #default="{ row }">
              <div class="software-cell">
                <div class="software-cell-icon">
                  <el-icon :size="20" color="#409eff"><Box /></el-icon>
                </div>
                <span class="software-cell-name">{{ row.name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="version" label="版本" width="100" />
          <el-table-column prop="size" label="大小" width="100" />
          <el-table-column label="可见性" width="100">
            <template #default="{ row }">
              <el-tag size="small" :type="row.visibility === 'public' ? 'success' : row.visibility === 'private' ? 'info' : 'warning'">
                {{ row.visibility === 'public' ? '公开' : row.visibility === 'private' ? '私有' : '链接' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="审核状态" width="100">
            <template #default="{ row }">
              <el-tag size="small" :type="getAuditType(row.status)">
                {{ getAuditText(row.status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="下载量" width="100">
            <template #default="{ row }">
              {{ formatCount(row.downloadCount) }}
            </template>
          </el-table-column>
          <el-table-column prop="createdAt" label="上传时间" width="160" />
          <el-table-column label="操作" width="160" fixed="right">
            <template #default="{ row }">
              <el-button size="small" text type="primary">编辑</el-button>
              <el-button size="small" text type="danger" @click="handleDeleteSoftware(row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!-- Upload Tasks -->
      <el-tab-pane label="上传任务" name="uploads">
        <div v-if="uploadTasks.length === 0" class="empty-state">
          <el-empty description="暂无上传任务" />
        </div>
        <div v-else class="task-list">
          <div v-for="task in uploadTasks" :key="task.id" class="task-item">
            <div class="task-icon">
              <el-icon :size="24" :color="task.status === 'completed' ? '#67c23a' : task.status === 'failed' ? '#f56c6c' : '#409eff'">
                <component :is="task.status === 'completed' ? 'CircleCheckFilled' : task.status === 'failed' ? 'CircleCloseFilled' : 'UploadFilled'" />
              </el-icon>
            </div>
            <div class="task-info">
              <div class="task-name">{{ task.name }} v{{ task.version }}</div>
              <el-progress v-if="task.status !== 'completed'" :percentage="task.progress" :stroke-width="4" :show-text="false" />
              <div class="task-meta">
                <span>{{ task.size }}</span>
                <span>{{ task.status === 'uploading' ? '上传中' : task.status === 'paused' ? '已暂停' : task.status === 'completed' ? '已完成' : '上传失败' }}</span>
              </div>
            </div>
            <div class="task-actions">
              <el-button v-if="task.status === 'uploading'" size="small" circle><el-icon><VideoPause /></el-icon></el-button>
              <el-button v-if="task.status === 'paused'" size="small" circle type="primary"><el-icon><VideoPlay /></el-icon></el-button>
              <el-button v-if="task.status === 'failed'" size="small" type="primary">重试</el-button>
              <el-button size="small" circle><el-icon><Close /></el-icon></el-button>
            </div>
          </div>
        </div>
      </el-tab-pane>

      <!-- Download Records -->
      <el-tab-pane label="下载记录" name="downloads">
        <div v-if="downloadRecords.length === 0" class="empty-state">
          <el-empty description="暂无下载记录" />
        </div>
        <div v-else class="task-list">
          <div v-for="record in downloadRecords" :key="record.id" class="task-item">
            <div class="task-icon">
              <el-icon :size="24" :color="record.status === 'completed' ? '#67c23a' : record.status === 'failed' ? '#f56c6c' : '#409eff'">
                <component :is="record.status === 'completed' ? 'CircleCheckFilled' : record.status === 'failed' ? 'CircleCloseFilled' : 'Loading'" />
              </el-icon>
            </div>
            <div class="task-info">
              <div class="task-name">{{ record.name }}</div>
              <el-progress v-if="record.status === 'downloading'" :percentage="record.progress" :stroke-width="4" :show-text="false" />
              <div class="task-meta">
                <span>{{ record.size }}</span>
                <span>{{ record.status === 'downloading' ? `${record.progress}%` : record.status === 'completed' ? '下载完成' : '下载失败' }}</span>
              </div>
            </div>
            <div class="task-actions">
              <el-button v-if="record.status === 'failed'" size="small" type="primary">重试</el-button>
              <el-button v-if="record.status === 'completed'" size="small">打开</el-button>
              <el-button size="small" circle><el-icon><Close /></el-icon></el-button>
            </div>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>

    <!-- Upload Dialog -->
    <UploadDialog v-model="showUpload" @uploaded="handleUploaded" />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import UploadDialog from '@/components/upload/UploadDialog.vue'

const activeTab = ref('software')
const showUpload = ref(false)

const stats = reactive({
  softwareCount: 8,
  publishedCount: 5,
  reviewingCount: 2,
  privateCount: 1,
  totalDownloads: 45800,
  weeklyDownloads: 1250,
  usedSpace: '2.4 GB',
  totalSpace: '10 GB',
  usedPercent: 24
})

const mySoftware = ref([
  { id: 1, name: 'MyTool Pro', icon: '', version: '2.1.0', size: '45 MB', visibility: 'public', status: 'published', downloadCount: 12500, createdAt: '2024-02-15 10:30' },
  { id: 2, name: 'DataHelper', icon: '', version: '1.3.2', size: '28 MB', visibility: 'public', status: 'published', downloadCount: 8900, createdAt: '2024-01-20 14:15' },
  { id: 3, name: 'QuickNote', icon: '', version: '1.0.0', size: '12 MB', visibility: 'public', status: 'reviewing', downloadCount: 0, createdAt: '2024-03-05 09:00' },
  { id: 4, name: 'FileSync', icon: '', version: '0.9.1', size: '35 MB', visibility: 'private', status: 'draft', downloadCount: 0, createdAt: '2024-03-08 16:45' },
  { id: 5, name: 'ScreenCapture', icon: '', version: '3.0.0', size: '18 MB', visibility: 'public', status: 'published', downloadCount: 15600, createdAt: '2023-11-10 08:20' }
])

const uploadTasks = ref([
  { id: 1, name: 'QuickNote', version: '1.1.0', size: '14 MB', progress: 65, status: 'uploading' },
  { id: 2, name: 'FileSync', version: '1.0.0', size: '38 MB', progress: 100, status: 'completed' }
])

const downloadRecords = ref([
  { id: 1, name: 'Visual Studio Code', size: '98.5 MB', progress: 100, status: 'completed' },
  { id: 2, name: 'Docker Desktop', size: '580 MB', progress: 35, status: 'downloading' },
  { id: 3, name: 'Figma', size: '120 MB', progress: 0, status: 'failed' }
])

function formatCount(count: number): string {
  if (count >= 1000000) return (count / 1000000).toFixed(1) + 'M'
  if (count >= 1000) return (count / 1000).toFixed(1) + 'K'
  return count.toString()
}

function getAuditType(status: string) {
  switch (status) {
    case 'published': return 'success'
    case 'reviewing': return 'warning'
    case 'rejected': return 'danger'
    case 'draft': return 'info'
    default: return 'info'
  }
}

function getAuditText(status: string) {
  switch (status) {
    case 'published': return '已上架'
    case 'reviewing': return '审核中'
    case 'rejected': return '已拒绝'
    case 'draft': return '草稿'
    default: return status
  }
}

function handleDeleteSoftware(row: any) {
  ElMessageBox.confirm(`确定要删除 ${row.name} 吗？此操作不可恢复。`, '确认删除', {
    type: 'warning'
  }).then(() => {
    const idx = mySoftware.value.findIndex(s => s.id === row.id)
    if (idx !== -1) mySoftware.value.splice(idx, 1)
    ElMessage.success('已删除')
  }).catch(() => {})
}

function handleUploaded() {
  ElMessage.success('上传完成')
}
</script>

<style scoped>
.space-page {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.page-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-bottom: 28px;
}

.stat-card {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 24px;
  display: flex;
  gap: 16px;
  border: 1px solid var(--border-color-light);
}

.stat-icon {
  width: 52px;
  height: 52px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-color-primary);
  line-height: 1.2;
}

.stat-label {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin: 4px 0;
}

.stat-detail {
  font-size: 12px;
  color: var(--text-color-placeholder);
}

.space-tabs :deep(.el-tabs__content) {
  padding-top: 8px;
}

.software-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.software-cell-icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.software-cell-name {
  font-weight: 500;
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.task-item {
  display: flex;
  align-items: center;
  gap: 14px;
  background: var(--bg-color-card);
  padding: 16px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color-light);
}

.task-icon {
  flex-shrink: 0;
}

.task-info {
  flex: 1;
  min-width: 0;
}

.task-name {
  font-weight: 500;
  font-size: 14px;
  margin-bottom: 4px;
}

.task-meta {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: var(--text-color-secondary);
  margin-top: 4px;
}

.task-actions {
  display: flex;
  gap: 6px;
  flex-shrink: 0;
}

.empty-state {
  padding: 60px 0;
}
</style>
