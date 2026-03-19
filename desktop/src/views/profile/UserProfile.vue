<template>
  <div class="profile-page">
    <div class="page-header">
      <h2 class="page-title">个人中心</h2>
    </div>

    <div class="profile-content">
      <!-- Avatar & Basic Info -->
      <div class="profile-card">
        <div class="avatar-section">
          <el-avatar :size="80" :src="userAvatar || undefined" class="user-avatar">
            {{ authStore.userName?.charAt(0) || 'U' }}
          </el-avatar>
          <el-upload
            :auto-upload="false"
            :show-file-list="false"
            accept="image/*"
            :on-change="handleAvatarChange"
          >
            <el-button size="small" type="primary" text>更换头像</el-button>
          </el-upload>
        </div>

        <el-form label-width="100px" class="profile-form" label-position="left">
          <el-form-item label="昵称">
            <div class="editable-field">
              <template v-if="!editingNickname">
                <span>{{ nickname }}</span>
                <el-button text type="primary" size="small" @click="editingNickname = true">
                  <el-icon><Edit /></el-icon>
                </el-button>
              </template>
              <template v-else>
                <el-input v-model="nickname" style="width: 200px" />
                <el-button type="primary" size="small" @click="saveNickname">保存</el-button>
                <el-button size="small" @click="editingNickname = false; nickname = authStore.userName">取消</el-button>
              </template>
            </div>
          </el-form-item>
          <el-form-item label="邮箱">
            <span class="readonly-text">{{ authStore.userEmail || 'user@example.com' }}</span>
          </el-form-item>
          <el-form-item label="注册时间">
            <span class="readonly-text">{{ authStore.userInfo?.createdAt || '2024-01-15 10:30:00' }}</span>
          </el-form-item>
        </el-form>
      </div>

      <!-- Change Password -->
      <div class="profile-card">
        <h3 class="card-title">修改密码</h3>
        <el-form ref="passwordFormRef" :model="passwordForm" :rules="passwordRules" label-width="100px" label-position="left" style="max-width: 500px">
          <el-form-item label="原密码" prop="oldPassword">
            <el-input v-model="passwordForm.oldPassword" type="password" show-password placeholder="输入当前密码" />
          </el-form-item>
          <el-form-item label="新密码" prop="newPassword">
            <el-input v-model="passwordForm.newPassword" type="password" show-password placeholder="8位以上，含字母和数字" />
          </el-form-item>
          <el-form-item label="确认密码" prop="confirmPassword">
            <el-input v-model="passwordForm.confirmPassword" type="password" show-password placeholder="再次输入新密码" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="changingPassword" @click="handleChangePassword">确认修改</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- Account Actions -->
      <div class="profile-card">
        <h3 class="card-title">账户操作</h3>
        <div class="account-actions">
          <div class="action-row">
            <div>
              <div class="action-label">退出登录</div>
              <div class="action-desc">退出当前账户，返回登录页面</div>
            </div>
            <el-button @click="handleLogout">退出登录</el-button>
          </div>
          <el-divider />
          <div class="action-row">
            <div>
              <div class="action-label danger-text">注销账户</div>
              <div class="action-desc">永久删除账户及所有数据，此操作不可恢复</div>
            </div>
            <el-button type="danger" @click="handleDeleteAccount">注销账户</el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance, FormRules, UploadFile } from 'element-plus'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const nickname = ref(authStore.userName || '用户')
const userAvatar = ref(authStore.userAvatar)
const editingNickname = ref(false)
const changingPassword = ref(false)

const passwordFormRef = ref<FormInstance>()
const passwordForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const passwordRules: FormRules = {
  oldPassword: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 8, message: '密码长度不能少于8位', trigger: 'blur' },
    { pattern: /(?=.*[a-zA-Z])(?=.*\d)/, message: '密码必须包含字母和数字', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    {
      validator: (_rule: any, value: string, callback: any) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error('两次输入密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

function handleAvatarChange(file: UploadFile) {
  if (file.raw) {
    const url = URL.createObjectURL(file.raw)
    userAvatar.value = url
    ElMessage.success('头像更新成功')
  }
}

function saveNickname() {
  authStore.updateUserInfo({ nickname: nickname.value })
  editingNickname.value = false
  ElMessage.success('昵称已更新')
}

async function handleChangePassword() {
  if (!passwordFormRef.value) return
  await passwordFormRef.value.validate(async (valid) => {
    if (!valid) return
    changingPassword.value = true
    setTimeout(() => {
      changingPassword.value = false
      ElMessage.success('密码修改成功')
      passwordForm.oldPassword = ''
      passwordForm.newPassword = ''
      passwordForm.confirmPassword = ''
    }, 1000)
  })
}

function handleLogout() {
  ElMessageBox.confirm('确定要退出登录吗？', '退出登录', {
    type: 'info'
  }).then(() => {
    authStore.logout()
  }).catch(() => {})
}

function handleDeleteAccount() {
  ElMessageBox.confirm(
    '注销账户后，所有数据将被永久删除，此操作不可恢复。确定要继续吗？',
    '注销账户',
    { type: 'error', confirmButtonText: '确认注销', cancelButtonText: '取消' }
  ).then(() => {
    ElMessage.success('账户已注销')
    authStore.logout()
  }).catch(() => {})
}
</script>

<style scoped>
.profile-page {
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 24px;
}

.page-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.profile-card {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 28px;
  margin-bottom: 20px;
  border: 1px solid var(--border-color-light);
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 20px;
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  margin-bottom: 24px;
}

.user-avatar {
  font-size: 28px;
}

.profile-form :deep(.el-form-item) {
  margin-bottom: 16px;
}

.editable-field {
  display: flex;
  align-items: center;
  gap: 8px;
}

.readonly-text {
  color: var(--text-color-secondary);
}

.account-actions {
  max-width: 600px;
}

.action-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.action-label {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-color-primary);
  margin-bottom: 4px;
}

.action-desc {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.danger-text {
  color: var(--color-danger);
}
</style>
