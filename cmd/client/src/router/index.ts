import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Signin from '@/views/Signin.vue'
import NotFound from '@/views/NotFound.vue'
import { config } from '@/config'

const router = createRouter({
  history: createWebHistory(""),
  // history: createWebHistory(config.web_client_url),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component:About,
    },
    {
      path: '/signin',
      name: 'signin',
      component:Signin,
    },
    {
      path: '/:catchAll(.*)',
      name: 'not-found',
      component:NotFound,
    }
  ]
})

export default router
