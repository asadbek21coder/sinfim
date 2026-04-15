<template>
  <section class="space-y-6">
    <div>
      <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Classroom</p>
      <h1 class="mt-1 font-display text-3xl font-bold text-primary">Classes and groups</h1>
      <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">Live cohorts with students, mentors, payment confirmation and access state.</p>
    </div>
    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
    <div class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
      <table class="data-table">
        <thead><tr><th>Class/group</th><th>Course</th><th>Students</th><th>Status</th><th>Action</th></tr></thead>
        <tbody>
          <tr v-if="isLoading"><td colspan="5" class="text-center text-on-surface-variant">Loading classes...</td></tr>
          <tr v-else-if="classes.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No classes yet. Create one from a course detail page.</td></tr>
          <tr v-for="item in classes" v-else :key="item.id">
            <td><p class="font-semibold text-primary">{{ item.name }}</p><p class="mt-1 text-xs text-on-surface-variant">{{ item.lesson_cadence }}</p></td>
            <td>{{ item.course_title }}</td>
            <td>{{ item.student_count }}</td>
            <td><span class="chip-approved">{{ item.status }}</span></td>
            <td><RouterLink class="btn-secondary" :to="`/app/classes/${item.id}`">Open</RouterLink></td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { classroomApi } from '@/api/classroom'
import { organizationApi } from '@/api/organization'
import type { ClassSummaryDto } from '@/types/classroom'

const classes = ref<ClassSummaryDto[]>([])
const isLoading = ref(true)
const errorMessage = ref('')

async function load() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const workspaces = await organizationApi.listMyWorkspaces()
    const workspace = workspaces.data.items[0]
    if (!workspace) { classes.value = []; return }
    const response = await classroomApi.listClasses({ organization_id: workspace.id, limit: 100 })
    classes.value = response.data.items
  } catch {
    errorMessage.value = 'Classes could not be loaded.'
  } finally {
    isLoading.value = false
  }
}

onMounted(load)
</script>
