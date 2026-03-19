<template>
  <div class="page-container">
    <!-- Search & Filter -->
    <el-card shadow="never" class="filter-card">
      <el-form :inline="true" :model="filters">
        <el-form-item>
          <el-input v-model="filters.keyword" placeholder="搜索软件名称/ID" prefix-icon="Search" clearable style="width: 220px" />
        </el-form-item>
        <el-form-item>
          <el-select v-model="filters.category" placeholder="选择分类" clearable style="width: 150px">
            <el-option v-for="c in categories" :key="c.value" :label="c.label" :value="c.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="filters.auditStatus" placeholder="审核状态" clearable style="width: 130px">
            <el-option label="待审核" value="pending" />
            <el-option label="已通过" value="approved" />
            <el-option label="已拒绝" value="rejected" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="filters.visibility" placeholder="可见性" clearable style="width: 120px">
            <el-option label="已上架" value="visible" />
            <el-option label="已下架" value="hidden" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-date-picker v-model="filters.dateRange" type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" style="width: 260px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
          <el-button @click="resetFilters"><el-icon><Refresh /></el-icon>重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Actions Bar -->
    <div class="action-bar">
      <div class="left">
        <el-button type="success" @click="handleExport"><el-icon><Download /></el-icon>导出Excel</el-button>
        <el-dropdown @command="handleBatch" v-show="selectedIds.length > 0">
          <el-button type="primary">
            批量操作 ({{ selectedIds.length }})<el-icon class="el-icon--right"><ArrowDown /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="publish">批量上架</el-dropdown-item>
              <el-dropdown-item command="unpublish">批量下架</el-dropdown-item>
              <el-dropdown-item command="delete" divided>批量删除</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
      <div class="right">
        <span class="total-text">共 {{ total }} 条</span>
      </div>
    </div>

    <!-- Table -->
    <el-card shadow="never">
      <el-table :data="tableData" @selection-change="handleSelectionChange" stripe border style="width: 100%">
        <el-table-column type="selection" width="45" />
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="软件信息" min-width="240">
          <template #default="{ row }">
            <div class="software-info">
              <el-avatar :size="40" shape="square" :src="row.icon">{{ row.name[0] }}</el-avatar>
              <div>
                <div class="sw-name">{{ row.name }}</div>
                <div class="sw-version">v{{ row.version }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="100" />
        <el-table-column prop="developer" label="开发者/上传者" width="120" />
        <el-table-column label="可见性" width="90" align="center">
          <template #default="{ row }">
            <el-tag :type="row.visible ? 'success' : 'info'" size="small">
              {{ row.visible ? '已上架' : '已下架' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="审核状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="auditStatusType(row.auditStatus)" size="small">
              {{ auditStatusText(row.auditStatus) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="downloads" label="下载量" width="90" align="right" sortable />
        <el-table-column label="评分" width="120" align="center">
          <template #default="{ row }">
            <div class="rating-cell">
              <el-rate v-model="row.rating" disabled :max="5" size="small" />
              <span class="rating-value">{{ row.rating.toFixed(1) }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="fileSize" label="文件大小" width="90" align="right" />
        <el-table-column prop="publishTime" label="上架时间" width="110" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="$router.push(`/software/detail/${row.id}`)">查看</el-button>
            <el-button link type="primary" size="small" @click="$router.push(`/software/detail/${row.id}`)">编辑</el-button>
            <el-button link :type="row.visible ? 'warning' : 'success'" size="small" @click="toggleVisibility(row)">
              {{ row.visible ? '下架' : '上架' }}
            </el-button>
            <el-popconfirm title="确定删除？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          background
        />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const categories = [
  { label: '工具软件', value: 'tools' },
  { label: '游戏娱乐', value: 'games' },
  { label: '效率办公', value: 'productivity' },
  { label: '社交通讯', value: 'social' },
  { label: '影音播放', value: 'media' },
  { label: '教育学习', value: 'education' }
]

const filters = reactive({
  keyword: '',
  category: '',
  auditStatus: '',
  visibility: '',
  dateRange: null as any
})

const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(286)
const selectedIds = ref<number[]>([])

const mockSoftware = [
  { id: 1001, name: '超级文件管理器', version: '3.2.1', icon: '', category: '工具软件', developer: '张三', visible: true, auditStatus: 'approved', downloads: 52340, rating: 4.5, fileSize: '28.5MB', publishTime: '2024-03-15' },
  { id: 1002, name: '极速浏览器', version: '5.0.0', icon: '', category: '工具软件', developer: '李四', visible: true, auditStatus: 'approved', downloads: 48210, rating: 4.3, fileSize: '45.2MB', publishTime: '2024-03-14' },
  { id: 1003, name: '智能笔记', version: '2.1.0', icon: '', category: '效率办公', developer: '王五', visible: false, auditStatus: 'pending', downloads: 12560, rating: 4.7, fileSize: '15.8MB', publishTime: '2024-03-13' },
  { id: 1004, name: '音乐播放器Pro', version: '4.5.2', icon: '', category: '影音播放', developer: '赵六', visible: true, auditStatus: 'approved', downloads: 38920, rating: 4.2, fileSize: '32.1MB', publishTime: '2024-03-12' },
  { id: 1005, name: '社交聊天', version: '1.8.0', icon: '', category: '社交通讯', developer: '孙七', visible: true, auditStatus: 'approved', downloads: 65100, rating: 4.6, fileSize: '52.3MB', publishTime: '2024-03-11' },
  { id: 1006, name: '在线学习平台', version: '2.3.1', icon: '', category: '教育学习', developer: '周八', visible: true, auditStatus: 'rejected', downloads: 8900, rating: 3.8, fileSize: '18.7MB', publishTime: '2024-03-10' },
  { id: 1007, name: '天气预报', version: '3.0.0', icon: '', category: '工具软件', developer: '吴九', visible: true, auditStatus: 'approved', downloads: 28700, rating: 4.1, fileSize: '12.4MB', publishTime: '2024-03-09' },
  { id: 1008, name: '密码管理器', version: '1.5.3', icon: '', category: '工具软件', developer: '郑十', visible: false, auditStatus: 'pending', downloads: 5200, rating: 4.8, fileSize: '8.9MB', publishTime: '2024-03-08' },
  { id: 1009, name: '视频编辑器', version: '2.0.1', icon: '', category: '影音播放', developer: '张三', visible: true, auditStatus: 'approved', downloads: 31240, rating: 4.4, fileSize: '68.5MB', publishTime: '2024-03-07' },
  { id: 1010, name: '健康计步器', version: '1.2.0', icon: '', category: '工具软件', developer: '李四', visible: true, auditStatus: 'approved', downloads: 15600, rating: 4.0, fileSize: '9.2MB', publishTime: '2024-03-06' }
]

const tableData = ref(mockSoftware)

function auditStatusType(status: string) {
  const map: Record<string, string> = { approved: 'success', pending: 'warning', rejected: 'danger' }
  return (map[status] || 'info') as any
}

function auditStatusText(status: string) {
  const map: Record<string, string> = { approved: '已通过', pending: '待审核', rejected: '已拒绝' }
  return map[status] || status
}

function handleSelectionChange(rows: any[]) {
  selectedIds.value = rows.map((r) => r.id)
}

function handleSearch() {
  ElMessage.info('搜索功能已触发')
}

function resetFilters() {
  filters.keyword = ''
  filters.category = ''
  filters.auditStatus = ''
  filters.visibility = ''
  filters.dateRange = null
}

function handleExport() {
  ElMessage.success('导出任务已提交')
}

function handleBatch(command: string) {
  ElMessage.info(`批量${command}操作，选中 ${selectedIds.value.length} 项`)
}

function toggleVisibility(row: any) {
  row.visible = !row.visible
  ElMessage.success(row.visible ? '已上架' : '已下架')
}

function handleDelete(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.filter-card {
  margin-bottom: 16px;
}

.action-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.action-bar .left {
  display: flex;
  gap: 12px;
}

.total-text {
  font-size: 14px;
  color: #909399;
}

.software-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.sw-name {
  font-weight: 500;
  color: #303133;
}

.sw-version {
  font-size: 12px;
  color: #909399;
}

.rating-cell {
  display: flex;
  align-items: center;
  gap: 4px;
}

.rating-value {
  font-size: 12px;
  color: #e6a23c;
  font-weight: 500;
}

.pagination-wrap {
  display: flex;
  justify-content: flex-end;
  margin-top: 16px;
}
</style>
