import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)

  async function checkAuth() {
    try {
      const res = await fetch('/api/auth/check')
      const data = await res.json()
      isAuthenticated.value = data.authenticated
      return data.authenticated
    } catch {
      isAuthenticated.value = false
      return false
    }
  }

  async function login(password) {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ password }),
    })
    const data = await res.json()
    if (!res.ok) throw new Error(data.error || 'Login failed')
    isAuthenticated.value = true
    return data
  }

  async function logout() {
    await fetch('/api/logout', { method: 'POST' })
    isAuthenticated.value = false
  }

  return { isAuthenticated, checkAuth, login, logout }
})
