import axios from 'axios'

const ACCESS_TOKEN_KEY  = 'app_access_token'
const REFRESH_TOKEN_KEY = 'app_refresh_token'

export const tokenStorage = {
  getAccess:  () => localStorage.getItem(ACCESS_TOKEN_KEY),
  getRefresh: () => localStorage.getItem(REFRESH_TOKEN_KEY),
  set: (access: string, refresh: string) => {
    localStorage.setItem(ACCESS_TOKEN_KEY, access)
    localStorage.setItem(REFRESH_TOKEN_KEY, refresh)
  },
  clear: () => {
    localStorage.removeItem(ACCESS_TOKEN_KEY)
    localStorage.removeItem(REFRESH_TOKEN_KEY)
  },
}

const client = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: { 'Content-Type': 'application/json' },
})

// Attach Bearer token to every request
client.interceptors.request.use((config) => {
  const token = tokenStorage.getAccess()
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

let isRefreshing = false
let refreshQueue: Array<(token: string) => void> = []

// On 401 → try refresh → retry original request
client.interceptors.response.use(
  (response) => response,
  async (error) => {
    const original = error.config

    if (error.response?.status !== 401 || original._retry) {
      return Promise.reject(error)
    }

    const refreshToken = tokenStorage.getRefresh()
    if (!refreshToken) {
      tokenStorage.clear()
      const onLoginPage = window.location.pathname === '/login' || window.location.pathname === '/auth/login'
      if (!onLoginPage) window.location.href = '/auth/login'
      return Promise.reject(error)
    }

    if (isRefreshing) {
      return new Promise((resolve) => {
        refreshQueue.push((newToken) => {
          original.headers.Authorization = `Bearer ${newToken}`
          resolve(client(original))
        })
      })
    }

    original._retry = true
    isRefreshing = true

    try {
      const res = await axios.post<{ accessToken: string; refreshToken: string }>(
        `${import.meta.env.VITE_API_BASE_URL}/auth/refresh-token`,
        { refreshToken },
      )
      const { accessToken, refreshToken: newRefresh } = res.data
      tokenStorage.set(accessToken, newRefresh)
      refreshQueue.forEach((cb) => cb(accessToken))
      refreshQueue = []
      original.headers.Authorization = `Bearer ${accessToken}`
      return client(original)
    } catch {
      tokenStorage.clear()
      window.location.href = '/auth/login'
      return Promise.reject(error)
    } finally {
      isRefreshing = false
    }
  },
)

export default client
