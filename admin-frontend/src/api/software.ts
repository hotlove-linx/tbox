import request from './request'

export function getSoftwareList(params: any) {
  return request.get('/admin/software', { params })
}

export function getSoftwareDetail(id: number) {
  return request.get(`/admin/software/${id}`)
}

export function updateSoftware(id: number, data: any) {
  return request.put(`/admin/software/${id}`, data)
}

export function deleteSoftware(id: number) {
  return request.delete(`/admin/software/${id}`)
}

export function toggleSoftwareVisibility(id: number, visible: boolean) {
  return request.put(`/admin/software/${id}/visibility`, { visible })
}

export function batchOperation(data: { ids: number[]; action: string }) {
  return request.post('/admin/software/batch', data)
}

export function exportSoftwareList(params: any) {
  return request.get('/admin/software/export', { params, responseType: 'blob' })
}
