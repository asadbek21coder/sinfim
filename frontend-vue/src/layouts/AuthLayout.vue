<template>
  <div class="min-h-screen bg-surface">
    <div class="mx-auto grid min-h-screen max-w-6xl items-center gap-10 px-6 py-10 lg:grid-cols-[1fr_420px]">
      <section class="hidden lg:block">
        <RouterLink to="/" class="inline-flex items-center gap-3 text-primary">
          <span class="flex h-10 w-10 items-center justify-center rounded-lg bg-primary text-sm font-bold text-white">S</span>
          <span class="font-display text-xl font-bold">Sinfim.uz</span>
        </RouterLink>
        <h1 class="mt-12 max-w-2xl font-display text-5xl font-bold leading-tight text-primary">
          Online school operations in one workspace.
        </h1>
        <p class="mt-5 max-w-xl text-lg leading-8 text-on-surface-variant">
          Courses, classes, mentors, students, access, lessons, homework and progress will come together here.
        </p>
      </section>

      <section class="rounded-lg border border-outline-variant bg-surface-container-lowest p-8 shadow-sm">
        <div class="mb-8 lg:hidden">
          <RouterLink to="/" class="inline-flex items-center gap-3 text-primary">
            <span class="flex h-10 w-10 items-center justify-center rounded-lg bg-primary text-sm font-bold text-white">S</span>
            <span class="font-display text-xl font-bold">Sinfim.uz</span>
          </RouterLink>
        </div>
        <slot />
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
