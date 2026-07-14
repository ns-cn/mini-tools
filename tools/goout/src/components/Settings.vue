<template>
  <v-dialog
    v-model="show"
    transition="dialog-bottom-transition"
    :scrim="false"
    content-class="settings-dialog"
    width="100%"
  >
    <v-card class="settings-card" rounded="t-lg">
      <div
        class="handle-wrapper"
        @touchstart="handleTouchStart"
        @touchmove="handleTouchMove"
        @touchend="handleTouchEnd"
      >
        <div class="handle"></div>
      </div>

      <v-card-title class="d-flex align-center px-4 py-2">
        <span>主题设置</span>
        <v-spacer></v-spacer>
        <v-btn icon="mdi-close" variant="text" density="compact" @click="close"></v-btn>
      </v-card-title>

      <v-divider></v-divider>

      <div class="pa-4">
        <div class="theme-settings">
          <h3 class="text-subtitle-1 font-weight-medium mb-3">主题颜色</h3>
          <div class="d-flex flex-wrap gap-3">
            <v-btn
              v-for="color in themeColors"
              :key="color.name"
              :color="color.value"
              icon
              size="large"
              class="mr-2 mb-2"
              :class="{ 'theme-selected': currentTheme === color.value }"
              @click="setTheme(color.value)"
            >
              <v-icon v-if="currentTheme === color.value">mdi-check</v-icon>
            </v-btn>
          </div>
        </div>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, defineProps, defineEmits, computed } from 'vue'

const props = defineProps({
  modelValue: Boolean,
})

const emit = defineEmits(['update:modelValue', 'theme-changed'])

const show = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const themeColors = [
  { name: '天空蓝', value: '#4A90E2' },
  { name: '薄荷绿', value: '#42B983' },
  { name: '珊瑚粉', value: '#FF7F7F' },
  { name: '薰衣草', value: '#9B7FFF' },
  { name: '柠檬黄', value: '#FFB900' },
  { name: '海洋蓝', value: '#48C0B6' },
]

// 从本地存储加载主题
const currentTheme = ref(localStorage.getItem('theme') || '#1867C0')

const close = () => {
  show.value = false
}

const setTheme = async (color) => {
  try {
    const token = localStorage.getItem('token')
    const response = await fetch('/api/users/profile', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify({
        theme: color,
        nickname: localStorage.getItem('nickname') // 需要保持昵称不变
      })
    })

    const result = await response.json()
    if (!response.ok || result.code !== 0) {
      throw new Error(result.message || '保存失败')
    }

    // 更新本地主题
    currentTheme.value = color
    localStorage.setItem('theme', color)
    emit('theme-changed', color)
  } catch (error) {
    ElMessage.error('保存主题设置失败')
  }
}

// 拖拽相关代码保持不变...
const touchStartY = ref(0)
const touchMoveY = ref(0)
const isDragging = ref(false)

const handleTouchStart = (event) => {
  touchStartY.value = event.touches[0].clientY
  isDragging.value = true
}

const handleTouchMove = (event) => {
  if (!isDragging.value) return

  touchMoveY.value = event.touches[0].clientY
  const deltaY = touchMoveY.value - touchStartY.value

  if (deltaY > 0) {
    event.preventDefault()
    const element = event.target.closest('.settings-card')
    if (element) {
      element.style.transform = `translateY(${deltaY}px)`
      element.style.transition = 'none'
    }
  }
}

const handleTouchEnd = (event) => {
  if (!isDragging.value) return

  const deltaY = touchMoveY.value - touchStartY.value
  const element = event.target.closest('.settings-card')

  if (element) {
    element.style.transition = 'transform 0.3s cubic-bezier(0.4, 0, 0.2, 1)'

    if (deltaY > 100) {
      element.style.transform = `translateY(100%)`
      setTimeout(() => {
        show.value = false
        element.style.transform = ''
        element.style.transition = ''
      }, 300)
    } else {
      element.style.transform = ''
    }
  }

  isDragging.value = false
  touchStartY.value = 0
  touchMoveY.value = 0
}
</script>

<style scoped>
.settings-dialog {
  position: fixed !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  margin: 0 !important;
  background: transparent !important;
}

.settings-card {
  position: relative !important;
  margin: 0 !important;
  height: 50vh !important;
  max-height: 50vh !important;
  overflow-y: auto;
  border-radius: 20px 20px 0 0 !important;
  display: flex;
  flex-direction: column;
  will-change: transform;
  touch-action: none;
}

.handle-wrapper {
  display: flex;
  justify-content: center;
  padding: 8px;
  cursor: grab;
  touch-action: none;
  user-select: none;
}

.handle {
  width: 36px;
  height: 4px;
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 2px;
}

.settings-tabs {
  position: sticky;
  top: 0;
  background: white;
  z-index: 1;
}

.theme-settings .v-btn {
  border-radius: 50%;
  width: 44px;
  height: 44px;
}

.theme-selected {
  border: 2px solid white;
  box-shadow: 0 0 0 2px currentColor;
}

.gap-3 {
  gap: 12px;
}

:deep(.dialog-bottom-transition-enter-active),
:deep(.dialog-bottom-transition-leave-active) {
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

:deep(.dialog-bottom-transition-enter-from),
:deep(.dialog-bottom-transition-leave-to) {
  transform: translateY(100%);
}

:deep(.v-overlay__scrim) {
  opacity: 0.3 !important;
}

:deep(.v-card) {
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 0.2) transparent;
}

:deep(.v-card::-webkit-scrollbar) {
  width: 6px;
}

:deep(.v-card::-webkit-scrollbar-track) {
  background: transparent;
}

:deep(.v-card::-webkit-scrollbar-thumb) {
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 3px;
}

:deep(.v-overlay__content) {
  position: fixed !important;
  bottom: 0 !important;
  margin: 0 !important;
  width: 100%;
}

:deep(.v-dialog) {
  margin: 0 !important;
  position: fixed !important;
  bottom: 0 !important;
}
</style>
