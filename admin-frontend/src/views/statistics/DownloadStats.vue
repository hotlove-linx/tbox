<template>
  <div class="page-container">
    <!-- Overview Cards -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6" v-for="stat in overviewStats" :key="stat.label">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-label">{{ stat.label }}</div>
            <div class="stat-value">{{ stat.value }}</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Download Trend -->
    <el-card shadow="never" class="section-card">
      <div class="chart-header">
        <span class="chart-title">下载趋势</span>
        <div class="chart-controls">
          <el-radio-group v-model="trendDimension" size="small" @change="updateTrendChart">
            <el-radio-button value="day">日</el-radio-button>
            <el-radio-button value="week">周</el-radio-button>
            <el-radio-button value="month">月</el-radio-button>
          </el-radio-group>
          <el-date-picker v-model="trendDateRange" type="daterange" range-separator="~" start-placeholder="开始" end-placeholder="结束" size="small" style="width: 240px; margin-left: 12px;" />
        </div>
      </div>
      <div ref="trendChartRef" style="height: 350px;"></div>
    </el-card>

    <!-- Category Distribution -->
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card shadow="never" class="section-card">
          <div class="chart-header"><span class="chart-title">分类下载分布</span></div>
          <div ref="categoryBarChartRef" style="height: 350px;"></div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="never" class="section-card">
          <div class="chart-header"><span class="chart-title">时段分布</span></div>
          <div ref="hourChartRef" style="height: 350px;"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'

const trendDimension = ref('day')
const trendDateRange = ref(null)

const overviewStats = ref([
  { label: '今日下载量', value: '15,728' },
  { label: '本周下载量', value: '98,342' },
  { label: '本月下载量', value: '425,680' },
  { label: '累计下载量', value: '3,256,890' }
])

const trendChartRef = ref<HTMLElement>()
const categoryBarChartRef = ref<HTMLElement>()
const hourChartRef = ref<HTMLElement>()
let trendChart: echarts.ECharts
let categoryBarChart: echarts.ECharts
let hourChart: echarts.ECharts

function generateDates(n: number) {
  const dates = []
  const now = new Date()
  for (let i = n - 1; i >= 0; i--) {
    const d = new Date(now)
    d.setDate(d.getDate() - i)
    dates.push(`${d.getMonth() + 1}/${d.getDate()}`)
  }
  return dates
}

function initTrendChart() {
  if (!trendChartRef.value) return
  trendChart = echarts.init(trendChartRef.value)
  updateTrendChart()
}

function updateTrendChart() {
  const days = trendDimension.value === 'day' ? 30 : trendDimension.value === 'week' ? 12 : 12
  const dates = trendDimension.value === 'day'
    ? generateDates(days)
    : Array.from({ length: days }, (_, i) => trendDimension.value === 'week' ? `第${i + 1}周` : `${i + 1}月`)

  trendChart.setOption({
    tooltip: { trigger: 'axis' },
    grid: { left: 60, right: 20, top: 20, bottom: 30 },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: { type: 'value' },
    series: [{
      data: Array.from({ length: days }, () => Math.floor(Math.random() * 15000 + 10000)),
      type: 'line', smooth: true,
      areaStyle: { color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
        { offset: 0, color: 'rgba(64,158,255,0.4)' }, { offset: 1, color: 'rgba(64,158,255,0.05)' }
      ])},
      lineStyle: { color: '#409eff', width: 2 },
      itemStyle: { color: '#409eff' }
    }]
  })
}

function initCategoryBarChart() {
  if (!categoryBarChartRef.value) return
  categoryBarChart = echarts.init(categoryBarChartRef.value)
  categoryBarChart.setOption({
    tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 80, right: 20, top: 10, bottom: 30 },
    xAxis: { type: 'value' },
    yAxis: { type: 'category', data: ['其他', '系统工具', '教育学习', '影音播放', '社交通讯', '效率办公', '游戏娱乐', '工具软件'] },
    series: [{
      type: 'bar', barWidth: 18,
      data: [45200, 78300, 112000, 186500, 248000, 325600, 410000, 520800],
      itemStyle: { borderRadius: [0, 4, 4, 0], color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
        { offset: 0, color: '#409eff' }, { offset: 1, color: '#67c23a' }
      ])}
    }]
  })
}

function initHourChart() {
  if (!hourChartRef.value) return
  hourChart = echarts.init(hourChartRef.value)
  const hours = Array.from({ length: 24 }, (_, i) => `${i}:00`)
  const data = [200, 120, 80, 50, 30, 60, 180, 450, 720, 980, 1100, 1250, 1380, 1200, 1050, 920, 850, 980, 1150, 1320, 1180, 980, 650, 380]
  hourChart.setOption({
    tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 50, right: 20, top: 10, bottom: 30 },
    xAxis: { type: 'category', data: hours, axisLabel: { interval: 2 } },
    yAxis: { type: 'value' },
    series: [{
      type: 'bar', data,
      itemStyle: { borderRadius: [4, 4, 0, 0], color: '#e6a23c' }
    }]
  })
}

function handleResize() {
  trendChart?.resize()
  categoryBarChart?.resize()
  hourChart?.resize()
}

onMounted(() => {
  initTrendChart()
  initCategoryBarChart()
  initHourChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  trendChart?.dispose()
  categoryBarChart?.dispose()
  hourChart?.dispose()
})
</script>

<style scoped>
.stats-row { margin-bottom: 20px; }
.section-card { margin-bottom: 20px; }
.chart-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; }
.chart-title { font-size: 16px; font-weight: 600; color: #303133; }
.chart-controls { display: flex; align-items: center; }
.stat-card { text-align: center; }
.stat-label { font-size: 14px; color: #909399; margin-bottom: 8px; }
.stat-value { font-size: 28px; font-weight: 700; color: #303133; }
</style>
