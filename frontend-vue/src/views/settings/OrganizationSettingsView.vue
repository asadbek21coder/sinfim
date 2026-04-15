<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Settings</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">Organization settings</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Update the school profile that will power the workspace and public page.
        </p>
      </div>
      <a v-if="workspace" class="btn-secondary" :href="`/${workspace.slug}`" target="_blank">Public page</a>
    </div>

    <div v-if="isLoading" class="card text-sm text-on-surface-variant">Loading workspace...</div>

    <div v-else-if="!workspace" class="card">
      <h2 class="font-display text-xl font-bold text-primary">No workspace assigned</h2>
      <p class="mt-2 text-sm text-on-surface-variant">Ask a platform admin to create an organization and assign you as owner.</p>
    </div>

    <form v-else class="grid gap-6 lg:grid-cols-[1.1fr_0.9fr]" @submit.prevent="submit">
      <section class="card space-y-5">
        <div class="grid gap-5 sm:grid-cols-2">
          <div>
            <label class="form-label" for="name">School name</label>
            <input id="name" v-model.trim="form.name" class="form-input" required>
          </div>
          <div>
            <label class="form-label" for="slug">Slug</label>
            <input id="slug" v-model="workspace.slug" class="form-input" disabled>
          </div>
          <div class="sm:col-span-2">
            <label class="form-label" for="description">Description</label>
            <textarea id="description" v-model.trim="form.description" class="form-textarea" rows="5" />
          </div>
          <div>
            <label class="form-label" for="category">Category</label>
            <select id="category" v-model="form.category" class="form-select">
              <option value="">Not selected</option>
              <option value="language">Language</option>
              <option value="school">School subjects</option>
              <option value="it">IT and profession</option>
              <option value="other">Other</option>
            </select>
          </div>
          <div>
            <label class="form-label" for="public-status">Public status</label>
            <select id="public-status" v-model="form.public_status" class="form-select">
              <option value="draft">Draft</option>
              <option value="public">Public</option>
              <option value="hidden">Hidden</option>
            </select>
          </div>
        </div>
      </section>

      <section class="card space-y-5">
        <div>
          <label class="form-label" for="contact-phone">Contact phone</label>
          <input id="contact-phone" v-model.trim="form.contact_phone" class="form-input" placeholder="+998 90 123 45 67">
        </div>
        <div>
          <label class="form-label" for="telegram-url">Telegram URL</label>
          <input id="telegram-url" v-model.trim="form.telegram_url" class="form-input" placeholder="https://t.me/school">
        </div>
        <div>
          <label class="form-label" for="logo-url">Logo URL</label>
          <input id="logo-url" v-model.trim="form.logo_url" class="form-input" placeholder="https://...">
        </div>
        <label class="flex items-center gap-3 rounded-lg border border-outline-variant bg-surface-container-low px-4 py-3 text-sm font-medium text-on-surface">
          <input v-model="form.is_demo" class="h-4 w-4" type="checkbox">
          Demo school
        </label>

        <div v-if="successMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">
          {{ successMessage }}
        </div>
        <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
          {{ errorMessage }}
        </div>
        <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? 'Saving...' : 'Save settings' }}
        </button>
      </section>
    </form>
  </section>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { organizationApi } from '@/api/organization'
import type { UpdateOrganizationPayload, WorkspaceDto } from '@/types/organization'

const isLoading = ref(true)
const isSubmitting = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const workspace = ref<WorkspaceDto | null>(null)

const form = reactive<UpdateOrganizationPayload>({
  id: '',
  name: '',
  description: '',
  logo_url: '',
  category: '',
  contact_phone: '',
  telegram_url: '',
  public_status: 'draft',
  is_demo: false,
})

function fillForm(item: WorkspaceDto) {
  form.id = item.id
  form.name = item.name
  form.description = item.description ?? ''
  form.logo_url = item.logo_url ?? ''
  form.category = item.category ?? ''
  form.contact_phone = item.contact_phone ?? ''
  form.telegram_url = item.telegram_url ?? ''
  form.public_status = item.public_status
  form.is_demo = item.is_demo
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

async function loadWorkspace() {
  isLoading.value = true
  const response = await organizationApi.listMyWorkspaces()
  workspace.value = response.data.items[0] ?? null
  if (workspace.value) fillForm(workspace.value)
  isLoading.value = false
}

async function submit() {
  isSubmitting.value = true
  errorMessage.value = ''
  successMessage.value = ''
  try {
    const response = await organizationApi.updateOrganization({
      ...form,
      description: optional(form.description),
      logo_url: optional(form.logo_url),
      category: optional(form.category),
      contact_phone: optional(form.contact_phone),
      telegram_url: optional(form.telegram_url),
    })
    if (workspace.value) {
      workspace.value = { ...workspace.value, ...response.data.item }
      fillForm(workspace.value)
    }
    successMessage.value = 'Organization settings saved.'
  } catch {
    errorMessage.value = 'Settings could not be saved.'
  } finally {
    isSubmitting.value = false
  }
}

onMounted(loadWorkspace)
</script>
