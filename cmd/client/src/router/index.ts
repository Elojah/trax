import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/Dashboard.vue'
import Sign from '@/views/Sign.vue'
import User from '@/views/User.vue'
import Map from '@/views/Map.vue'
import Entity from '@/views/Entity.vue'
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
      meta: { requiresAuth: true }
    },
    {
      path: '/map',
      name: 'map',
      component: Map,
      meta: { requiresAuth: true }
    },
    {
      path: '/sign',
      name: 'sign',
      component: Sign,
      meta: { missingAuth: true }
    },
    {
      path: '/user',
      name: 'user',
      component: User,
      meta: { requiresAuth: true }
    },
    {
      path: '/entity',
      name: 'entity',
      component: Entity,
      meta: { requiresAuth: true }
    },
    {
      path: '/:catchAll(.*)',
      name: 'not-found',
      component: NotFound
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.matched.some((record) => record.meta.requiresAuth) && !authStore.user) {
    authStore.refreshToken().then(() => {
      if (!authStore.user) {
        next({ name: 'sign' })
      } else {
        next()
      }
    })
  } else if (to.matched.some((record) => record.meta.missingAuth)) {
    authStore.refreshToken().then(() => {
      if (!authStore.user) {
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
