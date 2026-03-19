<template>
  <div class="search-filter">
    <slot />
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.search-filter {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
  margin-bottom: 16px;
  padding: 16px;
  background: #fff;
  border-radius: 8px;
}
</style>
