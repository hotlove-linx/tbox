import request from './request'

export interface LoginParams {
  email: string
  password: string
}

export interface RegisterParams {
  email: string
  code: string
  password: string
  confirmPassword: string
}

export interface LoginResult {
  token: string
  user: {
    id: number
    email: string
    nickname: string
    avatar: string
    createdAt: string
  }
}

export function login(data: LoginParams) {
  return request.post<any, { data: LoginResult }>('/auth/login', data)
}

export function register(data: RegisterParams) {
  return request.post('/auth/register', data)
}

export function sendVerifyCode(email: string) {
  return request.post('/auth/send-code', { email })
}

export function logout() {
  return request.post('/auth/logout')
}

export function getUserInfo() {
  return request.get('/auth/user-info')
}

export function changePassword(data: { oldPassword: string; newPassword: string }) {
  return request.put('/auth/change-password', data)
}

export function deleteAccount() {
  return request.delete('/auth/account')
}
