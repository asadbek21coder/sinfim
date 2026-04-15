<template>
  <main class="mx-auto max-w-7xl px-5 py-14 lg:px-8">
    <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Entry point</p>
    <h1 class="mt-2 max-w-3xl font-display text-4xl font-bold text-primary lg:text-5xl">Where do you want to go?</h1>
    <p class="mt-4 max-w-2xl text-base leading-7 text-on-surface-variant">
      Step 0 keeps these paths visible. Each one will receive its real backend flow in the planned implementation steps.
    </p>

    <div class="mt-10 grid gap-5 lg:grid-cols-3">
      <RouterLink v-for="item in items" :key="item.to" :to="item.to" class="rounded-lg border border-outline-variant bg-surface-container-lowest p-6 shadow-sm transition hover:border-secondary hover:bg-secondary-container/40">
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">{{ item.kicker }}</p>
        <h2 class="mt-3 font-display text-2xl font-bold text-primary">{{ item.title }}</h2>
        <p class="mt-3 text-sm leading-6 text-on-surface-variant">{{ item.text }}</p>
      </RouterLink>
    </div>
  </main>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

const items = [
  { kicker: 'School team', title: 'Enter my school', text: 'Owner, teacher or mentor login preview.', to: '/auth/login' },
  { kicker: 'New school', title: 'Request access', text: 'Leave a request to create a school workspace.', to: '/apply-school' },
  { kicker: 'Explore', title: 'Try demo school', text: 'Preview the product without real account setup.', to: '/demo' },
]
</script>
