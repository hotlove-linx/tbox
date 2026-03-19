<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="filter-bar">
        <el-input v-model="searchText" placeholder="搜索评论内容" prefix-icon="Search" clearable style="width: 200px" />
        <el-select v-model="statusFilter" placeholder="状态" clearable style="width: 120px">
          <el-option label="正常" value="normal" />
          <el-option label="已屏蔽" value="blocked" />
        </el-select>
        <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="评论内容" min-width="250">
          <template #default="{ row }">
            <el-tooltip :content="row.content" placement="top" :show-after="300">
              <span class="ellipsis-text">{{ row.content }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="评分" width="140" align="center">
          <template #default="{ row }">
            <el-rate v-model="row.rating" disabled size="small" />
          </template>
        </el-table-column>
        <el-table-column label="评论用户" width="120">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 6px;">
              <el-avatar :size="24">{{ row.user[0] }}</el-avatar>
              <span>{{ row.user }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="software" label="关联软件" width="150" />
        <el-table-column prop="publishTime" label="发布时间" width="170" />
        <el-table-column label="状态" width="90" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status === 'normal' ? 'success' : 'danger'" size="small">
              {{ row.status === 'normal' ? '正常' : '已屏蔽' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button v-if="row.status === 'normal'" link type="warning" size="small" @click="blockReview(row)">屏蔽</el-button>
            <el-button v-else link type="success" size="small" @click="restoreReview(row)">恢复</el-button>
            <el-popconfirm title="确定删除此评论？" @confirm="deleteReview(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" :page-size="20" :total="total" layout="total, prev, pager, next" background />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const searchText = ref('')
const statusFilter = ref('')
const currentPage = ref(1)
const total = ref(86)

const tableData = ref([
  { id: 1, content: '非常好用的文件管理器，推荐大家使用！功能强大，界面简洁。', rating: 5, user: '张三', software: '超级文件管理器', publishTime: '2024-03-15 10:30:00', status: 'normal' },
  { id: 2, content: '浏览器速度很快，但是有时候会崩溃。', rating: 3, user: '李四', software: '极速浏览器', publishTime: '2024-03-15 09:20:00', status: 'normal' },
  { id: 3, content: '垃圾软件，不要下载！全是广告！', rating: 1, user: '匿名用户', software: '天气预报', publishTime: '2024-03-14 18:45:00', status: 'blocked' },
  { id: 4, content: '笔记功能很实用，云同步也很方便。', rating: 4, user: '王五', software: '智能笔记', publishTime: '2024-03-14 15:00:00', status: 'normal' },
  { id: 5, content: '音质不错，歌曲库也很全面。', rating: 4, user: '赵六', software: '音乐播放器Pro', publishTime: '2024-03-14 12:30:00', status: 'normal' },
  { id: 6, content: '这是广告内容，请勿相信。联系方式：xxx', rating: 1, user: '广告号', software: '社交聊天', publishTime: '2024-03-13 20:00:00', status: 'blocked' },
  { id: 7, content: '学习效果很好，课程内容丰富。', rating: 5, user: '周八', software: '在线学习平台', publishTime: '2024-03-13 16:15:00', status: 'normal' },
  { id: 8, content: '视频编辑功能比较基础，希望能增加更多特效。', rating: 3, user: '孙七', software: '视频编辑器', publishTime: '2024-03-13 11:00:00', status: 'normal' }
])

function handleSearch() {
  ElMessage.info('搜索功能已触发')
}

function blockReview(row: any) {
  row.status = 'blocked'
  ElMessage.success('已屏蔽')
}

function restoreReview(row: any) {
  row.status = 'normal'
  ElMessage.success('已恢复')
}

function deleteReview(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; }
.ellipsis-text { display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; text-overflow: ellipsis; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
</style>
