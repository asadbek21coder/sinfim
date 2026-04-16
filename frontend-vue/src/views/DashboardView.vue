<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Owner workspace</p>
        <h1 class="mt-2 font-display text-3xl font-bold text-primary lg:text-4xl">{{ dashboard?.organization.name || 'School dashboard' }}</h1>
        <p class="mt-3 max-w-3xl text-base leading-7 text-on-surface-variant">Courses, classes, access confirmations, leads and homework queues in one place.</p>
      </div>
      <div class="flex flex-wrap gap-3">
        <RouterLink class="btn-secondary" to="/app/homework/review">Review homework</RouterLink>
        <RouterLink class="btn-primary" to="/app/courses">Create course</RouterLink>
      </div>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>

    <div class="grid gap-5 md:grid-cols-2 xl:grid-cols-4">
      <div v-for="metric in metrics" :key="metric.label" class="card border-l-4" :class="metric.border">
        <p class="text-sm text-on-surface-variant">{{ metric.label }}</p>
        <p class="mt-3 font-display text-3xl font-bold text-primary">{{ metric.value }}</p>
        <p class="mt-1 text-xs text-on-surface-variant">{{ metric.note }}</p>
      </div>
    </div>

    <section class="grid gap-5 xl:grid-cols-[1.25fr_0.75fr]">
      <div class="card space-y-4">
        <div class="flex items-center justify-between gap-3">
          <div>
            <h2 class="font-display text-xl font-bold text-primary">Course and class progress</h2>
            <p class="mt-1 text-sm text-on-surface-variant">Active course packages with student progress and pending homework.</p>
          </div>
          <RouterLink class="text-sm font-semibold text-secondary" to="/app/courses">Courses</RouterLink>
        </div>
        <div v-if="!dashboard?.course_progress.length" class="rounded-lg border border-dashed border-outline-variant p-4 text-sm text-on-surface-variant">No courses yet.</div>
        <div v-for="course in dashboard?.course_progress ?? []" :key="course.course_id" class="rounded-lg border border-outline-variant p-4">
          <div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <p class="font-display text-lg font-bold text-primary">{{ course.course_title }}</p>
              <p class="mt-1 text-sm text-on-surface-variant">{{ course.class_count }} classes · {{ course.student_count }} students · {{ course.lesson_count }} lessons</p>
            </div>
            <span class="rounded-lg bg-surface-container px-3 py-2 text-sm font-bold text-primary">{{ course.progress_percent }}%</span>
          </div>
          <div class="mt-4 h-2 rounded-full bg-outline-variant">
            <div class="h-2 rounded-full bg-secondary" :style="{ width: `${Math.min(course.progress_percent, 100)}%` }" />
          </div>
          <p class="mt-2 text-xs text-on-surface-variant">{{ course.completion_count }} lesson completions · {{ course.pending_homework }} pending homework</p>
        </div>
      </div>

      <div class="card space-y-4">
        <div class="flex items-center justify-between gap-3">
          <h2 class="font-display text-xl font-bold text-primary">Quick actions</h2>
          <button class="text-sm font-semibold text-secondary" type="button" @click="loadDashboard">Refresh</button>
        </div>
        <div class="grid gap-3">
          <RouterLink v-for="action in quickActions" :key="action.to" :to="action.to" class="rounded-lg border border-outline-variant p-4 transition hover:border-secondary hover:bg-secondary-container/50">
            <p class="font-semibold text-primary">{{ action.label }}</p>
            <p class="mt-1 text-sm text-on-surface-variant">{{ action.note }}</p>
          </RouterLink>
        </div>
      </div>
    </section>

    <section class="grid gap-5 xl:grid-cols-3">
      <div class="card space-y-4">
        <div class="flex items-center justify-between gap-3">
          <h2 class="font-display text-xl font-bold text-primary">Pending homework</h2>
          <RouterLink class="text-sm font-semibold text-secondary" to="/app/homework/review">Open</RouterLink>
        </div>
        <p v-if="!dashboard?.pending_homework.length" class="text-sm text-on-surface-variant">No pending homework.</p>
        <div v-for="item in dashboard?.pending_homework ?? []" :key="item.submission_id" class="rounded-lg border border-outline-variant p-3">
          <p class="font-semibold text-primary">{{ item.student_full_name }}</p>
          <p class="mt-1 text-sm text-on-surface-variant">{{ item.homework_title }} · {{ item.class_name }}</p>
          <p class="mt-1 text-xs uppercase text-on-surface-variant">{{ item.submission_type }}</p>
        </div>
      </div>

      <div class="card space-y-4">
        <div class="flex items-center justify-between gap-3">
          <h2 class="font-display text-xl font-bold text-primary">Access confirmations</h2>
          <RouterLink class="text-sm font-semibold text-secondary" to="/app/classes">Classes</RouterLink>
        </div>
        <p v-if="!dashboard?.pending_access.length" class="text-sm text-on-surface-variant">No pending access confirmations.</p>
        <div v-for="item in dashboard?.pending_access ?? []" :key="`${item.class_id}-${item.student_user_id}`" class="rounded-lg border border-outline-variant p-3">
          <p class="font-semibold text-primary">{{ item.student_name }}</p>
          <p class="mt-1 text-sm text-on-surface-variant">{{ item.class_name }} · {{ item.phone_number }}</p>
          <p class="mt-1 text-xs uppercase text-on-surface-variant">{{ item.access_status }} / {{ item.payment_status }}</p>
        </div>
      </div>

      <div class="card space-y-4">
        <h2 class="font-display text-xl font-bold text-primary">Recent activity</h2>
        <p v-if="!dashboard?.recent_activity.length" class="text-sm text-on-surface-variant">No recent activity yet.</p>
        <div v-for="item in dashboard?.recent_activity ?? []" :key="`${item.type}-${item.created_at}-${item.title}`" class="rounded-lg border border-outline-variant p-3">
          <p class="text-xs font-bold uppercase text-secondary">{{ item.type }}</p>
          <p class="mt-1 font-semibold text-primary">{{ item.title }}</p>
          <p class="mt-1 text-sm text-on-surface-variant">{{ item.subtitle }}</p>
        </div>
      </div>
    </section>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { organizationApi } from '@/api/organization'
