import request from './request'

export function getRecommendList(params?: any) {
  return request.get('/admin/recommendations', { params })
}

export function addRecommend(data: any) {
  return request.post('/admin/recommendations', data)
}

export function removeRecommend(id: number) {
  return request.delete(`/admin/recommendations/${id}`)
}

export function updateRecommendSort(data: { ids: number[] }) {
  return request.put('/admin/recommendations/sort', data)
}

export function getRankingConfig() {
  return request.get('/admin/recommendations/ranking-config')
}

export function updateRankingConfig(data: any) {
  return request.put('/admin/recommendations/ranking-config', data)
}
