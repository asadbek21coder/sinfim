<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Growth</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">Leads</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Potential students from your public school page appear here.
        </p>
      </div>
      <a v-if="workspace" class="btn-secondary" :href="`/${workspace.slug}`" target="_blank">Open public page</a>
    </div>

    <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>

    <div class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
      <table class="data-table">
        <thead>
          <tr>
            <th>Lead</th>
            <th>Note</th>
            <th>Status</th>
            <th>Created</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="isLoading"><td colspan="5" class="text-center text-on-surface-variant">Loading leads...</td></tr>
          <tr v-else-if="leads.length === 0"><td colspan="5" class="text-center text-on-surface-variant">No leads yet.</td></tr>
          <tr v-for="lead in leads" v-else :key="lead.id">
            <td>
              <p class="font-semibold text-primary">{{ lead.full_name }}</p>
              <p class="mt-1 text-xs text-on-surface-variant">{{ lead.phone_number }}</p>
            </td>
            <td class="max-w-sm"><p class="truncate">{{ lead.note || 'No note' }}</p></td>
            <td><span :class="chipClass(lead.status)">{{ labelStatus(lead.status) }}</span></td>
            <td>{{ formatDate(lead.created_at) }}</td>
            <td>
              <select class="form-select min-w-36" :value="lead.status" :disabled="updatingId === lead.id" @change="updateStatus(lead.id, ($event.target as HTMLSelectElement).value)">
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
import type { LeadDto, LeadStatus, WorkspaceDto } from '@/types/organization'

const statuses: LeadStatus[] = ['new', 'contacted', 'converted', 'archived']
const workspace = ref<WorkspaceDto | null>(null)
const leads = ref<LeadDto[]>([])
const isLoading = ref(true)
const updatingId = ref<string | null>(null)
const errorMessage = ref('')

async function load() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const workspaces = await organizationApi.listMyWorkspaces()
    workspace.value = workspaces.data.items[0] ?? null
    if (!workspace.value) {
      leads.value = []
      return
    }
    const response = await organizationApi.listLeads({ organization_id: workspace.value.id, limit: 100 })
    leads.value = response.data.items
  } catch {
    errorMessage.value = 'Leads could not be loaded.'
  } finally {
    isLoading.value = false
  }
}

async function updateStatus(id: string, status: string) {
  updatingId.value = id
  errorMessage.value = ''
  try {
    const response = await organizationApi.updateLeadStatus({ id, status: status as LeadStatus })
    leads.value = leads.value.map((item) => item.id === id ? response.data.item : item)
  } catch {
    errorMessage.value = 'Lead status could not be updated.'
  } finally {
    updatingId.value = null
  }
}

function labelStatus(status: LeadStatus) {
  return {
    new: 'New',
    contacted: 'Contacted',
    converted: 'Converted',
    archived: 'Archived',
  }[status]
}

function chipClass(status: LeadStatus) {
  return {
    new: 'chip-pending',
    contacted: 'chip-info',
    converted: 'chip-approved',
    archived: 'chip-rejected',
  }[status]
}

function formatDate(value: string) {
  return new Intl.DateTimeFormat('uz-UZ', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(value))
}

onMounted(load)
</script>
