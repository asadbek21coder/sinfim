<template>
  <section>
    <div class="mb-8 flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Owner workspace</p>
        <h1 class="mt-2 font-display text-3xl font-bold text-primary lg:text-4xl">School dashboard</h1>
        <p class="mt-3 max-w-3xl text-base leading-7 text-on-surface-variant">
          Step 0 shell for course operations, classes, mentors, students, leads, homework, and manual access confirmation.
        </p>
      </div>
      <RouterLink class="btn-primary" to="/app/courses">Create course later</RouterLink>
    </div>

    <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-4">
      <div v-for="metric in metrics" :key="metric.label" class="card border-l-4 border-l-secondary">
        <p class="text-sm text-on-surface-variant">{{ metric.label }}</p>
        <p class="mt-3 font-display text-3xl font-bold text-primary">{{ metric.value }}</p>
        <p class="mt-1 text-xs text-on-surface-variant">{{ metric.note }}</p>
      </div>
    </div>

    <div class="mt-6 grid gap-5 lg:grid-cols-[1fr_380px]">
      <div class="card">
        <h2 class="font-display text-xl font-bold text-primary">Step 0 route map</h2>
        <div class="mt-5 grid gap-3 sm:grid-cols-2">
          <RouterLink v-for="item in routeMap" :key="item.to" :to="item.to" class="rounded-lg border border-outline-variant p-4 transition hover:border-secondary hover:bg-secondary-container/50">
            <p class="font-semibold text-primary">{{ item.label }}</p>
            <p class="mt-1 text-sm text-on-surface-variant">{{ item.note }}</p>
          </RouterLink>
        </div>
      </div>
      <HealthStatus />
    </div>
  </section>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import HealthStatus from '@/components/HealthStatus.vue'

const metrics = [
  { label: 'Active courses', value: '0', note: 'Step 5 will connect real data.' },
  { label: 'Classes', value: '0', note: 'Step 6 will add groups.' },
  { label: 'Students', value: '0', note: 'Manual enrollment comes later.' },
  { label: 'Pending homework', value: '0', note: 'Step 10 will power review.' },
]

const routeMap = [
  { label: 'Courses', to: '/app/courses', note: 'Reusable course packages.' },
  { label: 'Classes', to: '/app/classes', note: 'Live cohorts and access.' },
  { label: 'Leads', to: '/app/leads', note: 'Public requests from school pages.' },
  { label: 'Student dashboard', to: '/learn/dashboard', note: 'Learning experience shell.' },
]
</script>
