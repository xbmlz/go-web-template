import { useAuthStore } from '@/store'
import { createRouter, createWebHashHistory } from 'vue-router'

const Layout = () => import('@/layout/index.vue')

/** @type {import('vue-router').RouteRecordRaw} */
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Layout,
    children: [
      {
        path: '',
        name: 'Index',
        component: () => import('@/views/home/index.vue'),
      },
    ],
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()
  if (authStore.isLoggedIn) {
    if (to.path === '/login') {
      next('/')
    }
    else {
      await authStore.getUserInfo()
      next()
    }
  }
  else {
    if (to.path === '/login') {
      next()
    }
    else {
      next('/login')
    }
  }
})

export default router
