<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <el-icon :size="40" color="#409eff"><Box /></el-icon>
        <h1>TBox Admin</h1>
        <p>后台运营管理平台</p>
      </div>
      <el-form ref="formRef" :model="loginForm" :rules="rules" size="large">
        <el-form-item prop="email">
          <el-input v-model="loginForm.email" placeholder="邮箱地址" prefix-icon="Message">
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="密码"
            prefix-icon="Lock"
            show-password
            @keyup.enter="handleLogin"
          />
        </el-form-item>
        <el-form-item prop="captcha">
          <div class="captcha-row">
            <el-input v-model="loginForm.captcha" placeholder="验证码" prefix-icon="Key" />
            <div class="captcha-img" @click="refreshCaptcha">
              <canvas ref="captchaCanvas" width="120" height="40"></canvas>
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <div class="remember-row">
            <el-checkbox v-model="loginForm.remember">记住我</el-checkbox>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="login-btn" :loading="loading" @click="handleLogin">
            登 录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="login-footer">
      TBox Admin &copy; 2024
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePermissionStore } from '@/stores/permission'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const router = useRouter()
const authStore = useAuthStore()
const permissionStore = usePermissionStore()

const formRef = ref<FormInstance>()
const captchaCanvas = ref<HTMLCanvasElement>()
const loading = ref(false)
const captchaText = ref('')

const loginForm = reactive({
  email: '',
  password: '',
  captcha: '',
  remember: false
})

const rules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' }
  ],
  captcha: [
    { required: true, message: '请输入验证码', trigger: 'blur' }
  ]
}

function generateCaptcha() {
  const chars = 'ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789'
  let text = ''
  for (let i = 0; i < 4; i++) {
    text += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  captchaText.value = text
  drawCaptcha(text)
}

function drawCaptcha(text: string) {
  const canvas = captchaCanvas.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  ctx.fillStyle = '#f0f0f0'
  ctx.fillRect(0, 0, 120, 40)

  // Noise lines
  for (let i = 0; i < 4; i++) {
    ctx.strokeStyle = `rgba(${Math.random() * 255},${Math.random() * 255},${Math.random() * 255},0.5)`
    ctx.beginPath()
    ctx.moveTo(Math.random() * 120, Math.random() * 40)
    ctx.lineTo(Math.random() * 120, Math.random() * 40)
    ctx.stroke()
  }

  // Text
  for (let i = 0; i < text.length; i++) {
    ctx.font = `${20 + Math.random() * 8}px Arial`
    ctx.fillStyle = `rgb(${Math.random() * 100},${Math.random() * 100},${Math.random() * 150})`
    ctx.save()
    ctx.translate(25 * i + 15, 28)
    ctx.rotate((Math.random() - 0.5) * 0.4)
    ctx.fillText(text[i], 0, 0)
    ctx.restore()
  }

  // Noise dots
  for (let i = 0; i < 30; i++) {
    ctx.fillStyle = `rgba(${Math.random() * 255},${Math.random() * 255},${Math.random() * 255},0.5)`
    ctx.beginPath()
    ctx.arc(Math.random() * 120, Math.random() * 40, 1, 0, Math.PI * 2)
    ctx.fill()
  }
}

function refreshCaptcha() {
  generateCaptcha()
}

async function handleLogin() {
  const form = formRef.value
  if (!form) return
  await form.validate()

  if (loginForm.captcha.toLowerCase() !== captchaText.value.toLowerCase()) {
    ElMessage.error('验证码错误')
    refreshCaptcha()
    return
  }

  loading.value = true
  try {
    await authStore.login({
      email: loginForm.email,
      password: loginForm.password,
      captcha: loginForm.captcha
    })
    permissionStore.generateMenus(authStore.permissions)
    ElMessage.success('登录成功')
    router.push('/dashboard')
  } catch (e) {
    refreshCaptcha()
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  nextTick(() => {
    generateCaptcha()
  })
})
</script>

<style scoped>
.login-container {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 420px;
  padding: 40px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-header h1 {
  margin: 12px 0 4px;
  font-size: 28px;
  color: #303133;
}

.login-header p {
  color: #909399;
  font-size: 14px;
}

.captcha-row {
  display: flex;
  gap: 12px;
  width: 100%;
}

.captcha-row .el-input {
  flex: 1;
}

.captcha-img {
  cursor: pointer;
  border-radius: 4px;
  overflow: hidden;
  flex-shrink: 0;
}

.captcha-img canvas {
  display: block;
}

.remember-row {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.login-btn {
  width: 100%;
}

.login-footer {
  margin-top: 24px;
  color: rgba(255, 255, 255, 0.7);
  font-size: 13px;
}
</style>
