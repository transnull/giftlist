import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ItemDetailView from '../views/ItemDetailView.vue'
import AdminLoginView from '../views/AdminLoginView.vue'
import AdminView from '../views/AdminView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: HomeView },
    { path: '/item/:id', component: ItemDetailView },
    { path: '/admin/login', component: AdminLoginView },
    { path: '/admin', component: AdminView, meta: { requiresAuth: true } },
  ],
})

router.beforeEach(async (to) => {
  if (to.meta.requiresAuth) {
    try {
      const res = await fetch('/api/auth/check')
      const data = await res.json()
      if (!data.authenticated) {
        return '/admin/login'
      }
    } catch {
      return '/admin/login'
    }
  }
})

export default router
