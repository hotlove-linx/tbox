import request from './request'

export function getUserList(params?: any) {
  return request.get('/admin/users', { params })
}

export function getUserDetail(id: number) {
  return request.get(`/admin/users/${id}`)
}

export function toggleUserStatus(id: number, enabled: boolean) {
  return request.put(`/admin/users/${id}/status`, { enabled })
}

export function adjustUserSpace(id: number, data: { space: number }) {
  return request.put(`/admin/users/${id}/space`, data)
}

export function resetUserPassword(id: number) {
  return request.put(`/admin/users/${id}/reset-password`)
}

export function exportUsers(params?: any) {
  return request.get('/admin/users/export', { params, responseType: 'blob' })
}
