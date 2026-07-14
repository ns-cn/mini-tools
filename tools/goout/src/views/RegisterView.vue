<template>
  <v-container fluid class="fill-height bg-grey-lighten-3">
    <v-row align="center" justify="center" class="fill-height">
      <v-col cols="12" md="6" class="pa-md-12 pa-4">
        <v-card class="mx-auto" max-width="400" elevation="8">
          <v-card-text class="pt-8">
            <h2 class="text-h4 font-weight-bold text-center mb-8">注册新账号</h2>

            <v-form @submit.prevent="handleRegister" ref="form" v-model="valid">
              <v-text-field
                v-model="registerForm.username"
                label="用户名"
                prepend-inner-icon="mdi-account"
                variant="outlined"
                class="mb-4"
                hide-details
                :rules="[v => !!v || '用户名是必填项', v => v.length >= 3 || '用户名至少需要3个字符']"
              />

              <v-text-field
                v-model="registerForm.nickname"
                label="昵称"
                prepend-inner-icon="mdi-account-circle"
                variant="outlined"
                class="mb-4"
                hide-details
                :rules="[v => !!v || '昵称是必填项', v => v.length >= 2 || '昵称至少需要2个字符']"
              />

              <v-text-field
                v-model="registerForm.password"
                label="密码"
                prepend-inner-icon="mdi-lock"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                variant="outlined"
                class="mb-4"
                hide-details
                :type="showPassword ? 'text' : 'password'"
                @click:append-inner="showPassword = !showPassword"
                :rules="[v => !!v || '密码是必填项', v => v.length >= 6 || '密码至少需要6个字符']"
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
                注册
              </v-btn>
            </v-form>

            <div class="text-center">
              <span>已有账号？</span>
              <v-btn text color="primary" @click="goToLogin">立即登录</v-btn>
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

const registerForm = reactive({
  username: '',
  nickname: '',
  password: ''
})

const handleRegister = async () => {
  if (!valid.value) return

  loading.value = true
  try {
    const response = await fetch('/api/users/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(registerForm)
    })

    const result = await response.json()
    if (!response.ok) {
      throw new Error(result.error || '注册失败')
    }

    ElMessage.success('注册成功，请登录')
    router.push('/login')
  } catch (error) {
    ElMessage.error(error.message)
  } finally {
    loading.value = false
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<style scoped>
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
