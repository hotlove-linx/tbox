<template>
  <div class="dashboard-page">
    <!-- Stats Cards -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6" v-for="stat in statsCards" :key="stat.title">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-card-inner">
            <div class="stat-info">
              <div class="stat-label">{{ stat.title }}</div>
              <div class="stat-value">{{ stat.value }}</div>
              <div :class="['stat-change', stat.changeType]">
                <el-icon v-if="stat.changeType === 'up'"><Top /></el-icon>
                <el-icon v-else><Bottom /></el-icon>
                {{ stat.change }} {{ stat.subText }}
              </div>
            </div>
            <div class="stat-icon" :style="{ background: stat.iconBg }">
              <el-icon :size="28" color="#fff"><component :is="stat.icon" /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Charts 2x2 Grid -->
    <el-row :gutter="20" class="chart-row">
      <el-col :span="12">
        <el-card shadow="hover" class="chart-card">
          <div class="chart-header">
            <span class="chart-title">下载量趋势</span>
            <el-radio-group v-model="downloadPeriod" size="small">
              <el-radio-button value="7">近7天</el-radio-button>
              <el-radio-button value="30">近30天</el-radio-button>
            </el-radio-group>
          </div>
          <div ref="downloadChartRef" class="chart-body"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover" class="chart-card">
          <div class="chart-header">
            <span class="chart-title">用户增长趋势</span>
          </div>
          <div ref="userGrowthChartRef" class="chart-body"></div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="chart-row">
      <el-col :span="12">
        <el-card shadow="hover" class="chart-card">
          <div class="chart-header">
            <span class="chart-title">软件分类分布</span>
          </div>
          <div ref="categoryChartRef" class="chart-body"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover" class="chart-card">
          <div class="chart-header">
            <span class="chart-title">热门软件 Top 10</span>
          </div>
          <div ref="topSoftwareChartRef" class="chart-body"></div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Todo Section -->
    <el-row :gutter="20" class="todo-row">
      <el-col :span="8" v-for="todo in todoItems" :key="todo.title">
        <el-card shadow="hover" class="todo-card" @click="$router.push(todo.link)">
          <div class="todo-inner">
            <el-icon :size="24" :color="todo.color"><component :is="todo.icon" /></el-icon>
            <div class="todo-info">
              <div class="todo-title">{{ todo.title }}</div>
              <div class="todo-count" :style="{ color: todo.color }">{{ todo.count }}</div>
            </div>
            <el-icon color="#c0c4cc"><ArrowRight /></el-icon>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import * as echarts from 'echarts'

const downloadPeriod = ref('7')

const statsCards = ref([
  { title: '总用户数', value: '128,365', change: '+286', subText: '今日新增', changeType: 'up', icon: 'User', iconBg: '#409eff' },
  { title: '总软件数', value: '3,842', change: '12', subText: '待审核', changeType: 'up', icon: 'Box', iconBg: '#67c23a' },
  { title: '今日下载量', value: '15,728', change: '+12.5%', subText: '较昨日', changeType: 'up', icon: 'Download', iconBg: '#e6a23c' },
  { title: '活跃用户(DAU)', value: '8,432', change: '-3.2%', subText: '较昨日', changeType: 'down', icon: 'TrendCharts', iconBg: '#f56c6c' }
])

const todoItems = ref([
  { title: '待审核软件', count: 12, icon: 'Document', color: '#e6a23c', link: '/audit/software' },
  { title: '待处理举报', count: 5, icon: 'Warning', color: '#f56c6c', link: '/audit/reports' },
  { title: '待回复反馈', count: 8, icon: 'ChatLineSquare', color: '#409eff', link: '/user/feedback' }
])

const downloadChartRef = ref<HTMLElement>()
const userGrowthChartRef = ref<HTMLElement>()
const categoryChartRef = ref<HTMLElement>()
const topSoftwareChartRef = ref<HTMLElement>()

let downloadChart: echarts.ECharts
let userGrowthChart: echarts.ECharts
let categoryChart: echarts.ECharts
let topSoftwareChart: echarts.ECharts

function generateDates(days: number) {
  const dates = []
  const now = new Date()
  for (let i = days - 1; i >= 0; i--) {
    const d = new Date(now)
    d.setDate(d.getDate() - i)
    dates.push(`${d.getMonth() + 1}/${d.getDate()}`)
  }
  return dates
}

function generateRandomData(len: number, min: number, max: number) {
  return Array.from({ length: len }, () => Math.floor(Math.random() * (max - min) + min))
}

function initDownloadChart() {
  if (!downloadChartRef.value) return
  downloadChart = echarts.init(downloadChartRef.value)
  updateDownloadChart()
}

