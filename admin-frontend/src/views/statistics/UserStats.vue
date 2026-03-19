<template>
  <div class="page-container">
    <!-- Overview Cards -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6" v-for="stat in overviewStats" :key="stat.label">
        <el-card shadow="hover">
          <div class="stat-card">
            <div class="stat-label">{{ stat.label }}</div>
            <div class="stat-value">{{ stat.value }}</div>
            <div class="stat-sub">{{ stat.sub }}</div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- User Growth Trend -->
    <el-card shadow="never" class="section-card">
      <div class="chart-header">
        <span class="chart-title">用户增长趋势</span>
        <el-radio-group v-model="growthPeriod" size="small">
          <el-radio-button value="30">近30天</el-radio-button>
          <el-radio-button value="90">近90天</el-radio-button>
        </el-radio-group>
      </div>
      <div ref="growthChartRef" style="height: 350px;"></div>
    </el-card>

    <!-- DAU/WAU/MAU -->
    <el-card shadow="never" class="section-card">
      <div class="chart-header">
        <span class="chart-title">活跃用户趋势 (DAU/WAU/MAU)</span>
      </div>
      <div ref="activeChartRef" style="height: 350px;"></div>
    </el-card>

    <!-- Retention Table -->
    <el-card shadow="never" class="section-card">
      <div class="chart-header"><span class="chart-title">用户留存率</span></div>
      <el-table :data="retentionData" stripe border>
        <el-table-column prop="date" label="日期" width="120" />
        <el-table-column prop="newUsers" label="新增用户" width="100" align="center" />
        <el-table-column label="次日留存" width="110" align="center">
          <template #default="{ row }">
            <span :style="{ color: retentionColor(row.day1) }">{{ row.day1 }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="3日留存" width="110" align="center">
          <template #default="{ row }">
            <span :style="{ color: retentionColor(row.day3) }">{{ row.day3 }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="7日留存" width="110" align="center">
          <template #default="{ row }">
            <span :style="{ color: retentionColor(row.day7) }">{{ row.day7 }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="14日留存" width="110" align="center">
          <template #default="{ row }">
            <span :style="{ color: retentionColor(row.day14) }">{{ row.day14 }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="30日留存" width="110" align="center">
          <template #default="{ row }">
            <span :style="{ color: retentionColor(row.day30) }">{{ row.day30 }}%</span>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'

const growthPeriod = ref('30')

const overviewStats = ref([
  { label: '总注册用户', value: '128,365', sub: '' },
  { label: '今日新增', value: '286', sub: '+12.3% 较昨日' },
  { label: '本周新增', value: '1,842', sub: '+5.8% 较上周' },
  { label: '本月新增', value: '8,256', sub: '+15.2% 较上月' }
])

const retentionData = ref([
  { date: '3/15', newUsers: 286, day1: 45.2, day3: 32.1, day7: 22.5, day14: 15.8, day30: 10.2 },
  { date: '3/14', newUsers: 312, day1: 42.8, day3: 30.5, day7: 21.3, day14: 14.6, day30: 9.8 },
  { date: '3/13', newUsers: 258, day1: 48.1, day3: 35.2, day7: 24.8, day14: 17.2, day30: 11.5 },
  { date: '3/12', newUsers: 295, day1: 44.5, day3: 31.8, day7: 22.1, day14: 15.3, day30: 10.0 },
  { date: '3/11', newUsers: 275, day1: 46.3, day3: 33.6, day7: 23.4, day14: 16.1, day30: 10.8 },
  { date: '3/10', newUsers: 320, day1: 43.7, day3: 31.2, day7: 21.8, day14: 14.9, day30: 9.5 },
  { date: '3/9', newUsers: 248, day1: 47.5, day3: 34.1, day7: 23.9, day14: 16.5, day30: 11.2 }
])

function retentionColor(val: number) {
  if (val >= 40) return '#67c23a'
  if (val >= 20) return '#e6a23c'
  return '#f56c6c'
}

const growthChartRef = ref<HTMLElement>()
const activeChartRef = ref<HTMLElement>()
let growthChart: echarts.ECharts
let activeChart: echarts.ECharts

function generateDates(n: number) {
  const dates = []
  const now = new Date()
  for (let i = n - 1; i >= 0; i--) {
    const d = new Date(now); d.setDate(d.getDate() - i)
    dates.push(`${d.getMonth() + 1}/${d.getDate()}`)
  }
  return dates
}

function initGrowthChart() {
  if (!growthChartRef.value) return
  growthChart = echarts.init(growthChartRef.value)
  const dates = generateDates(30)
  growthChart.setOption({
    tooltip: { trigger: 'axis' },
    legend: { data: ['新增用户', '累计用户'], top: 0 },
    grid: { left: 60, right: 60, top: 40, bottom: 30 },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: [
      { type: 'value', name: '新增用户', position: 'left' },
      { type: 'value', name: '累计用户', position: 'right' }
    ],
    series: [
      {
        name: '新增用户', type: 'bar', yAxisIndex: 0,
        data: Array.from({ length: 30 }, () => Math.floor(Math.random() * 200 + 200)),
        itemStyle: { borderRadius: [4, 4, 0, 0], color: '#409eff' }
      },
      {
        name: '累计用户', type: 'line', yAxisIndex: 1, smooth: true,
        data: Array.from({ length: 30 }, (_, i) => 120000 + i * 300 + Math.floor(Math.random() * 100)),
        lineStyle: { color: '#67c23a', width: 2 }, itemStyle: { color: '#67c23a' }
      }
    ]
  })
}

function initActiveChart() {
  if (!activeChartRef.value) return
  activeChart = echarts.init(activeChartRef.value)
  const dates = generateDates(30)
  activeChart.setOption({
    tooltip: { trigger: 'axis' },
    legend: { data: ['DAU', 'WAU', 'MAU'], top: 0 },
    grid: { left: 60, right: 20, top: 40, bottom: 30 },
    xAxis: { type: 'category', data: dates, boundaryGap: false },
    yAxis: { type: 'value' },
    series: [
      {
        name: 'DAU', type: 'line', smooth: true,
        data: Array.from({ length: 30 }, () => Math.floor(Math.random() * 2000 + 7000)),
        lineStyle: { color: '#409eff' }, itemStyle: { color: '#409eff' }
      },
      {
        name: 'WAU', type: 'line', smooth: true,
        data: Array.from({ length: 30 }, () => Math.floor(Math.random() * 5000 + 25000)),
        lineStyle: { color: '#67c23a' }, itemStyle: { color: '#67c23a' }
      },
      {
        name: 'MAU', type: 'line', smooth: true,
        data: Array.from({ length: 30 }, () => Math.floor(Math.random() * 10000 + 60000)),
        lineStyle: { color: '#e6a23c' }, itemStyle: { color: '#e6a23c' }
      }
    ]
  })
}

function handleResize() { growthChart?.resize(); activeChart?.resize() }

onMounted(() => {
  initGrowthChart(); initActiveChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  growthChart?.dispose(); activeChart?.dispose()
})
</script>

<style scoped>
.stats-row { margin-bottom: 20px; }
.section-card { margin-bottom: 20px; }
.chart-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; }
.chart-title { font-size: 16px; font-weight: 600; color: #303133; }
.stat-card { text-align: center; }
.stat-label { font-size: 14px; color: #909399; margin-bottom: 8px; }
.stat-value { font-size: 28px; font-weight: 700; color: #303133; }
.stat-sub { font-size: 13px; color: #67c23a; margin-top: 4px; }
</style>
