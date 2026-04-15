<template>
  <main class="mx-auto grid max-w-7xl gap-10 px-5 py-14 lg:grid-cols-[0.9fr_1.1fr] lg:px-8">
    <section>
      <p class="text-sm font-semibold uppercase tracking-wide text-secondary">Maktab arizasi</p>
      <h1 class="mt-2 font-display text-4xl font-bold text-primary lg:text-5xl">Sinfim.uz ish joyini ochish uchun ariza qoldiring.</h1>
      <p class="mt-4 text-base leading-7 text-on-surface-variant">
        MVP davrida yangi maktablarni platforma administratori tasdiqlaydi. Telefon raqamingizni qoldiring, keyin siz bilan bog'lanamiz.
      </p>
      <div class="mt-8 grid gap-3 text-sm text-on-surface-variant">
        <p class="rounded-lg border border-outline-variant bg-surface-container-lowest p-4">To'lov va hujjatlar avvalgidek tashqarida kelishiladi.</p>
        <p class="rounded-lg border border-outline-variant bg-surface-container-lowest p-4">Tasdiqdan keyin kurs, sinf, mentor va o'quvchilarni boshqarish paneli ochiladi.</p>
      </div>
    </section>

    <form class="rounded-lg border border-outline-variant bg-surface-container-lowest p-6 shadow-sm" @submit.prevent="submit">
      <div class="grid gap-5 sm:grid-cols-2">
        <div>
          <label class="form-label" for="full-name">Ism familiya</label>
          <input id="full-name" v-model.trim="form.full_name" class="form-input" placeholder="Ali Valiyev" required>
        </div>
        <div>
          <label class="form-label" for="phone-number">Telefon raqam</label>
          <input id="phone-number" v-model.trim="form.phone_number" class="form-input" placeholder="+998 90 000 00 00" required>
        </div>
        <div class="sm:col-span-2">
          <label class="form-label" for="school-name">Maktab yoki brend nomi</label>
          <input id="school-name" v-model.trim="form.school_name" class="form-input" placeholder="Sinfim Academy" required>
        </div>
        <div>
          <label class="form-label" for="category">Yo'nalish</label>
          <select id="category" v-model="form.category" class="form-select">
            <option value="">Tanlanmagan</option>
            <option value="language">Til kurslari</option>
            <option value="school">Maktab fanlari</option>
            <option value="it">IT va kasb</option>
            <option value="other">Boshqa</option>
          </select>
        </div>
        <div>
          <label class="form-label" for="student-count">Taxminiy o'quvchi soni</label>
          <input id="student-count" v-model.number="form.student_count" class="form-input" min="0" placeholder="120" type="number">
        </div>
        <div class="sm:col-span-2">
          <label class="form-label" for="note">Izoh</label>
          <textarea id="note" v-model.trim="form.note" class="form-textarea" placeholder="Qanday kurs sotasiz, hozir Telegramda qanday ishlaysiz?" rows="5" />
        </div>
      </div>
      <div v-if="successMessage" class="mt-5 rounded-lg border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm font-medium text-emerald-800">
        {{ successMessage }}
      </div>
      <div v-if="errorMessage" class="mt-5 rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
        {{ errorMessage }}
      </div>
      <button class="btn-primary mt-6" type="submit" :disabled="isSubmitting">
        {{ isSubmitting ? 'Yuborilmoqda...' : 'Ariza yuborish' }}
      </button>
    </form>
  </main>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { AxiosError } from 'axios'
import { organizationApi } from '@/api/organization'
import type { CreateSchoolRequestPayload } from '@/types/organization'

const form = reactive<CreateSchoolRequestPayload>({
  full_name: '',
  phone_number: '',
  school_name: '',
  category: '',
  student_count: undefined,
  note: '',
})

const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

async function submit() {
  isSubmitting.value = true
  successMessage.value = ''
  errorMessage.value = ''

  try {
    const payload: CreateSchoolRequestPayload = {
      full_name: form.full_name,
      phone_number: form.phone_number,
      school_name: form.school_name,
      category: form.category || undefined,
      student_count: typeof form.student_count === 'number' ? form.student_count : undefined,
      note: form.note || undefined,
    }
    const response = await organizationApi.createSchoolRequest(payload)
    successMessage.value = response.data.message
    form.full_name = ''
    form.phone_number = ''
    form.school_name = ''
    form.category = ''
    form.student_count = undefined
    form.note = ''
  } catch (error) {
    if (error instanceof AxiosError && error.response?.status === 400) {
      errorMessage.value = 'Ma\'lumotlarni tekshirib qayta yuboring.'
    } else {
      errorMessage.value = 'Arizani yuborishda xatolik bo\'ldi. Birozdan keyin qayta urinib ko\'ring.'
    }
  } finally {
    isSubmitting.value = false
  }
}
</script>
