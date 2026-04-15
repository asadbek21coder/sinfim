<template>
  <section class="space-y-6">
    <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Platform admin</p>
        <h1 class="mt-1 font-display text-3xl font-bold text-primary">Create organization</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Create a school workspace and give the owner a temporary first-login password.
        </p>
      </div>
      <RouterLink class="btn-secondary" to="/admin/school-requests">School requests</RouterLink>
    </div>

    <form class="grid gap-6 xl:grid-cols-[1.1fr_0.9fr]" @submit.prevent="submit">
      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">School workspace</h2>
          <p class="mt-1 text-sm text-on-surface-variant">This becomes the school profile at sinfim.uz/{slug}.</p>
        </div>

        <div class="grid gap-5 sm:grid-cols-2">
          <div>
            <label class="form-label" for="name">School name</label>
            <input id="name" v-model.trim="form.name" class="form-input" placeholder="Sinfim Academy" required @input="syncSlug">
          </div>
          <div>
            <label class="form-label" for="slug">Slug</label>
            <input id="slug" v-model.trim="form.slug" class="form-input" placeholder="sinfim-academy" required>
          </div>
          <div class="sm:col-span-2">
            <label class="form-label" for="description">Description</label>
            <textarea id="description" v-model.trim="form.description" class="form-textarea" placeholder="Online kurs markazi" rows="4" />
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
          <label class="flex items-center gap-3 rounded-lg border border-outline-variant bg-surface-container-low px-4 py-3 text-sm font-medium text-on-surface sm:col-span-2">
            <input v-model="form.is_demo" class="h-4 w-4" type="checkbox">
            Demo school
          </label>
        </div>
      </section>

      <section class="card space-y-5">
        <div>
          <h2 class="font-display text-xl font-bold text-primary">Owner access</h2>
          <p class="mt-1 text-sm text-on-surface-variant">The owner will change this password on first login.</p>
        </div>

        <div>
          <label class="form-label" for="owner-name">Owner full name</label>
          <input id="owner-name" v-model.trim="form.owner.full_name" class="form-input" placeholder="Ali Valiyev" required>
        </div>
        <div>
          <label class="form-label" for="owner-phone">Owner phone</label>
          <input id="owner-phone" v-model.trim="form.owner.phone_number" class="form-input" placeholder="+998 90 111 22 33" required>
        </div>
        <div>
          <label class="form-label" for="temp-password">Temporary password</label>
          <input id="temp-password" v-model="form.owner.temporary_password" class="form-input" minlength="8" placeholder="TempPass123" required type="text">
        </div>

        <div v-if="errorMessage" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
          {{ errorMessage }}
        </div>

        <div v-if="created" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-800">
          <p class="font-semibold">Organization created.</p>
          <p class="mt-1">{{ created.organization.name }} owner: {{ created.owner.phone_number }}</p>
          <p class="mt-1">Temporary password must be changed on first login.</p>
        </div>

        <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? 'Creating...' : 'Create organization' }}
        </button>
      </section>
    </form>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { RouterLink } from 'vue-router'
import { AxiosError } from 'axios'
import { organizationApi } from '@/api/organization'
import type { CreateOrganizationPayload, CreateOrganizationResponse } from '@/types/organization'

const form = reactive<CreateOrganizationPayload>({
  name: '',
  slug: '',
  description: '',
  logo_url: '',
  category: '',
  contact_phone: '',
  telegram_url: '',
  is_demo: false,
  owner: {
    full_name: '',
    phone_number: '',
    temporary_password: 'TempPass123',
  },
})

const isSubmitting = ref(false)
const errorMessage = ref('')
const created = ref<CreateOrganizationResponse | null>(null)
const touchedSlug = ref(false)

function syncSlug() {
  if (touchedSlug.value || form.slug) return
  form.slug = slugify(form.name)
}

function slugify(value: string) {
  return value
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '')
}

function optional(value?: string) {
  return value && value.trim() ? value.trim() : undefined
}

async function submit() {
  isSubmitting.value = true
  errorMessage.value = ''
  created.value = null

  try {
    const payload: CreateOrganizationPayload = {
      name: form.name,
      slug: slugify(form.slug),
      description: optional(form.description),
      logo_url: optional(form.logo_url),
      category: optional(form.category),
      contact_phone: optional(form.contact_phone),
      telegram_url: optional(form.telegram_url),
      is_demo: form.is_demo,
      owner: {
        full_name: form.owner.full_name,
        phone_number: form.owner.phone_number,
        temporary_password: form.owner.temporary_password,
      },
    }
    const response = await organizationApi.createOrganization(payload)
    created.value = response.data
  } catch (error) {
    if (error instanceof AxiosError && error.response?.status === 409) {
      errorMessage.value = 'This slug is already taken.'
    } else if (error instanceof AxiosError && error.response?.status === 401) {
      errorMessage.value = 'Please log in again.'
    } else if (error instanceof AxiosError && error.response?.status === 403) {
      errorMessage.value = 'Only platform admins can create organizations.'
    } else {
      errorMessage.value = 'Organization could not be created. Check the fields and try again.'
    }
  } finally {
    isSubmitting.value = false
  }
}
</script>
