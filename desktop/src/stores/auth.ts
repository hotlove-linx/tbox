import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as apiLogin, logout as apiLogout, type LoginParams } from '@/api/auth'
import router from '@/router'

export interface UserInfo {
  id: number
  email: string
  nickname: string
  avatar: string
  createdAt: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string>(localStorage.getItem('tbox_token') || '')
  const userInfo = ref<UserInfo | null>(
    JSON.parse(localStorage.getItem('tbox_user') || 'null')
  )

  const isLoggedIn = computed(() => !!token.value)
  const userName = computed(() => userInfo.value?.nickname || '用户')
  const userEmail = computed(() => userInfo.value?.email || '')
  const userAvatar = computed(() => userInfo.value?.avatar || '')

  async function login(params: LoginParams) {
    try {
      const res = await apiLogin(params)
      const data = res.data
      token.value = data.token
      userInfo.value = data.user
      localStorage.setItem('tbox_token', data.token)
      localStorage.setItem('tbox_user', JSON.stringify(data.user))
      return true
    } catch (error) {
      return false
    }
  }

  async function logout() {
    try {
      await apiLogout()
    } catch (e) {
      // ignore
    }
    clearAuth()
    router.push('/login')
  }

  function clearAuth() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('tbox_token')
    localStorage.removeItem('tbox_user')
  }

  function updateUserInfo(info: Partial<UserInfo>) {
    if (userInfo.value) {
      userInfo.value = { ...userInfo.value, ...info }
      localStorage.setItem('tbox_user', JSON.stringify(userInfo.value))
    }
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    userName,
    userEmail,
    userAvatar,
    login,
    logout,
    clearAuth,
    updateUserInfo
  }
})
