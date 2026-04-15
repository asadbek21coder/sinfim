<template>
  <div class="min-h-screen bg-surface text-on-surface">
    <header class="border-b border-outline-variant/70 bg-surface-container-lowest">
      <div class="mx-auto flex h-16 max-w-6xl items-center justify-between px-5 lg:px-8">
        <RouterLink to="/learn/dashboard" class="flex items-center gap-3 text-primary">
          <span class="flex h-9 w-9 items-center justify-center rounded-lg bg-secondary text-sm font-bold text-white">S</span>
          <div>
            <p class="font-display text-base font-bold">Sinfim.uz</p>
            <p class="text-xs text-on-surface-variant">Student learning</p>
          </div>
        </RouterLink>
        <nav class="flex items-center gap-3 text-sm font-medium">
          <RouterLink class="btn-secondary" to="/learn/dashboard">Lessons</RouterLink>
          <RouterLink class="btn-secondary hidden sm:inline-flex" to="/app/dashboard">School admin</RouterLink>
        </nav>
      </div>
    </header>
    <main class="mx-auto max-w-6xl px-5 py-8 lg:px-8">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
