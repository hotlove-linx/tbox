import request from './request'

export function login(data: { email: string; password: string; captcha: string }) {
  return request.post('/admin/auth/login', data)
}

export function logout() {
  return request.post('/admin/auth/logout')
}

export function getAdminInfo() {
  return request.get('/admin/auth/info')
}

export function changePassword(data: { oldPassword: string; newPassword: string }) {
  return request.put('/admin/auth/password', data)
}

export function getCaptcha() {
  return request.get('/admin/auth/captcha')
}
