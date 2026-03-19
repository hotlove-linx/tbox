import request from './request'

export function getAdminList(params?: any) {
  return request.get('/admin/admins', { params })
}

export function createAdmin(data: any) {
  return request.post('/admin/admins', data)
}

export function updateAdmin(id: number, data: any) {
  return request.put(`/admin/admins/${id}`, data)
}

export function deleteAdmin(id: number) {
  return request.delete(`/admin/admins/${id}`)
}

export function resetAdminPassword(id: number) {
  return request.put(`/admin/admins/${id}/reset-password`)
}

export function toggleAdminStatus(id: number, enabled: boolean) {
  return request.put(`/admin/admins/${id}/status`, { enabled })
}
