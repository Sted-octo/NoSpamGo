import { createRouter, createWebHistory } from 'vue-router'
import TwoFactorAuth from '@/components/TwoFactorAuth.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import EmailConfig from '@/views/EmailConfig.vue'
//import Dashboard from '../views/Dashboard.vue'

const routes = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/email-config/:email',
    name: 'EmailConfig',
    component: EmailConfig,
    //meta: { requiresAuth: true },
    props: true,
  },
  /*{
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true },
  },*/
  {
    path: '/two-factor-auth',
    name: 'TwoFactorAuth',
    component: TwoFactorAuth,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
