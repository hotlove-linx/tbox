import request from './request'

export function getOverviewStats() {
  return request.get('/admin/dashboard/overview')
}

export function getDownloadTrend(params: { period: string }) {
  return request.get('/admin/dashboard/download-trend', { params })
}

export function getUserGrowthTrend() {
  return request.get('/admin/dashboard/user-growth')
}

export function getCategoryDistribution() {
  return request.get('/admin/dashboard/category-distribution')
}

export function getTopSoftware() {
  return request.get('/admin/dashboard/top-software')
}

export function getTodoItems() {
  return request.get('/admin/dashboard/todos')
}
