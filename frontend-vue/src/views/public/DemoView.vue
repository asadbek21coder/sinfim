<template>
  <main>
    <section class="bg-primary text-white">
      <div class="mx-auto grid min-h-[58vh] max-w-7xl gap-10 px-5 py-14 lg:grid-cols-[1fr_420px] lg:px-8">
        <div class="flex flex-col justify-center">
          <p class="text-sm font-semibold uppercase tracking-wide text-white/70">Demo school</p>
          <h1 class="mt-3 max-w-3xl font-display text-4xl font-bold leading-tight lg:text-6xl">Explore Sinfim.uz as a real web workspace.</h1>
          <p class="mt-5 max-w-2xl text-base leading-7 text-white/78">
            Open the public school page, sign in as the owner, or enter the student lesson flow with prepared demo data.
          </p>
          <div class="mt-7 flex flex-col gap-3 sm:flex-row">
            <RouterLink class="btn-primary justify-center" :to="demo?.public_url || '/demo-school'">Open public school</RouterLink>
            <RouterLink class="inline-flex items-center justify-center rounded-lg bg-white px-5 py-2.5 text-sm font-semibold text-primary transition hover:bg-white/90" :to="ownerLoginUrl">Owner login</RouterLink>
          </div>
        </div>

        <div class="self-center rounded-lg border border-white/15 bg-white p-6 text-on-surface shadow-xl">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Safe playground</p>
          <h2 class="mt-2 font-display text-2xl font-bold text-primary">Demo data is resettable.</h2>
          <p class="mt-3 text-sm leading-6 text-on-surface-variant">
            This school is marked as demo and can be reseeded by visiting this page. Do not add real student data here.
          </p>
          <div v-if="isLoading" class="mt-5 rounded-lg bg-surface-container p-4 text-sm text-on-surface-variant">Preparing demo school...</div>
          <div v-else-if="errorMessage" class="mt-5 rounded-lg border border-red-200 bg-red-50 p-4 text-sm font-medium text-red-700">{{ errorMessage }}</div>
          <div v-else class="mt-5 space-y-3">
            <div v-for="account in accounts" :key="account.role" class="rounded-lg border border-outline-variant bg-surface-container-lowest p-4">
              <p class="text-xs font-semibold uppercase tracking-wide text-secondary">{{ account.label }}</p>
              <p class="mt-2 text-sm"><span class="font-semibold text-primary">Phone:</span> {{ account.phone }}</p>
              <p class="mt-1 text-sm"><span class="font-semibold text-primary">Password:</span> {{ account.password }}</p>
              <RouterLink class="btn-secondary mt-3 w-full justify-center" :to="account.loginUrl">Use this account</RouterLink>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-5 py-12 lg:px-8">
      <div class="grid gap-5 lg:grid-cols-3">
        <RouterLink class="rounded-lg border border-outline-variant bg-surface-container-lowest p-6 shadow-sm transition hover:border-secondary" :to="demo?.public_url || '/demo-school'">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Public page</p>
          <h2 class="mt-2 font-display text-2xl font-bold text-primary">Demo school profile</h2>
          <p class="mt-3 text-sm leading-6 text-on-surface-variant">Check the public page and lead form that students see first.</p>
        </RouterLink>
        <RouterLink class="rounded-lg border border-outline-variant bg-surface-container-lowest p-6 shadow-sm transition hover:border-secondary" :to="ownerLoginUrl">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Owner</p>
          <h2 class="mt-2 font-display text-2xl font-bold text-primary">Dashboard view</h2>
          <p class="mt-3 text-sm leading-6 text-on-surface-variant">Sign in as the owner to inspect courses, classes, access, leads and activity.</p>
        </RouterLink>
        <RouterLink class="rounded-lg border border-outline-variant bg-surface-container-lowest p-6 shadow-sm transition hover:border-secondary" :to="studentLoginUrl">
          <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Student</p>
          <h2 class="mt-2 font-display text-2xl font-bold text-primary">Lesson flow</h2>
          <p class="mt-3 text-sm leading-6 text-on-surface-variant">Sign in as the student and continue straight to the prepared lesson detail.</p>
        </RouterLink>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { organizationApi } from '@/api/organization'
import type { DemoAccessResponse } from '@/types/organization'

const demo = ref<DemoAccessResponse | null>(null)
const isLoading = ref(true)
const errorMessage = ref('')

const accounts = computed(() => {
  if (!demo.value) return []
  return [
    {
      role: demo.value.owner.role,
      label: 'Owner account',
      phone: demo.value.owner.phone_number,
      password: demo.value.owner.password,
      loginUrl: ownerLoginUrl.value,
    },
    {
      role: demo.value.student.role,
      label: 'Student account',
      phone: demo.value.student.phone_number,
      password: demo.value.student.password,
      loginUrl: studentLoginUrl.value,
    },
    {
      role: demo.value.mentor.role,
      label: 'Mentor account',
      phone: demo.value.mentor.phone_number,
      password: demo.value.mentor.password,
      loginUrl: mentorLoginUrl.value,
    },
  ]
})

const ownerLoginUrl = computed(() => {
  const phone = demo.value?.owner.phone_number || '+998900000777'
  const password = demo.value?.owner.password || 'DemoPass123'
  const redirect = demo.value?.owner_url || '/app/dashboard'
  return `/auth/login?phone=${encodeURIComponent(phone)}&password=${encodeURIComponent(password)}&redirect=${encodeURIComponent(redirect)}`
})

const studentLoginUrl = computed(() => {
  const phone = demo.value?.student.phone_number || '+998900000778'
  const password = demo.value?.student.password || 'DemoPass123'
  const redirect = demo.value?.student_url || '/learn/dashboard'
  return `/auth/login?phone=${encodeURIComponent(phone)}&password=${encodeURIComponent(password)}&redirect=${encodeURIComponent(redirect)}`
})

const mentorLoginUrl = computed(() => {
  const phone = demo.value?.mentor.phone_number || '+998900000779'
  const password = demo.value?.mentor.password || 'DemoPass123'
  return `/auth/login?phone=${encodeURIComponent(phone)}&password=${encodeURIComponent(password)}&redirect=${encodeURIComponent('/app/homework/review')}`
})

async function loadDemo() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const response = await organizationApi.getDemoAccess()
    demo.value = response.data
  } catch {
    errorMessage.value = 'Demo school could not be prepared. Make sure the backend is running.'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadDemo)
</script>
