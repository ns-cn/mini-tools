<template>
  <div class="flappy-bird" @click="jump" @keydown.space="jump" tabindex="0">
    <canvas ref="gameCanvas"></canvas>
    <div class="score">分数: {{ score }}</div>
    <div v-if="gameOver" class="game-over">
      游戏结束
      <p>得分: {{ score }}</p>
      <button @click="restartGame">重新开始</button>
      <p class="restart-hint">按空格或点击屏幕重新开始</p>
    </div>
    <div v-if="!gameStarted" class="start-screen">
      点击或按空格开始游戏
    </div>
  </div>
</template>

<script>
export default {
  name: 'FlappyBird',
  data() {
    return {
      canvas: null,
      ctx: null,
      bird: {
        x: 50,
        y: 150,
        width: 34,
        height: 24,
        velocity: 0,
        baseGravity: 0.2,
        gravity: 0.2,
        jump: -5,
        rotation: 0,
        targetRotation: 0,
      },
      maxRotation: Math.PI / 4,
      minRotation: -Math.PI / 3,
      rotationSpeed: 0.08,
      rotationLerpFactor: 0.1,
      pipes: [],
      gameLoop: null,
      gameOver: false,
      gameStarted: false,
      score: 0,
      bgImage: null,
      birdImage: null,
      scale: 1,
      pipeGap: 150,
      pipeWidth: 52,
      basePipeSpawnInterval: 2500,
      minPipeSpawnInterval: 1500,
      pipeSpawnInterval: 2500,
      lastPipeSpawn: 0,
      nextPipeSpawn: 0,
      difficultyFactor: 0.002,
      maxGravity: 0.4,
      collidedPipe: null,
      bgX: 0,
      bgSpeed: 0.5,
      pipeSpeed: 1.5,
    };
  },
  mounted() {
    window.addEventListener('resize', this.handleResize);
    this.initGame();
    this.$el.focus();
  },
  methods: {
    initGame() {
      this.canvas = this.$refs.gameCanvas;
      this.ctx = this.canvas.getContext('2d');
      this.handleResize();

      this.bgImage = new Image();
      this.bgImage.src = require('@/assets/flappy-bird/background.png');

      this.birdImage = new Image();
      this.birdImage.src = require('@/assets/flappy-bird/bird.png');

      this.draw();
    },
    handleResize() {
      this.canvas.width = window.innerWidth;
      this.canvas.height = window.innerHeight;
      this.scale = Math.min(this.canvas.width / 288, this.canvas.height / 512);
      this.bird.x = 50 * this.scale;
      this.bird.width = 34 * this.scale;
      this.bird.height = 24 * this.scale;
      this.pipeGap = 150 * this.scale;
      this.pipeWidth = 52 * this.scale;
    },
    startGame() {
      if (!this.gameStarted && !this.gameOver) {
        this.gameStarted = true;
        this.bird.velocity = 0;
        this.bird.y = this.canvas.height / 2;
        this.jump();
        this.lastPipeSpawn = performance.now();
        this.pipeSpawnInterval = this.basePipeSpawnInterval;
        this.gameLoop = requestAnimationFrame(this.update);
      }
    },
    update(timestamp) {
      if (!this.lastPipeSpawn) this.lastPipeSpawn = timestamp;

      this.updateBird();
      this.updateBackground();
      this.updatePipes(timestamp);
      this.checkCollisions();
      this.draw();

      if (!this.gameOver) {
        this.gameLoop = requestAnimationFrame(this.update);
      }
    },
    updateBird() {
      // 根据得分调整重力
      this.bird.gravity = Math.min(
        this.bird.baseGravity + this.score * this.difficultyFactor,
        this.maxGravity
      );

      this.bird.velocity += this.bird.gravity;
      this.bird.y += this.bird.velocity;

      // 更新目标旋转角度
      if (this.bird.velocity < 0) {
        // 上升时快速抬头
        this.bird.targetRotation = this.minRotation;
      } else {
        // 下降时逐渐低头，速度越快低头越快
        const fallRotation = Math.min(
          this.maxRotation,
          (this.bird.velocity / 10) * this.maxRotation
        );
        this.bird.targetRotation = fallRotation;
      }

      // 使用线性插值(LERP)平滑过渡到目标旋转角度
      const rotationDiff = this.bird.targetRotation - this.bird.rotation;
      this.bird.rotation += rotationDiff * this.rotationLerpFactor;

      // 限制最大下降速度
      if (this.bird.velocity > 10) {
        this.bird.velocity = 10;
      }

      if (this.bird.y < 0) {
        this.bird.y = 0;
        this.bird.velocity = 0;
      }
    },
    updateBackground() {
      this.bgX -= this.bgSpeed * this.scale;
      if (this.bgX <= -this.canvas.width) {
        this.bgX = 0;
      }
    },
    updatePipes(timestamp) {
      // 移动现有的管道
      this.pipes.forEach(pipe => {
        pipe.x -= this.pipeSpeed * this.scale;
      });

      // 移除离开屏幕的管道
      this.pipes = this.pipes.filter(pipe => pipe.x > -this.pipeWidth);

      // 检查是否需要生成新管道
      if (timestamp - this.lastPipeSpawn >= this.pipeSpawnInterval) {
        this.addPipe();
        this.lastPipeSpawn = timestamp;
        this.pipeSpawnInterval = this.calculatePipeSpawnInterval();
      }

      // 更新分数
      this.pipes.forEach(pipe => {
        if (!pipe.passed && pipe.x + this.pipeWidth < this.bird.x) {
          this.score++;
          pipe.passed = true;
        }
      });
    },
    addPipe() {
      const minHeight = 50 * this.scale;
      const maxHeight = this.canvas.height - this.pipeGap - minHeight;
      const height = Math.floor(Math.random() * (maxHeight - minHeight + 1) + minHeight);

      this.pipes.push({
        x: this.canvas.width,
        topHeight: height,
        bottomY: height + this.pipeGap,
        passed: false,
      });
    },
    checkCollisions() {
      if (this.bird.y + this.bird.height > this.canvas.height) {
        this.endGame();
      }

      this.collidedPipe = null; // 重置碰撞的柱子

      this.pipes.forEach(pipe => {
        if (
          this.bird.x < pipe.x + this.pipeWidth &&
          this.bird.x + this.bird.width > pipe.x &&
          (this.bird.y < pipe.topHeight || this.bird.y + this.bird.height > pipe.bottomY)
        ) {
          this.collidedPipe = pipe; // 记录碰撞的柱子
          this.endGame();
        }
      });
    },
    draw() {
      this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

      // 绘制背景
      this.ctx.drawImage(this.bgImage, this.bgX, 0, this.canvas.width, this.canvas.height);
      this.ctx.drawImage(this.bgImage, this.bgX + this.canvas.width, 0, this.canvas.width, this.canvas.height);

      // 绘制管道
      this.pipes.forEach(pipe => {
        this.ctx.fillStyle = pipe === this.collidedPipe ? '#FF0000' : '#5EAD4A';
        this.ctx.fillRect(pipe.x, 0, this.pipeWidth, pipe.topHeight);
        this.ctx.fillRect(pipe.x, pipe.bottomY, this.pipeWidth, this.canvas.height - pipe.bottomY);
      });

      // 绘制小鸟
      this.ctx.save();
      this.ctx.translate(this.bird.x + this.bird.width / 2, this.bird.y + this.bird.height / 2);
      this.ctx.rotate(this.bird.rotation);
      this.ctx.drawImage(
        this.birdImage,
        -this.bird.width / 2,
        -this.bird.height / 2,
        this.bird.width,
        this.bird.height
      );
      this.ctx.restore();
    },
    jump() {
      if (this.gameOver) {
        this.restartGame();
      } else if (!this.gameStarted) {
        this.startGame();
      } else {
        this.bird.velocity = this.bird.jump * this.scale;
        // 跳跃时立即开始向上旋转，但通过插值系统实现平滑过渡
        this.bird.targetRotation = this.minRotation;
      }
    },
    endGame() {
      this.gameOver = true;
      cancelAnimationFrame(this.gameLoop);
    },
    restartGame() {
      // 重置小鸟状态
      this.bird.y = this.canvas.height / 2;
      this.bird.velocity = 0;
      this.bird.gravity = this.bird.baseGravity;
      this.bird.rotation = 0;

      // 重置游戏状态
      this.pipes = [];
      this.score = 0;
      this.gameOver = false;
      this.gameStarted = false;
      this.collidedPipe = null;

      // 重置管道生成相关参数
      this.lastPipeSpawn = 0;
      this.pipeSpawnInterval = this.basePipeSpawnInterval;

      // 重置背景位置
      this.bgX = 0;

      // 重绘初始画面
      this.draw();

      // 重新开始游戏循环
      cancelAnimationFrame(this.gameLoop);
      this.gameLoop = null;

      // 使用 startGame 方法来确保一致的开始状态
      this.startGame();
    },
    calculatePipeSpawnInterval() {
      if (this.score <= 30) {
        // 前30分，缓慢减少间隔
        return this.basePipeSpawnInterval - this.score * 10;
      } else {
        // 30分以后，加快减少速度
        const baseDecrease = 300; // 前30分的总减少量
        const extraDecrease = Math.pow(this.score - 30, 1.5) * 2; // 使用指数函数加快减少速度，但减小系数
        return Math.max(
          this.minPipeSpawnInterval,
          this.basePipeSpawnInterval - baseDecrease - extraDecrease
        );
      }
    },
  },
  beforeUnmount() {
    cancelAnimationFrame(this.gameLoop);
    window.removeEventListener('resize', this.handleResize);
  },
};
</script>

<style scoped>
.flappy-bird {
  width: 100vw;
  height: 100vh;
  position: relative;
  overflow: hidden;
  outline: none;
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

.restart-hint {
  margin-top: 10px;
  font-size: 14px;
  color: #aaa;
}
</style>
