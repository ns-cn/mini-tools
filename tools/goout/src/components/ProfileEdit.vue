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
        <span>个人信息</span>
        <v-spacer></v-spacer>
        <v-btn icon="mdi-close" variant="text" density="compact" @click="close"></v-btn>
      </v-card-title>

      <v-divider></v-divider>

      <div class="pa-4">
        <v-form ref="profileForm">
          <v-text-field
            v-model="profile.nickname"
            label="昵称"
            variant="outlined"
            density="comfortable"
            class="mb-3"
            :rules="[v => !!v || '昵称不能为空']"
          ></v-text-field>

          <v-expansion-panels variant="accordion" class="mb-3">
            <v-expansion-panel title="修改密码">
              <template v-slot:text>
                <v-text-field
                  v-model="profile.oldPassword"
                  label="当前密码"
                  type="password"
                  variant="outlined"
                  density="comfortable"
                  class="mb-3"
                  :rules="passwordRules"
                ></v-text-field>

                <v-text-field
                  v-model="profile.newPassword"
                  label="新密码"
                  type="password"
                  variant="outlined"
                  density="comfortable"
                  class="mb-3"
                  :rules="passwordRules"
                ></v-text-field>

                <v-text-field
                  v-model="profile.confirmPassword"
                  label="确认新密码"
                  type="password"
                  variant="outlined"
                  density="comfortable"
                  class="mb-3"
                  :rules="[
                    ...passwordRules,
                    v => v === profile.newPassword || '两次输入的密码不一致'
                  ]"
                ></v-text-field>
              </template>
            </v-expansion-panel>
          </v-expansion-panels>

          <v-btn
            block
            color="primary"
            :loading="loading"
            @click="saveProfile"
          >
            保存修改
          </v-btn>
        </v-form>
      </div>
    </v-card>
  </v-dialog>
</template>

<script setup>
// ... 这里使用之前 Settings.vue 中的个人信息相关代码
</script>

<style scoped>
/* ... 使用相同的样式 */
</style>
