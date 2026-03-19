<template>
  <el-dialog
    v-model="visible"
    title="上传软件"
    width="680px"
    destroy-on-close
    @close="handleClose"
  >
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" label-position="top">
      <!-- File upload area -->
      <el-form-item label="软件文件" prop="file">
        <el-upload
          drag
          :auto-upload="false"
          :limit="1"
          :on-change="handleFileChange"
          accept=".exe,.msi,.zip,.dmg,.deb,.AppImage"
          class="upload-dragger"
        >
          <el-icon class="el-icon--upload" :size="40"><UploadFilled /></el-icon>
          <div class="el-upload__text">
            拖拽文件到此处，或 <em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip">支持 .exe, .msi, .zip, .dmg, .deb 格式，文件大小不超过 2GB</div>
          </template>
        </el-upload>
      </el-form-item>

      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item label="软件名称" prop="name">
            <el-input v-model="form.name" placeholder="输入软件名称" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="版本号" prop="version">
            <el-input v-model="form.version" placeholder="例如 1.0.0" />
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item label="分类" prop="categoryId">
        <el-select v-model="form.categoryId" placeholder="选择分类" style="width: 100%">
          <el-option label="办公工具" :value="1" />
          <el-option label="开发工具" :value="2" />
          <el-option label="影音娱乐" :value="3" />
          <el-option label="系统工具" :value="4" />
          <el-option label="安全防护" :value="5" />
          <el-option label="设计创意" :value="6" />
          <el-option label="网络通讯" :value="7" />
        </el-select>
      </el-form-item>

      <el-form-item label="简介" prop="shortDesc">
        <el-input v-model="form.shortDesc" placeholder="一句话描述你的软件" maxlength="100" show-word-limit />
      </el-form-item>

      <el-form-item label="详细描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="4"
          placeholder="详细介绍你的软件功能和特点"
          maxlength="2000"
          show-word-limit
        />
      </el-form-item>

      <el-row :gutter="16">
        <el-col :span="12">
          <el-form-item label="软件图标">
            <el-upload
              :auto-upload="false"
              :limit="1"
              accept="image/*"
              list-type="picture-card"
              :on-change="handleIconChange"
            >
              <el-icon><Plus /></el-icon>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="可见性" prop="visibility">
            <el-radio-group v-model="form.visibility">
              <el-radio value="public">公开</el-radio>
              <el-radio value="private">私有</el-radio>
              <el-radio value="link">链接分享</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item label="软件截图（最多5张）">
        <el-upload
          :auto-upload="false"
          :limit="5"
          accept="image/*"
          list-type="picture-card"
          :on-change="handleScreenshotChange"
          multiple
        >
          <el-icon><Plus /></el-icon>
        </el-upload>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" :loading="uploading" @click="handleSubmit">
        {{ uploading ? '上传中...' : '确认上传' }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance, FormRules, UploadFile } from 'element-plus'
import { ElMessage } from 'element-plus'

const visible = defineModel<boolean>({ default: false })
const emit = defineEmits(['uploaded'])

const formRef = ref<FormInstance>()
const uploading = ref(false)

const form = reactive({
  name: '',
  version: '',
  categoryId: undefined as number | undefined,
  shortDesc: '',
  description: '',
  visibility: 'public' as 'public' | 'private' | 'link',
  file: null as File | null,
  icon: null as File | null,
  screenshots: [] as File[]
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入软件名称', trigger: 'blur' }],
  version: [{ required: true, message: '请输入版本号', trigger: 'blur' }],
  categoryId: [{ required: true, message: '请选择分类', trigger: 'change' }],
  shortDesc: [{ required: true, message: '请输入简介', trigger: 'blur' }],
  description: [{ required: true, message: '请输入详细描述', trigger: 'blur' }],
  visibility: [{ required: true, message: '请选择可见性', trigger: 'change' }]
}

function handleFileChange(file: UploadFile) {
  form.file = file.raw || null
}

function handleIconChange(file: UploadFile) {
  form.icon = file.raw || null
}

function handleScreenshotChange(file: UploadFile) {
  if (file.raw) {
    form.screenshots.push(file.raw)
  }
}

async function handleSubmit() {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    uploading.value = true
    // Simulate upload
    setTimeout(() => {
      uploading.value = false
      ElMessage.success('软件上传成功，等待审核')
      emit('uploaded')
      visible.value = false
    }, 2000)
  })
}

function handleClose() {
  form.name = ''
  form.version = ''
  form.categoryId = undefined
  form.shortDesc = ''
  form.description = ''
  form.visibility = 'public'
  form.file = null
  form.icon = null
  form.screenshots = []
}
</script>

<style scoped>
.upload-dragger {
  width: 100%;
}

.upload-dragger :deep(.el-upload-dragger) {
  width: 100%;
}

.upload-dragger :deep(.el-upload) {
  width: 100%;
}
</style>
