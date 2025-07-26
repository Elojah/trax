import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Sign from '@/views/Sign.vue'
import User from '@/views/User.vue'
import Map from '@/views/Map.vue'
import Group from '@/views/Group.vue'
import GroupDetails from '@/components/group/Details.vue'
import RoleDetails from '@/components/role/Details.vue'
import NotFound from '@/views/NotFound.vue'
import { useAuthStore } from '@/stores/auth'
// import { config } from '@/config'

const router = createRouter({
  // history: createWebHistory(config.web_client_url),
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
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
      path: '/group',
      name: 'group',
      component: Group,
      meta: { requiresAuth: true }
    },
    {
      path: '/group/:id',
      name: 'group-details',
      component: GroupDetails,
      meta: { requiresAuth: true }
    },
    {
      path: '/group/:groupId/role/:roleId',
      name: 'role-details',
      component: RoleDetails,
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
        next({ name: 'home' })
      }
    })
  } else {
    next()
  }
})

export default router
