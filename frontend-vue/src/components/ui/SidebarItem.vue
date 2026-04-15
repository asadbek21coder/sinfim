<template>
  <RouterLink
    :to="to"
    class="flex items-center gap-3 mx-3 px-3 py-3 rounded-lg text-sm transition-colors"
    :class="isActive
      ? 'bg-sidebar-active border-l-4 border-sidebar-active-indicator text-white font-semibold pl-2'
      : 'text-white/70 hover:bg-white/10 hover:text-white'"
  >
    <span class="flex h-6 w-6 flex-shrink-0 items-center justify-center rounded bg-white/10 text-[11px] font-bold">{{ icon }}</span>
    <span class="truncate">{{ label }}</span>
  </RouterLink>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const props = defineProps<{ icon: string; label: string; to: string }>()
const route = useRoute()

const isActive = computed(() =>
  route.path === props.to || route.path.startsWith(props.to + '/'),
)
</script>
