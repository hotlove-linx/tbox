import request from './request'

export function getReportList(params?: any) {
  return request.get('/admin/reports', { params })
}

export function handleReport(id: number, data: { valid: boolean; note: string }) {
  return request.put(`/admin/reports/${id}/handle`, data)
}

export function getReportDetail(id: number) {
  return request.get(`/admin/reports/${id}`)
}
