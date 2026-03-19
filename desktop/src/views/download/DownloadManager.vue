<template>
  <div class="download-page">
    <div class="page-header">
      <h2 class="page-title">下载管理</h2>
      <div class="header-actions">
        <el-button @click="downloadStore.pauseAll()" :disabled="downloadStore.downloadingTasks.filter(t => t.status === 'downloading').length === 0">
          <el-icon><VideoPause /></el-icon> 全部暂停
        </el-button>
        <el-button @click="downloadStore.resumeAll()" :disabled="downloadStore.downloadingTasks.filter(t => t.status === 'paused').length === 0">
          <el-icon><VideoPlay /></el-icon> 全部继续
        </el-button>
        <el-button @click="downloadStore.clearCompleted()" :disabled="downloadStore.completedCount === 0">
          <el-icon><Delete /></el-icon> 清空已完成
        </el-button>
      </div>
    </div>

    <el-tabs v-model="activeTab" class="download-tabs">
      <el-tab-pane name="downloading">
        <template #label>
          下载中
          <el-badge :value="downloadStore.downloadingCount" :hidden="downloadStore.downloadingCount === 0" class="tab-badge" />
        </template>
        <div v-if="downloadStore.downloadingTasks.length === 0" class="empty-state">
          <el-empty description="暂无下载任务" />
        </div>
        <div v-else class="task-list">
          <div v-for="task in downloadStore.downloadingTasks" :key="task.id" class="task-item">
            <div class="task-icon">
              <el-icon :size="28" color="#409eff"><Box /></el-icon>
            </div>
            <div class="task-info">
              <div class="task-name">
                <span class="name-text">{{ task.softwareName }}</span>
                <span class="version-text">v{{ task.version }}</span>
              </div>
              <el-progress :percentage="task.progress" :stroke-width="6" :show-text="false" class="task-progress" />
              <div class="task-meta">
                <span>{{ task.progress }}% &middot; {{ formatBytes(task.downloadedBytes) }} / {{ task.fileSize }}</span>
                <span v-if="task.status === 'downloading'">{{ task.speed }} &middot; 剩余 {{ task.remainingTime }}</span>
                <span v-else class="paused-text">已暂停</span>
              </div>
            </div>
            <div class="task-actions">
              <el-button
                v-if="task.status === 'downloading'"
                circle
                @click="downloadStore.pauseTask(task.id)"
                title="暂停"
              >
                <el-icon><VideoPause /></el-icon>
              </el-button>
              <el-button
                v-if="task.status === 'paused'"
                circle
                type="primary"
                @click="downloadStore.resumeTask(task.id)"
                title="继续"
              >
                <el-icon><VideoPlay /></el-icon>
              </el-button>
              <el-button circle @click="downloadStore.cancelTask(task.id)" title="取消">
                <el-icon><Close /></el-icon>
              </el-button>
            </div>
          </div>
        </div>
      </el-tab-pane>

      <el-tab-pane name="completed">
        <template #label>
          已完成
          <el-badge :value="downloadStore.completedCount" :hidden="downloadStore.completedCount === 0" class="tab-badge" />
        </template>
        <div v-if="downloadStore.completedTasks.length === 0" class="empty-state">
          <el-empty description="暂无已完成任务" />
        </div>
        <div v-else class="task-list">
          <div v-for="task in downloadStore.completedTasks" :key="task.id" class="task-item">
            <div class="task-icon completed-icon">
              <el-icon :size="28" color="#67c23a"><CircleCheckFilled /></el-icon>
            </div>
            <div class="task-info">
              <div class="task-name">
                <span class="name-text">{{ task.softwareName }}</span>
                <span class="version-text">v{{ task.version }}</span>
              </div>
              <div class="task-meta">
                <span>{{ task.fileSize }} &middot; {{ task.completedAt }}</span>
              </div>
            </div>
            <div class="task-actions">
              <el-button size="small" @click="handleOpenFile(task)">打开文件</el-button>
              <el-button size="small" @click="handleOpenFolder(task)">打开文件夹</el-button>
              <el-button size="small" type="primary" @click="handleInstall(task)">一键安装</el-button>
              <el-button size="small" type="danger" text @click="downloadStore.removeTask(task.id)">删除记录</el-button>
            </div>
          </div>
        </div>
      </el-tab-pane>

      <el-tab-pane name="failed">
        <template #label>
          已失败
          <el-badge :value="downloadStore.failedCount" :hidden="downloadStore.failedCount === 0" class="tab-badge" type="danger" />
        </template>
        <div v-if="downloadStore.failedTasks.length === 0" class="empty-state">
          <el-empty description="暂无失败任务" />
        </div>
        <div v-else class="task-list">
          <div v-for="task in downloadStore.failedTasks" :key="task.id" class="task-item">
            <div class="task-icon failed-icon">
              <el-icon :size="28" color="#f56c6c"><CircleCloseFilled /></el-icon>
            </div>
            <div class="task-info">
              <div class="task-name">
                <span class="name-text">{{ task.softwareName }}</span>
                <span class="version-text">v{{ task.version }}</span>
              </div>
              <div class="task-meta">
                <span class="failed-text">下载失败 &middot; {{ formatBytes(task.downloadedBytes) }} / {{ task.fileSize }}</span>
              </div>
            </div>
            <div class="task-actions">
              <el-button size="small" type="primary" @click="downloadStore.retryTask(task.id)">
                <el-icon><RefreshRight /></el-icon> 重试
              </el-button>
              <el-button size="small" type="danger" text @click="downloadStore.removeTask(task.id)">删除</el-button>
            </div>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useDownloadStore } from '@/stores/download'