import type { OwnerDashboardResponse } from '@/types/organization'

const dashboard = ref<OwnerDashboardResponse | null>(null)
const errorMessage = ref('')

const metrics = computed(() => {
  const data = dashboard.value?.metrics
  return [
    { label: 'Active courses', value: data?.active_courses ?? 0, note: 'Course packages not archived', border: 'border-l-secondary' },
    { label: 'Classes', value: data?.active_classes ?? 0, note: 'Running groups and cohorts', border: 'border-l-primary' },
    { label: 'Students', value: data?.active_students ?? 0, note: 'Active enrolled students', border: 'border-l-emerald-600' },
    { label: 'Pending homework', value: data?.pending_homework ?? 0, note: `${data?.pending_access ?? 0} access confirmations`, border: 'border-l-red-500' },
  ]
})

const quickActions = [
  { label: 'Create course', to: '/app/courses', note: 'Build a reusable course package.' },
  { label: 'Open classes', to: '/app/classes', note: 'Manage groups, students and access.' },
  { label: 'Review homework', to: '/app/homework/review', note: 'Clear pending student submissions.' },
  { label: 'Check leads', to: '/app/leads', note: 'Follow up public requests.' },
]

async function loadDashboard() {
  errorMessage.value = ''
  try {
    const response = await organizationApi.getOwnerDashboard()
    dashboard.value = response.data
  } catch {
    errorMessage.value = 'Dashboard could not be loaded. Make sure this user has an owner, teacher or mentor workspace.'
  }
}

onMounted(loadDashboard)
</script>
