<template>
  <div class="rounded-lg border border-outline-variant bg-surface-container-lowest p-4">
    <div class="flex items-start justify-between gap-4">
      <div>
        <p class="text-sm font-semibold text-on-surface">Backend health</p>
        <p class="mt-1 text-sm text-on-surface-variant">Checks the local API through Vite proxy at <span class="font-mono">/health</span>.</p>
      </div>
      <span
        class="shrink-0 rounded-full px-3 py-1 text-xs font-semibold"
        :class="badgeClass"
      >
        {{ label }}
      </span>
    </div>

    <p v-if="result?.error" class="mt-3 text-sm text-error">{{ result.error }}</p>
    <button class="btn-secondary mt-4" :disabled="loading" @click="loadHealth">
      {{ loading ? 'Checking...' : 'Check again' }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { checkBackendHealth, type HealthResult } from '@/api/health'

const loading = ref(false)
const result = ref<HealthResult | null>(null)

const label = computed(() => {
  if (loading.value) return 'Checking'
  if (!result.value) return 'Not checked'
  return result.value.ok ? 'Online' : 'Offline'
})

const badgeClass = computed(() => {
  if (loading.value || !result.value) return 'bg-surface-container text-on-surface-variant'
  return result.value.ok ? 'bg-emerald-100 text-emerald-700' : 'bg-error-container text-error'
})

async function loadHealth() {
  loading.value = true
  result.value = await checkBackendHealth()
  loading.value = false
}

onMounted(loadHealth)
</script>
