import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Sign from '@/views/Sign.vue'
import Profile from '@/views/Profile.vue'
import NotFound from '@/views/NotFound.vue'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
// import { config } from '@/config'

const router = createRouter({
  // history: createWebHistory(config.web_client_url),
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/sign',
      name: 'sign',
      component: Sign,
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
      meta: { requiresAuth: true },
    },
    {
      path: '/:catchAll(.*)',
      name: 'not-found',
      component: NotFound,
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.matched.some(record => record.meta.requiresAuth) && !authStore.profile) {
      authStore.refreshProfile().then(() => {
        if (!authStore.profile) {
          next({ name: 'sign' })
        } else {
          next()
        }
      })
  } else {
    next()
  }
})

export default router
