<template>
  <div class="game-2048" @touchstart="handleTouchStart" @touchmove="handleTouchMove" @touchend="handleTouchEnd">
    <div class="game-content">
      <div class="game-info">
        <h2>2048 游戏</h2>
        <div class="score">分数: {{ score }}</div>
        <button @click="newGame">新游戏</button>
      </div>
      <div class="grid-container">
        <div class="grid">
          <div v-for="(row, rowIndex) in grid" :key="rowIndex" class="row">
            <div v-for="(cell, cellIndex) in row" :key="cellIndex" class="cell" :class="['cell-' + cell, { 'new-tile': isNewTile(rowIndex, cellIndex) }]">
              {{ cell !== 0 ? cell : '' }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="gameOver" class="game-over-modal">
      <div class="modal-content">
        <h2>游戏结束</h2>
        <p>你的得分是: {{ score }}</p>
        <button @click="newGame">再来一局</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Game2048View',
  data() {
    return {
      grid: [],
      score: 0,
      touchStartX: null,
      touchStartY: null,
      newTiles: [],
      gameOver: false,
    }
  },
  methods: {
    newGame() {
      this.grid = Array(4).fill().map(() => Array(4).fill(0));
      this.score = 0;
      this.gameOver = false;
      this.initTilePositions();
      this.addNewTile();
      this.addNewTile();
    },
    initTilePositions() {
      this.tilePositions = this.grid.map((row, i) => 
        row.map((cell, j) => cell !== 0 ? { row: i, col: j } : null)
      );
    },
    addNewTile() {
      const emptyCells = [];
      for (let i = 0; i < 4; i++) {
        for (let j = 0; j < 4; j++) {
          if (this.grid[i][j] === 0) {
            emptyCells.push({i, j});
          }
        }
      }
      if (emptyCells.length > 0) {
        const {i, j} = emptyCells[Math.floor(Math.random() * emptyCells.length)];
        this.grid[i][j] = Math.random() < 0.9 ? 2 : 4;
        this.newTiles.push({i, j});
        this.tilePositions[i][j] = { row: i, col: j };  // 立即更新位置
      }
    },
    move(direction) {
      let moved = false;
      let newGrid = JSON.parse(JSON.stringify(this.grid));
      this.newTiles = [];

      const rotate = (grid) => {
        return grid[0].map((val, index) => grid.map(row => row[index]).reverse());
      };

      // 向左移动
      const moveLeft = (row) => {
        const filtered = row.filter(cell => cell !== 0);
        const newRow = [];
        for (let i = 0; i < filtered.length; i++) {
          if (newRow.length > 0 && newRow[newRow.length - 1] === filtered[i]) {
            newRow[newRow.length - 1] *= 2;
            this.score += newRow[newRow.length - 1];
            moved = true;
          } else {
            newRow.push(filtered[i]);
          }
        }
        while (newRow.length < 4) {
          newRow.push(0);
        }
        return newRow;
      };

      // 根据方向旋转网格
      const rotations = {'left': 0, 'up': 3, 'right': 2, 'down': 1}[direction];
      for (let i = 0; i < rotations; i++) {
        newGrid = rotate(newGrid);
      }

      // 移动
      newGrid = newGrid.map(row => {
        const oldRow = [...row];
        const newRow = moveLeft(row);
        if (!moved && JSON.stringify(oldRow) !== JSON.stringify(newRow)) {
          moved = true;
        }
        return newRow;
      });

      // 旋转回原来的方向
      const reverseRotations = {'left': 0, 'up': 1, 'right': 2, 'down': 3}[direction];
      for (let i = 0; i < reverseRotations; i++) {
        newGrid = rotate(newGrid);
      }

      if (moved) {
        this.grid = newGrid;
        this.addNewTile();
        if (this.isGameOver()) {
          this.gameOver = true;
        }
      }
    },
    handleKeydown(e) {
      switch(e.key) {
        case 'ArrowLeft': this.move('left'); break;
        case 'ArrowUp': this.move('up'); break;
        case 'ArrowRight': this.move('right'); break;
        case 'ArrowDown': this.move('down'); break;
      }
    },
    handleTouchStart(event) {
      this.touchStartX = event.touches[0].clientX;
      this.touchStartY = event.touches[0].clientY;
    },

    handleTouchMove(event) {
      event.preventDefault(); // 防止页面滚动
    },

    handleTouchEnd(event) {
      if (!this.touchStartX || !this.touchStartY) {
        return;
      }

      const touchEndX = event.changedTouches[0].clientX;
      const touchEndY = event.changedTouches[0].clientY;

      const deltaX = touchEndX - this.touchStartX;
      const deltaY = touchEndY - this.touchStartY;

      // 判断滑动方向
      if (Math.abs(deltaX) > Math.abs(deltaY)) {
        if (deltaX > 0) {
          this.move('right');
        } else {
          this.move('left');
        }
      } else {
        if (deltaY > 0) {
          this.move('down');
        } else {
          this.move('up');
        }
      }

      // 重置触摸起始点
      this.touchStartX = null;
      this.touchStartY = null;
    },
    isNewTile(row, col) {
      return this.newTiles.some(tile => tile.i === row && tile.j === col);
    },
    isGameOver() {
      // 检查是否还有空格
      for (let i = 0; i < 4; i++) {
        for (let j = 0; j < 4; j++) {
          if (this.grid[i][j] === 0) {
            return false;
          }
        }
      }
      // 检查是否还有可以合并的相邻格子
      for (let i = 0; i < 4; i++) {
        for (let j = 0; j < 4; j++) {
          if (
            (i < 3 && this.grid[i][j] === this.grid[i + 1][j]) ||
            (j < 3 && this.grid[i][j] === this.grid[i][j + 1])
          ) {
            return false;
          }
        }
      }
      return true;
    },
  },
  mounted() {
    this.newGame();
    window.addEventListener('keydown', this.handleKeydown);
  },
  beforeUnmount() {
    window.removeEventListener('keydown', this.handleKeydown);
  },
}
</script>

