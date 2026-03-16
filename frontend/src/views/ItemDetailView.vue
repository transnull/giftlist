<template>
  <div class="min-h-screen bg-[#f9f7f5]">
    <header class="bg-white border-b border-gray-100 shadow-sm">
      <div class="max-w-3xl mx-auto px-4 py-4 flex items-center gap-3">
        <router-link to="/" class="text-rose-400 hover:text-rose-600 transition-colors flex items-center gap-1 text-sm font-medium">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
          返回首页
        </router-link>
      </div>
    </header>

    <main class="max-w-3xl mx-auto px-4 py-10">
      <!-- Loading -->
      <div v-if="loading" class="flex justify-center py-20">
        <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-rose-400"></div>
      </div>

      <!-- Not found -->
      <div v-else-if="!item" class="text-center py-20">
        <div class="text-5xl mb-4">😕</div>
        <p class="text-gray-400">找不到这个礼物哦</p>
        <router-link to="/" class="mt-4 inline-block text-rose-400 hover:text-rose-600">回到首页</router-link>
      </div>

      <!-- Detail -->
      <div v-else class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <!-- Image -->
        <div class="relative bg-gray-50" style="padding-bottom: 75%; max-height: 450px;">
          <img
            :src="item.image_url"
            :alt="item.name"
            class="absolute inset-0 w-full h-full object-contain p-4"
            @error="onImgError"
          />
        </div>

        <!-- Info -->
        <div class="p-6 md:p-8">
          <div class="flex flex-wrap items-start gap-3 mb-4">
            <h1 class="text-2xl font-bold text-gray-800 flex-1">{{ item.name }}</h1>
            <span :class="priorityClass(item.priority)" class="px-3 py-1 rounded-full text-xs font-semibold">
              {{ priorityLabel(item.priority) }}
            </span>
          </div>

          <div class="space-y-3 text-sm text-gray-600">
            <div v-if="item.category" class="flex items-center gap-2">
              <span class="text-gray-400 w-16">分类</span>
              <span class="text-gray-700 font-medium">{{ item.category }}</span>
            </div>
            <div v-if="item.estimated_cost" class="flex items-center gap-2">
              <span class="text-gray-400 w-16">参考价格</span>
              <span class="text-rose-500 font-bold text-lg">¥{{ formatCost(item.estimated_cost) }}</span>
            </div>
            <div v-if="item.remark" class="pt-3 border-t border-gray-100">
              <p class="text-gray-400 mb-2 text-xs uppercase tracking-wide">备注</p>
              <p class="text-gray-700 leading-relaxed whitespace-pre-line">{{ item.remark }}</p>
            </div>
          </div>

          <!-- Actions -->
          <div class="mt-8 flex gap-3">
            <a
              v-if="item.product_link"
              :href="item.product_link"
              target="_blank"
              rel="noopener noreferrer"
              class="flex-1 text-center btn-primary py-3 text-base rounded-xl inline-flex items-center justify-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
              去购买
            </a>
            <router-link to="/" class="flex-1 text-center btn-secondary py-3 text-base rounded-xl inline-flex items-center justify-center">
              返回首页
            </router-link>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const item = ref(null)
const loading = ref(true)

const PLACEHOLDER = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='600' height='450' viewBox='0 0 600 450'%3E%3Crect width='600' height='450' fill='%23f3f4f6'/%3E%3Ctext x='50%25' y='50%25' text-anchor='middle' dy='.3em' fill='%239ca3af' font-size='18' font-family='sans-serif'%3E图片加载失败%3C/text%3E%3C/svg%3E`

function onImgError(e) {
  e.target.src = PLACEHOLDER
}

function priorityClass(p) {
  return { high: 'priority-high', medium: 'priority-medium', low: 'priority-low' }[p] || 'priority-low'
}

function priorityLabel(p) {
  return { high: '❤️ 高优先级', medium: '💛 中优先级', low: '💚 低优先级' }[p] || p
}

function formatCost(cost) {
  return Number(cost).toLocaleString('zh-CN', { minimumFractionDigits: 0, maximumFractionDigits: 2 })
}

async function loadItem() {
  try {
    const res = await fetch(`/api/items/${route.params.id}`)
    if (!res.ok) { item.value = null; return }
    item.value = await res.json()
    if (item.value) document.title = item.value.name + ' - 礼物愿望清单'
  } catch (e) {
    item.value = null
  } finally {
    loading.value = false
  }
}

onMounted(loadItem)
</script>
