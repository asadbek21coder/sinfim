<template>
  <main>
    <section class="relative min-h-[72vh] overflow-hidden bg-primary text-white">
      <img
        class="absolute inset-0 h-full w-full object-cover opacity-55"
        alt="Students learning together in a classroom"
        src="https://images.unsplash.com/photo-1523580846011-d3a5bc25702b?auto=format&fit=crop&w=1800&q=80"
      >
      <div class="absolute inset-0 bg-primary/70" />
      <div class="relative mx-auto flex min-h-[72vh] max-w-7xl flex-col justify-center px-5 py-16 lg:px-8">
        <p class="text-sm font-semibold uppercase tracking-wide text-white/75">Online school platform</p>
        <h1 class="mt-4 max-w-4xl font-display text-5xl font-bold leading-tight lg:text-6xl">
          Build a real online school around your courses, classes and mentors.
        </h1>
        <p class="mt-6 max-w-2xl text-lg leading-8 text-white/78">
          Sinfim.uz helps schools, education centers and course teams organize video lessons, groups, access, homework, tests, mentor review and student progress in one web workspace.
        </p>
        <div class="mt-8 flex flex-col gap-3 sm:flex-row">
          <RouterLink class="btn-primary justify-center" to="/enter">Choose entry point</RouterLink>
          <RouterLink class="inline-flex items-center justify-center gap-2 rounded-lg bg-white px-5 py-2.5 text-sm font-semibold text-primary transition hover:bg-white/90" to="/demo">Explore demo school</RouterLink>
        </div>
      </div>
    </section>

    <section class="section-band">
      <div class="mx-auto grid max-w-7xl gap-5 px-5 py-12 md:grid-cols-4 lg:px-8">
        <div v-for="step in steps" :key="step.title" class="rounded-lg bg-surface-container-lowest p-5 shadow-sm">
          <p class="text-sm font-semibold text-secondary">{{ step.kicker }}</p>
          <h2 class="mt-2 font-display text-xl font-bold text-primary">{{ step.title }}</h2>
          <p class="mt-3 text-sm leading-6 text-on-surface-variant">{{ step.text }}</p>
        </div>
      </div>
    </section>

    <section class="mx-auto grid max-w-7xl gap-5 px-5 py-12 lg:grid-cols-[1fr_380px] lg:px-8">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Step 0 check</p>
        <h2 class="mt-2 font-display text-3xl font-bold text-primary">Local web shell is ready for vertical slices.</h2>
        <p class="mt-3 max-w-2xl text-base leading-7 text-on-surface-variant">
          The first implementation slice should now have public, app, auth and student shells. The next step will connect the school request flow.
        </p>
      </div>
      <HealthStatus />
    </section>
  </main>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import HealthStatus from '@/components/HealthStatus.vue'

const steps = [
  { kicker: '01', title: 'Create school', text: 'Superadmin creates the school workspace and owner access.' },
  { kicker: '02', title: 'Add courses', text: 'Owner or teacher prepares reusable lessons, videos and materials.' },
  { kicker: '03', title: 'Open groups', text: 'Classes connect students, mentors, schedule and manual access.' },
  { kicker: '04', title: 'Review progress', text: 'Mentors review homework while owners track school operations.' },
]
</script>