import { ElMessage } from 'element-plus'
import type { DownloadItem } from '@/stores/download'

const downloadStore = useDownloadStore()
const activeTab = ref('downloading')

function formatBytes(bytes: number): string {
  if (bytes >= 1073741824) return (bytes / 1073741824).toFixed(1) + ' GB'
  if (bytes >= 1048576) return (bytes / 1048576).toFixed(1) + ' MB'
  if (bytes >= 1024) return (bytes / 1024).toFixed(1) + ' KB'
  return bytes + ' B'
}

function handleOpenFile(task: DownloadItem) {
  ElMessage.info(`打开文件: ${task.filePath}`)
}

function handleOpenFolder(task: DownloadItem) {
  ElMessage.info(`打开文件夹: ${task.filePath}`)
}

function handleInstall(task: DownloadItem) {
  ElMessage.success(`开始安装 ${task.softwareName}`)
}
</script>

<style scoped>
.download-page {
  max-width: 1000px;
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

.header-actions {
  display: flex;
  gap: 8px;
}

.download-tabs :deep(.el-tabs__content) {
  padding-top: 8px;
}

.tab-badge {
  margin-left: 6px;
}

.tab-badge :deep(.el-badge__content) {
  font-size: 10px;
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.task-item {
  display: flex;
  align-items: center;
  gap: 16px;
  background: var(--bg-color-card);
  padding: 20px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color-light);
  transition: box-shadow 0.2s;
}

.task-item:hover {
  box-shadow: var(--shadow-sm);
}

.task-icon {
  width: 48px;
  height: 48px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.completed-icon {
  background: linear-gradient(135deg, #f0f9eb 0%, #e1f3d8 100%);
}

.failed-icon {
  background: linear-gradient(135deg, #fef0f0 0%, #fde2e2 100%);
}

.task-info {
  flex: 1;
  min-width: 0;
}

.task-name {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.name-text {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.version-text {
  font-size: 12px;
  color: var(--text-color-secondary);
  background: var(--bg-color);
  padding: 1px 6px;
  border-radius: var(--radius-sm);
}

.task-progress {
  margin-bottom: 6px;
}

.task-meta {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: var(--text-color-secondary);
}

.paused-text {
  color: var(--color-warning);
}

.failed-text {
  color: var(--color-danger);
}

.task-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.empty-state {
  padding: 60px 0;
}
</style>
