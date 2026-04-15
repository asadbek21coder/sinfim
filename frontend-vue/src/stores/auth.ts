import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth'
import { tokenStorage } from '@/api/client'
import type { UserDto } from '@/types/auth'

export const useAuthStore = defineStore('auth', () => {
  const user         = ref<UserDto | null>(null)
  const accessToken  = ref<string | null>(tokenStorage.getAccess())
  const refreshToken = ref<string | null>(tokenStorage.getRefresh())

  const isLoggedIn = computed(() => !!accessToken.value && !!user.value)
  const role       = computed(() => user.value?.role ?? null)

  function hasRole(roles: string[]): boolean {
    return !!role.value && roles.includes(role.value)
  }

  function setTokens(access: string, refresh: string) {
    accessToken.value  = access
    refreshToken.value = refresh
    tokenStorage.set(access, refresh)
  }

  function clearAuth() {
    user.value         = null
    accessToken.value  = null
    refreshToken.value = null
    tokenStorage.clear()
  }

  async function login(phoneNumber: string, password: string) {
    const res  = await authApi.login({ phoneNumber, password })
    const data = res.data
    setTokens(data.accessToken, data.refreshToken)
    user.value = data.user
  }

  async function fetchCurrentUser() {
    const res  = await authApi.me()
    user.value = res.data
  }

  async function changeMyPassword(currentPassword: string, newPassword: string) {
    await authApi.changeMyPassword({ currentPassword, newPassword })
    await fetchCurrentUser()
  }

  async function logout() {
    if (refreshToken.value) {
      await authApi.logout().catch((err) => {
        console.warn('Token revocation failed:', err)
      })
    }
    clearAuth()
  }

  return {
    user,
    accessToken,
    refreshToken,
    isLoggedIn,
    role,
    hasRole,
    setTokens,
    clearAuth,
    login,
    fetchCurrentUser,
    changeMyPassword,
    logout,
  }
})
