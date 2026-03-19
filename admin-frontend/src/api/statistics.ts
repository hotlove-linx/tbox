import request from './request'

export function getDownloadOverview() {
  return request.get('/admin/statistics/downloads/overview')
}

export function getDownloadTrend(params: any) {
  return request.get('/admin/statistics/downloads/trend', { params })
}

export function getDownloadByCategory(params?: any) {
  return request.get('/admin/statistics/downloads/by-category', { params })
}

export function getDownloadByHour() {
  return request.get('/admin/statistics/downloads/by-hour')
}

export function getUserOverview() {
  return request.get('/admin/statistics/users/overview')
}

export function getUserGrowthTrend(params: any) {
  return request.get('/admin/statistics/users/growth', { params })
}

export function getActiveUserTrend(params: any) {
  return request.get('/admin/statistics/users/active', { params })
}

export function getUserRetention() {
  return request.get('/admin/statistics/users/retention')
}

export function getSoftwareRanking(params: any) {
  return request.get('/admin/statistics/software/ranking', { params })
}
