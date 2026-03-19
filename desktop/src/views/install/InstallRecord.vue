<template>
  <div class="install-page">
    <div class="page-header">
      <h2 class="page-title">安装记录</h2>
    </div>

    <!-- Filter bar -->
    <div class="filter-bar">
      <el-input
        v-model="searchKeyword"
        placeholder="搜索已安装软件..."
        :prefix-icon="Search"
        clearable
        style="width: 260px"
      />
      <el-select v-model="sortBy" placeholder="排序方式" style="width: 140px">
        <el-option label="安装时间" value="time" />
        <el-option label="名称" value="name" />
        <el-option label="大小" value="size" />
      </el-select>
      <el-select v-model="statusFilter" placeholder="状态筛选" style="width: 120px">
        <el-option label="全部" value="all" />
        <el-option label="已安装" value="installed" />
        <el-option label="可更新" value="updatable" />
        <el-option label="已卸载" value="uninstalled" />
      </el-select>
    </div>

    <!-- Table -->
    <el-table :data="filteredRecords" stripe style="width: 100%" class="install-table">
      <el-table-column label="软件名称" min-width="240">
        <template #default="{ row }">
          <div class="software-cell">
            <div class="software-cell-icon">
              <el-icon :size="24" color="#409eff"><Box /></el-icon>
            </div>
            <div>
              <div class="software-cell-name">{{ row.softwareName }}</div>
              <div class="software-cell-developer">{{ row.developer }}</div>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="版本" width="120">
        <template #default="{ row }">
          <span>{{ row.version }}</span>
          <el-tag v-if="row.status === 'updatable'" size="small" type="warning" class="update-tag">
            {{ row.latestVersion }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="size" label="大小" width="100" />
      <el-table-column prop="installPath" label="安装路径" min-width="200" show-overflow-tooltip />
      <el-table-column prop="installedAt" label="安装时间" width="160" />
      <el-table-column label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)" size="small">
            {{ getStatusText(row.status) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <template v-if="row.status === 'installed'">
            <el-button size="small" text type="primary" @click="handleOpen(row)">打开</el-button>
            <el-button size="small" text type="danger" @click="handleUninstall(row)">卸载</el-button>
          </template>
          <template v-else-if="row.status === 'updatable'">
            <el-button size="small" text type="primary" @click="handleUpdate(row)">更新</el-button>
            <el-button size="small" text type="danger" @click="handleUninstall(row)">卸载</el-button>
          </template>
          <template v-else>
            <el-button size="small" text type="primary" @click="handleReinstall(row)">重新安装</el-button>
            <el-button size="small" text type="info" @click="handleRemoveRecord(row)">删除记录</el-button>
          </template>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

interface InstallItem {
  id: number
  softwareId: number
  softwareName: string
  softwareIcon: string
  developer: string
  version: string
  size: string
  installPath: string
  installedAt: string
  status: 'installed' | 'updatable' | 'uninstalled'
  latestVersion?: string
}

const searchKeyword = ref('')
const sortBy = ref('time')
const statusFilter = ref('all')

const records = ref<InstallItem[]>([
  { id: 1, softwareId: 1, softwareName: 'Visual Studio Code', softwareIcon: '', developer: 'Microsoft', version: '1.86.0', size: '340 MB', installPath: 'C:\\Program Files\\VS Code', installedAt: '2024-03-01 10:30', status: 'updatable', latestVersion: 'v1.87.0' },
  { id: 2, softwareId: 2, softwareName: 'Node.js', softwareIcon: '', developer: 'OpenJS Foundation', version: '20.11.1', size: '75 MB', installPath: 'C:\\Program Files\\nodejs', installedAt: '2024-02-15 14:20', status: 'installed' },
  { id: 3, softwareId: 3, softwareName: 'Git', softwareIcon: '', developer: 'Git Community', version: '2.44.0', size: '320 MB', installPath: 'C:\\Program Files\\Git', installedAt: '2024-02-10 09:45', status: 'installed' },
  { id: 4, softwareId: 4, softwareName: 'Python', softwareIcon: '', developer: 'Python Software Foundation', version: '3.12.2', size: '105 MB', installPath: 'C:\\Python312', installedAt: '2024-01-20 16:00', status: 'installed' },
  { id: 5, softwareId: 5, softwareName: 'Docker Desktop', softwareIcon: '', developer: 'Docker Inc.', version: '4.27.0', size: '1.2 GB', installPath: 'C:\\Program Files\\Docker', installedAt: '2024-01-05 11:30', status: 'updatable', latestVersion: 'v4.28.0' },
  { id: 6, softwareId: 7, softwareName: 'Notion', softwareIcon: '', developer: 'Notion Labs', version: '2.0.42', size: '180 MB', installPath: 'C:\\Users\\AppData\\Local\\Notion', installedAt: '2023-12-20 08:15', status: 'uninstalled' },
  { id: 7, softwareId: 9, softwareName: 'VLC Player', softwareIcon: '', developer: 'VideoLAN', version: '3.0.20', size: '120 MB', installPath: 'C:\\Program Files\\VideoLAN', installedAt: '2023-11-10 15:40', status: 'installed' },
  { id: 8, softwareId: 11, softwareName: '7-Zip', softwareIcon: '', developer: 'Igor Pavlov', version: '23.01', size: '5 MB', installPath: 'C:\\Program Files\\7-Zip', installedAt: '2023-10-25 12:00', status: 'installed' }
])

const filteredRecords = computed(() => {
  let list = [...records.value]

  if (statusFilter.value !== 'all') {
    list = list.filter(r => r.status === statusFilter.value)
  }

  if (searchKeyword.value.trim()) {
    const kw = searchKeyword.value.toLowerCase()
    list = list.filter(r => r.softwareName.toLowerCase().includes(kw))
  }

  switch (sortBy.value) {
    case 'name': list.sort((a, b) => a.softwareName.localeCompare(b.softwareName)); break
    case 'time': list.sort((a, b) => b.installedAt.localeCompare(a.installedAt)); break
  }

  return list
})

function getStatusType(status: string) {
  switch (status) {
    case 'installed': return 'success'
    case 'updatable': return 'warning'
    case 'uninstalled': return 'info'
    default: return 'info'
  }
}

function getStatusText(status: string) {
  switch (status) {
    case 'installed': return '已安装'
    case 'updatable': return '可更新'
    case 'uninstalled': return '已卸载'
    default: return status
  }
}

function handleOpen(row: InstallItem) {
  ElMessage.info(`打开 ${row.softwareName}`)
}

function handleUninstall(row: InstallItem) {
  ElMessageBox.confirm(`确定要卸载 ${row.softwareName} 吗？`, '确认卸载', {
    type: 'warning'
  }).then(() => {
    row.status = 'uninstalled'
    ElMessage.success(`已卸载 ${row.softwareName}`)
  }).catch(() => {})
}

function handleUpdate(row: InstallItem) {
  ElMessage.success(`开始更新 ${row.softwareName}`)
  row.status = 'installed'
  row.version = row.latestVersion?.replace('v', '') || row.version
}

function handleReinstall(row: InstallItem) {
  ElMessage.success(`开始重新安装 ${row.softwareName}`)
}

function handleRemoveRecord(row: InstallItem) {
  const idx = records.value.findIndex(r => r.id === row.id)
  if (idx !== -1) records.value.splice(idx, 1)
  ElMessage.success('已删除记录')
}
</script>

<style scoped>
.install-page {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 20px;
}

.page-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.filter-bar {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
  background: var(--bg-color-card);
  padding: 16px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color-light);
}

.install-table {
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.software-cell {
  display: flex;
  align-items: center;
  gap: 10px;
}

.software-cell-icon {
  width: 40px;
  height: 40px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.software-cell-name {
  font-weight: 600;
  color: var(--text-color-primary);
  font-size: 14px;
}

.software-cell-developer {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.update-tag {
  margin-left: 4px;
}
</style>
