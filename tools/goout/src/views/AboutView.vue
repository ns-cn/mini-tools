<template>
  <div class="about-container">
    <!-- 背景字母 -->
    <div class="background-text">
      <div class="letter g">G</div>
      <div class="letter o1">O</div>
      <div class="letter o2">O</div>
      <div class="letter u">U</div>
      <div class="letter t">T</div>
    </div>

    <!-- 内容区域 -->
    <v-container class="content-container">
      <v-row justify="center">
        <v-col cols="12" md="8" lg="6">
          <v-card class="about-card" elevation="0">
            <v-card-title class="text-h3 font-weight-bold text-center mb-6">
              GoOut
            </v-card-title>
            <v-card-subtitle class="text-h6 text-center mb-8">
              探索你的生活圈
            </v-card-subtitle>
            <v-card-text class="text-body-1">
              <p class="mb-4">
                GoOut 是一个帮助你规划和记录生活的应用。无论是日常出行还是特别的旅程，
                我们都希望能让你的每一次出行变得更加轻松愉快。
              </p>
              <p class="mb-4">
                我们信，美好的生活不仅在于目的地，更在于一路上的点点滴滴。
              </p>
            </v-card-text>
            <Version class="version-info" />
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
import Version from '../components/Version.vue'
</script>

<style scoped>
.about-container {
  position: relative;
  min-height: 100vh;
  overflow: hidden;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e8eb 100%);
}

.background-text {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
  pointer-events: none;
}

.letter {
  position: absolute;
  font-size: 25vw;
  font-weight: 900;
  opacity: 0.03;
  color: var(--v-theme-primary);
  animation: float 20s ease-in-out infinite;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
}

.g {
  top: 10%;
  left: 5%;
  animation-delay: 0s;
  transform: rotate(-15deg);
}

.o1 {
  top: 30%;
  left: 25%;
  animation-delay: -4s;
  transform: rotate(10deg);
}

.o2 {
  top: 15%;
  right: 35%;
  animation-delay: -8s;
  transform: rotate(-5deg);
}

.u {
  bottom: 20%;
  left: 40%;
  animation-delay: -12s;
  transform: rotate(8deg);
}

.t {
  bottom: 25%;
  right: 10%;
  animation-delay: -16s;
  transform: rotate(-12deg);
}

@keyframes float {
  0%, 100% {
    transform: translateY(0) rotate(var(--rotation));
  }
  50% {
    transform: translateY(-20px) rotate(var(--rotation));
  }
}

.content-container {
  position: relative;
  z-index: 1;
  padding-top: 64px;
}

.about-card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 32px;
}

.version-info {
  margin-top: 32px;
  text-align: center;
}

/* 添加响应式设计 */
@media (max-width: 600px) {
  .letter {
    font-size: 40vw;
  }

  .about-card {
    padding: 16px;
  }
}

/* 添加悬停效果 */
.letter {
  transition: opacity 0.3s ease;
}

.about-container:hover .letter {
  opacity: 0.05;
}

/* 为每个字母设置不同的旋转角度 */
.g { --rotation: -15deg; }
.o1 { --rotation: 10deg; }
.o2 { --rotation: -5deg; }
.u { --rotation: 8deg; }
.t { --rotation: -12deg; }
</style>
