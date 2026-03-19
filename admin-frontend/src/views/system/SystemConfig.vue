<template>
  <div class="page-container">
    <!-- Basic Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">基础设置</span></template>
      <el-form :model="basicConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="平台名称">
          <el-input v-model="basicConfig.siteName" />
        </el-form-item>
        <el-form-item label="平台描述">
          <el-input v-model="basicConfig.siteDescription" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item label="ICP备案号">
          <el-input v-model="basicConfig.icp" />
        </el-form-item>
        <el-form-item label="联系邮箱">
          <el-input v-model="basicConfig.contactEmail" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="saveConfig('basic')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Registration Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">注册设置</span></template>
      <el-form :model="regConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="允许注册">
          <el-switch v-model="regConfig.allowRegister" />
        </el-form-item>
        <el-form-item label="邮箱验证">
          <el-switch v-model="regConfig.emailVerification" />
        </el-form-item>
        <el-form-item label="注册验证码">
          <el-switch v-model="regConfig.captchaRequired" />
        </el-form-item>
        <el-form-item label="默认用户空间">
          <el-select v-model="regConfig.defaultSpace" style="width: 100%">
            <el-option label="100 MB" value="100MB" />
            <el-option label="500 MB" value="500MB" />
            <el-option label="1 GB" value="1GB" />
            <el-option label="2 GB" value="2GB" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="saveConfig('register')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Upload Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">上传设置</span></template>
      <el-form :model="uploadConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="最大文件大小">
          <el-input-number v-model="uploadConfig.maxFileSize" :min="10" :max="500" :step="10" />
          <span style="margin-left: 8px; color: #909399;">MB</span>
        </el-form-item>
        <el-form-item label="允许的文件格式">
          <el-input v-model="uploadConfig.allowedFormats" placeholder="apk, ipa, exe" />
        </el-form-item>
        <el-form-item label="自动病毒扫描">
          <el-switch v-model="uploadConfig.autoVirusScan" />
        </el-form-item>
        <el-form-item label="上传需审核">
          <el-switch v-model="uploadConfig.requireAudit" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="saveConfig('upload')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Content Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">内容设置</span></template>
      <el-form :model="contentConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="评论需审核">
          <el-switch v-model="contentConfig.reviewAudit" />
        </el-form-item>
        <el-form-item label="敏感词过滤">
          <el-switch v-model="contentConfig.sensitiveFilter" />
        </el-form-item>
        <el-form-item label="每日下载限制">
          <el-input-number v-model="contentConfig.dailyDownloadLimit" :min="0" :max="1000" :step="10" />
          <span style="margin-left: 8px; color: #909399;">次 (0为不限制)</span>
        </el-form-item>
        <el-form-item label="排行榜展示数量">
          <el-input-number v-model="contentConfig.rankingCount" :min="5" :max="100" :step="5" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="saveConfig('content')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Email Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">邮件设置</span></template>
      <el-form :model="emailConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="SMTP服务器">
          <el-input v-model="emailConfig.smtpHost" />
        </el-form-item>
        <el-form-item label="SMTP端口">
          <el-input-number v-model="emailConfig.smtpPort" :min="1" :max="65535" />
        </el-form-item>
        <el-form-item label="发件人邮箱">
          <el-input v-model="emailConfig.fromEmail" />
        </el-form-item>
        <el-form-item label="邮箱密码/授权码">
          <el-input v-model="emailConfig.fromPassword" type="password" show-password />
        </el-form-item>
        <el-form-item label="SSL加密">
          <el-switch v-model="emailConfig.ssl" />
        </el-form-item>
        <el-form-item>
          <el-button @click="testEmail">发送测试邮件</el-button>
          <el-button type="primary" @click="saveConfig('email')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Storage Settings -->
    <el-card shadow="never" class="config-card">
      <template #header><span class="section-title">存储设置</span></template>
      <el-form :model="storageConfig" label-width="160px" style="max-width: 600px;">
        <el-form-item label="存储方式">
          <el-radio-group v-model="storageConfig.type">
            <el-radio value="local">本地存储</el-radio>
            <el-radio value="oss">阿里云OSS</el-radio>
            <el-radio value="cos">腾讯云COS</el-radio>
            <el-radio value="s3">AWS S3</el-radio>
          </el-radio-group>
        </el-form-item>
        <template v-if="storageConfig.type !== 'local'">
          <el-form-item label="Access Key">
            <el-input v-model="storageConfig.accessKey" />
          </el-form-item>
          <el-form-item label="Secret Key">
            <el-input v-model="storageConfig.secretKey" type="password" show-password />
          </el-form-item>
          <el-form-item label="Bucket">
            <el-input v-model="storageConfig.bucket" />
          </el-form-item>
          <el-form-item label="Region">
            <el-input v-model="storageConfig.region" />
          </el-form-item>
        </template>
        <el-form-item>
          <el-button type="primary" @click="saveConfig('storage')">保存</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { ElMessage } from 'element-plus'

const basicConfig = reactive({
  siteName: 'TBox',
  siteDescription: 'TBox - 您的软件工具箱',
  icp: '京ICP备XXXXXXXX号',
  contactEmail: 'support@tbox.com'
})

const regConfig = reactive({
  allowRegister: true,
  emailVerification: true,
  captchaRequired: true,
  defaultSpace: '500MB'
})

const uploadConfig = reactive({
  maxFileSize: 100,
  allowedFormats: 'apk, ipa, exe, dmg, zip',
  autoVirusScan: true,
  requireAudit: true
})

const contentConfig = reactive({
  reviewAudit: false,
  sensitiveFilter: true,
  dailyDownloadLimit: 0,
  rankingCount: 20
})

const emailConfig = reactive({
  smtpHost: 'smtp.example.com',
  smtpPort: 465,
  fromEmail: 'noreply@tbox.com',
  fromPassword: '',
  ssl: true
})

const storageConfig = reactive({
  type: 'local',
  accessKey: '',
  secretKey: '',
  bucket: '',
  region: ''
})

function saveConfig(group: string) {
  ElMessage.success(`${group} 配置已保存`)
}

function testEmail() {
  ElMessage.info('测试邮件已发送，请检查收件箱')
}
</script>

<style scoped>
.config-card { margin-bottom: 20px; }
.section-title { font-size: 16px; font-weight: 600; }
</style>
