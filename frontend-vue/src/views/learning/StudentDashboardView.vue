<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Learning</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">{{ dashboard?.class.course_title || 'My lessons' }}</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">Watch available lessons, open materials and keep your progress clear.</p>
      </div>
      <select v-if="dashboard?.classes.length" v-model="selectedClassId" class="form-select max-w-xs" @change="load">
        <option v-for="item in dashboard.classes" :key="item.id" :value="item.id">{{ item.name }}</option>
      </select>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>

    <section v-if="dashboard" class="grid gap-4 md:grid-cols-3">
      <div class="card"><p class="text-sm text-on-surface-variant">School</p><p class="mt-2 font-display text-xl font-bold text-primary">{{ dashboard.organization.name }}</p></div>
      <div class="card"><p class="text-sm text-on-surface-variant">Class</p><p class="mt-2 font-display text-xl font-bold text-primary">{{ dashboard.class.name }}</p></div>
      <div class="card"><p class="text-sm text-on-surface-variant">Progress</p><p class="mt-2 font-display text-xl font-bold text-primary">{{ dashboard.progress.percentage }}%</p><p class="mt-1 text-sm text-on-surface-variant">{{ dashboard.progress.completed_lessons }} / {{ dashboard.progress.total_lessons }} completed</p></div>
    </section>

    <section v-if="dashboard?.locked" class="rounded-lg border border-amber-200 bg-amber-50 p-5 text-amber-900">
      <h2 class="font-display text-xl font-bold">Access is not active yet</h2>
      <p class="mt-2 text-sm">Your class access is currently {{ dashboard.class.access_status }}. Lessons will open after the school confirms access.</p>
    </section>

    <div v-else class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
      <table class="data-table">
        <thead><tr><th>Lesson</th><th>Publish</th><th>Content</th><th>Status</th><th>Action</th></tr></thead>
        <tbody>
          <tr v-if="!dashboard || dashboard.lessons.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No published lessons yet.</td></tr>
          <tr v-for="item in dashboard?.lessons ?? []" :key="item.lesson_id">
            <td><p class="font-semibold text-primary">{{ item.order_number }}. {{ item.title }}</p><p class="mt-1 text-xs text-on-surface-variant">{{ item.description || 'Lesson note will appear here.' }}</p></td>
            <td>Day {{ item.publish_day }}</td>
            <td>{{ item.has_video ? 'Video' : 'No video' }} · {{ item.material_count }} material</td>
            <td><span :class="item.status === 'available' ? 'chip-approved' : 'chip-pending'">{{ item.completed ? 'completed' : item.status }}</span></td>
            <td><RouterLink class="btn-secondary" :to="`/learn/lessons/${item.lesson_id}?class_id=${dashboard?.class.id}`">Open</RouterLink></td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { learningApi } from '@/api/learning'
import type { StudentDashboardResponse } from '@/types/learning'

const dashboard = ref<StudentDashboardResponse | null>(null)
const selectedClassId = ref('')
const errorMessage = ref('')

async function load() {
  errorMessage.value = ''
  try {
    const response = await learningApi.getStudentDashboard(selectedClassId.value ? { class_id: selectedClassId.value } : undefined)
    dashboard.value = response.data
    selectedClassId.value = response.data.class.id
  } catch {
    errorMessage.value = 'Learning dashboard could not be loaded.'
  }
}

onMounted(load)
</script>
