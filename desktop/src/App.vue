<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

onMounted(() => {
  // Listen for navigation events from Electron main process
  if (window.electronAPI) {
    window.electronAPI.onNavigate((route: string) => {
      router.push(route)
    })
  }
})
</script>

<style>
#app {
  width: 100%;
  height: 100%;
}
</style>
