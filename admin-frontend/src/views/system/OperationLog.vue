<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="filter-bar">
        <el-select v-model="filters.operator" placeholder="操作人" clearable style="width: 140px">
          <el-option label="超级管理员" value="超级管理员" />
          <el-option label="运营小王" value="运营小王" />
          <el-option label="审核员张" value="审核员张" />
          <el-option label="客服小李" value="客服小李" />
        </el-select>
        <el-select v-model="filters.module" placeholder="操作模块" clearable style="width: 130px">
          <el-option label="软件管理" value="software" />
          <el-option label="用户管理" value="user" />
          <el-option label="内容运营" value="content" />
          <el-option label="审核中心" value="audit" />
          <el-option label="系统管理" value="system" />
        </el-select>
        <el-select v-model="filters.type" placeholder="操作类型" clearable style="width: 120px">
          <el-option label="新增" value="create" />
          <el-option label="编辑" value="update" />
          <el-option label="删除" value="delete" />
          <el-option label="审核" value="audit" />
          <el-option label="登录" value="login" />
        </el-select>
        <el-date-picker v-model="filters.dateRange" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" style="width: 360px" />
        <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
        <el-button type="success" @click="handleExport"><el-icon><Download /></el-icon>导出</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="time" label="时间" width="170" />
        <el-table-column label="操作人" width="160">
          <template #default="{ row }">
            <div>
              <span>{{ row.operator }}</span>
              <el-tag size="small" type="info" style="margin-left: 6px;">{{ row.role }}</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="module" label="操作模块" width="110" />
        <el-table-column label="操作类型" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="typeTagMap[row.type]" size="small">{{ typeTextMap[row.type] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="target" label="操作对象" width="160" show-overflow-tooltip />
        <el-table-column prop="detail" label="操作详情" min-width="250" show-overflow-tooltip />
        <el-table-column prop="ip" label="IP地址" width="130" />
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[20, 50, 100]" :total="total" layout="total, sizes, prev, pager, next" background />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const filters = reactive({ operator: '', module: '', type: '', dateRange: null as any })
const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(1286)

const typeTagMap: Record<string, any> = { create: 'success', update: 'primary', delete: 'danger', audit: 'warning', login: 'info' }
const typeTextMap: Record<string, string> = { create: '新增', update: '编辑', delete: '删除', audit: '审核', login: '登录' }

const tableData = ref([
  { time: '2024-03-15 20:30:15', operator: '超级管理员', role: '超级管理员', module: '系统管理', type: 'update', target: '系统配置', detail: '修改了注册邮件验证开关为开启', ip: '192.168.1.100' },
  { time: '2024-03-15 18:20:30', operator: '运营小王', role: '运营管理员', module: '内容运营', type: 'create', target: 'Banner #5', detail: '新增Banner：春季活动推荐', ip: '192.168.1.101' },
  { time: '2024-03-15 17:45:00', operator: '审核员张', role: '审核员', module: '审核中心', type: 'audit', target: '软件 #2001', detail: '通过软件审核：智能笔记 v2.1.0', ip: '192.168.1.102' },
  { time: '2024-03-15 17:30:20', operator: '审核员张', role: '审核员', module: '审核中心', type: 'audit', target: '评论 #3892', detail: '屏蔽违规评论', ip: '192.168.1.102' },
  { time: '2024-03-15 16:00:45', operator: '客服小李', role: '客服', module: '用户管理', type: 'update', target: '用户反馈 #1', detail: '回复用户反馈：下载功能异常', ip: '192.168.1.103' },
  { time: '2024-03-15 15:20:10', operator: '运营小王', role: '运营管理员', module: '软件管理', type: 'update', target: '软件 #1001', detail: '编辑软件信息：超级文件管理器', ip: '192.168.1.101' },
  { time: '2024-03-15 14:00:00', operator: '超级管理员', role: '超级管理员', module: '系统管理', type: 'create', target: '管理员 #6', detail: '新增管理员：新审核员', ip: '192.168.1.100' },
  { time: '2024-03-15 10:30:00', operator: '运营小王', role: '运营管理员', module: '内容运营', type: 'delete', target: '公告 #8', detail: '删除过期公告：年终总结', ip: '192.168.1.101' },
  { time: '2024-03-15 09:00:00', operator: '超级管理员', role: '超级管理员', module: '系统管理', type: 'login', target: '登录', detail: '管理员登录系统', ip: '192.168.1.100' },
  { time: '2024-03-15 08:55:00', operator: '审核员张', role: '审核员', module: '系统管理', type: 'login', target: '登录', detail: '管理员登录系统', ip: '192.168.1.102' }
])

function handleSearch() {
  ElMessage.info('搜索功能已触发')
}

function handleExport() {
  ElMessage.success('导出任务已提交')
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; flex-wrap: wrap; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
</style>
