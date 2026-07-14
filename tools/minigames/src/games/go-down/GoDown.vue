<template>
  <div class="go-down" @click="jump" @touchstart="handleTouchStart" @touchmove="handleTouchMove" @touchend="handleTouchEnd" tabindex="0">
    <canvas ref="gameCanvas"></canvas>
    <div class="score">层数: {{ level }} | 分数: {{ score }}</div>
    <div v-if="gameOver" class="game-over">
      游戏结束
      <p>你下到了第 {{ level }} 层</p>
      <p>总分: {{ score }}</p>
      <button @click="restartGame">重新开始</button>
    </div>
    <div v-if="!gameStarted" class="start-screen">
      点击或按空格开始游戏
    </div>
  </div>
</template>

<script>
export default {
  name: 'GoDown',
  data() {
    return {
      canvas: null,
      ctx: null,
      player: {
        x: 0,
        y: 0,
        width: 30,
        height: 30,
        speed: 5,
        jumpForce: -10,
        velocity: 0,
      },
      platforms: [],
      obstacles: [],
      items: [],
      level: 0,
      passedPlatforms: 0,
      score: 0,
      gameLoop: null,
      gameOver: false,
      gameStarted: false,
      gravity: 0.5,
      platformTypes: ['normal', 'moving', 'fragile', 'bouncy', 'spiky'],
      keys: {},
      touchStartX: null,
      gameSpeed: 0, // 初始游戏速度为0
      initialPlatformWidth: 200, // 初始平台宽度
      maxLevel: 100,
      invincible: false,
      invincibleTimer: null,
    };
  },
  mounted() {
    this.initGame();
    window.addEventListener('resize', this.handleResize);
    window.addEventListener('keydown', this.handleKeyDown);
    window.addEventListener('keyup', this.handleKeyUp);
    this.$el.focus();
  },
  methods: {
    initGame() {
      this.canvas = this.$refs.gameCanvas;
      this.ctx = this.canvas.getContext('2d');
      this.handleResize();
      this.draw();
    },
    handleResize() {
      this.canvas.width = window.innerWidth;
      this.canvas.height = window.innerHeight;
      this.player.x = this.canvas.width / 2 - this.player.width / 2;
      this.player.y = this.canvas.height / 2;
    },
    startGame() {
      if (!this.gameStarted && !this.gameOver) {
        this.gameStarted = true;
        this.level = 0;
        this.score = 0;
        this.platforms = [];
        this.obstacles = [];
        this.items = [];
        this.generateInitialPlatforms();
        this.gameLoop = requestAnimationFrame(this.update);
        
        // 开始后逐渐增加游戏速度
        setTimeout(() => {
          this.gameSpeed = 2;
        }, 1000);
      }
    },
    update() {
      this.movePlayer();
      this.updatePlatforms();
      this.updateObstacles();
      this.updateItems();
      this.checkCollisions();
      this.draw();

      if (!this.gameOver) {
        this.gameLoop = requestAnimationFrame(this.update);
      }
    },
    movePlayer() {
      this.player.velocity += this.gravity;
      this.player.y += this.player.velocity;

      if (this.keys['ArrowLeft'] && this.player.x > 0) {
        this.player.x -= this.player.speed;
      }
      if (this.keys['ArrowRight'] && this.player.x < this.canvas.width - this.player.width) {
        this.player.x += this.player.speed;
      }

      if (this.player.y < 0 || this.player.y + this.player.height > this.canvas.height) {
        this.endGame();
      }
    },
    generateInitialPlatforms() {
      // 生成初始平台
      const initialPlatform = {
        x: (this.canvas.width - this.initialPlatformWidth) / 2,
        y: this.canvas.height - 50,
        width: this.initialPlatformWidth,
        height: 20,
        type: 'normal',
      };
      this.platforms.push(initialPlatform);

      // 设置玩家初始位置
      this.player.x = initialPlatform.x + (initialPlatform.width - this.player.width) / 2;
      this.player.y = initialPlatform.y - this.player.height;

      // 生成其他平台
      const platformCount = Math.ceil(this.canvas.height / 100) + 1;
      for (let i = 1; i < platformCount; i++) {
        this.platforms.push(this.generatePlatform(this.canvas.height - i * 100));
      }
    },
    generatePlatform(y) {
      const type = this.platformTypes[Math.floor(Math.random() * this.platformTypes.length)];
      return {
        x: Math.random() * (this.canvas.width - 100),
        y: y,
        width: 100,
        height: 20,
        type: type,
        life: type === 'fragile' ? 3 : Infinity,
      };
    },
    updatePlatforms() {
      this.platforms.forEach(platform => {
        platform.y -= this.gameSpeed;
        if (platform.type === 'moving') {
          platform.x += Math.sin(Date.now() / 1000) * 2;
        }
        
        // 检查是否越过平台
        if (!platform.passed && platform.y + platform.height < this.player.y) {
          platform.passed = true;
          this.passedPlatforms++;
          this.level = this.passedPlatforms;
          this.score += 10;
        }
      });

      // 移除超出屏幕的平台
      this.platforms = this.platforms.filter(platform => platform.y + platform.height > 0);

      // 添加新的平台
      if (this.platforms[this.platforms.length - 1].y > 0) {
        this.platforms.push(this.generatePlatform(this.platforms[this.platforms.length - 1].y - 100));

        if (this.level >= this.maxLevel) {
          this.winGame();
        }

        if (this.level % 10 === 0 && this.level > 0) {
          this.gameSpeed += 0.5;
        }
      }
    },
    updateObstacles() {
      // 实现障碍物的更新逻辑
    },
    updateItems() {
      // 实现道具的更新逻辑
    },
    checkCollisions() {
      let onPlatform = false;
      this.platforms.forEach(platform => {
        if (this.player.y + this.player.height > platform.y &&
            this.player.y + this.player.height < platform.y + platform.height &&
            this.player.x < platform.x + platform.width &&
            this.player.x + this.player.width > platform.x &&
            this.player.velocity > 0) {
          
          if (platform.type === 'spiky' && !this.invincible) {
            this.endGame();
            return;
          }

          if (platform.type === 'fragile') {
            platform.life--;
            if (platform.life <= 0) {
              return;
            }
          }

          this.player.y = platform.y - this.player.height;
          this.player.velocity = 0;
          onPlatform = true;

          if (platform.type === 'bouncy') {
            this.player.velocity = this.player.jumpForce * 1.5;
          }
        }
      });

      if (!onPlatform && this.player.velocity === 0) {
        this.player.velocity = 1;
      }
    },
    draw() {
      this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

      // 绘制玩家
      this.ctx.fillStyle = this.invincible ? 'gold' : 'red';
      this.ctx.fillRect(this.player.x, this.player.y, this.player.width, this.player.height);

      // 绘制平台
      this.platforms.forEach(platform => {
        switch(platform.type) {
          case 'normal':
            this.ctx.fillStyle = 'green';
            break;
          case 'moving':
            this.ctx.fillStyle = 'blue';
            break;
          case 'fragile':
            this.ctx.fillStyle = `rgba(165, 42, 42, ${platform.life / 3})`;
            break;
          case 'bouncy':
            this.ctx.fillStyle = 'orange';
            break;
          case 'spiky':
            this.ctx.fillStyle = 'purple';
            break;
        }
        this.ctx.fillRect(platform.x, platform.y, platform.width, platform.height);
      });

      // 绘制障碍物和道具
      // ...
    },
    jump() {
      if (this.gameOver) {
        this.restartGame();
      } else if (!this.gameStarted) {
        this.startGame();
      } else if (this.player.velocity === 0) {
        this.player.velocity = this.player.jumpForce;
      }
    },
    endGame() {
      this.gameOver = true;
      cancelAnimationFrame(this.gameLoop);
    },
    winGame() {
      this.gameOver = true;
      cancelAnimationFrame(this.gameLoop);
      alert("恭喜你成功下到100层！");
    },
    restartGame() {
      this.player.y = this.canvas.height - 100; // 将玩家放置在底部
      this.player.velocity = 0;
      this.platforms = [];
      this.obstacles = [];
      this.items = [];
      this.level = 0;
      this.passedPlatforms = 0;
      this.score = 0;
      this.gameOver = false;
      this.gameStarted = false;
      this.gameSpeed = 0;
      this.generateInitialPlatforms();
      this.draw();
      this.startGame();
    },
    handleKeyDown(event) {
      this.keys[event.code] = true;
      if (event.code === 'Space') {
        this.jump();
      }
    },
    handleKeyUp(event) {
      this.keys[event.code] = false;
    },
    handleTouchStart(event) {
      this.touchStartX = event.touches[0].clientX;
      this.jump();
    },
    handleTouchMove(event) {
      if (this.touchStartX === null) {
        return;
      }

      const touchEndX = event.touches[0].clientX;
      const diff = touchEndX - this.touchStartX;

      if (diff > 50) {
        this.keys['ArrowRight'] = true;
        this.keys['ArrowLeft'] = false;
      } else if (diff < -50) {
        this.keys['ArrowLeft'] = true;
        this.keys['ArrowRight'] = false;
      }
    },
    handleTouchEnd() {
      this.touchStartX = null;
      this.keys['ArrowLeft'] = false;
      this.keys['ArrowRight'] = false;
    },
  },
  beforeUnmount() {
    cancelAnimationFrame(this.gameLoop);
    window.removeEventListener('resize', this.handleResize);
    window.removeEventListener('keydown', this.handleKeyDown);
    window.removeEventListener('keyup', this.handleKeyUp);
  },
};
</script>

<style scoped>
.go-down {
  width: 100vw;
  height: 100vh;
  position: relative;
  overflow: hidden;
}

canvas {
  display: block;
}

.score {
  position: absolute;
  top: 20px;
  right: 20px;
  font-size: 24px;
  font-weight: bold;
  color: white;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
}

.game-over, .start-screen {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
}

button {
  margin-top: 10px;
  padding: 5px 10px;
  font-size: 16px;
  cursor: pointer;
}
</style>
