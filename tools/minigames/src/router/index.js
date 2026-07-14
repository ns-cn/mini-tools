import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../components/Home.vue'
import Game2048 from '../games/2048/Game2048.vue'
import FlappyBird from '../games/flappy-bird/FlappyBird.vue'
import GoDown from '../games/go-down/GoDown.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/2048',
    name: '2048',
    component: Game2048
  },
  {
    path: '/flappy-bird',
    name: 'FlappyBird',
    component: FlappyBird
  },
  {
    path: '/go-down',
    name: 'GoDown',
    component: GoDown
  }
]

const router = new VueRouter({
  routes
})

export default router
