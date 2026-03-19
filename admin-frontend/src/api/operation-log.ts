import request from './request'

export function getOperationLogs(params?: any) {
  return request.get('/admin/operation-logs', { params })
}

export function exportOperationLogs(params?: any) {
  return request.get('/admin/operation-logs/export', { params, responseType: 'blob' })
}
