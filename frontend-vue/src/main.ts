import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './assets/main.css'
import App from './App.vue'
import router from './router'
import { useAuthStore } from './stores/auth'

const app   = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

app.config.errorHandler = (err, _instance, info) => {
  console.error('[Global Error]', info, err)
}

// Role → home page mapping — adjust per your app
const roleHomeMap: Record<string, string> = {
  PLATFORM_ADMIN: '/admin/organizations/new',
  OWNER: '/app/dashboard',
  TEACHER: '/app/courses',
  MENTOR: '/app/homework/review',
  STUDENT: '/learn/dashboard',
}
function roleHome(role?: string | null): string {
  return role ? (roleHomeMap[role] ?? '/app/dashboard') : '/auth/login'
}

// Navigation guard
router.beforeEach(async (to, _from, next) => {
  const auth = useAuthStore()

  if (!to.meta.requiresAuth) {
    if ((to.path === '/login' || to.path === '/auth/login') && auth.isLoggedIn) {
      return next(roleHome(auth.role))
    }
    return next()
  }

  if (!auth.accessToken) {
    return next('/auth/login')
  }

  if (!auth.user) {
    try {
      await auth.fetchCurrentUser()
    } catch {
      auth.clearAuth()
      return next('/auth/login')
    }
  }

  if (auth.user?.mustChangePassword && to.path !== '/auth/change-password') {
    return next('/auth/change-password')
  }

  const requiredRoles = to.meta.roles
  if (requiredRoles && requiredRoles.length > 0 && !auth.hasRole(requiredRoles)) {
    return next(roleHome(auth.role))
  }

  next()
})

app.mount('#app')