<style scoped>
.game-2048 {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  padding: 10px;
  box-sizing: border-box;
  touch-action: none;
  overflow: hidden;
}

.game-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 500px;
  width: 100%;
  height: 100%;
  overflow-y: auto;
}

.game-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  width: 100%;
  margin-bottom: 20px;
}

h2 {
  margin: 0;
  font-size: 1.5rem;
}

.score {
  font-size: 1.2rem;
}

.grid-container {
  width: 100%;
  padding-bottom: 100%; /* 创建一个正方形容器 */
  position: relative;
}

.grid {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: grid;
  grid-template-rows: repeat(4, 1fr);
  grid-gap: 2%;
  background-color: #bbada0;
  border-radius: 5px;
  padding: 2%;
}

.row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 2%;
}

.cell {
  background-color: #ccc0b3;
  border-radius: 5px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-weight: bold;
  position: relative;
  font-size: 1.5rem;
}

.cell-2 { background-color: #eee4da; color: #776e65; }
.cell-4 { background-color: #ede0c8; color: #776e65; }
.cell-8 { background-color: #f2b179; color: #f9f6f2; }
.cell-16 { background-color: #f59563; color: #f9f6f2; }
.cell-32 { background-color: #f67c5f; color: #f9f6f2; }
.cell-64 { background-color: #f65e3b; color: #f9f6f2; }
.cell-128 { background-color: #edcf72; color: #f9f6f2; }
.cell-256 { background-color: #edcc61; color: #f9f6f2; }
.cell-512 { background-color: #edc850; color: #f9f6f2; }
.cell-1024 { background-color: #edc53f; color: #f9f6f2; }
.cell-2048 { background-color: #edc22e; color: #f9f6f2; }

.new-tile {
  animation: appear 0.2s ease-in-out;
}

@keyframes appear {
  0% {
    opacity: 0;
    transform: scale(0);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}

button {
  font-size: 1rem;
  padding: 10px 20px;
  cursor: pointer;
  margin-bottom: 10px;
}

@media (orientation: landscape) {
  .game-content {
    flex-direction: row;
    justify-content: center;
    align-items: center;
    max-width: none;
    height: 100%;
    overflow: hidden;
  }

  .game-info {
    width: auto;
    margin-right: 20px;
    margin-bottom: 0;
  }

  .grid-container {
    width: auto;
    height: 80vh;
    max-height: 80vmin;
    padding-bottom: 0;
    aspect-ratio: 1 / 1;
  }
}

@media (max-height: 600px), (max-width: 400px) {
  h2 {
    font-size: 1.2rem;
  }

  .score {
    font-size: 1rem;
  }

  .cell {
    font-size: 1rem;
  }

  button {
    font-size: 0.9rem;
    padding: 8px 16px;
  }
}

.game-over-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
}

.modal-content h2 {
  margin-top: 0;
}

.modal-content button {
  margin-top: 20px;
}
</style>
