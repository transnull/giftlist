<template>
  <div class="min-h-screen bg-[#f9f7f5]">
    <!-- Header -->
    <header class="bg-white border-b border-gray-100 shadow-sm">
      <div class="max-w-6xl mx-auto px-4 py-6 text-center">
        <h1 class="text-3xl md:text-4xl font-bold text-gray-800 tracking-tight">
          🎁 {{ siteName }}
        </h1>
      </div>
    </header>

    <main class="max-w-6xl mx-auto px-4 py-8">
      <!-- Intro block -->
      <div v-if="introText" class="mb-10 relative">
        <div class="bg-amber-50 border border-amber-100 rounded-2xl px-6 py-5 md:px-8 md:py-6">
          <div class="flex items-start gap-3">
            <span class="text-amber-300 text-3xl leading-none mt-1 select-none">"</span>
            <p class="text-gray-600 text-sm md:text-base leading-relaxed whitespace-pre-line flex-1">{{ introText }}</p>
            <span class="text-amber-300 text-3xl leading-none mt-auto select-none self-end">"</span>
          </div>
        </div>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center items-center py-20">
        <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-rose-400"></div>
      </div>

      <!-- Empty state -->
      <div v-else-if="items.length === 0" class="text-center py-20">
        <div class="text-6xl mb-4">🎀</div>
        <p class="text-gray-400 text-lg">愿望清单还是空的，期待满满的心意～</p>
      </div>

      <!-- Gift grid -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        <GiftCard
          v-for="item in items"
          :key="item.id"
          :item="item"
          @click="goToDetail(item.id)"
        />
      </div>
    </main>

    <!-- Footer -->
    <footer class="text-center py-8 text-gray-300 text-xs">
      <a href="/admin" class="hover:text-gray-400 transition-colors">后台管理</a>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import GiftCard from '../components/GiftCard.vue'

const router = useRouter()
const siteName = ref('我的礼物愿望清单')
const introText = ref('')
const items = ref([])
const loading = ref(true)

async function loadData() {
  try {
    const [configRes, itemsRes] = await Promise.all([
      fetch('/api/config'),
      fetch('/api/items?limit=100'),
    ])
    const config = await configRes.json()
    const itemsData = await itemsRes.json()
    siteName.value = config.site_name || '我的礼物愿望清单'
    introText.value = config.intro_text || ''
    items.value = itemsData.items || []
    document.title = siteName.value
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

function goToDetail(id) {
  router.push(`/item/${id}`)
}

onMounted(loadData)
</script>
