<template>
  <div id="app">
    <div class="home-icon" @click="goHome" v-if="showHomeIcon">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="28" height="28">
        <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z"/>
      </svg>
    </div>
    <router-view class="router-view"></router-view>
    <div class="game-transition" :style="transitionStyle" v-if="showTransition"></div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      showHomeIcon: false,
      showTransition: false,
      transitionStyle: {}
    }
  },
  methods: {
    goHome() {
      this.$router.push('/');
    },
    updateHomeIconVisibility() {
      this.showHomeIcon = this.$route.path !== '/';
    },
    startGameAnimation({ game, centerX, centerY }) {
      this.showTransition = true;
      this.transitionStyle = {
        left: `${centerX}px`,
        top: `${centerY}px`,
        backgroundColor: game === '2048' ? '#edc22e' : '#4ec0ca'
      };
      setTimeout(() => {
        this.showTransition = false;
      }, 1000);
    }
  },
  watch: {
    $route() {
      this.updateHomeIconVisibility();
    }
  },
  created() {
    this.updateHomeIconVisibility();
    this.$root.$on('start-game-animation', this.startGameAnimation);
  },
  beforeUnmount() {
    this.$root.$off('start-game-animation', this.startGameAnimation);
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
}

#app {
  position: relative;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.home-icon {
  position: fixed;
  top: 12px;
  left: 12px;
  cursor: pointer;
  z-index: 1000;
  padding: 5px;
  background-color: rgba(255, 255, 255, 0.7);
  border-radius: 50%;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
}

.home-icon:hover {
  transform: scale(1.1);
}

.home-icon svg {
  fill: #776e65;
  display: block;
}

@media (orientation: landscape) {
  .home-icon {
    top: 20px;
    left: 20px;
  }
}

.router-view {
  flex: 1;
  overflow-y: auto;
}

.game-transition {
  position: fixed;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  z-index: 1000;
  pointer-events: none;
  animation: expandGame 1s ease-out forwards;
}

@keyframes expandGame {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  100% {
    transform: scale(200);
    opacity: 1;
  }
}
</style>
