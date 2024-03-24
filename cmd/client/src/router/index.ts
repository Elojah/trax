import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/Dashboard.vue'
import Sign from '@/views/Sign.vue'
import Profile from '@/views/Profile.vue'
import Map from '@/views/Map.vue'
import NotFound from '@/views/NotFound.vue'
import { useAuthStore } from '@/stores/auth'
// import { config } from '@/config'

const router = createRouter({
  // history: createWebHistory(config.web_client_url),
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
      meta: { requiresAuth: true },
    },
    {
      path: '/map',
      name: 'map',
      component: Map,
      meta: { requiresAuth: true },
    },
    {
      path: '/sign',
      name: 'sign',
      component: Sign,
      meta: { missingAuth: true },
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
  } else if (to.matched.some(record => record.meta.missingAuth)) {
    authStore.refreshProfile().then(() => {
      if (!authStore.profile) {
        next()
      } else {
        next({ name: 'dashboard' })
      }
    })
  } else {
    next()
  }
})

export default router
