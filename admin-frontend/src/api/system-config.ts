import request from './request'

export function getSystemConfig() {
  return request.get('/admin/system-config')
}

export function updateSystemConfig(group: string, data: any) {
  return request.put(`/admin/system-config/${group}`, data)
}
