<template>
  <div class="login-page">
    <div class="login-card">
      <!-- Logo -->
      <div class="login-header">
        <div class="login-logo">
          <el-icon :size="36" color="#409eff"><Box /></el-icon>
          <span class="logo-text">TBox</span>
        </div>
        <p class="login-slogan">发现、下载、管理你的软件世界</p>
      </div>

      <!-- Tabs -->
      <el-tabs v-model="activeTab" class="login-tabs" stretch>
        <el-tab-pane label="登录" name="login">
          <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" @keyup.enter="handleLogin">
            <el-form-item prop="email">
              <el-input v-model="loginForm.email" placeholder="邮箱地址" :prefix-icon="Message" size="large" />
            </el-form-item>
            <el-form-item prop="password">
              <el-input v-model="loginForm.password" type="password" placeholder="密码" :prefix-icon="Lock" size="large" show-password />
            </el-form-item>
            <div class="login-options">
              <el-checkbox v-model="rememberMe">记住我</el-checkbox>
              <el-link type="primary" :underline="false">忘记密码？</el-link>
            </div>
            <el-form-item>
              <el-button type="primary" size="large" :loading="loginLoading" @click="handleLogin" class="login-btn">
                登录
              </el-button>
            </el-form-item>
          </el-form>
          <div class="switch-text">
            还没有账号？<el-link type="primary" :underline="false" @click="activeTab = 'register'">立即注册</el-link>
          </div>
        </el-tab-pane>

        <el-tab-pane label="注册" name="register">
          <el-form ref="registerFormRef" :model="registerForm" :rules="registerRules">
            <el-form-item prop="email">
              <el-input v-model="registerForm.email" placeholder="邮箱地址" :prefix-icon="Message" size="large" />
            </el-form-item>
            <el-form-item prop="code">
              <el-input v-model="registerForm.code" placeholder="验证码" :prefix-icon="Key" size="large">
                <template #append>
                  <el-button :disabled="codeCountdown > 0" @click="handleSendCode">
                    {{ codeCountdown > 0 ? `${codeCountdown}s 后重试` : '发送验证码' }}
                  </el-button>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input v-model="registerForm.password" type="password" placeholder="密码（8位以上，含字母和数字）" :prefix-icon="Lock" size="large" show-password />
            </el-form-item>
            <el-form-item prop="confirmPassword">
              <el-input v-model="registerForm.confirmPassword" type="password" placeholder="确认密码" :prefix-icon="Lock" size="large" show-password />
            </el-form-item>
            <el-form-item prop="agree">
              <el-checkbox v-model="registerForm.agree">
                我已阅读并同意<el-link type="primary" :underline="false">《用户协议》</el-link>
              </el-checkbox>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" size="large" :loading="registerLoading" @click="handleRegister" class="login-btn">
                注册
              </el-button>
            </el-form-item>
          </el-form>
          <div class="switch-text">
            已有账号？<el-link type="primary" :underline="false" @click="activeTab = 'login'">立即登录</el-link>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- Window controls -->
    <div class="login-window-controls no-drag">
      <div class="win-btn" @click="handleMinimize" title="最小化">
        <el-icon :size="14"><Minus /></el-icon>
      </div>
      <div class="win-btn" @click="handleMaximize" title="最大化">
        <el-icon :size="14"><FullScreen /></el-icon>
      </div>
      <div class="win-btn win-btn-close" @click="handleForceClose" title="关闭">
        <el-icon :size="14"><Close /></el-icon>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from 'element-plus'
import { Message, Lock, Key } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const activeTab = ref('login')
const rememberMe = ref(false)
const loginLoading = ref(false)
const registerLoading = ref(false)
const codeCountdown = ref(0)

// Login form
const loginFormRef = ref<FormInstance>()
const loginForm = reactive({
  email: '',
  password: ''
})
const loginRules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

// Register form
const registerFormRef = ref<FormInstance>()
const registerForm = reactive({
  email: '',
  code: '',
  password: '',
  confirmPassword: '',
  agree: false
})

