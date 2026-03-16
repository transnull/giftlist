<template>
  <div class="min-h-screen bg-[#f9f7f5]">
    <!-- Top bar -->
    <header class="bg-white border-b border-gray-100 shadow-sm sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <span class="text-xl">🎁</span>
          <h1 class="font-bold text-gray-800 text-lg">愿望清单管理</h1>
        </div>
        <div class="flex items-center gap-3">
          <router-link to="/" class="text-sm text-gray-400 hover:text-gray-600 transition-colors">公开页面</router-link>
          <button @click="handleLogout" class="btn-secondary text-sm px-3 py-1.5">退出登录</button>
        </div>
      </div>
    </header>

    <main class="max-w-7xl mx-auto px-4 py-6">
      <!-- Config section -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
        <!-- Site name -->
        <div class="bg-white rounded-xl border border-gray-100 p-5 shadow-sm">
          <div class="flex items-start justify-between">
            <div class="flex-1 min-w-0">
              <p class="text-xs text-gray-400 mb-1 uppercase tracking-wide">站点名称</p>
              <p class="text-gray-800 font-semibold text-lg truncate">{{ config.site_name }}</p>
            </div>
            <button @click="openConfigModal('site_name')" class="ml-3 text-gray-400 hover:text-rose-500 transition-colors p-1.5 rounded-lg hover:bg-rose-50">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Intro text -->
        <div class="bg-white rounded-xl border border-gray-100 p-5 shadow-sm">
          <div class="flex items-start justify-between">
            <div class="flex-1 min-w-0">
              <p class="text-xs text-gray-400 mb-1 uppercase tracking-wide">自定义引言</p>
              <p class="text-gray-600 text-sm line-clamp-2">{{ config.intro_text }}</p>
            </div>
            <button @click="openConfigModal('intro_text')" class="ml-3 text-gray-400 hover:text-rose-500 transition-colors p-1.5 rounded-lg hover:bg-rose-50 flex-shrink-0">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Items section -->
      <div class="bg-white rounded-xl border border-gray-100 shadow-sm overflow-hidden">
        <!-- Toolbar -->
        <div class="p-4 border-b border-gray-100 flex flex-wrap gap-3 items-center">
          <div class="flex-1 flex flex-wrap gap-3">
            <input v-model="filter.category" @input="debouncedLoad" type="text" placeholder="按类别筛选" class="input-field w-36" />
            <select v-model="filter.priority" @change="loadItems" class="input-field w-32">
              <option value="">全部优先级</option>
              <option value="high">高优先级</option>
              <option value="medium">中优先级</option>
              <option value="low">低优先级</option>
            </select>
            <select v-model="filter.sort" @change="loadItems" class="input-field w-36">
              <option value="created_at">按添加时间</option>
              <option value="estimated_cost">按价格</option>
            </select>
            <select v-model="filter.order" @change="loadItems" class="input-field w-24">
              <option value="desc">降序</option>
              <option value="asc">升序</option>
            </select>
          </div>
          <button @click="openItemModal(null)" class="btn-primary flex items-center gap-2 whitespace-nowrap">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
            新增物品
          </button>
        </div>

        <!-- Table -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="bg-gray-50 border-b border-gray-100">
                <th class="text-left px-4 py-3 text-gray-500 font-medium">物品名称</th>
                <th class="text-left px-4 py-3 text-gray-500 font-medium hidden md:table-cell">类别</th>
                <th class="text-left px-4 py-3 text-gray-500 font-medium">优先级</th>
                <th class="text-left px-4 py-3 text-gray-500 font-medium hidden sm:table-cell">估计费用</th>
                <th class="text-left px-4 py-3 text-gray-500 font-medium hidden lg:table-cell">图片</th>
                <th class="text-right px-4 py-3 text-gray-500 font-medium">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loading">
                <td colspan="6" class="text-center py-10 text-gray-400">加载中...</td>
              </tr>
              <tr v-else-if="items.length === 0">
                <td colspan="6" class="text-center py-10 text-gray-400">暂无物品，点击"新增物品"添加</td>
              </tr>
              <tr
                v-else
                v-for="item in items"
                :key="item.id"
                class="border-b border-gray-50 hover:bg-gray-50 transition-colors"
              >
                <td class="px-4 py-3">
                  <span class="font-medium text-gray-800 max-w-[200px] truncate block">{{ item.name }}</span>
                </td>
                <td class="px-4 py-3 text-gray-500 hidden md:table-cell">{{ item.category || '—' }}</td>
                <td class="px-4 py-3">
                  <span :class="priorityClass(item.priority)" class="px-2 py-0.5 rounded-full text-xs font-medium">
                    {{ priorityLabel(item.priority) }}
                  </span>
                </td>
                <td class="px-4 py-3 text-gray-700 hidden sm:table-cell">
                  {{ item.estimated_cost ? '¥' + formatCost(item.estimated_cost) : '—' }}
                </td>
                <td class="px-4 py-3 hidden lg:table-cell">
                  <img
                    :src="item.image_url"
                    :alt="item.name"
                    class="w-14 h-14 rounded-lg object-cover border border-gray-100"
                    @error="(e) => e.target.src = placeholderSvg"
                  />
                </td>
                <td class="px-4 py-3 text-right">
                  <div class="flex items-center justify-end gap-2">
                    <button @click="openItemModal(item)" class="text-gray-400 hover:text-blue-500 transition-colors p-1.5 rounded-lg hover:bg-blue-50">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                      </svg>
                    </button>
                    <button @click="confirmDelete(item)" class="text-gray-400 hover:text-red-500 transition-colors p-1.5 rounded-lg hover:bg-red-50">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                      </svg>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="px-4 py-3 border-t border-gray-100 flex items-center justify-between text-sm text-gray-500">
          <span>共 {{ total }} 条记录</span>
          <div class="flex items-center gap-2">
            <button @click="prevPage" :disabled="filter.page <= 1" class="px-3 py-1 rounded-lg border border-gray-200 disabled:opacity-40 hover:bg-gray-50 transition-colors">上一页</button>
            <span>第 {{ filter.page }} / {{ totalPages }} 页</span>
            <button @click="nextPage" :disabled="filter.page >= totalPages" class="px-3 py-1 rounded-lg border border-gray-200 disabled:opacity-40 hover:bg-gray-50 transition-colors">下一页</button>
          </div>
        </div>
      </div>
    </main>

    <!-- Config edit modal -->
    <Teleport to="body">
      <div v-if="configModal.show" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="configModal.show = false">
        <div class="bg-white rounded-2xl shadow-xl w-full max-w-md p-6">
          <h3 class="font-bold text-lg text-gray-800 mb-4">
            {{ configModal.key === 'site_name' ? '修改站点名称' : '修改自定义引言' }}
          </h3>
          <div class="mb-4">
            <textarea
              v-if="configModal.key === 'intro_text'"
              v-model="configModal.value"
              rows="5"
              class="input-field resize-none"
              placeholder="输入引言内容..."
            ></textarea>
            <input
              v-else
              v-model="configModal.value"
              type="text"
              class="input-field"
              placeholder="输入站点名称..."
            />
          </div>
          <div class="flex gap-3 justify-end">
            <button @click="configModal.show = false" class="btn-secondary">取消</button>
            <button @click="saveConfig" class="btn-primary" :disabled="configModal.saving">
              {{ configModal.saving ? '保存中...' : '保存' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Item modal -->
    <Teleport to="body">
      <div v-if="itemModal.show" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-50 flex items-center justify-center p-4 overflow-y-auto" @click.self="itemModal.show = false">
        <div class="bg-white rounded-2xl shadow-xl w-full max-w-lg my-4">
          <div class="p-6 border-b border-gray-100">
            <h3 class="font-bold text-lg text-gray-800">{{ itemModal.isEdit ? '编辑物品' : '新增物品' }}</h3>
          </div>
          <div class="p-6 space-y-4">
            <!-- Name -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">物品名称 <span class="text-red-400">*</span></label>
              <input v-model="itemForm.name" type="text" class="input-field" placeholder="例：AirPods Pro" />
              <p v-if="formErrors.name" class="text-red-400 text-xs mt-1">{{ formErrors.name }}</p>
            </div>
            <!-- Category -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">类别</label>
              <input v-model="itemForm.category" type="text" class="input-field" placeholder="例：数码、护肤、书籍..." />
            </div>
            <!-- Priority -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">优先级 <span class="text-red-400">*</span></label>
              <div class="flex gap-3">
                <label v-for="opt in priorityOptions" :key="opt.value" class="flex-1 cursor-pointer">
                  <input v-model="itemForm.priority" type="radio" :value="opt.value" class="sr-only" />
                  <div :class="[
                    'text-center py-2 rounded-lg border-2 text-xs font-semibold transition-all',
                    itemForm.priority === opt.value ? opt.activeClass : 'border-gray-200 text-gray-500 hover:border-gray-300'
                  ]">{{ opt.label }}</div>
                </label>
              </div>
              <p v-if="formErrors.priority" class="text-red-400 text-xs mt-1">{{ formErrors.priority }}</p>
            </div>
            <!-- Cost -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">估计费用（¥）</label>
              <input v-model="itemForm.estimated_cost" type="number" min="0" step="0.01" class="input-field" placeholder="例：999" />
            </div>
            <!-- Image URL -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">图片 URL <span class="text-red-400">*</span></label>
              <input v-model="itemForm.image_url" type="url" class="input-field" placeholder="https://..." />
              <p v-if="formErrors.image_url" class="text-red-400 text-xs mt-1">{{ formErrors.image_url }}</p>
              <!-- Preview -->
              <div v-if="itemForm.image_url" class="mt-2">
                <img
                  :src="itemForm.image_url"
                  alt="预览"
                  class="w-[150px] h-[150px] object-cover rounded-lg border border-gray-200"
                  @error="(e) => e.target.src = placeholderSvg"
                />
              </div>
            </div>
            <!-- Product link -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">商品链接</label>
              <input v-model="itemForm.product_link" type="url" class="input-field" placeholder="https://..." />
              <p v-if="formErrors.product_link" class="text-red-400 text-xs mt-1">{{ formErrors.product_link }}</p>
            </div>
            <!-- Remark -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">备注</label>
              <textarea v-model="itemForm.remark" rows="3" class="input-field resize-none" placeholder="可选：颜色、尺码偏好等"></textarea>
            </div>
          </div>
          <!-- Error -->
          <div v-if="itemModal.error" class="px-6 pb-2">
            <p class="text-red-500 text-sm bg-red-50 px-3 py-2 rounded-lg">{{ itemModal.error }}</p>
          </div>
          <div class="p-6 border-t border-gray-100 flex gap-3 justify-end">
            <button @click="itemModal.show = false" class="btn-secondary">取消</button>
            <button @click="saveItem" class="btn-primary" :disabled="itemModal.saving">
              {{ itemModal.saving ? '保存中...' : (itemModal.isEdit ? '保存修改' : '添加物品') }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Delete confirm modal -->
    <Teleport to="body">
      <div v-if="deleteModal.show" class="fixed inset-0 bg-black/40 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="deleteModal.show = false">
        <div class="bg-white rounded-2xl shadow-xl w-full max-w-sm p-6">
          <div class="text-center mb-4">
            <div class="text-4xl mb-3">🗑️</div>
            <h3 class="font-bold text-lg text-gray-800">确认删除？</h3>
            <p class="text-gray-500 text-sm mt-1">将删除"{{ deleteModal.item?.name }}"，此操作不可撤销。</p>
          </div>
          <div class="flex gap-3">
            <button @click="deleteModal.show = false" class="btn-secondary flex-1">取消</button>
            <button @click="executeDelete" class="btn-danger flex-1" :disabled="deleteModal.loading">
              {{ deleteModal.loading ? '删除中...' : '确认删除' }}
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const auth = useAuthStore()

// ---- Data ----
const config = ref({ site_name: '', intro_text: '' })
const items = ref([])
const total = ref(0)
const loading = ref(false)

const filter = reactive({ category: '', priority: '', sort: 'created_at', order: 'desc', page: 1, limit: 10 })
const totalPages = computed(() => Math.max(1, Math.ceil(total.value / filter.limit)))

const placeholderSvg = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='60' height='60' viewBox='0 0 60 60'%3E%3Crect width='60' height='60' fill='%23f3f4f6'/%3E%3Ctext x='50%25' y='55%25' text-anchor='middle' fill='%239ca3af' font-size='9' font-family='sans-serif'%3E无图片%3C/text%3E%3C/svg%3E`

// ---- Config modal ----
const configModal = reactive({ show: false, key: '', value: '', saving: false })

function openConfigModal(key) {
  configModal.key = key
  configModal.value = config.value[key] || ''
  configModal.show = true
}

async function saveConfig() {
  configModal.saving = true
  try {
    const body = {}
    body[configModal.key] = configModal.value
    const res = await fetch('/api/config', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })
    if (!res.ok) throw new Error((await res.json()).error)
    const updated = await res.json()
    config.value = updated
    configModal.show = false
  } catch (e) {
    alert('保存失败：' + e.message)
  } finally {
    configModal.saving = false
  }
}

// ---- Item form ----
const defaultForm = () => ({ name: '', category: '', priority: 'medium', estimated_cost: '', image_url: '', product_link: '', remark: '' })
const itemForm = reactive(defaultForm())
const formErrors = reactive({})
const itemModal = reactive({ show: false, isEdit: false, editId: null, saving: false, error: '' })

const priorityOptions = [
  { value: 'high', label: '❤️ 高', activeClass: 'border-red-400 bg-red-50 text-red-600' },
  { value: 'medium', label: '💛 中', activeClass: 'border-yellow-400 bg-yellow-50 text-yellow-600' },
  { value: 'low', label: '💚 低', activeClass: 'border-green-400 bg-green-50 text-green-600' },
]

function openItemModal(item) {
  Object.assign(itemForm, defaultForm())
  Object.keys(formErrors).forEach(k => delete formErrors[k])
  itemModal.error = ''
  if (item) {
    itemModal.isEdit = true
    itemModal.editId = item.id
    Object.assign(itemForm, {
      name: item.name || '',
      category: item.category || '',
      priority: item.priority || 'medium',
      estimated_cost: item.estimated_cost ?? '',
      image_url: item.image_url || '',
      product_link: item.product_link || '',
      remark: item.remark || '',
    })
  } else {
    itemModal.isEdit = false
    itemModal.editId = null
  }
  itemModal.show = true
}

function validateForm() {
  Object.keys(formErrors).forEach(k => delete formErrors[k])
  let valid = true
  if (!itemForm.name.trim()) { formErrors.name = '物品名称不能为空'; valid = false }
  if (!itemForm.priority) { formErrors.priority = '请选择优先级'; valid = false }
  if (!itemForm.image_url.trim()) { formErrors.image_url = '图片 URL 不能为空'; valid = false }
  if (itemForm.product_link && !isValidUrl(itemForm.product_link)) { formErrors.product_link = '请输入有效的 URL'; valid = false }
  return valid
}

function isValidUrl(url) {
  try { new URL(url); return true } catch { return false }
}

async function saveItem() {
  if (!validateForm()) return
  itemModal.saving = true
  itemModal.error = ''
  try {
    const body = {
      name: itemForm.name.trim(),
      category: itemForm.category.trim() || null,
      priority: itemForm.priority,
      estimated_cost: itemForm.estimated_cost !== '' ? parseFloat(itemForm.estimated_cost) : null,
      image_url: itemForm.image_url.trim(),
      product_link: itemForm.product_link.trim() || null,
      remark: itemForm.remark.trim() || null,
    }
    const url = itemModal.isEdit ? `/api/items/${itemModal.editId}` : '/api/items'
    const method = itemModal.isEdit ? 'PUT' : 'POST'
    const res = await fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })
    if (!res.ok) { const d = await res.json(); throw new Error(d.error || '保存失败') }
    itemModal.show = false
    loadItems()
  } catch (e) {
    itemModal.error = e.message
  } finally {
    itemModal.saving = false
  }
}

// ---- Delete ----
const deleteModal = reactive({ show: false, item: null, loading: false })

function confirmDelete(item) {
  deleteModal.item = item
  deleteModal.show = true
  deleteModal.loading = false
}

async function executeDelete() {
  if (!deleteModal.item) return
  deleteModal.loading = true
  try {
    await fetch(`/api/items/${deleteModal.item.id}`, { method: 'DELETE' })
    deleteModal.show = false
    loadItems()
  } catch (e) {
    alert('删除失败：' + e.message)
  } finally {
    deleteModal.loading = false
  }
}

// ---- Load items ----
async function loadItems() {
  loading.value = true
  filter.page = filter.page || 1
  const params = new URLSearchParams({
    page: filter.page,
    limit: filter.limit,
    sort: filter.sort,
    order: filter.order,
  })
  if (filter.category) params.set('category', filter.category)
  if (filter.priority) params.set('priority', filter.priority)
  try {
    const res = await fetch('/api/items?' + params)
    const data = await res.json()
    items.value = data.items || []
    total.value = data.total || 0
  } catch (e) { console.error(e) }
  finally { loading.value = false }
}

function prevPage() { if (filter.page > 1) { filter.page--; loadItems() } }
function nextPage() { if (filter.page < totalPages.value) { filter.page++; loadItems() } }

// ---- Debounce for category input ----
let debounceTimer = null
function debouncedLoad() {
  clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => { filter.page = 1; loadItems() }, 350)
}

// ---- Helpers ----
function priorityClass(p) {
  return { high: 'priority-high', medium: 'priority-medium', low: 'priority-low' }[p] || 'priority-low'
}
function priorityLabel(p) {
  return { high: '高', medium: '中', low: '低' }[p] || p
}
function formatCost(cost) {
  return Number(cost).toLocaleString('zh-CN', { minimumFractionDigits: 0, maximumFractionDigits: 2 })
}

// ---- Logout ----
async function handleLogout() {
  await auth.logout()
  router.push('/admin/login')
}

// ---- Load config ----
async function loadConfig() {
  const res = await fetch('/api/config')
  config.value = await res.json()
}

onMounted(() => {
  loadConfig()
  loadItems()
})
</script>
