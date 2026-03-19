<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openDialog()"><el-icon><Plus /></el-icon>新增 Banner</el-button>
      </div>
      <el-table :data="tableData" stripe border row-key="id">
        <el-table-column width="60" align="center" label="排序">
          <template #default>
            <el-icon class="drag-handle" :size="16"><Rank /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="Banner预览" width="200">
          <template #default="{ row }">
            <div class="banner-thumb" :style="{ background: row.bgColor }">
              <span>{{ row.title }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题" width="180" />
        <el-table-column prop="linkedSoftware" label="关联软件" width="150" />
        <el-table-column label="展示时间段" width="220">
          <template #default="{ row }">
            {{ row.startTime }} ~ {{ row.endTime }}
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)" size="small">{{ statusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openDialog(row)">编辑</el-button>
            <el-popconfirm title="确定删除此Banner？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingItem ? '编辑 Banner' : '新增 Banner'" width="600px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="Banner图片" prop="image">
          <el-upload action="#" :auto-upload="false" :show-file-list="false" accept="image/*">
            <div class="upload-area">
              <el-icon :size="28"><Plus /></el-icon>
              <span>点击上传</span>
            </div>
          </el-upload>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="Banner标题" />
        </el-form-item>
        <el-form-item label="副标题">
          <el-input v-model="form.subtitle" placeholder="Banner副标题" />
        </el-form-item>
        <el-form-item label="关联软件">
          <el-select v-model="form.linkedSoftware" filterable placeholder="搜索并选择软件" style="width: 100%">
            <el-option label="超级文件管理器" value="超级文件管理器" />
            <el-option label="极速浏览器" value="极速浏览器" />
            <el-option label="智能笔记" value="智能笔记" />
          </el-select>
        </el-form-item>
        <el-form-item label="跳转类型">
          <el-select v-model="form.jumpType" style="width: 100%">
            <el-option label="软件详情" value="software" />
            <el-option label="外部链接" value="url" />
            <el-option label="无跳转" value="none" />
          </el-select>
        </el-form-item>
        <el-form-item label="展示时间" prop="timeRange">
          <el-date-picker v-model="form.timeRange" type="datetimerange" range-separator="至" start-placeholder="开始" end-placeholder="结束" style="width: 100%" />
        </el-form-item>
        <el-form-item label="排序权重">
          <el-input-number v-model="form.sort" :min="0" :max="999" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const dialogVisible = ref(false)
const editingItem = ref<any>(null)
const formRef = ref<FormInstance>()

const form = reactive({
  title: '', subtitle: '', linkedSoftware: '', jumpType: 'software',
  timeRange: null as any, sort: 0, image: ''
})

const rules: FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  timeRange: [{ required: true, message: '请选择展示时间', trigger: 'change' }]
}

const tableData = ref([
  { id: 1, title: '新春特惠 精品推荐', bgColor: '#ff6b6b', linkedSoftware: '超级文件管理器', startTime: '2024-03-01', endTime: '2024-03-31', status: 'active', sort: 1 },
  { id: 2, title: '效率工具专区', bgColor: '#4ecdc4', linkedSoftware: '智能笔记', startTime: '2024-03-10', endTime: '2024-04-10', status: 'active', sort: 2 },
  { id: 3, title: '游戏精选', bgColor: '#45b7d1', linkedSoftware: '', startTime: '2024-04-01', endTime: '2024-04-30', status: 'scheduled', sort: 3 },
  { id: 4, title: '旧活动Banner', bgColor: '#96ceb4', linkedSoftware: '极速浏览器', startTime: '2024-01-01', endTime: '2024-02-28', status: 'expired', sort: 4 }
])

function statusType(s: string) {
  const map: Record<string, any> = { active: 'success', scheduled: 'warning', expired: 'info' }
  return map[s] || 'info'
}

function statusText(s: string) {
  const map: Record<string, string> = { active: '展示中', scheduled: '待展示', expired: '已过期' }
  return map[s] || s
}

function openDialog(item?: any) {
  editingItem.value = item || null
  if (item) {
    Object.assign(form, { title: item.title, subtitle: '', linkedSoftware: item.linkedSoftware, jumpType: 'software', timeRange: null, sort: item.sort })
  } else {
    Object.assign(form, { title: '', subtitle: '', linkedSoftware: '', jumpType: 'software', timeRange: null, sort: 0 })
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  await formRef.value?.validate()
  ElMessage.success(editingItem.value ? '编辑成功' : '新增成功')
  dialogVisible.value = false
}

function handleDelete(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.action-bar { margin-bottom: 16px; }
.drag-handle { cursor: move; color: #909399; }
.banner-thumb {
  width: 160px; height: 60px; border-radius: 6px;
  display: flex; align-items: center; justify-content: center;
  color: #fff; font-size: 13px; font-weight: 500;
}
.upload-area {
  width: 200px; height: 80px;
  border: 2px dashed #dcdfe6; border-radius: 8px;
  display: flex; flex-direction: column; align-items: center; justify-content: center;
  gap: 4px; color: #909399; cursor: pointer;
}
.upload-area:hover { border-color: #409eff; color: #409eff; }
</style>
