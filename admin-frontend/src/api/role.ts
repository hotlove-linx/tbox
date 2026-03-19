import request from './request'

export function getRoleList(params?: any) {
  return request.get('/admin/roles', { params })
}

export function createRole(data: any) {
  return request.post('/admin/roles', data)
}

export function updateRole(id: number, data: any) {
  return request.put(`/admin/roles/${id}`, data)
}

export function deleteRole(id: number) {
  return request.delete(`/admin/roles/${id}`)
}

export function getRolePermissions(id: number) {
  return request.get(`/admin/roles/${id}/permissions`)
}

export function updateRolePermissions(id: number, data: { permissions: string[] }) {
  return request.put(`/admin/roles/${id}/permissions`, data)
}

export function getAllPermissions() {
  return request.get('/admin/permissions')
}
