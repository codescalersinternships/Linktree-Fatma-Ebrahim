import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SignupView from '@/views/SignupView.vue'
import LoginView from '@/views/LoginView.vue'
import EditView from '@/views/EditView.vue'
import TreeView from '@/views/TreeView.vue'
import DetailsView from '@/views/DetailsView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/signup',
    name: 'signup',
    component: SignupView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/tree',
    name: 'tree',
    component: TreeView
  },
  {
    path: '/details',
    name: 'details',
    component: DetailsView
  },
  {
    path: '/edit',
    name: 'edit',
    component: EditView
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