const validatePassword = (_rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入密码'))
  } else if (value.length < 8) {
    callback(new Error('密码长度不能少于8位'))
  } else if (!/(?=.*[a-zA-Z])(?=.*\d)/.test(value)) {
    callback(new Error('密码必须包含字母和数字'))
  } else {
    callback()
  }
}

const validateConfirmPassword = (_rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请确认密码'))
  } else if (value !== registerForm.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const validateAgree = (_rule: any, value: boolean, callback: any) => {
  if (!value) {
    callback(new Error('请同意用户协议'))
  } else {
    callback()
  }
}

const registerRules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
  password: [{ required: true, validator: validatePassword, trigger: 'blur' }],
  confirmPassword: [{ required: true, validator: validateConfirmPassword, trigger: 'blur' }],
  agree: [{ validator: validateAgree, trigger: 'change' }]
}

async function handleLogin() {
  if (!loginFormRef.value) return
  await loginFormRef.value.validate(async (valid) => {
    if (!valid) return
    loginLoading.value = true
    try {
      // For demo: simulate login
      // In production, use: const success = await authStore.login(loginForm)
      setTimeout(() => {
        localStorage.setItem('tbox_token', 'demo_token_123')
        localStorage.setItem('tbox_user', JSON.stringify({
          id: 1,
          email: loginForm.email,
          nickname: loginForm.email.split('@')[0],
          avatar: '',
          createdAt: '2024-01-15 10:30:00'
        }))
        authStore.$patch({
          token: 'demo_token_123',
          userInfo: {
            id: 1,
            email: loginForm.email,
            nickname: loginForm.email.split('@')[0],
            avatar: '',
            createdAt: '2024-01-15 10:30:00'
          }
        })
        ElMessage.success('登录成功')
        router.push('/home')
        loginLoading.value = false
      }, 1000)
    } catch {
      loginLoading.value = false
    }
  })
}

async function handleRegister() {
  if (!registerFormRef.value) return
  await registerFormRef.value.validate(async (valid) => {
    if (!valid) return
    registerLoading.value = true
    setTimeout(() => {
      ElMessage.success('注册成功，请登录')
      activeTab.value = 'login'
      loginForm.email = registerForm.email
      registerLoading.value = false
    }, 1500)
  })
}

function handleMinimize() {
  window.electronAPI?.minimize()
}
function handleMaximize() {
  window.electronAPI?.maximize()
}
function handleForceClose() {
  window.electronAPI?.forceClose()
}

function handleSendCode() {
  if (!registerForm.email) {
    ElMessage.warning('请先输入邮箱')
    return
  }
  codeCountdown.value = 60
  ElMessage.success('验证码已发送')
  const timer = setInterval(() => {
    codeCountdown.value--
    if (codeCountdown.value <= 0) clearInterval(timer)
  }, 1000)
}
</script>

<style scoped>
.login-page {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  -webkit-app-region: drag;
}

.login-card {
  width: 420px;
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 40px;
  box-shadow: var(--shadow-lg);
  -webkit-app-region: no-drag;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-bottom: 12px;
}

.logo-text {
  font-size: 32px;
  font-weight: 700;
  color: var(--text-color-primary);
  letter-spacing: 2px;
}

.login-slogan {
  font-size: 14px;
  color: var(--text-color-secondary);
}

.login-tabs :deep(.el-tabs__nav-wrap::after) {
  display: none;
}

.login-options {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.login-btn {
  width: 100%;
  height: 44px;
  font-size: 16px;
  border-radius: var(--radius-md);
}

.switch-text {
  text-align: center;
  font-size: 13px;
  color: var(--text-color-secondary);
  margin-top: 16px;
}

.login-window-controls {
  position: fixed;
  top: 8px;
  right: 8px;
  display: flex;
  gap: 4px;
}

.win-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  color: rgba(255, 255, 255, 0.7);
  transition: all 0.15s;
}

.win-btn:hover {
  background-color: rgba(255, 255, 255, 0.15);
  color: #fff;
}

.win-btn-close:hover {
  background-color: #e81123;
  color: #fff;
}
</style>
