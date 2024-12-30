import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/hello-world.vue'
import Login from '../components/login-user.vue'
import Registration from '../components/register-user.vue'
import VerifyEmail from '@/components/verify-email.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HelloWorld,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/registration',
    name: 'Registration',
    component: Registration,
  },
  { 
    path: '/verify-email', 
    name: 'VerifyEmail',
    component: VerifyEmail 
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
