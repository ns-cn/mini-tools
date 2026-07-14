<template>
  <v-container fluid class="fill-height bg-grey-lighten-3">
    <v-row align="center" justify="center" class="fill-height">
      <!-- 左侧图片区域（仅在大屏幕显示） -->
      <v-col cols="12" md="6" class="d-none d-md-flex align-center justify-center" style="height: 100vh">
        <v-img
          src="/login-bg.jpg"
          cover
          height="100%"
          class="login-bg"
        >
          <div class="d-flex fill-height align-center justify-center">
            <div class="text-center">
              <h1 class="text-h2 font-weight-bold text-white">GoOut</h1>
              <p class="text-h5 text-white">探索你的生活圈</p>
            </div>
          </div>
        </v-img>
      </v-col>

      <!-- 登录表单区域 -->
      <v-col cols="12" md="6" class="pa-md-12 pa-4">
        <v-card class="mx-auto" max-width="400" elevation="8">
          <v-card-text class="pt-8">
            <h2 class="text-h4 font-weight-bold text-center mb-8">欢迎回来</h2>

            <v-form @submit.prevent="handleLogin" ref="form" v-model="valid">
              <v-text-field
                v-model="loginForm.username"
                label="用户名"
                prepend-inner-icon="mdi-account"
                variant="outlined"
                class="mb-4"
                hide-details
                :rules="[v => !!v || '用户名是必填项']"
              />

              <v-text-field
                v-model="loginForm.password"
                label="密码"
                prepend-inner-icon="mdi-lock"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                variant="outlined"
                class="mb-8"
                hide-details
                :type="showPassword ? 'text' : 'password'"
                @click:append-inner="showPassword = !showPassword"
                :rules="[v => !!v || '密码是必填项']"
              />

              <v-btn
                block
                color="primary"
                size="large"
                type="submit"
                :loading="loading"
                class="mb-4"
                :disabled="!valid"
              >
                登录
              </v-btn>
            </v-form>

            <div class="text-center">
              <span>还没有账号？</span>
              <v-btn text color="primary" @click="goToRegister">立即注册</v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const form = ref(null)
const loading = ref(false)
const showPassword = ref(false)
const valid = ref(false)

const loginForm = reactive({
  username: '',
  password: ''
})

const handleLogin = async () => {
  if (!valid.value) return

  loading.value = true
  try {
    const response = await fetch('/api/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: loginForm.username,
        password: loginForm.password
      })
    })

    const result = await response.json()
    if (!response.ok || result.code !== 0) {
      throw new Error(result.message || '登录失败')
    }

    const data = result.data
    localStorage.setItem('token', data.token)
    localStorage.setItem('username', data.username)
    localStorage.setItem('nickname', data.nickname)

    ElMessage.success(result.message)
    router.push('/dashboard')
  } catch (error) {
    ElMessage.error(error.message)
  } finally {
    loading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>

<style scoped>
.login-bg::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
}

:deep(.v-field__input) {
  font-size: 16px;
  padding: 12px;
}

:deep(.v-field__prepend-inner) {
  padding-inline-start: 12px;
}

@media (max-width: 600px) {
  :deep(.v-container) {
    padding: 12px;
  }
}
</style>
