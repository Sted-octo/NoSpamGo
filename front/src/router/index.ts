import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import EmailConfig from '@/views/EmailConfig.vue'
import Dashboard from '@/views/Dashboard.vue'
import { useAuthStore } from '@/stores/auth'

const routes = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { requiresAuth: false },
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { requiresAuth: false },
  },
  {
    path: '/email-config/:email',
    name: 'EmailConfig',
    component: EmailConfig,
    meta: { requiresAuth: true },
    props: true,
  },
  {
    path: '/dashboard/:email',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true },
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth !== true) {
    next()
    return
  }

  if (authStore.isAuthenticated()) {
    next()
    return
  }

  next({
    path: '/login',
    query: { redirect: to.fullPath },
  })
})

export default router
