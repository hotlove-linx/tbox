import request from './request'

export interface UserProfile {
  id: number
  email: string
  nickname: string
  avatar: string
  createdAt: string
}

export function getProfile() {
  return request.get<any, { data: UserProfile }>('/user/profile')
}

export function updateProfile(data: Partial<UserProfile>) {
  return request.put('/user/profile', data)
}

export function uploadAvatar(file: File) {
  const formData = new FormData()
  formData.append('avatar', file)
  return request.post('/user/avatar', formData, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}
