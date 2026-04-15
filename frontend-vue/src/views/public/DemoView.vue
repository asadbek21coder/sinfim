<template>
  <main class="mx-auto max-w-7xl px-5 py-14 lg:px-8">
    <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Demo school</p>
    <h1 class="mt-2 max-w-3xl font-display text-4xl font-bold text-primary lg:text-5xl">A guided demo will live here.</h1>
    <p class="mt-4 max-w-2xl text-base leading-7 text-on-surface-variant">
      Demo seed data and read-only demo behavior are planned for Step 12. For now, use these preview routes to inspect the main shells.
    </p>
    <div class="mt-8 flex flex-col gap-3 sm:flex-row">
      <RouterLink class="btn-primary justify-center" to="/app/dashboard">Owner dashboard preview</RouterLink>
      <RouterLink class="btn-secondary justify-center" to="/learn/dashboard">Student dashboard preview</RouterLink>
    </div>
  </main>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
