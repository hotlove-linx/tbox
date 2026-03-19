<template>
  <el-container class="admin-layout">
    <el-aside :width="isCollapsed ? '64px' : '220px'" class="admin-aside">
      <SideMenu :is-collapsed="isCollapsed" />
    </el-aside>
    <el-container>
      <el-header class="admin-header" height="60px">
        <TopNav :is-collapsed="isCollapsed" @toggle-collapse="isCollapsed = !isCollapsed" />
      </el-header>
      <el-main class="admin-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import SideMenu from './SideMenu.vue'
import TopNav from './TopNav.vue'

const isCollapsed = ref(false)
</script>

<style scoped>
.admin-layout {
  height: 100vh;
  overflow: hidden;
}

.admin-aside {
  background-color: #1d1e1f;
  transition: width 0.3s ease;
  overflow: hidden;
}

.admin-header {
  padding: 0;
  border-bottom: 1px solid #e4e7ed;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  z-index: 10;
}

.admin-main {
  background: #f0f2f5;
  overflow-y: auto;
  padding: 20px;
}
</style>
