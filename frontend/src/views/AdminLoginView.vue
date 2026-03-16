<template>
  <div class="min-h-screen bg-[#f9f7f5] flex items-center justify-center px-4">
    <div class="w-full max-w-sm">
      <!-- Logo / Title -->
      <div class="text-center mb-8">
        <div class="text-5xl mb-3">🎁</div>
        <h1 class="text-2xl font-bold text-gray-800">后台管理</h1>
        <p class="text-gray-400 text-sm mt-1">请输入管理密码继续</p>
      </div>

      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
        <form @submit.prevent="handleLogin">
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-1">管理密码</label>
            <input
              v-model="password"
              type="password"
              class="input-field"
              placeholder="请输入密码"
              autofocus
              :disabled="loading"
            />
          </div>

          <div v-if="errorMsg" class="mb-4 bg-red-50 border border-red-200 text-red-600 text-sm rounded-lg px-3 py-2">
            {{ errorMsg }}
          </div>

          <button type="submit" class="btn-primary w-full py-2.5" :disabled="loading">
            <span v-if="loading">登录中...</span>
            <span v-else>登 录</span>
          </button>
        </form>
      </div>

      <p class="text-center mt-6">
        <router-link to="/" class="text-gray-400 hover:text-gray-600 text-sm transition-colors">← 返回公开页面</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const auth = useAuthStore()
const password = ref('')
const loading = ref(false)
const errorMsg = ref('')

async function handleLogin() {
  if (!password.value) return
  loading.value = true
  errorMsg.value = ''
  try {
    await auth.login(password.value)
    router.push('/admin')
  } catch (e) {
    errorMsg.value = e.message || '密码错误，请重试'
  } finally {
    loading.value = false
  }
}
</script>
