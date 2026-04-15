<template>
  <section>
    <div class="mb-8 flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">{{ eyebrow }}</p>
        <h1 class="mt-2 font-display text-3xl font-bold text-primary lg:text-4xl">{{ title }}</h1>
        <p class="mt-3 max-w-3xl text-base leading-7 text-on-surface-variant">{{ description }}</p>
      </div>
      <span class="w-fit rounded-full bg-secondary-container px-4 py-2 text-sm font-semibold text-secondary">{{ status }}</span>
    </div>

    <div class="grid gap-5 lg:grid-cols-[1.3fr_0.7fr]">
      <div class="card">
        <h2 class="font-display text-xl font-bold text-primary">Manual test target</h2>
        <p class="mt-3 text-sm leading-6 text-on-surface-variant">
          This screen is part of the Step 0 route and layout skeleton. In its own implementation step, it will receive real backend data, forms, permissions, empty states, and validation.
        </p>
        <div class="mt-6 grid gap-3 sm:grid-cols-3">
          <div class="rounded-lg bg-surface-container-low p-4">
            <p class="text-xs font-semibold uppercase text-on-surface-variant">Route</p>
            <p class="mt-1 break-all font-mono text-sm text-primary">{{ route.path }}</p>
          </div>
          <div class="rounded-lg bg-surface-container-low p-4">
            <p class="text-xs font-semibold uppercase text-on-surface-variant">Layout</p>
            <p class="mt-1 text-sm font-semibold text-primary">{{ route.meta.layout ?? 'public' }}</p>
          </div>
          <div class="rounded-lg bg-surface-container-low p-4">
            <p class="text-xs font-semibold uppercase text-on-surface-variant">Status</p>
            <p class="mt-1 text-sm font-semibold text-primary">Ready for next slice</p>
          </div>
        </div>
      </div>

      <HealthStatus />
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import HealthStatus from '@/components/HealthStatus.vue'

const route = useRoute()

const eyebrow = computed(() => route.meta.eyebrow ?? 'Sinfim.uz')
const title = computed(() => route.meta.title ?? 'Screen placeholder')
const description = computed(() => route.meta.description ?? 'This route is reserved for a later implementation step.')
const status = computed(() => route.meta.status ?? 'Step 0')
</script>
