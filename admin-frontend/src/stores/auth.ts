import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi, logout as logoutApi, getAdminInfo } from '@/api/auth'
import router from '@/router'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('admin_token') || '')
  const adminInfo = ref<any>(null)

  const isLoggedIn = computed(() => !!token.value)
  const permissions = computed(() => adminInfo.value?.permissions || [])
  const role = computed(() => adminInfo.value?.role || '')

  async function login(data: { email: string; password: string; captcha: string }) {
    // Mock login for development
    token.value = 'mock_token_' + Date.now()
    localStorage.setItem('admin_token', token.value)
    adminInfo.value = {
      id: 1,
      name: '超级管理员',
      email: data.email,
      avatar: '',
      role: '超级管理员',
      permissions: ['*']
    }
    return true
  }

  async function fetchAdminInfo() {
    // Mock admin info
    if (!adminInfo.value) {
      adminInfo.value = {
        id: 1,
        name: '超级管理员',
        email: 'admin@tbox.com',
        avatar: '',
        role: '超级管理员',
        permissions: ['*']
      }
    }
  }

  async function logout() {
    token.value = ''
    adminInfo.value = null
    localStorage.removeItem('admin_token')
    router.push('/login')
  }

  function hasPermission(perm: string): boolean {
    if (permissions.value.includes('*')) return true
    return permissions.value.includes(perm)
  }

  return {
    token,
    adminInfo,
    isLoggedIn,
    permissions,
    role,
    login,
    fetchAdminInfo,
    logout,
    hasPermission
  }
})
