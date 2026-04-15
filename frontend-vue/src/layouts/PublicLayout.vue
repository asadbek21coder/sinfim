<template>
  <div class="min-h-screen bg-surface text-on-surface">
    <header class="sticky top-0 z-20 border-b border-outline-variant/70 bg-surface/95 backdrop-blur">
      <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-5 lg:px-8">
        <RouterLink to="/" class="flex items-center gap-3 text-primary">
          <span class="flex h-9 w-9 items-center justify-center rounded-lg bg-primary text-sm font-bold text-white">S</span>
          <span class="font-display text-lg font-bold">Sinfim.uz</span>
        </RouterLink>
        <nav class="hidden items-center gap-6 text-sm font-medium text-on-surface-variant md:flex">
          <RouterLink class="hover:text-primary" to="/enter">Enter</RouterLink>
          <RouterLink class="hover:text-primary" to="/demo">Demo</RouterLink>
          <RouterLink class="hover:text-primary" to="/apply-school">Apply</RouterLink>
        </nav>
        <RouterLink class="btn-primary" to="/enter">Get started</RouterLink>
      </div>
    </header>

    <slot />

    <footer class="border-t border-outline-variant/70 bg-surface-container-lowest">
      <div class="mx-auto flex max-w-7xl flex-col gap-3 px-5 py-8 text-sm text-on-surface-variant sm:flex-row sm:items-center sm:justify-between lg:px-8">
        <p>Sinfim.uz</p>
        <p>sinfim.uz/{school-slug}</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
