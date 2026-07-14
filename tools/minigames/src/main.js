import Vue from 'vue'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false

new Vue({
  router,  // 确保这里添加了 router
  render: h => h(App),
}).$mount('#app')
