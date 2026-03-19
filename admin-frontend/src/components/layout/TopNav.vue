<template>
  <div class="top-nav">
    <div class="left-section">
      <el-icon class="collapse-btn" @click="$emit('toggleCollapse')" :size="20">
        <Fold v-if="!isCollapsed" />
        <Expand v-else />
      </el-icon>
      <el-breadcrumb separator="/">
        <el-breadcrumb-item v-if="route.meta.parent">{{ route.meta.parent }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ route.meta.title }}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="right-section">
      <el-input
        v-model="searchText"
        placeholder="全局搜索..."
        prefix-icon="Search"
        class="global-search"
        clearable
        size="default"
        style="width: 200px"
      />
      <el-badge :value="8" :max="99" class="notification-badge">
        <el-icon :size="20" class="action-icon"><Bell /></el-icon>
      </el-badge>
      <el-dropdown trigger="click" @command="handleCommand">
        <div class="admin-info">
          <el-avatar :size="32" :src="adminInfo?.avatar || undefined">
            {{ adminInfo?.name?.[0] || 'A' }}
          </el-avatar>
          <span class="admin-name">{{ adminInfo?.name || '管理员' }}</span>
          <el-tag size="small" type="primary" effect="plain">{{ adminInfo?.role || '管理员' }}</el-tag>
          <el-icon><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile"><el-icon><User /></el-icon>个人设置</el-dropdown-item>
            <el-dropdown-item command="password"><el-icon><Lock /></el-icon>修改密码</el-dropdown-item>
            <el-dropdown-item divided command="logout"><el-icon><SwitchButton /></el-icon>退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>

    <!-- Change Password Dialog -->
    <el-dialog v-model="passwordDialogVisible" title="修改密码" width="420px">
      <el-form :model="passwordForm" :rules="passwordRules" ref="passwordFormRef" label-width="80px">
        <el-form-item label="旧密码" prop="oldPassword">
          <el-input v-model="passwordForm.oldPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input v-model="passwordForm.newPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="passwordForm.confirmPassword" type="password" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="passwordDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitPasswordChange">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

defineProps<{ isCollapsed: boolean }>()
defineEmits(['toggleCollapse'])

const route = useRoute()
const authStore = useAuthStore()
const adminInfo = computed(() => authStore.adminInfo)

const searchText = ref('')
const passwordDialogVisible = ref(false)
const passwordFormRef = ref<FormInstance>()
const passwordForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const passwordRules: FormRules = {
  oldPassword: [{ required: true, message: '请输入旧密码', trigger: 'blur' }],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入新密码', trigger: 'blur' },
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

function handleCommand(command: string) {
  switch (command) {
    case 'profile':
      ElMessage.info('个人设置功能开发中')
      break
    case 'password':
      passwordDialogVisible.value = true
      break
    case 'logout':
      ElMessageBox.confirm('确定要退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        authStore.logout()
      })
      break
  }
}

async function submitPasswordChange() {
  const form = passwordFormRef.value
  if (!form) return
  await form.validate()
  ElMessage.success('密码修改成功')
  passwordDialogVisible.value = false
}
</script>

<style scoped>
.top-nav {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.left-section {
  display: flex;
  align-items: center;
  gap: 16px;
}

.collapse-btn {
  cursor: pointer;
  color: #606266;
  transition: color 0.2s;
}

.collapse-btn:hover {
  color: #409eff;
}

.right-section {
  display: flex;
  align-items: center;
  gap: 20px;
}

.action-icon {
  cursor: pointer;
  color: #606266;
}

.action-icon:hover {
  color: #409eff;
}

.notification-badge {
  cursor: pointer;
}

.admin-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.admin-name {
  font-size: 14px;
  color: #303133;
}
</style>
