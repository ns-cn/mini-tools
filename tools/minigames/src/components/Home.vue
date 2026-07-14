<template>
  <div class="home">
    <h1 class="title">
      <span>M</span>
      <span>i</span>
      <span>n</span>
      <span>i</span>
      <span> </span>
      <span>G</span>
      <span>a</span>
      <span>m</span>
      <span>e</span>
      <span>s</span>
    </h1>
    <div class="game-list">
      <div class="game-item">
        <div class="game-card" @click="(event) => startGame('2048', null, event)">
          <div class="game-icon game-icon-2048">
            <div class="number">20</div>
            <div class="number">48</div>
          </div>
          <div class="game-info">
            <h2>2048</h2>
            <p>经典数字益智游戏</p>
          </div>
        </div>
      </div>
      <div class="game-item">
        <div class="game-card" @click="(event) => startGame('FlappyBird', null, event)">
          <div class="game-icon">🐦</div>
          <div class="game-info">
            <h2>Flappy Bird</h2>
            <p>考验反应的小鸟飞行游戏</p>
          </div>
        </div>
      </div>
      <div class="game-item">
        <div class="game-card" @click="(event) => startGame('GoDown', null, event)">
          <div class="game-icon">🏃</div>
          <div class="game-info">
            <h2>GoDown</h2>
            <p>是男人就下100层</p>
          </div>
        </div>
      </div>
      <!-- 可以继续添加更多游戏入口 -->
    </div>
    <div class="version">v{{ version }}</div>
  </div>
</template>

<script>
import { VERSION } from '../version.js';

export default {
  name: 'HomeView',
  data() {
    return {
      version: VERSION
    }
  },
  methods: {
    startGame(game, navigate, event) {
      // 阻止默认的导航行为
      if (event) {
        event.preventDefault();
      }

      // 获取卡片元素
      const card = event ? event.currentTarget : null;

      if (card) {
        const rect = card.getBoundingClientRect();
        const centerX = rect.left + rect.width / 2;
        const centerY = rect.top + rect.height / 2;

        this.$root.$emit('start-game-animation', { game, centerX, centerY });
        
        // 使用 setTimeout 来延迟导航，给动画一些时间开始
        setTimeout(() => {
          this.$router.push({ name: game });
        }, 500);
      } else {
        // 如果没有事件对象，直接导航
        this.$router.push({ name: game });
      }
    }
  }
}
</script>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start; /* 改为 flex-start */
  min-height: 100vh;
  background: linear-gradient(135deg, #e0f7fa, #b2ebf2);
  padding: 40px 20px; /* 增加顶部内边距 */
  box-sizing: border-box;
  overflow-y: auto;
  overflow-x: hidden;
  width: 100vw;
}

h1 {
  margin-top: 20px; /* 添加顶部外边距 */
  margin-bottom: 30px; /* 减少底部外边距 */
  color: #00838f;
  font-size: 3rem;
  text-shadow: 1px 1px 2px rgba(0,0,0,0.1);
  animation: fadeInDown 0.8s ease-out;
}

.game-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 100%;
  max-width: 600px;
  animation: fadeInUp 0.8s ease-out;
  padding: 0 20px; /* 添加左右内边距 */
  box-sizing: border-box; /* 确保内边距不会增加总宽度 */
}

.game-item {
  width: 100%;
}

.game-card {
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 15px;
  padding: 20px;
  display: flex;
  align-items: center;
  transition: all 0.3s ease;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.game-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
}

.game-icon {
  font-size: 2.5rem;
  margin-right: 20px;
  color: #00acc1;
  transition: all 0.3s ease;
  flex-shrink: 0;
  width: 60px;
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 172, 193, 0.1);
  border-radius: 12px;
}

.game-card:hover .game-icon {
  transform: scale(1.1);
}

.game-info {
  text-align: left;
  flex-grow: 1;
}

.game-info h2 {
  margin: 0 0 5px 0;
  font-size: 1.5rem;
  color: #00838f;
}

.game-info p {
  margin: 0;
  font-size: 0.9rem;
  color: #0097a7;
}

.version {
  position: fixed;
  bottom: 10px;
  right: 10px;
  font-size: 0.9rem;
  color: #00838f;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 768px) {
  h1 {
    font-size: 2.5rem;
  }

  .game-list {
    padding: 0 10px; /* 在小屏幕上减少内边距 */
  }

  .game-card {
    padding: 15px;
  }

  .game-icon {
    font-size: 2rem;
    margin-right: 15px;
    width: 50px;
    height: 50px;
  }

  .game-info h2 {
    font-size: 1.3rem;
  }

  .game-info p {
    font-size: 0.8rem;
  }
}

.game-icon-2048 {
  background-color: #edc22e;
  color: #776e65;
  font-size: 1.5rem;
  font-weight: 900;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 5px;
  line-height: 1;
}

.game-icon-2048 .number {
  width: 100%;
  text-align: center;
}

/* 为了确保图标大小一致，可以设置固定的宽高 */
.game-icon {
  width: 70px;
  height: 70px;
  font-size: 2.5rem;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .game-icon {
    width: 60px;
    height: 60px;
    font-size: 2rem;
  }

  .game-icon-2048 {
    font-size: 1.3rem;
  }
}

.title {
  margin-top: 20px;
  margin-bottom: 30px;
  font-size: 4rem;
  font-weight: bold;
  text-transform: uppercase;
}

.title span {
  display: inline-block;
  animation: letterAnimation 2s ease-in-out infinite;
  font-family: 'Arial', sans-serif;
}

.title span:nth-child(1) { color: #FF5252; animation-delay: 0s; }
.title span:nth-child(2) { color: #FF4081; animation-delay: 0.1s; }
.title span:nth-child(3) { color: #E040FB; animation-delay: 0.2s; }
.title span:nth-child(4) { color: #7C4DFF; animation-delay: 0.3s; }
.title span:nth-child(6) { color: #536DFE; animation-delay: 0.4s; }
.title span:nth-child(7) { color: #448AFF; animation-delay: 0.5s; }
.title span:nth-child(8) { color: #40C4FF; animation-delay: 0.6s; }
.title span:nth-child(9) { color: #18FFFF; animation-delay: 0.7s; }
.title span:nth-child(10) { color: #64FFDA; animation-delay: 0.8s; }

@keyframes letterAnimation {
  0%, 100% {
    transform: translateY(0) rotate(0deg) scale(1);
  }
  25% {
    transform: translateY(-20px) rotate(5deg) scale(1.1);
  }
  50% {
    transform: translateY(0) rotate(0deg) scale(1);
  }
  75% {
    transform: translateY(10px) rotate(-5deg) scale(0.9);
  }
}

@media (max-width: 768px) {
  .title {
    font-size: 3rem;
  }
}

/* 为 GoDown 游戏添加特定样式 */
.game-icon-godown {
  background-color: #4CAF50;
  color: white;
}
</style>
