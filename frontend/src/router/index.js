import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/hello-world.vue'
import Login from '../components/login-user.vue'

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
];

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
