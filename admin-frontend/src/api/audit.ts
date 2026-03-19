import request from './request'

export function getAuditList(params?: any) {
  return request.get('/admin/audit/software', { params })
}

export function getAuditDetail(id: number) {
  return request.get(`/admin/audit/software/${id}`)
}

export function approveSoftware(id: number, data?: any) {
  return request.put(`/admin/audit/software/${id}/approve`, data)
}

export function rejectSoftware(id: number, data: { reason: string }) {
  return request.put(`/admin/audit/software/${id}/reject`, data)
}

export function transferAudit(id: number, data: { adminId: number }) {
  return request.put(`/admin/audit/software/${id}/transfer`, data)
}
