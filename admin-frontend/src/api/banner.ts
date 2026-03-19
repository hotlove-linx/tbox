import request from './request'

export function getBannerList(params?: any) {
  return request.get('/admin/banners', { params })
}

export function createBanner(data: any) {
  return request.post('/admin/banners', data)
}

export function updateBanner(id: number, data: any) {
  return request.put(`/admin/banners/${id}`, data)
}

export function deleteBanner(id: number) {
  return request.delete(`/admin/banners/${id}`)
}

export function updateBannerSort(data: { ids: number[] }) {
  return request.put('/admin/banners/sort', data)
}
