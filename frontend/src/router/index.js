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
    path: '/tree/:id',
    name: 'tree',
    props: true,
    component: TreeView
  },
  {
    path: '/details',
    name: 'details',
    component: DetailsView
  },
  {
    path: '/edit/:id',
    name: 'edit',
    probs:true,
    component: EditView
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})


router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");

  if (to.path === "/login" || to.path === "/signup" || token) {
    next();
    console.log("Token found, allowing access");
  } else {
    console.log("No token found, redirecting to login");
    next("/login");
  }
});



export default router
