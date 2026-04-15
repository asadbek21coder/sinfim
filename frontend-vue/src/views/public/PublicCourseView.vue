<template>
  <section class="bg-surface text-on-surface">
    <div v-if="isLoading" class="mx-auto flex min-h-[60vh] max-w-6xl items-center px-6 text-on-surface-variant">Loading course...</div>
    <div v-else-if="errorMessage" class="mx-auto flex min-h-[60vh] max-w-6xl items-center px-6">
      <div>
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Sinfim.uz</p>
        <h1 class="mt-3 font-display text-4xl font-bold text-primary">Course is not available</h1>
        <p class="mt-3 text-on-surface-variant">{{ errorMessage }}</p>
      </div>
    </div>

    <div v-else-if="page" class="mx-auto max-w-6xl px-6 py-10 lg:py-14">
      <div class="grid gap-8 lg:grid-cols-[1fr_380px] lg:items-start">
        <section class="overflow-hidden rounded-lg border border-outline-variant bg-surface-container-lowest shadow-sm">
          <img
            class="h-72 w-full object-cover"
            alt="Students learning online"
            src="https://images.unsplash.com/photo-1522202176988-66273c2fd55f?auto=format&fit=crop&w=1600&q=80"
          >
          <div class="p-6 lg:p-8">
            <p class="text-sm font-semibold uppercase tracking-wide text-secondary">{{ page.organization.name }}</p>
            <h1 class="mt-3 font-display text-4xl font-bold tracking-tight text-primary lg:text-5xl">{{ page.course.title }}</h1>
            <p class="mt-4 max-w-3xl text-base leading-7 text-on-surface-variant">
              {{ page.course.description || 'Mentor nazorati, tartibli darslar va bosqichma-bosqich oquv jarayoni.' }}
            </p>
            <div class="mt-6 flex flex-wrap gap-2">
              <span v-if="page.course.category" class="chip-info">{{ page.course.category }}</span>
              <span v-if="page.course.level" class="chip-approved">{{ page.course.level }}</span>
              <span class="chip-pending">Online course</span>
            </div>
          </div>
        </section>

        <aside class="card sticky top-6 space-y-5">
          <div>
            <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Apply</p>
            <h2 class="mt-2 font-display text-2xl font-bold text-primary">Join this course</h2>
            <p class="mt-2 text-sm leading-6 text-on-surface-variant">Leave your phone number and the school team will contact you.</p>
          </div>

          <form class="space-y-4" @submit.prevent="submitLead">
            <div>
              <label class="form-label" for="full-name">Full name</label>
              <input id="full-name" v-model.trim="lead.full_name" class="form-input" required placeholder="Ali Valiyev">
            </div>
            <div>
              <label class="form-label" for="phone">Phone number</label>
              <input id="phone" v-model.trim="lead.phone_number" class="form-input" required placeholder="+998 90 123 45 67">
            </div>
            <div>
              <label class="form-label" for="note">Note</label>
              <textarea id="note" v-model.trim="lead.note" class="form-textarea" rows="3" placeholder="Preferred group or schedule" />
            </div>

            <div v-if="leadMessage" class="rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ leadMessage }}</div>
            <div v-if="leadError" class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ leadError }}</div>

            <button class="btn-primary w-full justify-center py-3" type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? 'Sending...' : 'Apply for course' }}
            </button>
          </form>
        </aside>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import { catalogApi } from '@/api/catalog'
import { organizationApi } from '@/api/organization'
import type { PublicCoursePageResponse } from '@/types/catalog'

const route = useRoute()
const page = ref<PublicCoursePageResponse | null>(null)
const isLoading = ref(true)
const isSubmitting = ref(false)
const errorMessage = ref('')
const leadMessage = ref('')
const leadError = ref('')

const lead = reactive({ full_name: '', phone_number: '', note: '' })

async function load() {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const response = await catalogApi.getPublicCoursePage(String(route.params.schoolSlug), String(route.params.courseSlug))
    page.value = response.data
  } catch {
    errorMessage.value = 'This course may be draft, hidden, or the link may be wrong.'
  } finally {
    isLoading.value = false
  }
}

async function submitLead() {
  if (!page.value) return
  isSubmitting.value = true
  leadMessage.value = ''
  leadError.value = ''
  try {
    const note = [lead.note, `Course: ${page.value.course.title}`].filter(Boolean).join('\n')
    const response = await organizationApi.createLead({
      organization_id: page.value.organization.id,
      full_name: lead.full_name,
      phone_number: lead.phone_number,
      note,
    })
    leadMessage.value = response.data.message
    lead.full_name = ''
    lead.phone_number = ''
    lead.note = ''
  } catch {
    leadError.value = 'Request could not be sent. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}

onMounted(load)
</script>
