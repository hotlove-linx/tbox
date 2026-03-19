import request from './request'

export function getTagList(params?: any) {
  return request.get('/admin/tags', { params })
}

export function createTag(data: any) {
  return request.post('/admin/tags', data)
}

export function updateTag(id: number, data: any) {
  return request.put(`/admin/tags/${id}`, data)
}

export function deleteTag(id: number) {
  return request.delete(`/admin/tags/${id}`)
}
