<template>
  <div class="version-info">
    <div v-if="version">
      <span>版本: {{ version.version }}</span>
      <span class="divider">|</span>
      <span>构建时间: {{ version.buildTime }}</span>
    </div>
    <div v-else>
      <span>加载中...</span>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const version = ref(null)

const fetchVersion = async () => {
  try {
    const response = await fetch('/api/version')
    const data = await response.json()
    version.value = data
  } catch (error) {
    console.error('获取版本信息失败:', error)
  }
}

onMounted(() => {
  fetchVersion()
})
</script>

<style scoped>
.version-info {
  margin-top: 32px;
  padding: 8px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 4px;
  font-size: 12px;
  color: #666;
  text-align: center;
}

.divider {
  margin: 0 8px;
  color: #999;
}
</style>
