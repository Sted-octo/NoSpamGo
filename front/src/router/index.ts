import { createRouter, createWebHistory } from 'vue-router'
import TwoFactorAuth from '@/components/TwoFactorAuth.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'TwoFactorAuth',
      component: TwoFactorAuth,
    },
  ],
})

export default router
