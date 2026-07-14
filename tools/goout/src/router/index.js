import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import DashboardView from '../views/DashboardView.vue'
import AboutView from '../views/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'root',
      component: {
        template: '<div></div>'  // 空组件
      },
      beforeEnter: async (to, from, next) => {
        const token = localStorage.getItem('token')
        if (!token) {
          next('/login')
          return
        }

        try {
          const response = await fetch('/api/users/validate', {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          })
          const result = await response.json()
          if (response.ok && result.code === 0) {
            next('/dashboard')
            return
          }
        } catch (error) {
          console.error('Token验证失败:', error)
        }

        // 如果验证失败，清除本地存储的token并跳转到登录页
        localStorage.removeItem('token')
        localStorage.removeItem('username')
        localStorage.removeItem('nickname')
        next('/login')
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      meta: { requiresAuth: true }
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
      meta: { requiresAuth: true }
    }
  ]
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('token')

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router
