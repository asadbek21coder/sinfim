<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Platform admin</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">School requests</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Public formdan kelgan maktab arizalari shu yerda ko'rinadi. Hozircha bu preview ekran; Step 2 auth bilan yopiladi.
        </p>
      </div>
      <div class="flex flex-wrap gap-2">
        <button class="btn-secondary" type="button" @click="selectedStatus = undefined; loadRequests()">All</button>
        <button v-for="status in statuses" :key="status" class="btn-secondary" type="button" @click="selectedStatus = status; loadRequests()">
          {{ labelStatus(status) }}
        </button>
      </div>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
      {{ errorMessage }}
    </div>

    <div class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
      <table class="data-table">
        <thead>
          <tr>
            <th>School</th>
            <th>Owner</th>
            <th>Category</th>
            <th>Status</th>
            <th>Created</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="isLoading">
            <td colspan="6" class="text-center text-on-surface-variant">Loading requests...</td>
          </tr>
          <tr v-else-if="requests.length === 0">
            <td colspan="6" class="text-center text-on-surface-variant">No school requests yet.</td>
          </tr>
          <tr v-for="request in requests" v-else :key="request.id">
            <td>
              <p class="font-semibold text-primary">{{ request.school_name }}</p>
              <p v-if="request.note" class="mt-1 max-w-xs truncate text-xs text-on-surface-variant">{{ request.note }}</p>
            </td>
            <td>
              <p class="font-medium">{{ request.full_name }}</p>
              <p class="mt-1 text-xs text-on-surface-variant">{{ request.phone_number }}</p>
            </td>
            <td>
              <p>{{ request.category || 'Not set' }}</p>
              <p v-if="request.student_count !== null && request.student_count !== undefined" class="mt-1 text-xs text-on-surface-variant">
                {{ request.student_count }} students
              </p>
            </td>
            <td><span :class="chipClass(request.status)">{{ labelStatus(request.status) }}</span></td>
            <td>{{ formatDate(request.created_at) }}</td>
            <td>
              <select class="form-select min-w-36" :value="request.status" :disabled="updatingId === request.id" @change="updateStatus(request.id, ($event.target as HTMLSelectElement).value)">
                <option v-for="status in statuses" :key="status" :value="status">{{ labelStatus(status) }}</option>
              </select>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { organizationApi } from '@/api/organization'
import type { SchoolRequestDto, SchoolRequestStatus } from '@/types/organization'

const statuses: SchoolRequestStatus[] = ['new', 'contacted', 'approved', 'rejected']
const requests = ref<SchoolRequestDto[]>([])
const selectedStatus = ref<SchoolRequestStatus | undefined>()
const isLoading = ref(false)
const updatingId = ref<string | null>(null)
const errorMessage = ref('')

async function loadRequests() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const response = await organizationApi.listSchoolRequests({ status: selectedStatus.value, limit: 100 })
    requests.value = response.data.items
  } catch {
    errorMessage.value = 'School request list could not be loaded.'
  } finally {
    isLoading.value = false
  }
}

async function updateStatus(id: string, status: string) {
  updatingId.value = id
  errorMessage.value = ''
  try {
    const response = await organizationApi.updateSchoolRequestStatus({ id, status: status as SchoolRequestStatus })
    requests.value = requests.value.map((item) => item.id === id ? response.data.item : item)
  } catch {
    errorMessage.value = 'Status could not be updated.'
  } finally {
    updatingId.value = null
  }
}

function labelStatus(status: SchoolRequestStatus) {
  return {
    new: 'New',
    contacted: 'Contacted',
    approved: 'Approved',
    rejected: 'Rejected',
  }[status]
}

function chipClass(status: SchoolRequestStatus) {
  return {
    new: 'chip-pending',
    contacted: 'chip-info',
    approved: 'chip-approved',
    rejected: 'chip-rejected',
  }[status]
}

function formatDate(value: string) {
  return new Intl.DateTimeFormat('uz-UZ', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(value))
}

onMounted(loadRequests)
</script>
