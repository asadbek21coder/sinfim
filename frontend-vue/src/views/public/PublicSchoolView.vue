<template>
  <main>
    <section v-if="isLoading" class="mx-auto max-w-7xl px-5 py-16 lg:px-8">
      <p class="text-sm text-on-surface-variant">Loading school page...</p>
    </section>

    <section v-else-if="!school" class="mx-auto max-w-7xl px-5 py-16 lg:px-8">
      <p class="text-sm font-semibold uppercase tracking-wide text-secondary">School page</p>
      <h1 class="mt-2 font-display text-4xl font-bold text-primary">This school page is not available.</h1>
      <RouterLink class="btn-primary mt-6" to="/">Back to Sinfim.uz</RouterLink>
    </section>

    <template v-else>
      <section class="relative overflow-hidden bg-primary text-white">
        <img
          class="absolute inset-0 h-full w-full object-cover opacity-35"
          alt="Students learning in a classroom"
          src="https://images.unsplash.com/photo-1523580846011-d3a5bc25702b?auto=format&fit=crop&w=1800&q=80"
        >
        <div class="absolute inset-0 bg-primary/80" />
        <div class="relative mx-auto grid min-h-[62vh] max-w-7xl gap-10 px-5 py-16 lg:grid-cols-[1fr_420px] lg:px-8">
          <div class="flex flex-col justify-center">
            <p class="text-sm font-semibold uppercase tracking-wide text-white/65">{{ school.category || 'Online school' }}</p>
            <h1 class="mt-4 font-display text-5xl font-bold leading-tight lg:text-6xl">{{ school.name }}</h1>
            <p class="mt-5 max-w-2xl text-lg leading-8 text-white/78">
              {{ school.description || 'Kurslar, guruhlar va mentorlar bilan tartibli online ta\'lim.' }}
            </p>
            <div class="mt-7 flex flex-wrap gap-3 text-sm text-white/75">
              <span v-if="school.contact_phone" class="rounded-lg bg-white/10 px-4 py-2">{{ school.contact_phone }}</span>
              <a v-if="school.telegram_url" class="rounded-lg bg-white/10 px-4 py-2 hover:bg-white/15" :href="school.telegram_url" target="_blank">Telegram</a>
            </div>
          </div>

          <form class="self-center rounded-lg border border-white/15 bg-white p-6 text-on-surface shadow-xl" @submit.prevent="submitLead">
            <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Join request</p>
            <h2 class="mt-2 font-display text-2xl font-bold text-primary">Leave your phone number</h2>
            <p class="mt-2 text-sm leading-6 text-on-surface-variant">The school team will contact you about available groups and courses.</p>
            <div class="mt-5 space-y-4">
              <div>
                <label class="form-label" for="lead-name">Full name</label>
                <input id="lead-name" v-model.trim="leadForm.full_name" class="form-input" required placeholder="Ali Valiyev">
              </div>
              <div>
                <label class="form-label" for="lead-phone">Phone number</label>
                <input id="lead-phone" v-model.trim="leadForm.phone_number" class="form-input" required placeholder="+998 90 123 45 67">
              </div>
              <div>
                <label class="form-label" for="lead-note">Note</label>
                <textarea id="lead-note" v-model.trim="leadForm.note" class="form-textarea" rows="3" placeholder="Which course are you interested in?" />
              </div>
            </div>
            <div v-if="successMessage" class="mt-4 rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">{{ successMessage }}</div>
            <div v-if="errorMessage" class="mt-4 rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">{{ errorMessage }}</div>
            <button class="btn-primary mt-5 w-full justify-center py-3" type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? 'Sending...' : 'Send request' }}
            </button>
          </form>
        </div>
      </section>

      <section class="mx-auto max-w-7xl px-5 py-12 lg:px-8">
        <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Courses</p>
        <h2 class="mt-2 font-display text-3xl font-bold text-primary">Public course list comes next.</h2>
        <p class="mt-3 max-w-2xl text-sm leading-6 text-on-surface-variant">
          Course cards will appear here in Step 5 after course management is connected.
        </p>
      </section>
    </template>
  </main>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { organizationApi } from '@/api/organization'
import type { OrganizationDto } from '@/types/organization'

const route = useRoute()
const isLoading = ref(true)
const isSubmitting = ref(false)
const school = ref<OrganizationDto | null>(null)
const successMessage = ref('')
const errorMessage = ref('')

const leadForm = reactive({
  full_name: '',
  phone_number: '',
  note: '',
})

async function loadSchool() {
  isLoading.value = true
  try {
    const response = await organizationApi.getPublicSchoolPage(String(route.params.schoolSlug))
    school.value = response.data.organization
  } catch {
    school.value = null
  } finally {
    isLoading.value = false
  }
}

async function submitLead() {
  if (!school.value) return
  isSubmitting.value = true
  successMessage.value = ''
  errorMessage.value = ''
  try {
    const response = await organizationApi.createLead({
      organization_id: school.value.id,
      full_name: leadForm.full_name,
      phone_number: leadForm.phone_number,
      note: leadForm.note || undefined,
    })
    successMessage.value = response.data.message
    leadForm.full_name = ''
    leadForm.phone_number = ''
    leadForm.note = ''
  } catch {
    errorMessage.value = 'Request could not be sent. Check the fields and try again.'
  } finally {
    isSubmitting.value = false
  }
}

onMounted(loadSchool)
</script>
