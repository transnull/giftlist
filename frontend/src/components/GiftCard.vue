<template>
  <div
    class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden cursor-pointer
           hover:-translate-y-1 hover:shadow-md transition-all duration-200 w-full"
    @click="$emit('click')"
  >
    <!-- Image -->
    <div class="relative overflow-hidden bg-gray-50" style="height: 300px;">
      <img
        :src="item.image_url"
        :alt="item.name"
        class="w-full h-full object-cover"
        @error="onImgError"
      />
    </div>
    <!-- Info -->
    <div class="p-4">
      <h3 class="font-semibold text-gray-800 truncate text-sm mb-2">{{ item.name }}</h3>
      <div class="flex items-center justify-between gap-2 flex-wrap">
        <span :class="priorityClass(item.priority)" class="px-2 py-0.5 rounded-full text-xs font-medium">
          {{ priorityLabel(item.priority) }}
        </span>
        <span v-if="item.estimated_cost" class="text-rose-500 font-bold text-sm">
          ¥{{ formatCost(item.estimated_cost) }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({ item: { type: Object, required: true } })
defineEmits(['click'])

const PLACEHOLDER = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='300' height='300' viewBox='0 0 300 300'%3E%3Crect width='300' height='300' fill='%23f3f4f6'/%3E%3Ctext x='50%25' y='50%25' text-anchor='middle' dy='.3em' fill='%239ca3af' font-size='14' font-family='sans-serif'%3E图片加载失败%3C/text%3E%3C/svg%3E`

function onImgError(e) { e.target.src = PLACEHOLDER }
function priorityClass(p) {
  return { high: 'priority-high', medium: 'priority-medium', low: 'priority-low' }[p] || 'priority-low'
}
function priorityLabel(p) {
  return { high: '❤️ 高', medium: '💛 中', low: '💚 低' }[p] || p
}
function formatCost(cost) {
  return Number(cost).toLocaleString('zh-CN', { minimumFractionDigits: 0, maximumFractionDigits: 2 })
}
</script>