function updateDownloadChart() {
  const days = parseInt(downloadPeriod.value)
  const dates = generateDates(days)
  downloadChart.setOption({
    tooltip: { trigger: 'axis' },
    grid: { left: 50, right: 20, top: 20, bottom: 30 },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: { type: 'value' },
    series: [{
      data: generateRandomData(days, 10000, 20000),
      type: 'line',
      smooth: true,
      areaStyle: { color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
        { offset: 0, color: 'rgba(64,158,255,0.3)' },
        { offset: 1, color: 'rgba(64,158,255,0.05)' }
      ])},
      lineStyle: { color: '#409eff', width: 2 },
      itemStyle: { color: '#409eff' }
    }]
  })
}

function initUserGrowthChart() {
  if (!userGrowthChartRef.value) return
  userGrowthChart = echarts.init(userGrowthChartRef.value)
  const dates = generateDates(30)
  userGrowthChart.setOption({
    tooltip: { trigger: 'axis' },
    grid: { left: 50, right: 20, top: 20, bottom: 30 },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: { type: 'value' },
    series: [{
      data: generateRandomData(30, 200, 500),
      type: 'line',
      smooth: true,
      areaStyle: { color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
        { offset: 0, color: 'rgba(103,194,58,0.3)' },
        { offset: 1, color: 'rgba(103,194,58,0.05)' }
      ])},
      lineStyle: { color: '#67c23a', width: 2 },
      itemStyle: { color: '#67c23a' }
    }]
  })
}

function initCategoryChart() {
  if (!categoryChartRef.value) return
  categoryChart = echarts.init(categoryChartRef.value)
  categoryChart.setOption({
    tooltip: { trigger: 'item', formatter: '{b}: {c} ({d}%)' },
    legend: { orient: 'vertical', right: 10, top: 'center' },
    series: [{
      type: 'pie',
      radius: ['40%', '70%'],
      center: ['40%', '50%'],
      avoidLabelOverlap: true,
      label: { show: false },
      data: [
        { value: 580, name: '工具软件' },
        { value: 420, name: '游戏娱乐' },
        { value: 350, name: '效率办公' },
        { value: 280, name: '社交通讯' },
        { value: 220, name: '影音播放' },
        { value: 180, name: '教育学习' },
        { value: 150, name: '系统工具' },
        { value: 120, name: '其他' }
      ]
    }]
  })
}

function initTopSoftwareChart() {
  if (!topSoftwareChartRef.value) return
  topSoftwareChart = echarts.init(topSoftwareChartRef.value)
  const names = ['微信', 'QQ', '抖音', '支付宝', '淘宝', 'WPS', '网易云音乐', '百度地图', '美团', '京东']
  const values = [52340, 48210, 43560, 38920, 35100, 31240, 28700, 25300, 22800, 20100]
  topSoftwareChart.setOption({
    tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 100, right: 30, top: 10, bottom: 10 },
    xAxis: { type: 'value' },
    yAxis: { type: 'category', data: names.reverse(), axisLabel: { fontSize: 12 } },
    series: [{
      type: 'bar',
      data: values.reverse(),
      barWidth: 16,
      itemStyle: {
        borderRadius: [0, 4, 4, 0],
        color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
          { offset: 0, color: '#409eff' },
          { offset: 1, color: '#67c23a' }
        ])
      }
    }]
  })
}

watch(downloadPeriod, () => {
  updateDownloadChart()
})

function handleResize() {
  downloadChart?.resize()
  userGrowthChart?.resize()
  categoryChart?.resize()
  topSoftwareChart?.resize()
}

onMounted(() => {
  initDownloadChart()
  initUserGrowthChart()
  initCategoryChart()
  initTopSoftwareChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  downloadChart?.dispose()
  userGrowthChart?.dispose()
  categoryChart?.dispose()
  topSoftwareChart?.dispose()
})
</script>

<style scoped>
.dashboard-page {
  min-width: 1000px;
}

.stats-row {
  margin-bottom: 20px;
}

.stat-card-inner {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: #303133;
  margin-bottom: 8px;
}

.stat-change {
  font-size: 13px;
  display: flex;
  align-items: center;
  gap: 4px;
}

.stat-change.up {
  color: #67c23a;
}

.stat-change.down {
  color: #f56c6c;
}

.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chart-row {
  margin-bottom: 20px;
}

.chart-card {
  height: 100%;
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.chart-title {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.chart-body {
  height: 300px;
}

.todo-row {
  margin-bottom: 20px;
}

.todo-card {
  cursor: pointer;
  transition: transform 0.2s;
}

.todo-card:hover {
  transform: translateY(-2px);
}

.todo-inner {
  display: flex;
  align-items: center;
  gap: 16px;
}

.todo-info {
  flex: 1;
}

.todo-title {
  font-size: 14px;
  color: #606266;
}

.todo-count {
  font-size: 24px;
  font-weight: 700;
  margin-top: 4px;
}
</style>
